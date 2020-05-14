package arrays

import ("fmt")

func ArraysUsos() string{
  var i [3]int
  i[0] = 1
  i[1] = 2
  i[2] = 3
  j := [...]int{1,2,3,4}
  return fmt.Sprintf("%v %v",i,j)
}
