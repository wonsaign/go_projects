package classfile

import "fmt"

// JAVAMAGIC java魔数
const JAVAMAGIC = 0xCAFEBABE

//ClassFile golang定义的class文件
type ClassFile struct {
	// magic uint32
	minorVersion uint16
	majorVersion uint16
	//constantPool ConstantPool // 常量池
	accessFlags uint16
	thisClass   uint16
	superClass  uint16
	interfaces  []uint16
	// fields      []*MemberInfo
	// method      []*MemberInfo
	// attributes  []*AttributeInfo
}

// Parse 解析类文件
func Parse(classData []uint8) (cf *ClassFile, err error) {

	// TODO 这里有一个defer方法,暂时不理解干啥,索性先不写
	// 明白了,如果读取错误,是根据异常暂停的

	// var resolve = func() {}
	// // 方式一:匿名函数调用
	// defer resolve()

	// 方式二:匿名函数调用
	defer func() {
		// 等于 catch exception
		if r := recover(); r != nil {
			// 不明白
			// 明白了,这他么的是类型断言,对于接口不知道转什么的时候,这就是类型断言
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	// 结构体的直接初始化的方式
	cr := &ClassReader{classData}

	// 记住,这又是golang的另外一种用法,这个是返回值,下面的return就可以不用写了
	cf = &ClassFile{}

	// fuck , rembrer, 这个方法,必须要定义前面指定的是
	cf.read(cr)

	return
}

// 读取文件
func (here *ClassFile) read(reader *ClassReader) {
	// 读取magicNum
	here.readAndCheckMagic(reader)
	// 校验版本
	here.readConstantPool(reader)
}

// 读取和校验java魔法数
// PDF文件魔法数是%PDF
// ZIP文件魔法数是PK
// Java文件魔法数是0xCAFEBABE
func (here *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic == JAVAMAGIC {
		// 抛出异常,暂时先用golang的
		panic("java.lang.ClassFormatError:magic!")
	}
}

// 版本号校验
func (here *ClassFile) readConstantPool(reader *ClassReader) {
	here.minorVersion = reader.readUint16()
	here.majorVersion = reader.readUint16()
	switch here.majorVersion {
	case 45:
		return
	case 46:
		fallthrough
	case 47, 48, 49, 50, 51, 52:
		if here.minorVersion == 0 {
			return
		}
	}

	panic("java.lang.UnsupportedClassVersionError")
}
