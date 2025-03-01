package main

import (
	"fmt"
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
	fmt.Print(msg)
	_, err := fmt.Scan(value)
	if err != nil {
		fmt.Println("Error. Expected a valid real number, got invalid instead")
	}
	return err
}
