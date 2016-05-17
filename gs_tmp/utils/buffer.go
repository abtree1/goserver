package utils

import (
	"math"
)

func BuffFactory(data []byte) *Buffer {
	buff := &Buffer{
		Cur_p: 0,
		Data:  data,
	}
	buff.WriteInt16(0)
	return buff
}

//不调用 CompleteBuff()的Buffer
func BuffNoClose() *Buffer {
	buff := &Buffer{
		Cur_p: 0,
		Data:  []byte{},
	}
	return buff
}

func (buff *Buffer) CompleteBuff() {
	buff.Replace(0, int16(len(buff.Data)-2))
}

func (buff *Buffer) Replace(pos int, data int16) {
	bytes := make([]byte, 2)
	bytes[0] = byte(data >> 8)
	bytes[1] = byte(data)
	copy(buff.Data[pos:pos+2], bytes)
}

func (buff *Buffer) WriteString(str string) {
	size := int16(len(str))
	buff.WriteInt16(size)
	buff.Data = append(buff.Data, str...)
}

func (buff *Buffer) WriteBool(b bool) {
	if b {
		buff.Data = append(buff.Data, byte(1))
	} else {
		buff.Data = append(buff.Data, byte(0))
	}
}

func (buff *Buffer) WriteData(data []byte) {
	size := int16(len(data))
	buff.WriteInt16(size)
	buff.Data = append(buff.Data, data...)
}

func (buff *Buffer) WriteInt8(i int8) {
	buff.Data = append(buff.Data, byte(i))
}

func (buff *Buffer) WriteInt16(i int16) {
	bytes := make([]byte, 2)
	bytes[0] = byte(i >> 8)
	bytes[1] = byte(i)
	buff.Data = append(buff.Data, bytes...)
}

func (buff *Buffer) WriteInt32(i int32) {
	bytes := make([]byte, 4)
	for j := range bytes {
		bytes[j] = byte(i >> uint((3-j)*8))
	}
	buff.Data = append(buff.Data, bytes...)
}

func (buff *Buffer) WriteInt64(i int64) {
	bytes := make([]byte, 8)
	for j := range bytes {
		bytes[j] = byte(i >> uint((7-j)*8))
	}
	buff.Data = append(buff.Data, bytes...)
}

func (buff *Buffer) WriteRune(c rune) {
	buff.Data = append(buff.Data, byte(c))
}

func (buff *Buffer) WriteUint16(i uint16) {
	buff.WriteInt16(int16(i))
}

func (buff *Buffer) WriteUint32(i uint32) {
	buff.WriteInt32(int32(i))
}

func (buff *Buffer) WriteUint64(i uint64) {
	buff.WriteInt64(int64(i))
}

func (buff *Buffer) WriteFloat32(f float32) {
	bytes := math.Float32bits(f)
	buff.WriteUint32(bytes)
}

func (buff *Buffer) WriteFloat64(f float64) {
	bytes := math.Float64bits(f)
	buff.WriteUint64(bytes)
}

func (buff *Buffer) ReadString() string {
	size := buff.ReadInt16()
	str := string(buff.Data[buff.Cur_p : buff.Cur_p+int(size)])
	buff.Cur_p += int(size)
	return str
}

func (buff *Buffer) ReadBool() bool {
	buff.Cur_p += 1
	if buff.Data[buff.Cur_p-1] == byte(1) {
		return true
	} else {
		return false
	}
}

func (buff *Buffer) ReadData(size int16) []byte {
	buff.Cur_p += int(size)
	return buff.Data[buff.Cur_p-int(size) : buff.Cur_p]
}

func (buff *Buffer) ReadInt8() int8 {
	buff.Cur_p += 1
	return int8(buff.Data[buff.Cur_p-1])
}

func (buff *Buffer) ReadInt16() int16 {
	ret := buff.ReadUint16()
	return int16(ret)
}

func (buff *Buffer) ReadInt32() int32 {
	ret := buff.ReadUint32()
	return int32(ret)
}

func (buff *Buffer) ReadInt64() int64 {
	ret := buff.ReadUint64()
	return int64(ret)
}

func (buff *Buffer) ReadRune() rune {
	buff.Cur_p += 1
	return rune(buff.Data[buff.Cur_p-1])
}

func (buff *Buffer) ReadUint16() uint16 {
	bytes := buff.Data[buff.Cur_p : buff.Cur_p+2]
	ret := uint16(bytes[0])<<8 | uint16(bytes[1])
	buff.Cur_p += 2
	return uint16(ret)
}

func (buff *Buffer) ReadUint32() uint32 {
	bytes := buff.Data[buff.Cur_p : buff.Cur_p+4]
	ret := uint32(0)
	for i, v := range bytes {
		ret |= uint32(v) << uint((3-i)*8)
	}
	buff.Cur_p += 4
	return uint32(ret)
}

func (buff *Buffer) ReadUint64() uint64 {
	bytes := buff.Data[buff.Cur_p : buff.Cur_p+8]
	ret := uint64(0)
	for i, v := range bytes {
		ret |= uint64(v) << uint((7-i)*8)
	}
	buff.Cur_p += 8
	return uint64(ret)
}

func (buff *Buffer) ReadFloat32() float32 {
	ret := buff.ReadUint32()
	return math.Float32frombits(ret)
}

func (buff *Buffer) ReadFloat64() float64 {
	ret := buff.ReadUint64()
	return math.Float64frombits(ret)
}
