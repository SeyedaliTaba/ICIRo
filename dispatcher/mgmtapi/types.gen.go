// Package mgmtapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package mgmtapi

// Defines values for LogLevelLevel.
const (
	LogLevelLevelDebug LogLevelLevel = "debug"

	LogLevelLevelError LogLevelLevel = "error"

	LogLevelLevelInfo LogLevelLevel = "info"
)

// LogLevel defines model for LogLevel.
type LogLevel struct {
	// Logging level
	Level LogLevelLevel `json:"level"`
}

// Logging level
type LogLevelLevel string

// StandardError defines model for StandardError.
type StandardError struct {
	// Error message
	Error string `json:"error"`
}

// BadRequest defines model for BadRequest.
type BadRequest StandardError

// SetLogLevelJSONBody defines parameters for SetLogLevel.
type SetLogLevelJSONBody LogLevel

// SetLogLevelJSONRequestBody defines body for SetLogLevel for application/json ContentType.
type SetLogLevelJSONRequestBody SetLogLevelJSONBody
