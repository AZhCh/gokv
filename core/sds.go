package core

import "fmt"

/**
* A simple dynamic string struct in order to store strings
 */

type SDS struct {
	len      uint   //长度
	buf      []byte // 值
	encode   string //编码方式，utf8、utf16、utf32
	capacity uint   //容量
}

func NewSDSFromValue(value []byte) *SDS {
	vlen := len(value)
	sds := &SDS{
		len:      uint(vlen),
		buf:      make([]byte, vlen<<1),
		capacity: uint(vlen),
	}
	copy(sds.buf, value)
	return sds
}

func NewSDSFromCapacity(capacity uint) *SDS {
	return &SDS{
		len:      0,
		buf:      make([]byte, capacity),
		capacity: capacity,
	}
}

func (sds *SDS) ToString() string {
	return string(sds.buf[:sds.len])
}

func (sds *SDS) Len() uint {
	return sds.len
}

func (sds *SDS) Capacity() uint {
	return sds.capacity
}

func (sds *SDS) Print() {
	fmt.Printf("SDS->Len: %d\n", sds.Len())
	fmt.Printf("SDS->Value: %v\n", sds.ToString())
	fmt.Printf("SDS->Capacity: %d\n", sds.Capacity())
}
