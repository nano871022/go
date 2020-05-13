package text

import(
  "fmt"
  "time"
)



func Ingreso (canal chan int){
  for{
    time.Sleep(1 * time.Second)
   var numero int
   fmt.Println("Ingrese opcion")

   fmt.Scanln(&numero)

   fmt.Printf("Opción ingresada %d \n",numero)
   canal <- numero
 }
}
