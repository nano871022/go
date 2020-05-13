package text

import(
  "fmt"
  "strconv"
)

const VALUE = "texto"
var numero int

func Texto(val1 string){
  fmt.Println(val1)
  longitud := len(val1)
  fmt.Println(longitud)
  caracter := val1[2]
  fmt.Println(caracter)
  substr := val1[2:5]
  fmt.Println(substr)
  val1 = val1 + "otra"
  fmt.Println(val1)
  val2 := `
      <html>
        <head></head>
        <body></body>
      </html>
      `
  fmt.Println(val2)
  edad := "Edad "+strconv.Itoa(30)
  fmt.Println(edad)
}

func Ejercicio1(){
  numero = 10
  fmt.Println("hi",VALUE," - ",numero)
}
