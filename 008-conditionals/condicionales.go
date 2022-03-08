package main

import (
	"fmt"
)

/*
	== igual a...
	!= diferente de...
	< menor que
	> mayor que
	>= mayor o igual que
	<= menor o igual que
	&& and
	|| or

*/

func main() {
	x, y, z := 10, 20, 10

	if x == z {
		fmt.Println(x, "es igual a", z)
	} else {
		fmt.Println("Entra el else")
	}

	if x >= y {
		fmt.Println("Entra y en accion")
	}

	var i int = 0

	for i = 0; i < 10; i++ {
		fmt.Println("Hello World #", i+1)
	}

	//for como while
	i = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	i = 0
	for {
		if i == 2 {
			fmt.Println("Aqui hay un 2 7u7r")
			i++
			continue
		}

		fmt.Println("bb #", i)
		i++
		if i == 10 {
			break
		}
	}

	for k := 0; k <= 10; k++ {
		fmt.Println(k)
	}

}
