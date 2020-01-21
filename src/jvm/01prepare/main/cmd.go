package main

import (
	"flag"
	"fmt"
	"os"
)

// Cmd 用来调用结构体的
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string // 这是切片,不是数组啊,菜鸟切记切记
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
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
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
