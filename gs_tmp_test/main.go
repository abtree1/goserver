package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	. "gs_tmp/utils"
)

func main() {

	Client()
}

func Client() {
	client, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(Show("服务端连接失败"), Show(err.Error()))
		return
	}
	defer client.Close()

	login(client, 1)

	test_message(client)
	test_wrapmsg(client)

	exit := BuffFactory([]byte{})
	exit.WriteInt32(PROTOCOL_EXIT_PARAM)
	exit.CompleteBuff()
	client.Write(exit.Data)

	client, err = net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(Show("服务端连接失败"), Show(err.Error()))
		return
	}
	login(client, 2)
}

func login(client net.Conn, id int32) {
	login := BuffFactory([]byte{})
	login.WriteInt32(PROTOCOL_LOGIN_PARAM)
	login.WriteInt32(id)
	login.CompleteBuff()
	client.Write(login.Data)

	head := make([]byte, 2)
	io.ReadFull(client, head)
	size := binary.BigEndian.Uint16(head)
	data := make([]byte, size)
	io.ReadFull(client, data)
	buff := BuffFactory(data)
	i32 := buff.ReadInt32()
	str := buff.ReadString()
	fmt.Println("category=", i32, " params=", str)
}

func test_message(client net.Conn) {
	buff := BuffFactory([]byte{})
	buff.WriteInt32(PROTOCOL_TEST_PARAM)
	buff.WriteString("你好，服务器!\r\n")
	buff.WriteBool(true)
	buff.WriteFloat32(1.23)
	buff.CompleteBuff()
	client.Write(buff.Data)
	fmt.Println("send hello to server")

	head := make([]byte, 2)
	io.ReadFull(client, head)
	size := binary.BigEndian.Uint16(head)
	data := make([]byte, size)
	io.ReadFull(client, data)
	buff = BuffFactory(data)
	i32 := buff.ReadInt32()
	str := buff.ReadString()
	fmt.Println("category=", i32, " params=", str)
}

func test_wrapmsg(client net.Conn) {
	buff := BuffFactory([]byte{})
	buff.WriteInt32(PROTOCOL_WRAP_PARAM)
	buff.WriteInt32(2)
	buff.CompleteBuff()
	client.Write(buff.Data)
	fmt.Println("send hello to server")

	head := make([]byte, 2)
	io.ReadFull(client, head)
	size := binary.BigEndian.Uint16(head)
	data := make([]byte, size)
	io.ReadFull(client, data)
	buff = BuffFactory(data)
	i32 := buff.ReadInt32()
	str := buff.ReadString()
	fmt.Println("category=", i32, " params=", str)
}
