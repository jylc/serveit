package connector

import "net"

type IConnector interface {
	Server(handler IHandler) error
}

type IHandler interface {
	Handle(conn net.Conn)
}
