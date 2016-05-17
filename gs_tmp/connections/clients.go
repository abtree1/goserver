package connections

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"gs_tmp/observer"
	. "gs_tmp/utils"
)

func ClientRead(conn *net.TCPConn) {
	head := make([]byte, 2)
	var handler chan<- *Msg
	for {
		io.ReadFull(conn, head)
		size := binary.BigEndian.Uint16(head)
		data := make([]byte, size)
		io.ReadFull(conn, data)
		buff := BuffFactory(data)
		category := buff.ReadInt32()
		if category == PROTOCOL_LOGIN_PARAM {
			id := int(buff.ReadInt32())
			handler = SessionLogin(conn, id, buff)
		} else if category == PROTOCOL_EXIT_PARAM {
			HandleRequest(handler, category, buff)
			close(handler)
			break
		} else {
			HandleRequest(handler, category, buff)
		}
	}
	fmt.Println(Show("客户端退出!"), Show(conn.RemoteAddr().String()))
	conn.Close()
}

func SessionLogin(client *net.TCPConn, id int, buff *Buffer) chan *Msg {
	back := make(chan *Msg, 1)
	msg := &Msg{
		Category: PROTOCOL_LOGIN_PARAM,
		PlayerId: id,
		Buff:     buff,
		Handler:  back,
	}
	observer.Proxy(msg)
	msg1 := <-back
	msg1.Handler.(chan *net.TCPConn) <- client
	return (<-back).Handler.(chan *Msg)
}

func HandleRequest(handler chan<- *Msg, category int32, buff *Buffer) {
	msg := &Msg{
		Category: category,
		Buff:     buff,
	}
	fmt.Println("Receive Type:", category)
	handler <- msg
}
