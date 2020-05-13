package arrays

import ("fmt")

func Mapas(){
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
