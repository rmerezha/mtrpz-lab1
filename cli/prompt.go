package cli

import (
	"fmt"
	"strconv"
)

func GetInput() (a, b, c float64) {
	type data struct {
		msg   string
		value *float64
	}
	dataSet := []data{
		{"a = ", &a},
		{"b = ", &b},
		{"c = ", &c},
	}
	for _, val := range dataSet {
		promptUntilValid(val.msg, val.value)
	}
	return
}

func promptUntilValid(msg string, value *float64) {
	for {
		if err := prompt(msg, value); err == nil {
			break
		}
	}
}

func prompt(msg string, value *float64) error {
	var scanValue string
	fmt.Print(msg)
	fmt.Scan(&scanValue)
	parsedValue, err := strconv.ParseFloat(scanValue, 64)
	if err != nil {
		fmt.Printf("Error. Expected a valid real number, got %s instead\n", scanValue)
	}
	*value = parsedValue
	return err
}
