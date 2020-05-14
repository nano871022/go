package servers

import(
  "fmt"
  "strconv"
  "net/http"
  "../operations"
  "../constants"
)

const const_valor = "valor"
const const_svc_suma = "suma"
const const_svc_resta = "resta"
const const_svc_multiplicacion = "multiplicacion"
const const_svc_divicion = "divicion"
const const_svc_path = "operaciones/"

func init(){
  http.HandleFunc(constants.PathAPI+const_svc_path+const_svc_suma, servicioSumar)
  http.HandleFunc(constants.PathAPI+const_svc_path+const_svc_divicion, servicioDividir)
  http.HandleFunc(constants.PathAPI+const_svc_path+const_svc_resta, servicioRestar)
  http.HandleFunc(constants.PathAPI+const_svc_path+const_svc_multiplicacion, servicioMultiplicar)
}

func servicioSumar(response http.ResponseWriter, request *http.Request) {
	x,y,existe := valuesInt(request)
	if existe {
		valor := operations.Suma(int8(x), int8(y))
		fmt.Fprintf(response, "Valor sumado :: %d + %d ",x,y,valor)
	}
}

func servicioDividir(response http.ResponseWriter, request *http.Request){
  listaValores,existe := request.URL.Query()[const_valor]
  if existe {
    x, _ := strconv.ParseFloat(listaValores[0], 32)
		y, _ := strconv.ParseFloat(listaValores[1], 32)
    valor := operations.Dividir(float32(x),float32(y))
    fmt.Fprintf(response,"Valor dividido %d / %d = %d",x,y,valor)
  }
}

func servicioRestar(response http.ResponseWriter,request *http.Request){
  x,y,existe := valuesInt(request)
  if existe {
    valor := operations.Restar(int8(x),int8(y))
    fmt.Fprintf(response,"Valor restado %d - %d = %d",x,y,valor)
  }
}

func servicioMultiplicar(response http.ResponseWriter,request *http.Request){
  x,y,existe := valuesInt(request)
  if existe {
    valor := operations.Multiplicar(int8(x),int8(y))
    fmt.Fprintf(response,"Valor multiplicado %d X %d = %d",x,y,valor)
  }
}

func valuesInt(request *http.Request)(x int,y int, existe bool){
  listaValores,existe2 := request.URL.Query()[const_valor]
  existe = existe2
  if existe {
    x2, _ := strconv.ParseInt(listaValores[0], 10,8)
		y2, _ := strconv.ParseInt(listaValores[1], 10,8)
    x = int(x2)
    y = int(y2)
    return
  }else{
    x = 0
    y = 0
    existe = false
    return
  }

}
