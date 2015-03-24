// Basic types.
package whitetree

import (
	"errors"
)

type ParserId string
type Parser func(TaskContext) (HandlerId, error)
type ParserPackage map[ParserId]Parser

type HandlerId string
type Handler func(TaskContext) error
type HandlerPackage map[HandlerId]Handler

type TaskContext struct {
	Data []byte
}

var (
	ErrUnableToParse   = errors.New("unable to parse")
	ErrHandlerNotFound = errors.New("handler not found")
)
