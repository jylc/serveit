package polemo

const (
	HeaderLength          = 4
	PackageTypeHandshake  = 0x01
	PackageTypeAck        = 0x02
	PackageTypeHeartBeat  = 0x03
	PackageTypeData       = 0x04
	PackageTypeDisconnect = 0x05
)

type Package struct {
	Type   byte
	Length int
	Body   []byte
}

type Message struct {
	Flag      byte
	MessageId []byte
	Route     []byte
}
