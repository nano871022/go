package arrays

import ("fmt")

func Mapas()string{
  var valor string = ""
  valor = valor + fmt.Sprintf("mapas")
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
     valor = valor + fmt.Sprintf("No esta el dia")
   }

  valor = valor + fmt.Sprintf("%v",m)
  valor = valor + fmt.Sprintf("%v",m2)
  valor = valor + fmt.Sprintf("%v",m2["1"])
  valor = valor + fmt.Sprintf("%v",weekday[1])

  for numero,dia := range weekday{
    valor = valor + fmt.Sprintf("El día %s tiene como numero el %d \n",dia,numero)
  }
  return valor
}
