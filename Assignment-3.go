package main

import (
    "fmt"
    "sync"
)
func squareNumber(num int, ch chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    sq := num * num
    ch <- sq
}
func aggregateSqnos(ch <-chan int, length int) {
    sum := 0
    for i := 0; i < length; i++ {
        sq := <-ch
        sum = sum + sq
        fmt.Printf("Received square: %d, Running sum: %d\n", sq, sum)
    }
    fmt.Printf("Final sum of squares: %d\n", sum)
}
func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
    ch := make(chan int)
    var wg sync.WaitGroup
    for _, num := range nums {
        wg.Add(1) //so this increases the count
        go squareNumber(num, ch, &wg)
    }
    go func() {
        wg.Wait()
        close(ch) 
    }()
    aggregateSqnos(ch, len(nums))
}
