package main

import (
    "fmt"
    "sync"
)

var mu1, mu2 sync.Mutex

func function1() {
    mu1.Lock()
    fmt.Println("Locked mu1 in function1")
    
    // Simulating some work
    // Deadlock will happen because function2 already locked mu2
    mu2.Lock()
    fmt.Println("Locked mu2 in function1")
    
    mu2.Unlock()
    mu1.Unlock()
}

func function2() {
    mu2.Lock()
    fmt.Println("Locked mu2 in function2")
    
    // Simulating some work
    // Deadlock will happen because function1 already locked mu1
    mu1.Lock()
    fmt.Println("Locked mu1 in function2")
    
    mu1.Unlock()
    mu2.Unlock()
}

func main() {
    go function1()
    go function2()

    // Wait to simulate the deadlock
    fmt.Scanln()
}
