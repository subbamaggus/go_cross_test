package main

import (
    "fmt"
    "mymod/mytestlib"
)

func main() {
    fmt.Println("hello world")
    result := mytestlib.MySum(1,2)
    fmt.Println("result: ", result)
    fmt.Scanln()
}    