package types

import ("fmt")

type Dinero int64

func (d Dinero) string()string{
  fmt.Printf("$ %d\n",d)
  return "$"
}

func Tipos(){
  dinero := Dinero(10000000)
  fmt.Printf("dinero %d o %s\n",dinero,dinero)
}
