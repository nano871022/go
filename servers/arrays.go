package servers

import (
  "net/http"
  "../constants"
  "../arrays"
  "io"
)

const const_path_arrays = "arrays/"
const const_svc_array   = "array"
const const_svc_mapas   = "mapas"
const const_svc_slice   = "slice"
const const_svc_copy   = "copy"

func init(){
  http.HandleFunc(constants.PathAPI+const_path_arrays+const_svc_array,arrayFunc)
  http.HandleFunc(constants.PathAPI+const_path_arrays+const_svc_mapas,mapasFunc)
  http.HandleFunc(constants.PathAPI+const_path_arrays+const_svc_slice,sliceFunc)
  http.HandleFunc(constants.PathAPI+const_path_arrays+const_svc_copy,copyFunc)
}

func arrayFunc(response http.ResponseWriter,request *http.Request){
  io.WriteString(response,arrays.ArraysUsos())
}

func mapasFunc(response http.ResponseWriter,request *http.Request){
 io.WriteString(response,arrays.Mapas())
}

func sliceFunc(response http.ResponseWriter,request *http.Request){
  io.WriteString(response,arrays.Slice())
}

func copyFunc(response http.ResponseWriter,request *http.Request){
  io.WriteString(response,arrays.CopySlice())
}
