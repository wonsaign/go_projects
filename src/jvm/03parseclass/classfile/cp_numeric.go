package classfile

import "math"

// 导包工具只要是写对包的名字就可以

/**

CONSTANT_Integer_info{
	u1 tag;
	u4 bytes;
}

*/

// ConstantIntegerInfo 常量 int
type ConstantIntegerInfo struct {
	val int32
}

// impl interface
func (here *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	here.val = int32(bytes)
}

/**

CONSTANT_Float_info{
	u1 tag;
	u4 bytes;
}

*/
type ConstantFloatInfo struct {
	val float32
}

func (here *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	here.val = math.Float32frombits(bytes)
}

/**

CONSTANT_Long_info{
	u1 tag;
	u4 high_bytes;
	u4 low_bytes;
}

*/
type ConstantLongInfo struct {
	val int64
}

func (here *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	here.val = int64(bytes)
}

/**

CONSTANT_Double_info{
	u1 tag;
	u4 high_bytes;
	u4 low_bytes;
}

*/
type ConstantDoubleInfo struct {
	val float64
}

func (here *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	here.val = math.Float64frombits(bytes)
}
