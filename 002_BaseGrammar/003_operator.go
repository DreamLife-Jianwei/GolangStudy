
//运算符

package main

import "fmt"

func main()  {


	var number1 int 
	number1 = 123
	var number2 = 15
	//算术运算符
	fmt.Println(number1 + number2)
	fmt.Println(number1 - number2)
	fmt.Println(number1 * number2)
	fmt.Println(number1 / number2)
	fmt.Println(number1 % number2)
	//比较运算符
	if number1 > number2 {
		fmt.Printf("%d > %d",number1,number2)
	}

	if number2 < number1 {
		fmt.Printf("%d < %d",number2,number1)
	}

	if number1 == number2 {
		fmt.Printf("%d == %d",number1,number2)
	}

	if number1 != number2 {
		fmt.Printf("%d != %d",number1,number2)
	}

	if number1 >= number2 {
		fmt.Printf("%d >= %d",number1,number2)
	}

	if number2 <= number1 {
		fmt.Printf("%d <= %d",number2,number1)
	}

	//逻辑运算符
	var flag1 bool
	flag1 = true
	var flag2 = false

	if flag1 && flag2 {
		fmt.Printf("flag1 和 flag2 均为真")
	}

	if flag1 || flag2 {
		fmt.Printf("flag1 和 flag2 有一个为真")
	}

	if !flag2 {
		fmt.Printf("flag2 为假")
	}

	//条件判断

	var number3 int
	number3 = 2
	var number4 = 3

	if (number3+number4) == 5 {

		fmt.Printf("%d + %d = 5",number3,number4)
		
	}else{
		fmt.Printf("%d + %d != 5",number3,number4)
	}


	//循环

	for i := 0; i < 100; i++ {
		fmt.Println("cout")
	}


	//格式化输出

	var number5 float32
	number5 = 3.1415926
	fmt.Printf("%2.3f",number5)
	
}