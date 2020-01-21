package impl

import "fmt"

// UsbImpl 结构体
type UsbImpl struct {
}

// Connect 链接
func (u UsbImpl) Connect() {
	fmt.Print("usb链接中...")
}

// Transform 传输
func (u UsbImpl) Transform() bool {
	fmt.Print("数据传输中...")
	return true
}
