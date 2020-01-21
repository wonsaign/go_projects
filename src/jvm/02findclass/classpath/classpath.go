package classpath

import "os"

import "path/filepath"

//import

// Classpath struct
type Classpath struct {
	// boot
	bootClasspath Entry
	// ext
	extClasspath Entry
	// user
	userClasspath Entry
}

// Parse 解析
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	// 解析boot and ext classpath
	cp.parseBootAndExtClasspath(jreOption)
	// 语法糖
	cp.parseUserClasspath(cpOption)
	return cp
}

func (here *Classpath) parseBootAndExtClasspath(jreOption string) {
	// 为此 我特意下了jre1.8  就是为了能找到jre
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	here.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	here.extClasspath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOptio string) string {
	if jreOptio == "" && exists(jreOptio) {
		return jreOptio
	}
	if exists("./jre") {
		return "./jre"
	}

	// 从环境变量里获取JavaHome,麻辣比,我mac就是获取不到,无奈自己手设置一下
	os.Setenv("JAVA_HOME_8", "/Library/Java/JavaVirtualMachines/jre1.8.0_241.jre/Contents/Home")
	javahome := os.Getenv("JAVA_HOME_8")
	if javahome != "" {
		//return filepath.Join(javahome, "jre")
		return javahome
	}
	// 类似java中 threw new Exception() ,目的是程序因为异常停下
	panic("Can not find jre folder")
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// ReadClass 读取classpath
func (here *Classpath) ReadClass(classname string) ([]byte, Entry, error) {
	className := classname + ".class"

	// i need to set jre in jdk13
	// i copy the code , copy wrong
	if data, entry, err := here.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := here.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return here.userClasspath.readClass(className)

}

// String 获取路径
func (here *Classpath) String() string {
	return here.userClasspath.String()
}

func (here *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	here.userClasspath = newEntry(cpOption)
}
