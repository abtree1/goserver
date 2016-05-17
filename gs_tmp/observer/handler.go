package observer

import (
	"gs_tmp/controllers"
	. "gs_tmp/utils"
)

func handler(msg *Msg) {
	switch msg.Category {
	case PROTOCOL_LOGIN_PARAM:
		has_player(msg)
	case PROXY_ADD_PLAYER:
		add_player(msg)
	case PROXY_DELETE_PLAYER:
		del_player(msg)
	default:
		send(msg)
	}
}

func add_player(msg *Msg) {
	clients[msg.PlayerId] = msg.Handler.(chan *Msg)
}

func del_player(msg *Msg) {
	delete(clients, msg.PlayerId)
}

func has_player(msg *Msg) {
	c, ok := clients[msg.PlayerId]
	if ok {
		c <- msg
	} else {
		go controllers.RunController(msg, writes)
	}
}

func send(msg *Msg) {
	c, ok := clients[msg.PlayerId]
	if ok == true {
		c <- msg
	} else {
		go controllers.RunController(msg, writes)
	}
}
