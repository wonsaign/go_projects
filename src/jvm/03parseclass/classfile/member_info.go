package classfile

/**
java class 中定义的字段结构
filed_info{
	u2 access_flags;
	u2 name_index;
	u2 descriptor_index;
	u2 attributes_count;
	attribute_info attributes[attributes_count]
}
*/
// MemberInfo 代表了字段的属性,copy 上面
type MemberInfo struct {
	//cp ConstantPool
	accessFlags     uint16
	nameInex        uint16
	descriptorIndex uint16
	//attributes []Attributes
}

func readMembers(reader *ClassReader)
