package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ZipEntry  zip或jar文件的类路径
type ZipEntry struct {
	absDir string
}

// readClass 从zip获取获取文件
func (zipEntry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 使用zip流读取文件,打开包文件
	readCloser, err := zip.OpenReader(zipEntry.absDir)
	// finally 方法,最后在最后关闭IO流
	if err != nil {
		return nil, nil, err
	}
	defer readCloser.Close()

	// readCloser.File和下面的写法是一样的,这是golang的继承,但是我这么写,是为了初学
	// zip包中的文件,遍历循环,找到其中的files
	for _, f := range readCloser.Reader.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, zipEntry, nil
		}

	}
	return nil, nil, errors.New("class not found :" + className)
}

// String toString方法
func (zipEntry *ZipEntry) String() string {
	return zipEntry.absDir
}

// newDirEntry 代替构造函数,另外我不是很喜欢这个名字,但是为了后面不出错,就将就着用吧.
func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absDir: absDir}
}
