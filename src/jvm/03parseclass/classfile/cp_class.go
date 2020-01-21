package classfile

/**

CONSTANT_Class_info{
	u1 tag;
	u2 name_index;
}
*/
// ConstantClassInfo 类信息
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (here *ConstantClassInfo) readInfo(reader *ClassReader) {
	here.nameIndex = reader.readUint16()
}

// Name 从常量池中获取名称
func (here *ConstantClassInfo) Name() string {
	return here.cp.getUtf8(here.nameIndex)
}
