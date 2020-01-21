package classfile

/*
CONSTANT_NameAndType_info{
	u1 tag;
	u2 name_index; // 指向常量池的索引
	u2 descriptor_index; // 指向常量池的索引
}
*/

/*
1.类型描述符
	yi基本类型 byte short char int long float double,分别对应的是
		B S C I J F D (注意long是J,不是L)

	er引用类型 L+类的完全限定名+分号 如: Ljava.lang.String 代表了 String

	san 数组类型 [+数组元素描述符号
2.字段描述符就是字段类型的描述符 如-> [I  代表了  int[]

3.方法描述符是(分号分割的参数类型描述符)+返回值类型描述符,其中void是V,如->  ()Ljava.lang.String; 代表了 String toString()
*/

// ConstantNameAndTypeInfo 名称和类型结构体类型
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

// 接口实现
func (here *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	here.nameIndex = reader.readUint16()
	here.descriptorIndex = reader.readUint16()
}
