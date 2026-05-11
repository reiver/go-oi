package oi

import (
	"fmt"
	"io"
)

func ReadSeeker(readerAt io.ReaderAt) io.ReadSeeker {
	if nil == readerAt {
		return nil
	}

	rs := internalReadSeeker{
		readerAt:readerAt,
	}

	return &rs
}

type internalReadSeeker struct {
	readerAt io.ReaderAt
	offset   int64
}

func (receiver *internalReadSeeker) Read(p []byte) (n int, err error) {
	if nil == receiver {
		return 0, ErrReceiverNil
	}

	readerAt := receiver.readerAt
	if nil == readerAt {
		return 0, ErrReaderAtNil
	}

	n, err = readerAt.ReadAt(p, receiver.offset)
	receiver.offset += int64(n)

	return n, err
}

func (receiver *internalReadSeeker) Seek(offset int64, whence int) (int64, error) {
	if nil == receiver {
		return 0, ErrReceiverNil
	}

	readerAt := receiver.readerAt
	if nil == readerAt {
		return 0, ErrReaderAtNil
	}

	var size int64 = -1
	func(){
		sizer, casted := readerAt.(interface{Size() int64})
		if !casted {
			return
		}

		size = sizer.Size()
	}()

	var absolute int64
	var whenceName string
	switch whence {
	default:
		return 0, fmt.Errorf("oi: Invalid Whence: %d", whence)

	case io.SeekStart:
		whenceName = "Seek Start"
		absolute = offset

	case io.SeekCurrent:
		whenceName = "Seek Current"
		absolute = receiver.offset + offset

	case io.SeekEnd:
		whenceName = "Seek End"
		if 0 > size {
			return 0, fmt.Errorf("oi: Unsupported Whence: %d (%s)", whence, whenceName)
		}
		absolute = size + offset
	}

	if absolute < 0 {
		return 0, errInvalidOffsetf(offset, whence, "resulting absolute offset (%d) is less than zero (0)", absolute)
	}
	if 0 <= size {
		if size < absolute {
			return 0, errInvalidOffsetf(offset, whence, "resulting absolute offset (%d) is larger than size (%d)", absolute, size)
		}
	}

	receiver.offset = absolute

	return receiver.offset, nil
}
