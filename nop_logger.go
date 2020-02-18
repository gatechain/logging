package logging

import (
	"io"

	"github.com/gatechain/logging/telemetryspec"
	"github.com/sirupsen/logrus"
)

type nopLogger struct{}

// Interface assertions
var _ Logger = (*nopLogger)(nil)

// NewNopLogger returns a logger that doesn't do anything.
func NewNopLogger() Logger { return &nopLogger{} }

//func (nopLogger) Info(string, ...interface{})  {}
//func (nopLogger) Debug(string, ...interface{}) {}
//func (nopLogger) Error(string, ...interface{}) {}
//
//func (l *nopLogger) With(...interface{}) Logger {
//	return l
//}

// Debug logs a message at level Debug.
func (nopLogger) Debug(...interface{})          {}
func (nopLogger) Debugln(...interface{})        {}
func (nopLogger) Debugf(string, ...interface{}) {}

// Info logs a message at level Info.
func (nopLogger) Info(...interface{})          {}
func (nopLogger) Infoln(...interface{})        {}
func (nopLogger) Infof(string, ...interface{}) {}

// Warn logs a message at level Warn.
func (nopLogger) Warn(...interface{})          {}
func (nopLogger) Warnln(...interface{})        {}
func (nopLogger) Warnf(string, ...interface{}) {}

// Error logs a message at level Error.
func (nopLogger) Error(...interface{})          {}
func (nopLogger) Errorln(...interface{})        {}
func (nopLogger) Errorf(string, ...interface{}) {}

// Fatal logs a message at level Fatal.
func (nopLogger) Fatal(...interface{})          {}
func (nopLogger) Fatalln(...interface{})        {}
func (nopLogger) Fatalf(string, ...interface{}) {}

// Panic logs a message at level Panic.
func (nopLogger) Panic(...interface{})          {}
func (nopLogger) Panicln(...interface{})        {}
func (nopLogger) Panicf(string, ...interface{}) {}

// Add one key-value to log
func (l *nopLogger) With(key string, value interface{}) Logger {
	return l
}

// WithFields logs a message with specific fields
func (l *nopLogger) WithFields(Fields) Logger {
	return l
}

// Set the logging version (Info by default)
func (nopLogger) SetLevel(Level) {}

// Sets the output target
func (nopLogger) SetOutput(io.Writer) {}

// Sets the logger to JSON Format
func (nopLogger) SetJSONFormatter() {}
func (nopLogger) IsLevelEnabled(level Level) bool {
	return false
}

// source adds file, line and function fields to the event
func (l *nopLogger) source() *logrus.Entry {
	return nil
}

// Adds a hook to the logger
func (nopLogger) AddHook(hook logrus.Hook) {}
func (nopLogger) EnableTelemetry(cfg TelemetryConfig) error {
	return nil
}
func (nopLogger) UpdateTelemetryURI(uri string) bool {
	return false
}
func (nopLogger) GetTelemetryEnabled() bool {
	return false
}
func (nopLogger) Metrics(category telemetryspec.Category, metrics telemetryspec.MetricDetails, details interface{}) {
}
func (nopLogger) Event(category telemetryspec.Category, identifier telemetryspec.Event) {}
func (nopLogger) EventWithDetails(category telemetryspec.Category, identifier telemetryspec.Event, details interface{}) {
}
func (nopLogger) StartOperation(category telemetryspec.Category, identifier telemetryspec.Operation) TelemetryOperation {
	return TelemetryOperation{}
}
func (nopLogger) GetTelemetrySession() string {
	return ""
}
func (nopLogger) GetTelemetryHostName() string {
	return ""
}
func (nopLogger) GetInstanceName() string {
	return ""
}
func (nopLogger) GetTelemetryURI() string {
	return ""
}
func (nopLogger) CloseTelemetry() {}
