package main

import (
	"fmt"
	"strconv"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
	var b int = 1
	b = b + 1
	b++
	fmt.Println(b)
	b++
	fmt.Println(b)
	var celcius int = 100
	var farenheit int = 212
	fmt.Println(strconv.Itoa(celcius) + "C expressed as F : " + strconv.Itoa(convertCelciusToFarenheit(celcius)))
	fmt.Println(strconv.Itoa(farenheit) + "F expressed as C : " + strconv.Itoa(convertFarenheitToCelcius(farenheit)))
}

func convertCelciusToFarenheit(c int) int {
	var f int
	f = c*9/5 + 32
	return f
}

func convertFarenheitToCelcius(f int) int {
	var c int
	c = (f - 32) * 5 / 9
	return c
}
