package classpath

import (
	"errors"
	"strings"
)

// CompositeEntry 组合模式
type CompositeEntry []Entry

// readClass 切片 不是指针
func (compositeEntry CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	// 使用zip流读取文件,打开包文件
	for _, entry := range compositeEntry {
		// 递归调用本身,非常经典
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found :" + className)
}

// String toString方法
func (compositeEntry CompositeEntry) String() string {
	strs := make([]string, len(compositeEntry))
	for index, entry := range compositeEntry {
		strs[index] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}

// newDirEntry 代替构造函数,另外我不是很喜欢这个名字,但是为了后面不出错,就将就着用吧.
func newCompositeEntry(pathList string) CompositeEntry {
	// 使用切片,切片的直接声明方式
	// var compositeEntry []CompositeEntry = []CompositeEntry{}
	compositeEntry := []Entry{}

	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}
