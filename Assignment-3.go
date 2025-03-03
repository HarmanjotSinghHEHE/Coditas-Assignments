package main

import (
    "fmt"
    "sync"
)

//now first we will make the squarenumber func

func squareNumber(num int, ch chan<- int, wg *sync.WaitGroup) {
    defer wg.Done() //Decreases the count after a particular job is completed
    sq := num * num
    ch <- sq
}

// then we will make the aggregator , that will add all the values recieved from the channel
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
        wg.Wait() // Wait for all squareWorker goroutines to finish
        close(ch) // Close the channel to signal no more data
    }()

    aggregateSqnos(ch, len(nums))

}

//-------------------------------------------------------------Notes--------------------------------------------------------------------
/*sync.waitgroups have 3 methods
wg.Wait()
--Blocks execution until the WaitGroup counter reaches 0.
--Does nothing if the counter is already 0.
wg.Add()
--Increases the WaitGroup counter by n, where n is the number of tasks (goroutines) you want to wait for.
wg.Done()
--Decrements the WaitGroup counter by 1.
--Signals that one goroutine (task) has completed.
--Typically called with defer to ensure it runs even if the goroutine panics.
*/
