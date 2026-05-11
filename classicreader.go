package oi

import (
	"context"
	"io"
)

func ClassicReader(ctx context.Context, reader Reader) io.Reader {
	if nil == reader {
		return nil
	}

	return internalClassicReader{
		ctx:    ctx,
		reader: reader,
	}
}

type internalClassicReader struct {
	ctx    context.Context
	reader Reader
}

var _ io.Reader = internalClassicReader{}

func (receiver internalClassicReader) Read(p []byte) (n int, err error) {
	if nil == receiver.reader {
		var nada int
		return nada, ErrReaderNil
	}

	ctx := receiver.ctx
	if nil == ctx {
		ctx = context.Background()
	}

	return receiver.reader.Read(ctx, p)
}
