package main

import (
	"fmt"
	"math/rand"
	"time"
)

//n int // возврат слайса
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Vector[T Number] []T

func sliceGen[T any](inpSlice []T) []T { // generation slice of different types random values
	outSlice := inpSlice
	return outSlice

}

func randFloats(min, max float64, n int) []float64 { // Generation random float64. n - length (number) of randoms
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

//func sliceConcatenation[T, F any](firstSlice []T, secondSlice []F) []T {
//	ConcatenatedSlice := append(firstSlice, []T(secondSlice)...)
//	return ConcatenatedSlice
//}

func AddVector[T, F Number](vec1 Vector[T], vec2 Vector[F]) Vector[T] {
	var result Vector[T]
	for i := range vec1 {
		result = append(result, vec1[i]+vec2[i])
	}
	return result
}

func main() {

	testArrayOfInts := Vector[int]{rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100)}
	sliceOfTestInts := sliceGen[int](testArrayOfInts)
	fmt.Println(sliceOfTestInts)

	rand.Seed(time.Now().UnixNano())
	testArrayOfFloats := randFloats(1.00, 100.00, 1)
	sliceOfTestFloats := sliceGen[float64](testArrayOfFloats)
	fmt.Println(sliceOfTestFloats)

	//fmt.Println(sliceConcatenation(sliceOfTestInts, sliceOfTestFloats))
	result := AddVector(sliceOfTestInts, sliceOfTestFloats)
	fmt.Println(result)
}

//

//package main
//
//import "fmt"
//
//func main() {
//	var a float64
//	fmt.Scan(&a)
//	if a <= 0 {
//		fmt.Printf("число %2.2f не подходит", a)
//	} else if a > 10000 {
//		fmt.Printf("%e", a)
//	} else {
//		a = a * a
//		fmt.Printf("%.4f", a)
//	}
//
//}

//var a rune = 'Ы'
//fmt.Printf("%U", a)
////fmt.Println(string(a))
//var a float64 = 100.123456
//fmt.Printf("это число %f типа %T", a, a)
// вывод: это число 100.123456 типа float64

//package main
//
//import (
//	"fmt"
//	"math"
//	"strconv"
//)
//
//func quantityOfNumbersCnt(a float64) int {
//	var i, y float64
//	var count int
//
//	for i = 10000; i >= 1; i = i / 10 {
//		y = a / i
//		if y <= 1 {
//			continue
//		} else {
//			count++
//		}
//	}
//	return count
//}
//
//func main() {
//	var x, y float64
//	fmt.Scan(&x, &y)
//	if x < 0 || x > 10000 || y < 0 || y > 10000 {
//		panic("Wrong numbers")
//	}
//	w := quantityOfNumbersCnt(x)
//	z := quantityOfNumbersCnt(y)
//
//	firstnum := int(x)
//	secondNum := int(y)
//
//	var count int
//	var answerString string
//
//	for i := (w - 1); i >= 0; i-- {
//		count = firstnum / int(math.Pow(10, float64(i)))
//		firstnum = firstnum % int(math.Pow(10, float64(i)))
//		secondNum2 := secondNum
//		for q := (z - 1); q >= 0; q-- {
//			var count2 int
//
//			count2 = secondNum2 / int(math.Pow(10, float64(q)))
//			secondNum2 = secondNum2 % int(math.Pow(10, float64(q)))
//			if count == count2 {
//				answerString = answerString + strconv.Itoa(count) + " "
//			}
//		}
//	}
//	fmt.Println(answerString)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var x, p, y int
//	fmt.Scan(&x, &p, &y)
//	x1 := x * 100
//
//	//p1 := p * 100
//	y1 := y * 100
//	for i := 1; i <= y; i++ {
//		x1 = x1 + x1*(p)/100
//		fmt.Println(x1)
//		if x1 >= y1 {
//			fmt.Println(i)
//			break
//		}
//	}
//}

// package main

// import "fmt"

// func main() {
// 	var v int
// 	for ; v < 101; fmt.Scan(&v) {
// 		if v < 10 {
// 			continue
// 		} else if v > 100 {
// 			break
// 		} else {
// 			fmt.Println(v)
// 		}
// 	}
// }

//package main
//
//import "fmt"
//
//func main() {
//	var n, c, d int
//	fmt.Scan(&n, &c, &d)
//	for i := 1; i <= n; i++ {
//		z := i%c == 0
//		y := i%d != 0
//		//fmt.Println(i, ":", z, y)
//		if (n > 0) && c > 0 && d > 0 && (n <= 10000) && (c <= 10000) && (d <= 10000) {
//			if z && y {
//				fmt.Println(i)
//				break
//
//			}
//		}
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	var a int
//	var i int
//	var count int
//	for fmt.Scan(&a); a != 0; {
//		if a > i {
//			i = a
//			count = 1
//		} else if a == i {
//			count = count + 1
//		}
//		fmt.Scan(&a)
//	}
//	fmt.Println(count)
//}

//package main
//
//import "fmt"
//
////func validator(a int) int {
////	if a > 9 && a < 100 && (a%8 == 0) {
////		return a
////	} else {
////		return 0
////	}
////}
//
//func main() {
//	var a, b, sum int
//	fmt.Scan(&a)
//
//	for i := 1; i <= a; i++ {
//		fmt.Scan(&b)
//		if b > 9 && b < 100 && (b%8 == 0) {
//			sum = sum + b
//		}
//		//fmt.Println(sum)
//	}
//	fmt.Println(sum)
//}

//var a, b, c, d, e, f int
//fmt.Scan(&a)
//fmt.Scan(&b, &c, &d, &e, &f)
////val := validatorStruct{b, c, d, e, f, 0, 0}
//
//bb := validator(b)
//cc := validator(c)
//dd := validator(d)
//ee := validator(e)
//ff := validator(f)
//
//sum := bb + cc + dd + ee + ff
//
//fmt.Println(sum)

//var sum int
//for i := 1; i < a; i++ {
//	sum = sum +
//}
//fmt.Scan(&b, &c, &d, &e, &f)
//fmt.Println(a, b, c, d, e, f)

//package main
//
//import "fmt"
//
//func main() {
//	var v, w int
//	count, err := fmt.Scan(&v, &w)
//	fmt.Println(count)
//	if !(v < w && w <= 100 && v <= 100) {
//		panic("wrong numbers")
//	}
//	if err != nil {
//		panic(err)
//	} else if count != 2 {
//		fmt.Println("right number of input values=2", "\nnumber of input values=", count)
//	} else {
//		fmt.Println("fmt.Scat: right number of input values and no errors")
//	}
//
//	var numb int
//
//	for i := v; i <= w; i++ {
//		numb = numb + i
//	}
//	fmt.Println(numb)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var v int
//	fmt.Scan(&v)
//
//	a := v % 400
//	n1 := a == 0
//
//	b := v % 4
//	n2 := b == 0
//	c := v % 100
//	n3 := c == 0
//	n4 := n2 && !n3
//
//	if n1 || n4 {
//		fmt.Println("YES")
//	} else {
//		fmt.Println("NO")
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	var v int
//	fmt.Scan(&v)
//	var a, b, c, d, e, f, g, h, j, k, l int
//	a = v / 100000
//	//fmt.Println(a)
//	b = v % 100000
//	c = b / 10000
//	//fmt.Println(c)
//	d = b % 10000
//	e = d / 1000
//	//fmt.Println(e)
//	f = d % 1000
//	g = f / 100
//	//fmt.Println(g)
//	h = f % 100
//	j = h / 10
//	//fmt.Println(j)
//	k = h % 10
//	l = k
//	//fmt.Println(l)
//
//	y := a + c + e
//	z := g + j + l
//
//	if z == y {
//		fmt.Println("YES")
//	} else {
//		fmt.Println("NO")
//	}
//	//switch {
//	//case v > 999999:
//	//	fmt.Println("wrong number. Input value cant be more than 10000")
//	//case v > 99999:
//	//	a = v / 100000
//	//	fmt.Println(a)
//	//case v > 9999:
//	//	b = v / 10000
//	//	fmt.Println(b)
//	//case v > 999:
//	//	c = v / 1000
//	//	fmt.Println(c)
//	//case v > 99:
//	//	d = v / 100
//	//	fmt.Println(d)
//	//case v > 9:
//	//	e = v / 10
//	//	fmt.Println(e)
//	//case v > 0:
//	//	f = v
//	//	fmt.Println(f)
//	//default:
//	//	fmt.Println("wrong number. Input value must be more than 0")
//	//}
//	//fmt.Println(a)
//	//fmt.Println(b)
//	//fmt.Println(c)
//	//fmt.Println(d)
//	//fmt.Println(e)
//	//fmt.Println(f)
//
//}

//package main
//
//import "fmt"
//
//func main() {
//	var v int
//	count, err := fmt.Scan(&v)
//
//	if err != nil {
//		panic(err)
//	} else if count != 1 {
//		//fmt.Println("right number of input values=1", "\nnumber of input values=", count)
//	} else {
//		//fmt.Println("fmt.Scat: right number of input values and no errors")
//	}
//	var a int
//	switch {
//	case v > 10000:
//		fmt.Println("wrong number. Input value cant be more than 10000")
//	case v == 1:
//		fmt.Println("1")
//	case v > 999:
//		a = v / 1000
//		fmt.Println(a)
//	case v > 99:
//		a = v / 100
//		fmt.Println(a)
//	case v > 9:
//		a = v / 10
//		fmt.Println(a)
//	case v > 0:
//		a = v
//		fmt.Println(a)
//	default:
//		fmt.Println("wrong number. Input value must be more than 0")
//	}
//
//}

//package main
//
//import "fmt"
//
//func main() {
//	var v int
//	fmt.Scan(&v)
//	a := v / 100
//	//fmt.Println(a)
//	b := v % 100
//	c := b / 10
//	//fmt.Println(c)
//	d := b % 10
//	//fmt.Println(d)
//	if a != c && c != d && a != d {
//		fmt.Println("YES")
//	} else {
//		fmt.Println("NO")
//	}
//}

//package main
//
//import (
//	"errors"
//	"fmt"
//	"math"
//)
//
//type InputDater interface {
//	verifyScanInputDataFormat() error
//}
//
//type typeInputData struct {
//	number float64
//}
//
//func (z *typeInputData) verifyScanInputDataFormat() error {
//	var a float64
//	a = math.Round(z.number)
//	//fmt.Println("input value=", z.number, "round value=", a)
//	if a != z.number {
//		return errors.New("invalid Scan data input type. Need: int")
//	} else {
//		return nil
//	}
//
//}
//
//func main() {
//	var v float64
//	count, err := fmt.Scan(&v)
//
//	if err != nil {
//		panic(err)
//	} else if count != 1 {
//		//fmt.Println("right number of input values=1", "\nnumber of input values=", count)
//	} else {
//		//fmt.Println("fmt.Scat: right number of input values and no errors")
//	}
//
//	var z InputDater = &typeInputData{v}
//
//	err = z.verifyScanInputDataFormat()
//	if err != nil {
//		panic(err)
//	}
//
//	if v > 0 {
//		fmt.Println("Число положительное")
//	} else if v < 0 {
//		fmt.Println("Число отрицательное")
//	} else {
//		fmt.Println("Ноль")
//	}
//
//}
