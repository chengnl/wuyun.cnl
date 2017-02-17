package thriftclient

import (
	"git.apache.org/thrift.git/lib/go/thrift"
)

type ctransport struct {
	n         *node
	transport thrift.TTransport
}

func NewCTransport(n *node, t thrift.TTransport) *ctransport {
	return &ctransport{n: n, transport: t}
}
