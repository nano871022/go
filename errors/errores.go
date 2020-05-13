package errors

import (
  "fmt"
  "errors"
)

var (
  Error1 = errors.New("Mensaje 1")
  Error2 = errors.New("Mensaje 2")
)

func TestErrors2(){
  test := Testing {Nombre:"Alejandro"}
  test2 := Testing {Errores: Error1}

  if(test2.Errores != nil){
    fmt.Println(test2.Errores)
  }
  fmt.Println(test.Nombre)
}

func TestErrors(val int8)(err error){
  if val == 2{
    return Error1
  }else if val == 3{
    return Error2
  }
  return nil
}
