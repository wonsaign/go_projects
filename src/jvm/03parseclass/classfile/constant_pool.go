package classfile

/**
常量池是一个表

注意1 表头给出的常量池大小比实际大1
2.有效大常量池索引是1~n-1, 0是无效索引,表示不指向任何常量
3.CONSTANT_Long_info,CONSTANT_Double_info各占两个位置
*/
// ConstantPool pool
type ConstantPool []ConstantInfo

// 读取常量池
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)

	// 注意索引从1开始,不是0
	for index := 1; index < cpCount; index++ {
		// 妈蛋的,golang要有一个好脑子才行啊
		cp[index] = readContantInfo(reader, cp)
		switch cp[index].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			index++ // 占两个位置
		}
	}
	return cp
}

// 按照索引查找常量
func (here ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	cpInfo := here[index]
	if cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

// 从常量池中找字段名称或方法的名字和描述符
func (here ConstantPool) getNameAndType(index uint16) (string, string) {
	// 这是啥语法啊,妈的,这是接口类型转换
	var ok bool
	ntInfo, ok := here.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	if !ok {
		panic("类型转化错误")
	}
	name := here.getUtf8(ntInfo.nameIndex)
	_type := here.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// 从常量池中找类名
func (here ConstantPool) getClassName(index uint16) string {
	var ok bool
	classInfo, ok := here.getConstantInfo(index).(*ConstantClassInfo)
	if !ok {
		panic("类型转化错误")
	}
	return here.getUtf8(classInfo.nameIndex)
}

// 从常量池中查找utf8字符串
func (here ConstantPool) getUtf8(index uint16) string {
	var ok bool
	utf8Info, ok := here.getConstantInfo(index).(*ConstantUtf8Info)
	if !ok {
		panic("类型转化错误")
	}
	return utf8Info.str
}
