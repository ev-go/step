package main

import (
	"errors"
	"fmt"
	"math"
)

type InputDater interface {
	verifyScanInputDataFormat() error
}

type typeInputData struct {
	number float64
}

func (z *typeInputData) verifyScanInputDataFormat() error {
	var a float64
	a = math.Round(z.number)
	//fmt.Println("input value=", z.number, "round value=", a)
	if a != z.number {
		return errors.New("invalid Scan data input type. Need: int")
	} else {
		return nil
	}

}

func main() {
	var v float64
	count, err := fmt.Scan(&v)

	if err != nil {
		panic(err)
	} else if count != 1 {
		//fmt.Println("right number of input values=1", "\nnumber of input values=", count)
	} else {
		//fmt.Println("fmt.Scat: right number of input values and no errors")
	}

	var z InputDater = &typeInputData{v}

	err = z.verifyScanInputDataFormat()
	if err != nil {
		panic(err)
	}

	if v > 0 {
		fmt.Println("Число положительное")
	} else if v < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}

}
