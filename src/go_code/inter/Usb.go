package inter

// Usb usb接口
type Usb interface {
	Connect()
	Transform() bool
}
