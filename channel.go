// cerner_2^5_2020

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func Call1(ch chan string) {
	fmt.Println("In Call1")
	defer wg.Done()
	total := 0
	a := []int{10, 20, 30, 40}
	for _, v := range a {
		total += v
		fmt.Println(total)
	}
	ch <- "-- Call1 Function finished --"
}
func call2(ch chan string) {
	fmt.Println("In call2")

	//	time.Sleep(time.Millisecond * 20)
	defer wg.Done()
	ch <- "-- call2 Function finished --"
}
func call3(ch chan string) {
	fmt.Println("In call3")

	//	time.Sleep(time.Millisecond * 30)
	defer wg.Done()
	ch <- "-- call3 Function finished --"
}
