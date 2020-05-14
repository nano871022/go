package arrays

import ("fmt")

func CopySlice() string{
  var valor string = ""
  a := []int {1,2,3,4,5,6,7,8,9,10,11}
  b := make([]int,2)
  c := make([]int,len(a))
  copy(b,a)
  copy(c,b)
  valor = valor + fmt.Sprintf("%v",a)
  valor = valor + fmt.Sprintf("%v",b)
  valor = valor + fmt.Sprintf("%v",c)
  return valor
}

func Slice() string{
  var valor string = ""
    b := [...]int {1,2,3,4,5,6,7,8,9,10,11}
    var a []int
    var c []int
    a = b[3:6]
    c = a[1:3]
    c[0] = 15

    valor = valor + fmt.Sprintf("a %v %v %v",a,len(a),cap(a))
    valor = valor + fmt.Sprintf("b %v %v %v",b,len(b),cap(b))
    valor = valor + fmt.Sprintf("c %v %v %v",c,len(c),cap(c))
    return valor
}
