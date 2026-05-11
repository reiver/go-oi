package oi

import (
	"context"
	"io"
	"time"

	"codeberg.org/reiver/go-erorr"
)

// DeadLinedWriter represents an [io.Writer] with a (particular type of) SetWriteDeadline method.
//
// One example of this is [net.Conn].
type DeadLinedWriter interface {
	io.Writer
	SetWriteDeadline(t time.Time) error
}

type internalDeadLinedWriterWrapper struct {
	deadLinedWriter DeadLinedWriter
}

var _ Writer = internalDeadLinedWriterWrapper{}

// CastDeadLinedWriter returns a [Writer] based on a [DeadLinedWriter] (such as [net.Conn]).
func CreateWriter(dw DeadLinedWriter) Writer {
	if nil == dw {
		return nil
	}

	return internalDeadLinedWriterWrapper{
		deadLinedWriter: dw,
	}
}

func (receiver internalDeadLinedWriterWrapper) Write(ctx context.Context, p []byte) (n int, err error) {
	if nil == receiver.deadLinedWriter {
		var nada int
		return nada, ErrDeadLinedWriterNil
	}

	if nil == ctx {
		ctx = context.Background()
	}

	if err := ctx.Err(); err != nil {
		var nada int
		err = erorr.Errors{ErrContextInvalid, err}
		return nada, erorr.Wrap(err, "failed to write due to invalid context")
	}

	if deadline, ok := ctx.Deadline(); ok {
		// Intentionally mostly ignoring the error from SetWriteDeadline —
		// not all net.Conn implementations support deadlines, and
		// the write itself will surface any real failures.
		//
		// Note: the defer clears the deadline to zero after the write completes.
		// This means any deadline previously set on the underlying conn by an
		// outer layer will be erased. This is acceptable because oi.Writer
		// owns the deadline for the duration of each Write call, and net.Conn
		// provides no way to read back the current deadline to restore it.
		if nil == receiver.deadLinedWriter.SetWriteDeadline(deadline) {
			defer receiver.deadLinedWriter.SetWriteDeadline(time.Time{})
		}
	}

	if err := ctx.Err(); err != nil {
		var nada int
		err = erorr.Errors{ErrContextInvalid, err}
		return nada, erorr.Wrap(err, "failed to write due to invalid context")
	}
	return receiver.deadLinedWriter.Write(p)
}
