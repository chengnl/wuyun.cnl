package thriftclient

import (
	"errors"
)

const (
	NO_NODE_SERVICE   = 1
	NO_AVAILABLE_NODE = 2
)

type NodeException interface {
	CException
	TypeID() int
	Err() error
}

type cNodeException struct {
	typeID int
	err    error
}

func (c *cNodeException) TypeID() int {
	return c.typeID
}
func (c *cNodeException) Err() error {
	return c.err
}
func (c *cNodeException) Error() string {
	return c.err.Error()
}

func NewNodeException(t int, e string) NodeException {
	return &cNodeException{typeID: t, err: errors.New(e)}
}
