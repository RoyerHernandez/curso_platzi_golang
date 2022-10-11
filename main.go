package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/text"
	//"time"
)

func say(text string, wg * sync.WaitGroup){
	fmt.Println(text)

	defer	wg.Done()

}

func main() {
    var wg sync.WaitGroup
	
	fmt.Println("Hello")
	wg.Add(1)

	go say("World", &wg)

	wg.Wait()

	go func(texto string){
		fmt.Println(texto)
	}("Adios")

	time.Sleep(time.Second * 1)
}
