package main

import (
    "fmt"
    "math/rand/v2"
    "sync"
    "time"
)

var goroutinesCount = 10000
var maxGoroutineWorkTimeSeconds = 5

func work() int {
    workTimeSeconds := rand.IntN(maxGoroutineWorkTimeSeconds + 1)

    time.Sleep(time.Duration(workTimeSeconds) * time.Second)

    return workTimeSeconds
}

func main() {
    wg := &sync.WaitGroup{}
    mutex := &sync.Mutex{}

    totalWorkTime := 0
    mainStartWorkTimePoint := time.Now()

    wg.Add(goroutinesCount)
    for range goroutinesCount {
        go func() {
            defer wg.Done()

            workTime := work()

            mutex.Lock()
            totalWorkTime += workTime
            mutex.Unlock()
        }()
    }

    wg.Wait()

    mainWorkTime := time.Since(mainStartWorkTimePoint).Seconds()

    fmt.Printf("Main функция запустила %d горутин и дождалась их выполнения за %f сек\n", goroutinesCount, mainWorkTime)
    fmt.Printf("Общее время выполнения всех горутин составило %d сек\n", totalWorkTime)
}