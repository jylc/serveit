package polemo

import (
	"io"
	"log"
	"net"
	"serveit/base"
	"serveit/net/connector"
)

// Parser polemo协议解析
type Parser struct {
	app  *base.Application
	conn connector.IConnector
}

func NewPolemoParser(app *base.Application) *Parser {
	return &Parser{app: app}
}

func (p *Parser) SetConnector(conn connector.IConnector) {
	p.conn = conn
}

// Start 启动解析器
func (p *Parser) Start() {
	if err := p.conn.Server(p); err != nil {
		log.Printf("[ERROR] start parser failed, %v", err)
	}
}

func (p *Parser) Handle(conn net.Conn) {
	for {
		header, err := io.ReadAll(io.LimitReader(conn, HeaderLength))
		if err != nil {
			log.Printf("[ERROR] load header failed, %v", err)
			return
		}
		if len(header) == 0 {
			log.Printf("[ERROR] header length is %v", len(header))
			return
		}

	}
}

// ParseHeader 解析header部分
func (p *Parser) ParseHeader(header []byte) error {
	typ := int(header[0])
	switch typ {
	case PackageTypeHandshake:
	case PackageTypeAck:
	case PackageTypeHeartBeat:
	case PackageTypeData:
	case PackageTypeDisconnect:
	default:

	}
	return nil
}
