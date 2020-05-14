package types

import ("fmt")

type Dinero int64

func (d *Dinero) string()string{
  return fmt.Sprintf("$ %v",d)
}

func Tipos()(valor string){
  dinero := Dinero(10000000)
  valor = "dinero "+dinero.string()+" \n"
  return
}
