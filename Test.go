package main

import (
    "fmt"
    "time"
)


func one(y chan <- int){
    for i := 0; i < 20; i++ {
        time.Sleep(time.Second)
        y <- i
    }
}

func hundred(y chan<- int){
    for i := 0; i < 10; i++ {
        time.Sleep(time.Second * 2)
        y <- i * 100
    }
}

func main(){
    
    numbers := make(chan int)

    go one(numbers)
    go hundred(numbers)
    /*
    for i := 0; i < 35; i++ {
        select{
            case result := <- numbers:
                fmt.Println(result)
            default:
                fmt.Println("waiting")
        }
        time.Sleep(time.Second)
    } */
    
    for num := range numbers {
        fmt.Println(num)
    }
}