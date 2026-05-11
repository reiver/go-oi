package oi

import (
	"codeberg.org/reiver/go-erorr"
)


const (
	ErrReaderAtNil = erorr.Error("nil io.ReaderAt")
	ErrReceiverNil = erorr.Error("nil receiver")
)
