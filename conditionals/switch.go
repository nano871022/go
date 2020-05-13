package conditionals

import("fmt")

func SwitchSentence(dia int8){
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
