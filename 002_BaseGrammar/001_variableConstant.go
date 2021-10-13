//简单的说，变量就是就是程序在运行时会发生改变的量，常量就是在程序运行中不改变的量，就这么理解吧

package main
import "fmt"

func  main()  {
	//变量声明与赋值
	//声明一个变量 (关键字)var + 变量名 + 数据类型
	var number1 int
	var number2 float32
	//变量赋值 变量名 = 变量的值
	number1 = 23
	number2 = 12.12
	fmt.Println(number1,number2)

	//常量声明与赋值，因为是常量，不可改变，所以在声明的时候就得赋值
	//常量声明 (关键字)const + 常量名 + 数据类型 = 常量值
	const pi1 float32 = 3.1415926	//显式声明
	const pi2 = 3.141956			//隐式声明
	fmt.Println(pi1,pi2)

}

