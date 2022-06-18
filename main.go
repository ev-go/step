package main

import (
	"errors"
	"fmt"
)

type typeInputData struct{}

func (v *typeInputData) verifyScanInputDataFormat() (string, error) {
	var a float64
	a = v % 1
	if a != 0 {
		return "", errors.New("invalid Scan data input type. Need: int")
	}
}

func main() {
	var v float32
	count, err := fmt.Scan(&v)

	if err != nil {
		panic(err)
	} else if count != 1 {
		fmt.Println("right number of input values=1", "\nnumber of input values=", count)
	} else {
		fmt.Println("fmt.Scat: right number of input values and no errors")
	}

	if v > 0 {
		fmt.Println("число положительное")
	} else if v < 0 {
		fmt.Println("число отрицательное")
	} else {
		fmt.Println("число не полодительное и не отрицательное")
	}

}
