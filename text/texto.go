package text

import(

  "strconv"
)

const VALUE = "texto"
var numero int

func Texto(val1 string) string{
  var texto string = ""
  texto = val1 + "\n"
  longitud := len(val1)
  texto = texto + strconv.Itoa(longitud) + "\n"
  caracter := string(val1[2] )
  texto = caracter + "\n"
  substr := val1[2:5]
  texto = texto +  substr + "\n"
  val1 = val1 + "otra"
  texto = texto + val1 + "\n"
  val2 := `
      <html>
        <head></head>
        <body>Cuerpo Mesaje</body>
      </html>
      `
  texto = texto + val2 + "\n"
  edad := "Edad "+strconv.Itoa(30)
  texto = texto + edad + "\n"
  return texto
}

func Ejercicio1() string{
  numero = 10
  return "hi" + VALUE + " - " + strconv.Itoa(numero)
}
