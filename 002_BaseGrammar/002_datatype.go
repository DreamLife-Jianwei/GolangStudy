//基本数据类型
	//整数 
		// 有符号整数：int int8 int16 int32 int64
		// 无符号整数：uint uint8 uint16 uing32 uint64
	//浮点数 float32 float64 (注意 Go里面没有float)
	//字符串：string
	//布尔值：bool
	//内置虚数类型：complex64 complex128
	//固定长度的组数：array

package main

import "fmt"

func  main()  {

	var number1 int
	number1 = 46
	var number2 = 32			//隐式声明
	fmt.Println(number1,number2)

	var flag1 bool
	flag1 = true
	var flag2 = false
	fmt.Println(flag1,flag2)

	var floatnumber1 float32
	floatnumber1 = 2.234
	var floatnumber2 = 45.54
	fmt.Println(floatnumber1,floatnumber2)

	var string1 string
	string1 = "dreamlife"
	var string2 = "DREAMLIFE"
	fmt.Println(string1,string2)

}