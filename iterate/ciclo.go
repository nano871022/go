package iterate

import(
  "net/http"
  "io"
  "strconv"
)

func Ciclo(response http.ResponseWriter,target int8){
  var i int8 = 1
  io.WriteString(response,"ciclo for tipo while\n")
  for  i < target {
    io.WriteString(response,strconv.Itoa(int(i)))
    i++
  }
io.WriteString(response,"ciclo for clasico\n")
  for i = 0 ; i < target ; i++ {
    io.WriteString(response,strconv.Itoa(int(i)))
  }
  io.WriteString(response,"do while\n")
  i = 0;
  for {
    io.WriteString(response,"Test\n")
    i++
    if i > 2 {
      break
    }
  }
}
