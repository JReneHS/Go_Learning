package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 22

	edad_str := strconv.Itoa(edad)

	edad_int, _ := strconv.Atoi(edad_str)

	fmt.Println("mi edad es ", edad)
	fmt.Println("Mi edad es => ", edad_str)
	fmt.Println("Mi edad + 10 es: ", edad_int+10)

	fmt.Printf("Mi edad es %d, %T, %v\n", edad, edad, edad)
}
