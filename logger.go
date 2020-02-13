package log

import (
	"io"

	kitlog "github.com/go-kit/kit/log"
)

// Logger is what any Tendermint library should take.
type Logger interface {
	Debug(msg string, keyvals ...interface{})
	Info(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})

	With(keyvals ...interface{}) Logger

	// add
	Debugf(msg string, vals ...interface{})
	Infof(msg string, vals ...interface{})
	Errorf(msg string, vals ...interface{})
}

// NewSyncWriter returns a new writer that is safe for concurrent use by
// multiple goroutines. Writes to the returned writer are passed on to w. If
// another write is already in progress, the calling goroutine blocks until
// the writer is available.
//
// If w implements the following interface, so does the returned writer.
//
//    interface {
//        Fd() uintptr
//    }

func NewSyncWriter(w io.Writer) io.Writer {
	return kitlog.NewSyncWriter(w)
}

// add SetLevel
func SetLevel(l Logger, lvl uint32) (Logger, error) {

	// algorand log level : *Panic,Fatal,Error,Warn,Info,Debug*
	var ls string
	if lvl == 2 { //error
		ls = "error"
	} else if lvl == 4 { // Info
		ls = "info"
	} else if lvl == 5 { // debug
		ls = "debug"
	} else {
		ls = "none"
	}

	option, err := AllowLevel(ls)
	if err != nil {
		return nil, err
	}

	nl = NewFilter(l, option)

	return nl, nil
}
