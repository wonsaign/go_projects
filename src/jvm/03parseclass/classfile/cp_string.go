package classfile

/**

CONSTANT_String_info{
	u1 tag;
	u4 bytes;
}
*/
// ConstantStringInfo 结构体
type ConstantStringInfo struct {
	cp          ConstantPool // 本身不存贮字符串数据,只是存贮了常量池的引用
	stringIndex uint16       // 索引
}

// 读取常量池索引
func (here *ConstantStringInfo) readInfo(reader *ClassReader) {
	here.stringIndex = reader.readUint16()
}

// String 按照索引从常量池中查找字符串
func (here *ConstantStringInfo) String() string {
	str := here.cp.getUtf8(here.stringIndex)
	return str
}
