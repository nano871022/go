package servers

import(
  "net/http"
  "../constants"
  "../types"
  "io"
)
const const_path_types = "tipos/"
const const_svc_type   = "tipo"

func init(){
  http.HandleFunc(constants.PathAPI+const_path_types+const_svc_type,typeFunc)
}

func typeFunc(response http.ResponseWriter, request *http.Request){
  valor := types.Tipos()
  io.WriteString(response,valor)
}
