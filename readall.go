package oi

import (
	"context"
	"io"

	"codeberg.org/reiver/go-erorr"
)

func ReadAll(ctx context.Context, reader Reader) ([]byte, error) {
	if nil == reader {
		return nil, ErrReaderNil
	}
	if nil == ctx {
		ctx = context.Background()
	}

	if err := ctx.Err(); err != nil {
		err = erorr.Errors{ErrContextInvalid, err}
		return nil, erorr.Wrap(err, "failed to read due to invalid context")
	}

	return io.ReadAll(ClassicReader(ctx, reader))
}
