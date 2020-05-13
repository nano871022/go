package iterate

import("fmt")

func Ciclo(target int8){
  var i int8 = 1
  fmt.Println("ciclo for tipo while")
  for  i < target {
    fmt.Println(i)
    i++
  }
fmt.Println("ciclo for clasico")
  for i = 0 ; i < target ; i++ {
    fmt.Println(i)
  }
  fmt.Println("do while")
  i = 0;
  for {
    fmt.Println("Test")
    i++
    if i > 2 {
      break
    }
  }
}
