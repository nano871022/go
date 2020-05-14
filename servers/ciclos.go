package servers

import(
  "../constants"
  "net/http"
  "strconv"
  "../iterate"
)

const const_svc_path_ciclo = "ciclos/"
const const_svc = "ciclo"


func init(){
  http.HandleFunc(constants.PathAPI+const_svc_path_ciclo+const_svc,ciclos)

}

func ciclos(response http.ResponseWriter,request *http.Request){
  listaValores,existe := request.URL.Query()[const_valor]
  if existe {
    x,_ := strconv.ParseInt(listaValores[0],10,8)
    iterate.Ciclo(response,int8(x))
  }
}
