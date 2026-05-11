package oi

import (
	"codeberg.org/reiver/go-erorr"
)


const (
	ErrContextInvalid     = erorr.Error("context invalid")
	ErrDeadLinedWriterNil = erorr.Error("nil oi.DeadLinedWriter")
	ErrReaderAtNil        = erorr.Error("nil io.ReaderAt")
	ErrReceiverNil        = erorr.Error("nil receiver")
)
