package main

import (
	"fmt"
)

func main() {
	s1 := "hello world" //las strings son inmutables
	//son cadenas de bytes (int8)
	//puedes hacer operaciones de cadenas
	s2 := "Good morning"
	s3 := s1 + " " + s2

	sbytes := []byte(s1) //las cadenas son arreglos de bytes
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", s3)
	fmt.Printf("length of s3 is: %d, %v\n", len(s3), len(s3))
	fmt.Printf("%v, %T\n", s1[2], s1[2])
	fmt.Printf("%v, %T\n", string(s1[2]), s1[2])
	fmt.Printf("%v, %T\n", sbytes, sbytes)

	//los strings no se pueden modificar...
	//pero los arreglos de bytes si... 7u7
	sbytes[2] = 109
	//Con el constructor (o metodo, no se) de string pasas de arreglo a string
	s4 := string(sbytes)

	fmt.Printf("%v, %T\n", sbytes, sbytes)
	fmt.Printf("%v, %T\n", string(sbytes), sbytes) //casteo de byte a string
	fmt.Printf("%v, %T\n", s4, s4)

	r := 'm'          //int32
	var rr rune = 'm' //es un sinonimo de int32 y reprensenta codigo Unicode
	var r3 byte = 'm' //es un sinonimo de uint8
	fmt.Printf("%v, %T\n", r, r)
	fmt.Printf("%v, %T\n", rr, rr)
	fmt.Printf("%v, %T\n", r3, r3)

}
