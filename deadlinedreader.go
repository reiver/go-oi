package oi

import (
	"context"
	"io"
	"time"

	"codeberg.org/reiver/go-erorr"
)

// DeadLinedReader represents an [io.Reader] with a (particular type of) SetReadDeadline method.
//
// One example of this is [net.Conn].
type DeadLinedReader interface {
	io.Reader
	SetReadDeadline(t time.Time) error
}

type internalDeadLinedReaderWrapper struct {
	deadLinedReader DeadLinedReader
}

var _ Reader = internalDeadLinedReaderWrapper{}

// CastDeadLinedReader returns a [Reader] based on a [DeadLinedReder] (such as [net.Conn]).
func CastDeadLinedReader(dr DeadLinedReader) Reader {
        if nil == dr {
                return nil
        }

        return internalDeadLinedReaderWrapper{
                deadLinedReader: dr,
        }
}

func (receiver internalDeadLinedReaderWrapper) Read(ctx context.Context, p []byte) (n int, err error) {
	if nil == receiver.deadLinedReader {
		var nada int
		return nada, ErrDeadLinedReaderNil
	}

	if nil == ctx {
		ctx = context.Background()
	}

	if err := ctx.Err(); err != nil {
		var nada int
		err = erorr.Errors{ErrContextInvalid, err}
		return nada, erorr.Wrap(err, "failed to read due to invalid context")
	}

	if deadline, ok := ctx.Deadline(); ok {
		// Intentionally mostly ignoring the error from SetReadDeadline —
		// not all net.Conn implementations support deadlines, and
		// the read itself will surface any real failures.
		//
		// Note: the defer clears the deadline to zero after the read completes.
		// This means any deadline previously set on the underlying conn by an
		// outer layer will be erased. This is acceptable because oi.Reader
		// owns the deadline for the duration of each Read call, and net.Conn
		// provides no way to read back the current deadline to restore it.
		if nil == receiver.deadLinedReader.SetReadDeadline(deadline) {
			defer receiver.deadLinedReader.SetReadDeadline(time.Time{})
		}
	}

	if err := ctx.Err(); err != nil {
		var nada int
		err = erorr.Errors{ErrContextInvalid, err}
		return nada, erorr.Wrap(err, "failed to read due to invalid context")
	}
	return receiver.deadLinedReader.Read(p)
}
