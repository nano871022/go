package servers

import (
"net/http"
"io"
  "../constants"
  "../text"
)

const const_svc_path_text = "text/"
const const_svc_text = "texto"

func init(){
    http.HandleFunc(constants.PathAPI+const_svc_path_text+const_svc_text,texto)
}

func texto(response http.ResponseWriter, request *http.Request){
  listaValores,existe := request.URL.Query()[const_valor]
  if existe {
    var eje1 = text.Ejercicio1()+"///"+text.Texto(listaValores[0])
    io.WriteString(response,eje1)
  }
}
