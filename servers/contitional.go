package servers

import (
  "net/http"
  "../constants"
  "../conditionals"
  "strconv"
  "io"
)

const const_path_conditional = "conditionals/"
const const_svc_conditional  = "conditional"
const const_svc_switch  = "switch"

func init(){
  http.HandleFunc(constants.PathAPI+const_path_conditional+const_svc_conditional, conditionales)
  http.HandleFunc(constants.PathAPI+const_path_conditional+const_svc_switch, switchFuc)
}

func conditionales(response http.ResponseWriter,request *http.Request){
  listaValores,existe := request.URL.Query()[const_valor]
  if existe {
    x,_ := strconv.ParseInt(listaValores[0],10,64)
    result := conditionals.CondicionalSi(int(x))
    io.WriteString(response,result)
  }
}

func switchFuc(response http.ResponseWriter,request *http.Request){
  listaValores,existe := request.URL.Query()[const_valor]
  if existe {
    x,_ := strconv.ParseInt(listaValores[0],10,64)
    result := conditionals.SwitchSentence(int8(x))
    io.WriteString(response,result)
  }
}
