package classfile

// ConstantMemberrefInfo 成员
type ConstantMemberrefInfo struct {
	cp               ConstantPool
	calssIndex       uint16
	nameAndTypeIndex uint16
}

// 接口实现
func (here *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	here.calssIndex = reader.readUint16()
	here.nameAndTypeIndex = reader.readUint16()
}

// ClassName 获取字符串
func (here *ConstantMemberrefInfo) ClassName() string {
	return here.cp.getUtf8(here.calssIndex)
}

// NameAndDescriptor 获取字符串
func (here *ConstantMemberrefInfo) NameAndDescriptor() string {
	return here.cp.getUtf8(here.nameAndTypeIndex)
}

// ConstantFieldrefInfo 字段符号引用
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo //继承
}

// ConstantMethodrefInfo 普通(非接口)方法符号应用
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo //继承
}

// ConstantInterfaceMethodrefInfo 接口方法引用
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo //继承
}
