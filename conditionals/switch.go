package conditionals



func SwitchSentence(dia int8)string{
  var texto string = ""
  switch dia {
    case 1: texto = "Domingo"
    case 2:  texto = "Lunes"
    case 3:  texto = "Martes"
    case 4:  texto = "Miercoles"
  default:  texto = "Opcion Invalida"
  }

  switch {
  case dia == 1: texto = texto + "Domingo"
  case dia == 2: texto =  texto +"Lunes"
      fallthrough
  case dia >= 3 && dia <= 4:  texto = texto + "Martes"
  default: texto = texto + "Opcion Invalida"
  }
  return texto
}
