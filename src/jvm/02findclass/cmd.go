package main

import (
	"flag"
	"fmt"
	"jvm/02findclass/classpath"
	"os"
	"strings"
)

// Cmd 用来调用结构体的
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string // 这是切片,不是数组啊,菜鸟切记切记
	XjarOption  string
}

// parseCmd 解析cmd结构体
func parseCmd() *Cmd {
	var cmd *Cmd = &Cmd{}
	flag.Usage = printUsage

	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message") // .的运算符 要高于&/*
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjarOption, "Xjre", "", "path to jar")
	flag.Parse()

	var args []string = flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:] // 取第二个到最后一个到省略写法,等同于下面
		//cmd.args = args[1:len(args)]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

func startJVM(cmd *Cmd) {
	// 解析xjreOption和cpOption
	cp := classpath.Parse(cmd.XjarOption, cmd.cpOption)

	// fmt.Printf("classpath:%s class:%s args:%v\n",
	// 	cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	//className := cmd.class
	// fmt.Printf("this is className is %v", className)

	data, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s \n", cmd.class)
		return
	}
	fmt.Printf("class data: %v\n", data)
}

func main() {
	var cmd = parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
