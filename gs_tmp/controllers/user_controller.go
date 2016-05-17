package controllers

import (
	"fmt"
	"net"

	"gs_tmp/config"
	"gs_tmp/models"
	. "gs_tmp/utils"
)

type Client struct {
	Client     *net.TCPConn
	PlayerData map[string]*models.TableDb
	inited     bool
	Handler    chan *Msg
	Proxy      chan<- *Msg
}

func (client *Client) Login(msg *Msg) {
	client.check_loaded(msg.PlayerId)
	user := client.PlayerData["users"].Data[0]
	id := user["id"].(int)
	name := user["name"].(string)
	pwd := user["pwd"].(string)
	age := user["age"].(int)
	fmt.Println("resecive user: id = ", id, " name=", name, " pwd=", pwd, " age=", age)

	user_conn := client.PlayerData["user_conns"].Data[0]
	id = user_conn["id"].(int)
	phone := user_conn["phone"].(string)
	mobile := user_conn["mobile"].(string)
	email := user_conn["email"].(string)
	qq := user_conn["qq"].(string)
	user_id := user_conn["user_id"].(int)
	fmt.Println("resecive user_conn: id = ", id, " phone=", phone, " mobile=", mobile, " email=", email, " qq=", qq, " user_id=", user_id)

	bak := BuffFactory([]byte{})
	bak.WriteInt32(PROTOCOL_LOGIN_BAK)
	bak.WriteString("你好，客户端!\r\n")
	client.SendClient(bak)
}

func (client *Client) LoginOut(msg *Msg) {
	user := client.PlayerData["users"].Data[0]
	id := user["id"].(int)
	del := &Msg{
		PlayerId: id,
		Category: PROXY_DELETE_PLAYER,
	}
	client.Proxy <- del
}

func (client *Client) Wrap(msg *Msg) {
	id := int(msg.Buff.ReadInt32())
	back := make(chan *models.TableDb, 1)
	buff := BuffNoClose()
	buff.WriteString("users")
	m := &Msg{
		PlayerId: id,
		Category: PROXY_GET_INFO,
		Buff:     buff,
		Handler:  back,
	}
	client.Proxy <- m
	user := (<-back).Data[0]
	id = user["id"].(int)
	name := user["name"].(string)
	pwd := user["pwd"].(string)
	age := user["age"].(int)
	fmt.Println("wrap user: id = ", id, " name=", name, " pwd=", pwd, " age=", age)

	bak := BuffFactory([]byte{})
	bak.WriteInt32(PROTOCOL_LOGIN_BAK)
	bak.WriteString("你好，客户端!\r\n")
	client.SendClient(bak)
}

func (client *Client) Test(msg *Msg) {
	buff := msg.Buff
	str := buff.ReadString()
	b := buff.ReadBool()
	f32 := buff.ReadFloat32()
	fmt.Println("resecive: str= ", str, " b=", b, " f32=", f32)

	row, _ := config.Find("excel_config", 1)
	fmt.Println("find row: id:", row["id"].(int), " name:", row["name"].(string), " rate:", row["rate"].(float32))
	name, _ := config.GetValueString("test", 2, "name")
	fmt.Println("get value: name:", name)

	bak := BuffFactory([]byte{})
	bak.WriteInt32(PROTOCOL_LOGIN_BAK)
	bak.WriteString("你好，客户端!\r\n")
	client.SendClient(bak)
}

func (client *Client) check_loaded(id int) {
	if !client.inited {
		fmt.Println("load data", id)
		client.PlayerData = models.LoadAllPlayerData(id)
		client.inited = true
	}
}
