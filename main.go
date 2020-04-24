package main

import (
  "fmt"
  "strconv"
  "errors"
)

var (
  Error1 = errors.New("Mensaje 1")
  Error2 = errors.New("Mensaje 2")
)
type Dinero int64
type Persona struct{
  Nombre string
  Edad int
}

type Estudiante struct{
  Persona
  Semestre int8
}
type Profesor struct{
  Estudiante
  Semestre int8
}

type Testing struct{
  Nombre string
  Errores error
}

const VALUE = "texto"
var numero int
func main(){
 numero = 10
 fmt.Println("hi",VALUE," - ",numero)

 fmt.Println(suma(10,40))
 fmt.Println(restar(10,20))
 fmt.Println(dividir(10.2,2.4))

 texto("pruebastextos")
 ciclo(5)
 condicionalSi()
 //ingreso()
 switchSentence(2)
 arraysUsos()
 slice()
 copySlice()
 mapas()
 tipos()
 structuras()
 structur2()
 e1:=testErrors(1)
 e2:=testErrors(2)
 e3:=testErrors(3)
 if e1 != nil {
   fmt.Println(e1)
 }
 if e2 != nil {
   fmt.Println(e2)
 }
 if e3 != nil {
   fmt.Println(e3)
 }
 testErrors2()
}

func testErrors2(){
  test := Testing {Nombre:"Alejandro"}
  test2 := Testing {Errores: Error1}

  if(test2.Errores != nil){
    fmt.Println(test2.Errores)
  }
  fmt.Println(test.Nombre)
}

func testErrors(val int8)(err error){
  if val == 2{
    return Error1
  }else if val == 3{
    return Error2
  }
  return nil
}

func structur2(){
  alejandro := Estudiante{
    Persona{
      Nombre: "Alejandro",
      Edad: 32,
      },
     2,
  }
  fisica := Profesor{
    Estudiante{
      Persona{
        Nombre: "Camilo",
        Edad: 21,
      },
      5,
    },
    10,
  }
  fmt.Println(alejandro)
  fmt.Println(alejandro.Nombre)
  fmt.Println(alejandro.Edad)
  fmt.Println(alejandro.Semestre)
  fmt.Println(fisica)
  fmt.Println(fisica.Nombre)
  fmt.Println(fisica.Edad)
  fmt.Println(fisica.Semestre)
}

func diferenciaEdad(persona1,persona2 Persona)(mayor string,diff int){
  diff = persona1.Edad - persona2.Edad
  if(diff >= 0){
    mayor = persona2.Nombre
  }else{
    mayor = persona2.Nombre
    diff = -1 * diff
  }
  return
}

func structuras(){
  var persona Persona
  persona.Nombre = "Alejandro"
  persona.Edad = 32

  persona2 := Persona{Nombre: "Camilo", Edad: 10}

  persona3 := Persona{"Andres",11}

  mayor,diferencia := diferenciaEdad(persona2,persona3)

  fmt.Println(persona)
  fmt.Println(persona2)
  fmt.Println(persona3)
  fmt.Printf("El mayor es %s y la diferencia de edad es %d ",mayor,diferencia)
}

func (d Dinero) string()string{
  fmt.Printf("$ %d\n",d)
  return "$"
}

func tipos(){
  dinero := Dinero(10000000)
  fmt.Printf("dinero %d o %s\n",dinero,dinero)
}

func mapas(){
  fmt.Println("mapas")
  m := make(map[string]string)
  m2 := make(map[string]string,7)

  m["nombre"] = "valor"
  m["apellidos"] = "apellido"

  m2["1"] = "lunes"

  weekday := map[int]string{
    1:"Domingo",
    2:"Lunes",
    3:"Martes",
  }

  delete(m,"nombre")

   _, esta := weekday[4]
   if(!esta){
     fmt.Println("No esta el dia")
   }

  fmt.Println(m)
  fmt.Println(m2)
  fmt.Println(m2["1"])
  fmt.Println(weekday[1])

  for numero,dia := range weekday{
    fmt.Printf("El día %s tiene como numero el %d \n",dia,numero)
  }
}

func copySlice(){
  a := []int {1,2,3,4,5,6,7,8,9,10,11}
  b := make([]int,2)
  c := make([]int,len(a))
  copy(b,a)
  copy(c,b)
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

func slice(){
    b := [...]int {1,2,3,4,5,6,7,8,9,10,11}
    var a []int
    var c []int
    a = b[3:6]
    c = a[1:3]
    c[0] = 15

    fmt.Println("a",a,len(a),cap(a))
    fmt.Println("b",b,len(b),cap(b))
    fmt.Println("c",c,len(c),cap(c))
}


func arraysUsos(){
  var i [3]int
  i[0] = 1
  i[1] = 2
  i[2] = 3
  j := [...]int{1,2,3,4}
  fmt.Println(i)
  fmt.Println(j)
}

func switchSentence(dia int8){
  switch dia {
    case 1: fmt.Println("Domingo")
    case 2: fmt.Println("Lunes")
    case 3: fmt.Println("Martes")
    case 4: fmt.Println("Miercoles")
  default: fmt.Println("Opcion Invalida")
  }

  switch {
  case dia == 1:fmt.Println("Domingo")
  case dia == 2:fmt.Println("Lunes")
      fallthrough
  case dia >= 3 && dia <= 4: fmt.Println("Martes")
  default: fmt.Println("Opcion Invalida")
  }
}

func ingreso (){
   var numero int
   fmt.Println("Ingrese numero")

   fmt.Scanln(&numero)

   fmt.Printf("Numero ingresado %d \n",numero)
}

func condicionalSi(){
  a:=3
  if a = a + 2 ; a > 6 {
    fmt.Println ("algo")
  }else if a < 6 {
    fmt.Println ("else if")
  }else{
    fmt.Println ("else")
  }
}

func ciclo(target int8){
  var i int8 = 1
  fmt.Println("ciclo for tipo while")
  for  i < target {
    fmt.Println(i)
    i++
  }
fmt.Println("ciclo for clasico")
  for i = 0 ; i < target ; i++ {
    fmt.Println(i)
  }
  fmt.Println("do while")
  i = 0;
  for {
    fmt.Println("Test")
    i++
    if i > 2 {
      break
    }
  }
}

func texto(val1 string){
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


func suma(val1,val2 int8) int8{
  return val1 + val2
}

func restar (val1,val2 int8) int8{
  return val1 - val2
}

func dividir (val1, val2 float32)float32{
  return val1 / val2
}
