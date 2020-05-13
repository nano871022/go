package folder

import (
  "fmt"
//  "errors"
  "sync"
)

var waitGroup sync.WaitGroup

func Paralel(ciclos int8,name string){
   var i int8
  for i = 0; i < ciclos ;i++{
    fmt.Printf("%d del proceso %s",i,name)
  }
  waitGroup.Done()
}
