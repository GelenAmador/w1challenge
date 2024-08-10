package temperature

import "fmt"

func RunTemperatureConverter() {
	var temp float64
	var unit string

	fmt.Println("What's the temperature?")
	fmt.Scanf("%f", &temp)
	fmt.Scanln()
	fmt.Println("To which system? (C/F)")
	fmt.Scanf("%s", &unit)

	if unit == "F" {
		// Convertir de Celsius a Fahrenheit
		converted := temp*9/5 + 32
		fmt.Printf("%.2f Celsius equals %.2f Fahrenheit\n", temp, converted)
	} else if unit == "C" {
		// Convertir de Fahrenheit a Celsius
		converted := (temp - 32) * 5 / 9
		fmt.Printf("%.2f Fahrenheit equals %.2f Celsius\n", temp, converted)
	} else {
		fmt.Println("Unknown unit. Please specify C or F.")
	}
}
