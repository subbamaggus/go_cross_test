package main

import (
    "fmt"
    "mymod/mytestlib"
    "mymod2"
)

func main() {
    fmt.Println("hello world")
    
    result := mytestlib.MySum(1,2)
    fmt.Println("MySum: ", result)
    
    result = mytestlib.MyProd(1,2)
    fmt.Println("MyProd: ", result)
    
    fmt.Println("lets see: ", myanotherone.MyString("eins", "zwei"))
    
    fmt.Scanln()
}    