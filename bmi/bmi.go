package bmi

import "fmt"

func RunBMICalculator() {
	var weight, height float64

	fmt.Println("How much do you weigh? (don't lie)")
	fmt.Scanf("%f", &weight)
	fmt.Scanln()
	fmt.Println("How tall are you? (barefoot)")
	fmt.Scanf("%f", &height)

	bmi := weight / (height * height)
	fmt.Printf("Right now your BMI is %.2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("You are underweight, add more potato to the broth")
	} else if bmi >= 18.5 && bmi < 25 {
		fmt.Println("You have a normal weight, I have healthy envy of you")
	} else {
		fmt.Println("You are overweight, I know, the pandemic has affected us all")
	}
}
