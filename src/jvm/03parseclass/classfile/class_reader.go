package classfile

import "encoding/binary"

// ClassReader 字节码读取
type ClassReader struct {
	data []uint8 // byte
}

func (here *ClassReader) readUint8() uint8 {
	val := here.data[0]
	// 弹出1个字节
	here.data = here.data[1:] // reslice语法,re-切片
	return val
}

func (here *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(here.data)
	// 弹出2个字节
	here.data = here.data[2:] // reslice语法,re-切片
	return val
}

func (here *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(here.data)
	// 弹出4个字节
	here.data = here.data[4:] // reslice语法,re-切片
	return val
}

func (here *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(here.data)
	// 弹出8个字节
	here.data = here.data[8:] // reslice语法,re-切片
	return val
}

// 读取unit16表
func (here *ClassReader) readUint16s() []uint16 {
	n := here.readUint16()
	// 切片
	s := make([]uint16, n)

	// 下标记
	for index := range s {
		s[index] = here.readUint16()
	}
	return s
}

// 读取字节么? 为啥传递的是4个字节呢?
func (here *ClassReader) readBytes(n uint32) []uint8 {
	bytes := here.data[:n]
	here.data = here.data[n:]
	return bytes
}
