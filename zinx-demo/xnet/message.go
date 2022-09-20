package xnet

type Message interface {
	GetMsgId() uint32
	GetMsgLength() uint32
	GetData() []byte

	SetMsgId(uint32)
	SetMsgLength(uint32)
	SetData([]byte)
}

type MessageImpl struct {
	Id     uint32
	Length uint32
	Data   []byte
}

func (m *MessageImpl) GetMsgId() uint32 {
	return m.Id
}

func (m *MessageImpl) GetMsgLength() uint32 {
	return m.Length
}

func (m *MessageImpl) GetData() []byte {
	return m.Data
}

func (m *MessageImpl) SetMsgId(u uint32) {
	m.Id = u
}

func (m *MessageImpl) SetMsgLength(u uint32) {
	m.Length = u
}

func (m *MessageImpl) SetData(bytes []byte) {
	m.Data = bytes
}
