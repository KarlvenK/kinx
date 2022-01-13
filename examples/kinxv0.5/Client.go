package main

import (
	"fmt"
	"github.com/KarlvenK/kinx/knet"
	"io"
	"net"
	"time"
)

//imitate client
func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit")
		return
	}
	for {
		dp := knet.NewDataPack()
		binnaryMsg, err := dp.Pack(knet.NewMsgPackage(0, []byte("kinxv0.5 client test message")))
		if err != nil {
			fmt.Println("pack err", err)
			return
		}
		_, err = conn.Write(binnaryMsg)
		if err != nil {
			fmt.Println("conn write msg err", err)
			return
		}

		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error", err)
			break
		}

		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack msgHead err", err)
			break
		}
		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*knet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read smg data err", err)
				return
			}

			fmt.Println("--> recv Server Msg Id:", msg.GetMsgId(), "data = ", string(msg.GetData()))
		}

		time.Sleep(1 * time.Second)
	}

}
