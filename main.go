package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type task struct {
	A      int64
	B      int64
	Result int64
	Worker int
}

var wg sync.WaitGroup

func queue(c chan task, a, b int64) {
	t := task{a, b, 0, -1}
	c <- t
}

func worker(id int, in, out chan task) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for {
		t := <-in
		t.Result = t.A + t.B
		t.Worker = id
		i := r1.Intn(4000)

		//log.Println("sleeping", i)
		time.Sleep(time.Duration(i) * time.Millisecond)
		log.Println(fmt.Sprintf("worker #%d: %v %dms", id, t, i))
		out <- t
	}
}

func statticker(c chan task) {
	i := 0
	startTime := float64(time.Now().UnixNano() / 1000000000)
	workerticker := map[int]int{}
	for {
		t := <-c
		i++
		timeElapsed := float64(time.Now().UnixNano()/1000000000) - startTime
		rate := float64(float64(i) / timeElapsed)
		if val, ok := workerticker[t.Worker]; ok {
			val++
			workerticker[t.Worker] = val
		} else {
			workerticker[t.Worker] = 1
		}
		log.Println("fan in:", i, timeElapsed, rate, t.Result, workerticker)
		wg.Done()
	}
}

func main() {
	log.Println("Program Started")
	timeStart := time.Now().Unix()
	type abStruct struct {
		A int64
		B int64
	}
	ab := []abStruct{}

	for i := int64(0); i < 1000; i++ {
		ab = append(ab, abStruct{i, i + 1})
	}

	in := make(chan task)
	out := make(chan task)

	go statticker(out)

	numworkers := 10
	for i := 0; i < numworkers; i++ {
		go worker(i, in, out)
	}

	for _, v := range ab {
		wg.Add(1)
		queue(in, v.A, v.B)
	}

	log.Println("waiting for all tasks to complete")
	wg.Wait()
	timeEnd := time.Now().Unix()
	log.Println(fmt.Sprintf("time elapsed: %ds", timeEnd-timeStart))
	log.Println("Program Ended")

}
