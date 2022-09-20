package xnet

import (
	"bytes"
	"encoding/binary"
)

/*
* TLV格式的编解码
 */

type Codec interface {
	GetHeaderLen() uint32
	Encoder(msg Message) ([]byte, error)
	Decoder([]byte) (Message, error)
}

type CodeCImpl struct {
}

func NewCodeCImpl() Codec {
	return &CodeCImpl{}
}

func (c *CodeCImpl) GetHeaderLen() uint32 {
	// dataLen: 4字节
	// id：4字节
	return 8
}

func (c *CodeCImpl) Encoder(msg Message) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.BigEndian, msg.GetMsgLength()); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (c *CodeCImpl) Decoder(bData []byte) (Message, error) {
	buffer := bytes.NewReader(bData)
	msg := &MessageImpl{}

	if err := binary.Read(buffer, binary.BigEndian, &msg.Length); err != nil {
		return nil, err
	}

	if err := binary.Read(buffer, binary.BigEndian, &msg.Id); err != nil {
		return nil, err
	}

	return msg, nil
}
