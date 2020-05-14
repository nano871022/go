package servers

import(
  "net/http"
  "../structs"
  "io"
  "../constants"
)


const const_path_struct = "structs/"
const const_svc_struct1 = "struct1"
const const_svc_struct2 = "struct2"

func init(){
  http.HandleFunc(constants.PathAPI+const_path_struct+const_svc_struct1,runStruct1)
  http.HandleFunc(constants.PathAPI+const_path_struct+const_svc_struct2,runStruct2)
}

func runStruct1(response http.ResponseWriter,request *http.Request) {
	io.WriteString(response, structs.Structuras())
}

func runStruct2(response http.ResponseWriter,request *http.Request) {
	io.WriteString(response, structs.Structur2())
}
