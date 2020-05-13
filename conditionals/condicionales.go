package conditionals

import("fmt")

func CondicionalSi(){
  a:=3
  if a = a + 2 ; a > 6 {
    fmt.Println ("algo")
  }else if a < 6 {
    fmt.Println ("else if")
  }else{
    fmt.Println ("else")
  }
}
