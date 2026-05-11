package oi

import (
	"context"
)

// Writer represents a modernized version of [io.Writer], where the Write method has a [context.Context] as its first parameter.
type Writer interface {
	Write(ctx context.Context, p []byte) (n int, err error)
}
