package structs

import (
  "fmt"
  "strconv"
)

func Structur2()string{
  var texto string = ""
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
  texto = texto + fmt.Sprintf("%v",alejandro)
  texto = texto + alejandro.Nombre
  texto = texto + strconv.Itoa(alejandro.Edad)
  texto = texto + strconv.Itoa(int(alejandro.Semestre))
  texto = texto + fmt.Sprintf("%v",fisica)
  texto = texto + fisica.Nombre
  texto = texto + strconv.Itoa(fisica.Edad)
  texto = texto + strconv.Itoa(int(fisica.Semestre))
  return texto
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

func Structuras()(texto string){
  var persona Persona
  persona.Nombre = "Alejandro"
  persona.Edad = 32

  persona2 := Persona{Nombre: "Camilo", Edad: 10}

  persona3 := Persona{"Andres",11}

  mayor,diferencia := diferenciaEdad(persona2,persona3)

  texto = texto + fmt.Sprintln("%v",persona)
  texto = texto + fmt.Sprintln("%v",persona2)
  texto = texto + fmt.Sprintln("%v",persona3)
  texto = texto + fmt.Sprintf("El mayor es %s y la diferencia de edad es %d ",mayor,diferencia)
  return
}
