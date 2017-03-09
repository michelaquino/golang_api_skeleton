package logger

// Logger is a interface to log object
type Logger interface {
	// Debug write a debug log level
	Debug(class, method, requestID, user, ip, action, result, message string)

	// Info write a info log level
	Info(class, method, requestID, user, ip, action, result, message string)

	// Warn write a warning log level
	Warn(class, method, requestID, user, ip, action, result, message string)

	// Error write a error log level
	Error(class, method, requestID, user, ip, action, result, message string)
}
