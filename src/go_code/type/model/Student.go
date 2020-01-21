package model

// Student 学生结构体
type Student struct {
	name string
	Age  int8
	Sex  bool
}

// SetName 设置名字
func (student *Student) SetName(name string) {
	(*student).name = name
}

// GetName 获取名字
func (student *Student) GetName(name string) string {
	return (*student).name
}
