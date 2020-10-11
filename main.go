package main

import (
	"log"
	"sync"
)

type task struct {
	A      int64
	B      int64
	Result int64
}

var wg sync.WaitGroup

func queue(c chan task, a, b int64) {
	t := task{a, b, 0}
	c <- t
}

func worker(c chan task) {
	for {
		t := <-c
		t.Result = t.A + t.B
		log.Println(t)
		wg.Done()
	}
}

func main() {
	log.Println("Program Started")
	type abStruct struct {
		A int64
		B int64
	}
	ab := []abStruct{}

	for i := int64(0); i < 1000; i++ {
		ab = append(ab, abStruct{i, i + 1})
	}

	in := make(chan task)
	go worker(in)

	for _, v := range ab {
		wg.Add(1)
		queue(in, v.A, v.B)
	}

	log.Println("waiting for all tasks to complete")
	wg.Wait()
	log.Println("Program Ended")

}
