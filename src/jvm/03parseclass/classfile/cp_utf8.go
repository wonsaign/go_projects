package classfile

/**

CONSTANT_Utf8_info{
	u1 tag;
	u4 bytes;
}

*/
// ConstantUtf8Info utf8结构体
type ConstantUtf8Info struct {
	str string
}

func (here *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	here.str = string(bytes)
}
