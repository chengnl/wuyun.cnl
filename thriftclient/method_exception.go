package thriftclient

import (
	"errors"
)

type MethodException interface {
	CException
	TypeID() int
	Err() error
}

const (
	NO_SUCH_METHOD = 1
	NO_MACH_PARAMS = 2
	NO_MACH_RESULT = 3
)

type cMethodException struct {
	typeID int
	err    error
}

func (m *cMethodException) TypeID() int {
	return m.typeID
}
func (m *cMethodException) Error() string {
	return m.err.Error()
}
func (m *cMethodException) Err() error {
	return m.err
}
func NewMethodException(t int, e string) MethodException {
	return &cMethodException{typeID: t, err: errors.New(e)}
}
