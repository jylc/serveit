package polemo

import (
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
