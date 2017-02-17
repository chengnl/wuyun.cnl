package thriftclient

import (
	"git.apache.org/thrift.git/lib/go/thrift"
)

type connectionProvider interface {
	getConnection(node *node, timeOut int64) (thrift.TTransport, error)
	returnConnection(ct *ctransport) error
	distoryConnection(ct *ctransport) error
	clearConnection(node *node) error
}
