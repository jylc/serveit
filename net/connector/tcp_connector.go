package connector

import (
	"fmt"
	"log"
	"net"
	"os"
	"serveit/base"
	"strings"
	"sync"
)

type TcpConnector struct {
	app   *base.Application
	conns sync.Map
}

func NewTcpConnector(app *base.Application) IConnector {
	return &TcpConnector{app: app}
}

func (t *TcpConnector) Server() error {
	profile := t.app.GetProfile()
	listener, err := net.Listen("tcp", profile.TCPAddress)
	if err != nil {
		log.Printf("[ERROR] listen on %v failed", profile.TCPAddress)
		os.Exit(-1)
	}
	var wg sync.WaitGroup
	for {
		conn, err := listener.Accept()
		if err != nil {
			if !strings.Contains(err.Error(), "use of closed network connection") {
				return fmt.Errorf("listener.Accept() error - %s", err)
			}
			break
		}

		go func() {
			wg.Add(1)
			t.handler(conn)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}

func (t *TcpConnector) handler(conn net.Conn) {
	log.Printf("[INFO] new connection, %v", conn.RemoteAddr())
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	if err != nil {
		log.Printf("[ERROR] read connection failed, %v", err)
		return
	}
}
