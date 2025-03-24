package main

import (
	"fmt"
)

func main() {

	fmt.Println(2 + 3)

	fmt.Println("Hello aizen")

	fmt.Println(2 * 3)

	foo()

}

func foo() {
	//Boolean type
	var guess bool = true
	fmt.Println(guess)

	//integer type
	var size int8 = 55
	var distance int8 = 10
	var total int8 = size / distance
	var quant byte = 72
	//var total int = int(size) / distance
	fmt.Println(total)
	fmt.Println(quant)

	//float type
	var height1 float32 = 5.6
	var height2 float64 = 3442424234.32
	var ttt float64 = float64(height1) + height2
	var so float32 = 0
	fmt.Println(so / 0) //prints NaN

	fmt.Println(ttt)

	//complex type
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x / y)

	//rune literal
	var initial rune = 'A'
	var sssa int32 = 'b'
	fmt.Println(initial)
	fmt.Println(sssa)

	//data type conversion
	var rr int = 10
	var tt float64 = 30.2
	var sum1 float64 = float64(rr) + tt

	fmt.Println(sum1)
	fmt.Println(x + y)

}
