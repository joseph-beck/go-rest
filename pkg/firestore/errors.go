package firestore

import "errors"

var (
	ErrEmptyRead     = errors.New("empty read occurred")
	ErrClosedConn    = errors.New("no connection open")
	ErrMappingStruct = errors.New("error mapping struct")
)
