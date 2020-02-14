package logging

type nopLogger struct{}

// Interface assertions
var _ TmLogger = (*nopLogger)(nil)

// NewNopLogger returns a logger that doesn't do anything.
func NewNopLogger() TmLogger { return &nopLogger{} }

func (nopLogger) Info(string, ...interface{})  {}
func (nopLogger) Debug(string, ...interface{}) {}
func (nopLogger) Error(string, ...interface{}) {}

func (l *nopLogger) With(...interface{}) TmLogger {
	return l
}
