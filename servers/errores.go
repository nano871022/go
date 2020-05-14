package servers

import(
  "../errors"
  "io"
  "net/http"
  "../constants"
  "fmt"
)

const const_path_error = "errores/"
const const_svc_error  = "error"

func init(){
  http.HandleFunc(constants.PathAPI+const_path_error+const_svc_error,runErrors)
}

func runErrors(response http.ResponseWriter, request *http.Request) {
  texto := ""
	e1 := errors.TestErrors(1)
	e2 := errors.TestErrors(2)
	e3 := errors.TestErrors(3)
	if e1 != nil {
		texto = texto + fmt.Sprintf("%v",e1)
	}
	if e2 != nil {
		texto = texto + fmt.Sprintf("%v",e2)
	}
	if e3 != nil {
		texto = texto + fmt.Sprintf("%v",e3)
	}
	texto = texto + errors.TestErrors2()
  io.WriteString(response,texto)
}
