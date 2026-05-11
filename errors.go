package oi

import (
	"codeberg.org/reiver/go-erorr"
)


const (
	ErrContextInvalid     = erorr.Error("context invalid")
	ErrDeadLinedReaderNil = erorr.Error("nil oi.DeadLinedReader")
	ErrDeadLinedWriterNil = erorr.Error("nil oi.DeadLinedWriter")
	ErrReaderAtNil        = erorr.Error("nil io.ReaderAt")
	ErrReceiverNil        = erorr.Error("nil receiver")
)
