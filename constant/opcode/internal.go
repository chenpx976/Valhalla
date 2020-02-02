package opcode

const (
	WorldNew              byte = 0x01
	WorldRequestOk        byte = 0x02
	WorldRequestBad       byte = 0x03
	WorldInfo             byte = 0x03
	ChannelNew            byte = 0x04
	ChannelOk             byte = 0x05
	ChannelBad            byte = 0x06
	ChannelInfo           byte = 0x07
	ChannelConnectionInfo byte = 0x08
)
