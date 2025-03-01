package main

import (
	"fmt"
	"strconv"
)

func getInput() (a, b, c float64) {
	data := map[string]*float64 {
		"a = ": &a,
		"b = ": &b,
		"c = ": &c,
	}
    for msg, value := range data {
		promptUntilValid(msg, value)

	}
	return
}

func promptUntilValid(msg string, value *float64) {
	for {
		if err := prompt(msg, value); err == nil {
			break
		}
		fmt.Println("Некорректный ввод, попробуйте еще раз.")
	}
}

func prompt(msg string, value *float64) error {
	var scanValue string
	fmt.Print(msg)
	fmt.Scan(&scanValue)
	parsedValue, err := strconv.ParseFloat(scanValue, 64)
	if err != nil {
		fmt.Println("Error. Expected a valid real number, got invalid instead")
	}
	*value = parsedValue
	return err
}
