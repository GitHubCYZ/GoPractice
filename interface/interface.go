package main

import (
	"fmt"
)

type USB interface {
	Name() string
	//Connect()
	Connecter //接口嵌入
}

type Connecter interface {
	Connect()
}

type PC struct {
	name string
}

func (pc PC) Name() string {
	return pc.name
}

func (pc PC) Connect() {
	fmt.Println("Connect:", pc.Name())
}

// //此方法只适用于一种情况
// func Disconnect(usb USB) {
// 	//OK-patten方法,适用于type情况较小的情况
// 	if pc, ok := usb.(PC); ok {
// 		fmt.Println("Disconnected:", pc.name)
// 	}
// }

//此方法为统一接口，适用于多种情况
func Disconnect(usb interface{}) {
	//OK-patten方法,适用于type情况较小的情况
	// if pc, ok := usb.(PC); ok {
	// 	fmt.Println("Disconnected:", pc.name)
	// }

	//type-switch方法，适用于有多种type的情况
	switch v := usb.(type) {
	case PC:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown device.")
	}
}

func main() {
	// var a USB
	pc := PC{"Andorid"}
	pc.Connect()
	Disconnect(pc)

	//超级接口只能转成子接口
	var con Connecter
	con = Connecter(pc)
	con.Connect()
	//超级转成的子接口对象为拷贝对象，再修改超级接口中的对象时，子接口对象无影响
	pc.name = "IPhone" //修改超级接口中的Name由Android为IPhone
	con.Connect()      //输出仍未Connect: Andorid
}
