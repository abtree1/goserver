package observer

import (
	. "gs_tmp/utils"
)

var clients = map[int]chan<- *Msg{}
var writes = make(chan *Msg)

func RunObserver() {
	for {
		write := <-writes
		handler(write)
	}
}

func Proxy(msg *Msg) {
	writes <- msg
}
