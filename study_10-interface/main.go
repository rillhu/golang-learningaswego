package main

import (
	"fmt"
	)


/*---Interface example 1---*/
//接口， 方法的集合，下面定义了一组方法的集合：手机都可以打电话，发短信
type Phone interface {
	call()
	sms()
}


//nokia 电话的结构体定义
type NokiaPhone struct {

}
//nokia电话的方法，即接口的代码实现
func (NokiaPhone1 NokiaPhone) call() {
	fmt.Println("I am a nokia phone call")
}

func (NokiaPhone1 NokiaPhone) sms()  {
	fmt.Println("I am a nokia phone sms")

}

type Iphone struct {

}


func(Iphone1 Iphone) call(){
	fmt.Println("I am a iPhone call")
}
/*
func (Iphone1 Iphone) sms(){
	fmt.Println("I am a iPhone sms")
}
*/

/*---------------------------------------*/

/*---Interface example 2---*/
type namer interface {
	area() int
}

type rect struct {
	width,height int
}

type square struct {
	side int
}

func (r rect) area() int {
	return r.height * r.width
}

func (s square) area() int {
	return s.side * s.side
}
/*---------------------------------------*/

/*Inner type alias for adding method*/
type INTEGER int

func (i INTEGER) print() {
	fmt.Println(i)
}
/*---------------------------------------*/

func main() {

	/*---Give inner type a alias,then we can add new method for it---*/
	var intVar INTEGER
	intVar = 100
	intVar.print()

	/*---Interface example 1---*/
	var phone1 Phone
	phone1 = new(NokiaPhone)
	phone1.call()
	phone1.sms()

	phone2 := new(Iphone) //Iphone 可以不完全实现接口中声明的方法
	phone2.call()


	//phone2.sms()
	/*---Interface example 2---*/
	var a = rect{4,3}
	var b  = square{6}

	fmt.Println(a.area())
	fmt.Println(b.area())


	/*interface query*/
	//测试某个值是否实现了某个接口, such as test if phone1 implements the interface
	if sPhone, ok := phone1.(Phone); ok{
		fmt.Println("ok:", ok, "ret:",sPhone)
		sPhone.call()
		sPhone.sms()
	}
	//基于口语类型查询，暂时还有点问题
	var phoneTmp Phone
	//fmt.Println(phoneTmp.(type))
	//switch Phone.(type) {
	switch phoneTmp.(type) {
		case interface{}:
			fmt.Println("interface")

		default:
			fmt.Println("default")


	}


}
