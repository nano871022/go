package arrays

import ("fmt")

func CopySlice(){
  a := []int {1,2,3,4,5,6,7,8,9,10,11}
  b := make([]int,2)
  c := make([]int,len(a))
  copy(b,a)
  copy(c,b)
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

func Slice(){
    b := [...]int {1,2,3,4,5,6,7,8,9,10,11}
    var a []int
    var c []int
    a = b[3:6]
    c = a[1:3]
    c[0] = 15

    fmt.Println("a",a,len(a),cap(a))
    fmt.Println("b",b,len(b),cap(b))
    fmt.Println("c",c,len(c),cap(c))
}
