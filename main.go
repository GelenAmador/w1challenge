package main

import (
	"fmt"
	"os"

	"github.com/GelenAmador/w1challenge/bmi"
	"github.com/GelenAmador/w1challenge/mario"
	"github.com/GelenAmador/w1challenge/temperature"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: temperature, bmi, or mario")
		return
	}

	command := os.Args[1]

	switch command {
	case "temperature":
		temperature.RunTemperatureConverter()
	case "bmi":
		bmi.RunBMICalculator()
	case "mario":
		mario.RunMarioPyramid()
	default:
		fmt.Println("Unknown command:", command)
	}
}
