package utils

type Buffer struct {
	Cur_p int
	Data  []byte
}

type Msg struct {
	PlayerId int
	Category int32
	Buff     *Buffer
	Handler  interface{}
}
