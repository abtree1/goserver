package connections

import (
	"fmt"
	"net"

	. "gs_tmp/utils"
)

func Server(exit chan<- bool) {
	ip := net.ParseIP("127.0.0.1")
	addr := net.TCPAddr{ip, 8888, "0:0:0:0:0:0:0:1"}
	listen, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		fmt.Println(Show("初始化失败"), Show(err.Error()))
		exit <- true
		return
	}
	fmt.Println(Show("正在监听..."))
	for {
		client, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println(Show("客户端连接"), Show(client.RemoteAddr().String()))

		go ClientRead(client)
	}
}
