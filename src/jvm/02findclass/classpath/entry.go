package classpath

import "os"

import "strings"

// 路径分隔符
const pathListSeparator = string(os.PathListSeparator)

// Entry 实体接口
type Entry interface {
	// 加载和寻找class文件
	readClass(className string) ([]byte, Entry, error)
	// 类似toString方法
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
