package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// wildcard 通配符的意思
func newWildcardEntry(path string) CompositeEntry {
	// string 可以直接转切片
	baseDir := path[:len(path)-1]

	compositeEntry := []Entry{}

	// 匿名接口实现
	// 为了初学,所以我写的很全,与java不一样的是,需要写出包名

	// 匿名函数可以直接调用函数中的变量
	// 我这个变量声明也只是一个壳,只是告诉你,goFunc是一个自定义的(type)的类型
	// 类似java8 中的function函数  ,比如 (e1,e2)->return e1.compare(e2)
	// goFunc的赋值

	// 绝对是jdk8中的function, 抽象模版的总结
	var walkFn filepath.WalkFunc = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") ||
			strings.HasSuffix(path, ".JAR") {
			// 主要是这两个变量,这里确实是函数式编程
			// 为什么这两个变量可以存在不报错,因为,这个方法,是在这个包下调用的,并对这个两个包进行赋值,如果在filepath包下肯定不行
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
