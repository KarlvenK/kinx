package knet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {
	/*
		tiny server
	*/
	listenner, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("Server listen err:", err)
		return
	}

	go func() {
		for {
			conn, err := listenner.Accept()
			if err != nil {
				fmt.Println("server accept error", err)
			}
			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					//read head
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head error")
						break
					}
					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack err", err)
						return
					}
					if msgHead.GetMsgLen() > 0 {
						//read data according to datalen in head
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetMsgLen())
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server unpack err", err)
							return
						}

						fmt.Println("---> Recv MsgID: ", msg.Id, "datalen = ", msg.DataLen, "data = ", string(msg.Data))
					}
				}
			}(conn)
		}
	}()

	/*
		tiny client
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err", err)
		return
	}
	dp := NewDataPack()

	msg1 := &Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte("kinx!"),
	}
	data1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 err", err)
		return
	}

	msg2 := &Message{
		Id:      2,
		DataLen: 7,
		Data:    []byte("abcdefg"),
	}
	data2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 err", err)
		return
	}

	data1 = append(data1, data2...)
	_, _ = conn.Write(data1)
	select {}

}
