package ifaces

import "net"

type ISession interface {
	GetId() int64
	GetConn() net.Conn
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
}
