package connector

import (
	"fmt"
	"net"
	"serveit/base"
	"strings"
	"sync"
)

type TcpConnector struct {
	app *base.Application
}

func NewTcpConnector(app *base.Application) IConnector {
	return &TcpConnector{app: app}
}

func (t *TcpConnector) Server(handler IHandler) error {
	profile := t.app.GetProfile()
	listener, err := net.Listen("tcp", profile.TCPAddress)
	if err != nil {
		return fmt.Errorf("listen on %v failed, %v", profile.TCPAddress, err)
	}
	var wg sync.WaitGroup
	for {
		conn, err := listener.Accept()
		if err != nil {
			if !strings.Contains(err.Error(), "use of closed network connection") {
				return fmt.Errorf("accept listener failed, %s", err)
			}
			break
		}

		go func() {
			wg.Add(1)
			handler.Handle(conn)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}
