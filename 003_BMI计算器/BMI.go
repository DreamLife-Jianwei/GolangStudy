/*
BMI = 体重（公斤） / (身高*身高)(米)
体制率： （1.2 * BMI + 0.23*年龄 - 5.4 -10.8*性别(男1女0)） / 100

*/

package main

import "fmt"

func  main()  {
	fmt.Println("BMI is Running…")

	var(
		weight float64	//体重(公斤)
		height float64		//身高(米)
		age int				//年龄
		sex string 
		sexRate float64
	)

	for isContinue := true; isContinue;  {
		fmt.Printf("性别：")
		fmt.Scanln(&sex)
		fmt.Printf("体重（公斤）：")
		fmt.Scanln(&weight)
		fmt.Printf("身高（米）：")
		fmt.Scanln(&height)
		fmt.Printf("年龄：")
		fmt.Scanln(&age)

	if sex == "man" {
		sexRate = 1
	}else{
		sexRate = 0
	}

	var bmi float64 = weight/(height * height)
	var fatRate float64 = (1.2 * bmi + 0.23 * float64(age) - 5.4 - 10.8 * sexRate) /100
	// fmt.Printf("BMI:%2.3f 体脂率：%2.3f",bmi,fatRate)		//输出BMI 体脂率

	if sex == "man" {
		if age < 18 {
			if fatRate <= 0.1 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,太瘦了",bmi,fatRate)
			}else if fatRate >0.1 &&fatRate <= 0.16 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,标准，太棒了",bmi,fatRate)
			}else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,偏重，不影响把妹，加油",bmi,fatRate)
			}else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,稍微有点胖了，需要控制",bmi,fatRate)
			}else {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,危险！危险！危险！要减肥了",bmi,fatRate)
			}
		}else if age >= 18 && age <=39 {
			if fatRate <= 0.1 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,太瘦了",bmi,fatRate)
			}else if fatRate >0.1 &&fatRate <= 0.16 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,标准，太棒了",bmi,fatRate)
			}else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,偏重，不影响把妹，加油",bmi,fatRate)
			}else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,稍微有点胖了，需要控制",bmi,fatRate)
			}else {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,危险！危险！危险！要减肥了",bmi,fatRate)
			}
		}else if age >= 40 && age <= 59{
			if fatRate <= 0.1 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,太瘦了",bmi,fatRate)
			}else if fatRate >0.1 &&fatRate <= 0.16 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,标准，太棒了",bmi,fatRate)
			}else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,偏重，不影响把妹，加油",bmi,fatRate)
			}else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,稍微有点胖了，需要控制",bmi,fatRate)
			}else {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,危险！危险！危险！要减肥了",bmi,fatRate)
			}
		}else{
			if fatRate <= 0.1 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,太瘦了",bmi,fatRate)
			}else if fatRate >0.1 &&fatRate <= 0.16 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,标准，太棒了",bmi,fatRate)
			}else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,偏重，不影响把妹，加油",bmi,fatRate)
			}else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,稍微有点胖了，需要控制",bmi,fatRate)
			}else {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,危险！危险！危险！要减肥了",bmi,fatRate)
			}	
		}
	}else{
		if age < 18 {
			if fatRate <= 0.1 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,太瘦了",bmi,fatRate)
			}else if fatRate >0.1 &&fatRate <= 0.16 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,标准，太棒了",bmi,fatRate)
			}else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,偏重，不影响把妹，加油",bmi,fatRate)
			}else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,稍微有点胖了，需要控制",bmi,fatRate)
			}else {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,危险！危险！危险！要减肥了",bmi,fatRate)
			}
		}else if age >= 18 && age <=39 {
			if fatRate <= 0.1 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,太瘦了",bmi,fatRate)
			}else if fatRate >0.1 &&fatRate <= 0.16 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,标准，太棒了",bmi,fatRate)
			}else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,偏重，不影响把妹，加油",bmi,fatRate)
			}else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,稍微有点胖了，需要控制",bmi,fatRate)
			}else {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,危险！危险！危险！要减肥了",bmi,fatRate)
			}
		}else if age >= 40 && age <= 59{
			if fatRate <= 0.1 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,太瘦了",bmi,fatRate)
			}else if fatRate >0.1 &&fatRate <= 0.16 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,标准，太棒了",bmi,fatRate)
			}else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,偏重，不影响把妹，加油",bmi,fatRate)
			}else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,稍微有点胖了，需要控制",bmi,fatRate)
			}else {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,危险！危险！危险！要减肥了",bmi,fatRate)
			}
		}else{
			if fatRate <= 0.1 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,太瘦了",bmi,fatRate)
			}else if fatRate >0.1 &&fatRate <= 0.16 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,标准，太棒了",bmi,fatRate)
			}else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,偏重，不影响把妹，加油",bmi,fatRate)
			}else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,稍微有点胖了，需要控制",bmi,fatRate)
			}else {
				fmt.Printf("BMI:%2.3f 体脂率：%2.3f,危险！危险！危险！要减肥了",bmi,fatRate)
			}	
		}
	}
	fmt.Printf("Y 继续，Q 退出")
	var tmep string
	fmt.Scanln(&tmep)
	if tmep == "Y" || tmep == "y" {
		isContinue = true
		
	}else if tmep == "Q" || tmep == "q"  {
		isContinue = false
	}

	}

}