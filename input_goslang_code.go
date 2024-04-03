package main

import (
    "fmt"
)

func GetFoo() {
    hello := make(chan int, 10)
    test := foo()
    fmt.Println(test)
}

func foo() int {
    return 0
}