package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/rmerezha/mtrpz-lab1/cli"
	"github.com/rmerezha/mtrpz-lab1/mathutils"
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	if len(args) > 2 {
		fmt.Fprintln(os.Stderr, errors.New("too many args"))
		return 1
	}
	var a, b, c float64
	if len(args) == 1 {
		a, b, c = cli.GetInput()
	} else {
		var err error
		a, b, c, err = cli.ReadFileFromArg(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 255
		}
	}
	fmt.Printf("Equation is: (%.2f) x^2 + (%.2f) x + (%.2f) = 0\n", a, b, c)
	return calculate(a, b, c)
}

func calculate(a, b, c float64) int {
	if a == 0 {
		fmt.Fprintln(os.Stderr, errors.New("a cannot be 0"))
		return 2
	}

	roots, hasSolution := mathutils.SolveQuadratic(a, b, c)
	if !hasSolution {
		fmt.Println("There are 0 roots")
		return 0
	}

	fmt.Println("There are", len(roots), "roots")
	for index, root := range roots {
		fmt.Printf("x%d = %.2f\n", index+1, root)
	}
	return 0
}
