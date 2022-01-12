package knet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/KarlvenK/kinx/kiface"
	"github.com/KarlvenK/kinx/utils"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	//DataLen (uint32) 4Byte, ID (uint32) 4Byte
	//Sum = 8 Byte
	return 8
}

func (dp *DataPack) Pack(msg kiface.IMessage) ([]byte, error) {
	//create a bytes buffer
	dataBuff := bytes.NewBuffer([]byte{})

	err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen())
	if err != nil {
		return nil, err
	}
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId())
	if err != nil {
		return nil, err
	}
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetData())
	if err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

func (dp *DataPack) Unpack(binaryData []byte) (kiface.IMessage, error) {
	dataBuff := bytes.NewReader(binaryData)

	msg := &Message{}

	//read DataLen
	err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen)
	if err != nil {
		return nil, err
	}
	//read MsgId
	err = binary.Read(dataBuff, binary.LittleEndian, &msg.Id)
	if err != nil {
		return nil, err
	}
	//read data
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too Large msg data received")
	}

	return msg, nil
}
