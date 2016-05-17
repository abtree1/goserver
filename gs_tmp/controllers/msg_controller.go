package controllers

import (
	"net"

	. "gs_tmp/utils"
)

func RunController(msg *Msg, proxy chan<- *Msg) {
	hand := make(chan *Msg)
	c := &Client{
		Handler: hand,
		Proxy:   proxy,
	}
	c.conn_observer(msg.PlayerId)
	c.HandleMsg(msg)
	for {
		select {
		case msg := <-c.Handler:
			if c.HandleMsg(msg) {
				return
			}
		default: // do nothing
		}
	}
}

func (client *Client) conn_client(msg *Msg) {
	cli := make(chan *net.TCPConn, 1)
	back := &Msg{
		Handler: cli,
	}
	msg.Handler.(chan *Msg) <- back
	back1 := &Msg{
		Handler: client.Handler,
	}
	msg.Handler.(chan *Msg) <- back1
	client.Client = <-cli
}

func (client *Client) conn_observer(player_id int) {
	add := &Msg{
		PlayerId: player_id,
		Category: PROXY_ADD_PLAYER,
		Handler:  client.Handler,
	}
	client.Proxy <- add
}
