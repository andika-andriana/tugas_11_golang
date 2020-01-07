package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str1, str2 string = "125", "25"
	var num1, _ = strconv.Atoi(str1)
	var num2, _ = strconv.Atoi(str2)

	fmt.Println("Menghitung...")
	fmt.Println(num1, " + ", num2, " = ", num1+num2)
	fmt.Println(num1, " - ", num2, " = ", num1-num2)
	fmt.Println(num1, " x ", num2, " = ", num1*num2)
	fmt.Println(num1, " / ", num2, " = ", num1/num2)
	fmt.Println("Selesai...")
}
