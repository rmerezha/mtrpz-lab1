package cli

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func readFile(arg string) ([]byte, error) {
	data, err := os.ReadFile(arg)
	if err != nil {
		return nil, errors.New("file read error (file does not exist, permission denied, etc.)")
	}
	return data, nil

}

func ReadFileFromArg(arg string) (a, b, c float64, err error) {
	bytes, err := readFile(arg)
	if err != nil {
		return
	}
	valuesAsStr := strings.SplitN(string(bytes), " ", 3)

	a, err = strconv.ParseFloat(strings.TrimSpace(valuesAsStr[0]), 64)
	if err != nil {
		err = errors.New("invalid file format")
		return
	}
	b, err = strconv.ParseFloat(strings.TrimSpace(valuesAsStr[1]), 64)
	if err != nil {
		err = errors.New("invalid file format")
		return
	}
	c, err = strconv.ParseFloat(strings.TrimSpace(valuesAsStr[2]), 64)
	if err != nil {
		err = errors.New("invalid file format")
		return
	}
	return
}
