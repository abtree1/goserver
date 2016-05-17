package controllers

import (
	"gs_tmp/models"
	. "gs_tmp/utils"
)

func (client *Client) GetInfo(msg *Msg) {
	client.check_loaded(msg.PlayerId)
	name := msg.Buff.ReadString()
	data := client.PlayerData[name]
	msg.Handler.(chan *models.TableDb) <- data
}
