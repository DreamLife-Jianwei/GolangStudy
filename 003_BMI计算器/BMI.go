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

		avearageFat float64 = 0.0	//平均体脂
		peoplecount int = 0

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
	peoplecount ++
	var bmi float64 = weight/(height * height)
	var fatRate float64 = (1.2 * bmi + 0.23 * float64(age) - 5.4 - 10.8 * sexRate) /100
	// fmt.Printf("BMI:%2.3f 体脂率：%2.3f",bmi,fatRate)		//输出BMI 体脂率
	fmt.Println("+**************************************+")
	if sex == "man" {
		 if age >= 18 && age <=39 {
			if fatRate <= 0.1 {
				fmt.Printf("BMI:%2.3f\n 体脂率：%2.3f\n太瘦了\n",bmi,fatRate)
			}else if fatRate >0.1 &&fatRate <= 0.16 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n标准，太棒了\n",bmi,fatRate)
			}else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n偏重，不影响把妹，加油\n",bmi,fatRate)
			}else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Printf("BMI:%2.3f\n 体脂率：%2.3f\n稍微有点胖了，需要控制\n",bmi,fatRate)
			}else {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n危险！危险！危险！要减肥了\n",bmi,fatRate)
			}
		}else if age >= 40 && age <= 59{
			if fatRate <= 0.11 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n太瘦了\n",bmi,fatRate)
			}else if fatRate >0.11 &&fatRate <= 0.17 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n标准，太棒了\n",bmi,fatRate)
			}else if fatRate > 0.17 && fatRate <= 0.22 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n偏重，不影响把妹，加油\n",bmi,fatRate)
			}else if fatRate > 0.22 && fatRate <= 0.27 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n稍微有点胖了，需要控制\n",bmi,fatRate)
			}else {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n危险！危险！危险！要减肥了\n",bmi,fatRate)
			}
		}else{
			if fatRate <= 0.13 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n太瘦了\n",bmi,fatRate)
			}else if fatRate >0.13 &&fatRate <= 0.19 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n标准，太棒了\n",bmi,fatRate)
			}else if fatRate > 0.19 && fatRate <= 0.24 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n偏重，不影响把妹，加油\n",bmi,fatRate)
			}else if fatRate > 0.24 && fatRate <= 0.29 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n稍微有点胖了，需要控制\n",bmi,fatRate)
			}else {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n危险！危险！危险！要减肥了\n",bmi,fatRate)
			}	
		}
	}else{
		if age >= 18 && age <=39 {
			if fatRate <= 0.2 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n太瘦了\n",bmi,fatRate)
			}else if fatRate >0.2 &&fatRate <= 0.27 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n标准，太棒了\n",bmi,fatRate)
			}else if fatRate > 0.27 && fatRate <= 0.34 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n偏重，不影响，加油\n",bmi,fatRate)
			}else if fatRate > 0.34 && fatRate <= 0.39 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n稍微有点胖了，需要控制\n",bmi,fatRate)
			}else {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n危险！危险！危险！要减肥了\n",bmi,fatRate)
			}
		}else if age >= 40 && age <= 59{
			if fatRate <= 0.21 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n太瘦了\n",bmi,fatRate)
			}else if fatRate >0.21 &&fatRate <= 0.28 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n标准，太棒了\n",bmi,fatRate)
			}else if fatRate > 0.28 && fatRate <= 0.35 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n偏重，不影响，加油\n",bmi,fatRate)
			}else if fatRate > 0.35 && fatRate <= 0.40 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n稍微有点胖了，需要控制\n",bmi,fatRate)
			}else {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n危险！危险！危险！要减肥了\n",bmi,fatRate)
			}
		}else{
			if fatRate <= 0.22 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n太瘦了\n",bmi,fatRate)
			}else if fatRate >0.22&&fatRate <= 0.29 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n标准，太棒了\n",bmi,fatRate)
			}else if fatRate > 0.29&& fatRate <= 0.36 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n偏重，不影响，加油\n",bmi,fatRate)
			}else if fatRate > 0.36 && fatRate <= 0.41 {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n稍微有点胖了，需要控制\n",bmi,fatRate)
			}else {
				fmt.Printf("BMI:%2.3f \n体脂率：%2.3f\n危险！危险！危险！要减肥了\n",bmi,fatRate)
			}	
		}
	}

	

fmt.Printf("+**************************************+\n共录入%d人记录，平均体脂为%2.3f \n+**************************************+\n",peoplecount,(avearageFat*float64(peoplecount-1)+fatRate)/float64(peoplecount))


	//循环控制
	fmt.Printf("Y 继续，Q 退出\n")
	var tmep string
	fmt.Scanln(&tmep)
	if tmep == "Y" || tmep == "y" {
		isContinue = true
		
	}else if tmep == "Q" || tmep == "q"  {
		isContinue = false
	}

	}

}