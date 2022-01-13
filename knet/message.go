package knet

type Message struct {
	//head
	Id      uint32
	DataLen uint32

	//data
	Data []byte
}

func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		DataLen: uint32(len(data)),
		Id:      id,
		Data:    data,
	}
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

func (m *Message) SetDataLen(length uint32) {
	m.DataLen = length
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}
