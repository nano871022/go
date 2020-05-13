package structs

import ("fmt")

func Structur2(){
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

func Structuras(){
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
