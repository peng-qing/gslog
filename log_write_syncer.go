package gslog

import "io"

type WriteSyncer interface {
	io.WriteCloser
	Sync() error
}
