package main

import "fmt"

// Step 1: Định nghĩa giao diện State
type State interface {
	Handle()
}

// Step 2: Triển khai các Concrete States
type OffState struct{}
type OnState struct{}

func (off *OffState) Handle() {
	fmt.Println("Turn on the device.")
}

func (on *OnState) Handle() {
	fmt.Println("Turn off the device.")
}

// Step 3: Định nghĩa Context
type Device struct {
	state State
}

func (d *Device) SetState(state State) {
	d.state = state
}

func (d *Device) PressButton() {
	d.state.Handle()
}

func main() {
	// Tạo đối tượng Context
	device := &Device{}

	// Sử dụng Concrete States khác nhau
	offState := &OffState{}
	onState := &OnState{}

	device.SetState(offState)
	device.PressButton()

	device.SetState(onState)
	device.PressButton()
}
