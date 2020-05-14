package main

import (
	"fmt"

	"./arrays"
	"./conditionals"
	"./errors"
	"./servers"

	//  "./folder"

	"./iterate"
	"./operations"
	"./structs"
	"./text"
	"./types"
)

var port = 9191
var canal chan int

func main() {
	servers.StartServer(port)
	old()
}

func old() {
	canal = make(chan int)
	var option int
	go text.Ingreso(canal)
	for {
		menu()
		option = <-canal
		fmt.Println("_______________________________________________________________")
		switch option {
		case 1:
			runOperations()
		case 2:
			runText()
		case 3:
			runErrors()
		case 4:
			runArrays()
		case 5:
			runTypes()
		case 6:
			runIterates()
		case 7:
			runConditionals()
		case 8:
			runStructs()
		default:
			fmt.Println("Opcion no valida")
		}
		if option == 0 {
			break
		}
	}
}

func menu() {
	fmt.Println("===============================================================")
	fmt.Println("1. Operaciones")
	fmt.Println("2. Texto")
	fmt.Println("3. Errores")
	fmt.Println("4. Arrays")
	fmt.Println("5. Tipos")
	fmt.Println("6. Iteraciones")
	fmt.Println("7. Condicionales")
	fmt.Println("8. Estructuras")
	fmt.Println("0. Salida")
}

func runStructs() {
	structs.Structuras()
	structs.Structur2()
}

func runConditionals() {
	conditionals.CondicionalSi()
	conditionals.SwitchSentence(2)
}

func runIterates() {
	iterate.Ciclo(5)
}

func runTypes() {
	types.Tipos()
}

func runArrays() {
	arrays.ArraysUsos()
	arrays.Slice()
	arrays.CopySlice()
	arrays.Mapas()
}

func runErrors() {
	e1 := errors.TestErrors(1)
	e2 := errors.TestErrors(2)
	e3 := errors.TestErrors(3)
	if e1 != nil {
		fmt.Println(e1)
	}
	if e2 != nil {
		fmt.Println(e2)
	}
	if e3 != nil {
		fmt.Println(e3)
	}
	errors.TestErrors2()
}

func runText() {
	text.Ejercicio1()
	text.Texto("pruebastextos")
}

func runOperations() {
	fmt.Println(operations.Suma(10, 40))
	fmt.Println(operations.Restar(10, 20))
	fmt.Println(operations.Dividir(10.2, 2.4))
}
