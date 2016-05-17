package controllers

import (
	. "gs_tmp/utils"
)

func (client *Client) HandleMsg(msg *Msg) bool {
	switch msg.Category {
	case PROTOCOL_LOGIN_PARAM:
		client.conn_client(msg)
		client.Login(msg)
	case PROTOCOL_TEST_PARAM:
		client.Test(msg)
	case PROTOCOL_EXIT_PARAM:
		client.LoginOut(msg)
		return true
	case PROTOCOL_WRAP_PARAM:
		client.Wrap(msg)
	case PROXY_GET_INFO:
		client.GetInfo(msg)
	}
	return false
}

func (client *Client) SendClient(buff *Buffer) {
	buff.CompleteBuff()
	client.Client.Write(buff.Data)
}
