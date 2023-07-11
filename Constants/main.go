package main

import (
	"fmt"
	"math"
)

const userName string = "Rahul"

func main() {

	fmt.Println(userName)

	const num = 500000000

	const d = 3e20 / num
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(num))

}
