package polemo

const (
	PackageTypeReqRsp     = 0x01
	PackageTypeAck        = 0x02
	PackageTypeHeartBeat  = 0x03
	PackageTypeData       = 0x04
	PackageTypeDisconnect = 0x05
)

type Package struct {
	pType   byte
	pLength [3]byte
	pBody   []byte
}

type Message struct {
	mFlag      byte
	mMessageId []byte
	mRoute     []byte
}
