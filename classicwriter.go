package oi

import (
	"context"
	"io"
)

func ClassicWriter(ctx context.Context, writer Writer) io.Writer {
	if nil == writer {
		return nil
	}

	return internalClassicWriter{
		ctx:    ctx,
		writer: writer,
	}
}

type internalClassicWriter struct {
	ctx    context.Context
	writer Writer
}

var _ io.Writer = internalClassicWriter{}

func (receiver internalClassicWriter) Write(p []byte) (n int, err error) {
	if nil == receiver.writer {
		var nada int
		return nada, ErrWriterNil
	}

	ctx := receiver.ctx
	if nil == ctx {
		ctx = context.Background()
	}

	return receiver.writer.Write(ctx, p)
}
