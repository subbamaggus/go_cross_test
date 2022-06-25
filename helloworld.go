package main

import (
    "fmt"
    "mymod/mytestlib"
)

func main() {
    fmt.Println("hello world")
    
    result := mytestlib.MySum(1,2)
    fmt.Println("MySum: ", result)
    
    result = mytestlib.MyProd(1,2)
    fmt.Println("MyProd: ", result)
    
    fmt.Scanln()
}    