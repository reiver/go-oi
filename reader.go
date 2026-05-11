package oi

import (
	"context"
)

// Reader represents a modernized version of [io.Reader], where the Read method has a [context.Context] as its first parameter.
type Reader interface {
	Read(ctx context.Context, p []byte) (n int, err error)
}

