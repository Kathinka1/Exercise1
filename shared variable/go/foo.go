// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
)

// Definerer hvilke kommandoer serveren kan motta
type Command string

const (
    Increment Command = "inc"  // Øker verdien til i
    Decrement Command = "dec"  //
    Get      Command = "get"  //Henter verdien til i
)
//Må gi de en string-verdi sånn at de ikke er like..... skjønner ikke helt. 

// Server-funksjonen som håndterer all tilgang til variabelen i
func numberServer(cmdChan <-chan Command, resultChan chan<- int) {
    i := 0  // Den delte variabelen håndteres bare av serveren

    for {
        select {
        case cmd := <-cmdChan: //cmdChan = Kanal som skal motta ulike kommandoer, som inc, dec osv..
                                //Når cmdChan mottar en verdi/kommando blir denne lagret i cmd. 
                                // Sjekker så under hvilken kommando dette er.  
            switch cmd {
            case Increment:
                i++
            case Decrement:
                i--
            case Get:   //Get skal hente verdien til i. 
                        //Sender verdien til i til kanalen "resultChan":
                resultChan <- i  // Sender verdien tilbake til main
                return           // Avslutter serveren etter å ha sendt verdien
            }
        }
    }
}

func main() {
    runtime.GOMAXPROCS(2)  // Tillater programmet å bruke to CPU-kjerner

    cmdChan := make(chan Command)  // Kanal for kommandoer til serveren
    resultChan := make(chan int)   // Kanal for å hente resultatet
    done := make(chan bool)        // Kanal for å signalere at en goroutine er ferdig

    // Starter serveren som håndterer i
    go numberServer(cmdChan, resultChan)

    // Goroutine for å inkrementere i
    go func() {
        for j := 0; j < 1000000; j++ {
            cmdChan <- Increment        // Istedet for å skrive "i++", sender vi kanalen "Increment" til cmdChan. 
                                        // Da vil numberServer "kjøre". Switchen vil altså fungere når cmdChan blir tilsendt en kanal/verdi. 
        }
        done <- true  // Signaliserer at inkrementering er ferdig
        // done er kanalen som passer på om en gorutine er ferdig eller ikke. Setter den til å være true. 
    }()

    // Goroutine for å dekrementere i
    go func() {
        for j := 0; j < 1000000; j++ {
            cmdChan <- Decrement
        }
        done <- true  // Signaliserer at dekrementering er ferdig
    }()

    // Venter på at begge goroutines er ferdige
    <-done
    <-done  //done blokkerer programmet helt til det mottar noe. Så når det mottar "true" går vi videre. 

    // Sender en "Get"-kommando for å hente sluttverdien av i
    cmdChan <- Get  // Kommandoen Get blir sendt, og resultChan får da verdien til i. 
    result := <-resultChan  //Lagrer i i result. 

    Println("The magic number is:", result)
}
