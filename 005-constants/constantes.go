package main

import (
	"fmt"
	//	"math"
)

const (
	S1 = iota //iota es un contador que genera constantes en bloque
	S2 = iota
	S3 = iota
	S4 = iota
	S5 = iota
)

const ( //el compilador detecta el bloque y rellena las demas constantes
	Z1 = iota
	Z2
	Z3
	Z4
)

const ( //el compilador detecta el bloque y rellena las demas constantes
	L1 = iota + 8
	L2
	L3
	L4
)

const ( //el compilador detecta el bloque y rellena las demas constantes
	_  = iota             //se pone una linea ya que muchas veces no se usa el cero 0
	KB = 1 << (10 * iota) //lo que hace es recorrer un x numero de bytes
	MB                    //Para dar los valores de kilobyte, megabyte, etc...
	GB
	TB
	PB
	EB
	ZB
	YB
)

const ( //Tambien iota se puede usar como identificador de bits
	isPrincipal = 1 << iota
	isPublic
	isTxt
	isVideo
	isGroup
	isProtected
	isModificable
	isFree
)

func main() {
	//Local scope tiene mayor prioridad
	const FirstConstant int = 57
	// genera un error ya que tiene que generar la constante antes de el metodo
	//	const ConFloat float64 = math.Tan(float64(FirstConstant))
	const A = 23
	var b int16 = 55
	fmt.Printf("%v, %T\n", FirstConstant, FirstConstant)
	fmt.Printf("%v, %T\n\n", A+b, A+b)

	fmt.Printf("S1 = %v, %T\n", S1, S1)
	fmt.Printf("S2 = %v, %T\n", S2, S2)
	fmt.Printf("S3 = %v, %T\n", S3, S3)
	fmt.Printf("S4 = %v, %T\n\n", S4, S4)

	fmt.Printf("Z1 = %v, %T\n", Z1, Z1)
	fmt.Printf("Z2 = %v, %T\n", Z2, Z2)
	fmt.Printf("Z3 = %v, %T\n", Z3, Z3)
	fmt.Printf("Z4 = %v, %T\n\n", Z4, Z4)

	fmt.Printf("Z1 = %v, %T\n", L1, L1)
	fmt.Printf("Z2 = %v, %T\n", L2, L2)
	fmt.Printf("Z3 = %v, %T\n", L3, L3)
	fmt.Printf("Z4 = %v, %T\n\n", L4, L4)

	fileSize := 4000000000.0

	fmt.Printf("%.2fGB\n\n", fileSize/GB)

	fmt.Printf("KiloByte  = %v bytes\n", KB)
	fmt.Printf("MegaByte  = %v KB\n", MB/KB)
	fmt.Printf("GigaByte  = %v MB\n", GB/MB)
	fmt.Printf("TeraByte  = %v GB\n", TB/GB)
	fmt.Printf("PetaByte  = %v TB\n", PB/TB)
	fmt.Printf("ExaByte   = %v PB\n", EB/PB)
	fmt.Printf("ZettaByte = %v EB\n", ZB/EB)
	fmt.Printf("YottaByte = %v ZB\n\n", YB/ZB)

	var roles byte = isPublic | isPrincipal | isProtected
	fmt.Printf("Roles: %08b\n\n", roles)
	fmt.Printf("is Principal?    %v\n", isPrincipal&roles == isPrincipal)
	fmt.Printf("is Public?       %v\n", isPublic&roles == isPublic)
	fmt.Printf("is Txt?          %v\n", isTxt&roles == isTxt)
	fmt.Printf("is Video?        %v\n", isVideo&roles == isVideo)
	fmt.Printf("is Group?        %v\n", isGroup&roles == isGroup)
	fmt.Printf("is Protected?    %v\n", isProtected&roles == isProtected)
	fmt.Printf("is Modificable?  %v\n", isModificable&roles == isModificable)
	fmt.Printf("is Free?         %v\n\n", isFree&roles == isFree)
}
