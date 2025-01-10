
package main

import "fmt"
import "time"


func producer(buffer chan int){

    for i := 0; i < 10; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("[producer]: pushing %d\n", i)
        // TODO: push real value to buffer

        buffer <- i // bufferen (som er en kanal) blir til sendt verdien i
    }
    close(buffer) //Må lukke bufferen

}

func consumer(buffer chan int){

    time.Sleep(1 * time.Second)
    for i := range buffer {  //Vil nå lese fra bufferet helt til det er tomt
         //TODO: get real value from buffer
        fmt.Printf("[consumer]: %d\n", i)
        time.Sleep(50 * time.Millisecond)
    }
    
}


func main(){
    
    // TODO: make a bounded buffer
    buffer := make(chan int, 5) //Hvordan man lager en buffer med kapasitet på 5
    
    go consumer(buffer)
    go producer(buffer)
    
    select {}
}