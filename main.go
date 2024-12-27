package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Poduct struct {
	Id   string
	Name string
}

func (p Poduct) Vender(priece float64) {
	fmt.Printf("Vendido %s por %f\n", p.Name, priece)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))

	product := Poduct{
		Id:   "1",
		Name: "Carro",
	}

	json.NewEncoder(w).Encode(product)
}

func contador(v int) {
	for i := range v {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func publish(ch chan int) {
	for i := range 10 {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int, workerID int) {
	for x := range ch {
		time.Sleep(time.Second)
		fmt.Printf("Worker ID:  %d - Recebido %d\n", workerID, x)
	}
}

func main() {
	// canal := make(chan int)

	// go publish(canal)

	// for i := range 3 {
	// 	go consumer(canal, i)
	// }

	// time.Sleep(time.Minute)
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
