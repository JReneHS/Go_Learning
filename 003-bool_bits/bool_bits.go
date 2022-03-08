package main

import "fmt"

func main() {
	a := 10 //1010
	b := 3  //0011
	c := 8  //1000

	fmt.Printf("%v\n", a&b)  //00010 and
	fmt.Printf("%v\n", a|b)  //01011 or
	fmt.Printf("%v\n", a^b)  //01001 or exclusiva
	fmt.Printf("%v\n", a&^b) //01000 ?? ??
	fmt.Printf("%v\n", c)    //01000
	c = c << 1
	fmt.Printf("%v\n", c) //10000
	c = c >> 2
	fmt.Printf("%v\n", c) //00100
}
