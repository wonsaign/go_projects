package classpath

import "path/filepath"

import "io/ioutil"

// DirEntry 文件夹实体解析
// golang接口不需要显示的使用implements关键字,问题是没有
type DirEntry struct {
	absDir string
}

// readClass 读取class文件
func (dirEntry *DirEntry) readClass(className string) ([]byte, Entry, error) {
	//  filepath,链接文件使用的
	fileName := filepath.Join((*dirEntry).absDir, className)

	data, err := ioutil.ReadFile(fileName)
	return data, dirEntry, err
}

// String toString方法
func (dirEntry *DirEntry) String() string {

	return dirEntry.absDir
}

// newDirEntry 代替构造函数,另外我不是很喜欢这个名字,但是为了后面不出错,就将就着用吧.
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir: absDir}
}
