package main

import "fmt"
import "reflect"

//eg:1
type User struct { //反射结构
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello world!")
}

func (u User) TestMethod() {
	fmt.Println("Hello Method!")
}

//eg:2
/******************************************************
 go语言里面struct里面变量如果大写则是public,如果是小写则是
 private的，private的时候通过反射不能获取其值。
******************************************************/
type Manager struct { //包含匿名字段的结构
	//<私有字段>
	// title string
	// userI User
	Title string
	//<1>非匿名字段
	UserI User
	//<2>匿名字段
	//User
}

func main() {
	eg := 2
	switch eg {
	case 1:
		u := User{1, "OK", 12}
		//Info(&u)
		Info(u)
	case 2:
		//<1>非匿名字段
		m := Manager{Title: "1314", UserI: User{1, "OK", 12}}
		//<2>匿名字段
		//m := Manager{Title: "1314", User: User{1, "OK", 12}}
		Info(m)
	}
}

func Info(o interface{}) {
	ty := reflect.TypeOf(o)                   //获取接口参数类型
	fmt.Println("Interface Type:", ty.Name()) //打印参数类型

	//判断传入参数的类型是否为结构体
	if kind := ty.Kind(); kind != reflect.Struct {
		fmt.Printf("%s is not a struct!\n", ty.Name())
		return
	}

	fmt.Println("Fields:")
	value := reflect.ValueOf(o) //获取接口参数值

	for i := 0; i < ty.NumField(); i++ {
		field := ty.Field(i)                 //获取字段信息
		fType := field.Type.Kind()           //获取字段的数据类型
		fValue := value.Field(i).Interface() //获取字段值
		if reflect.Struct == fType {         //判断当前地段数据类型是否为结构体类型
			Info(fValue) //递归打印嵌套结构数据信息
		} else {
			fmt.Printf("%8v:%6s = %v\n", field.Type, field.Name, fValue)
		}
	}

	//获取方法信息
	for i := 0; i < ty.NumMethod(); i++ {
		method := ty.Method(i)
		fmt.Printf("%6s:%v\n", method.Name, method.Type)
	}
}
