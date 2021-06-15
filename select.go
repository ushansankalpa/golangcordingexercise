package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	defaultcase()
	fmt.Println()
	reacecondition()

	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

func server1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}

//Default case
func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func defaultcase() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}
}

//Program with a race condition
var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func reacecondition() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

//Solving the race condition using a mutex
var y = 0

func increment2(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	y = y + 1
	m.Unlock()
	wg.Done()
}

func reacecondition2() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment2(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", y)
}
