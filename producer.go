// cerner_2^5_2020

import "fmt"
type Consumer struct {
	msg *chan int
}
type Producer struct {
	msg  *chan int
	done *chan bool
}
func NewConsumer(msg *chan int) *Consumer {
	fmt.Println("NewConsumer called")
	return &Consumer{msg: msg}
}
func (c *Consumer) consume() {
	fmt.Println("consume: Started")
	for {
		msg := <-*c.msg
		fmt.Println("consume: Received:", msg)
	}
}
func NewProducer(msg *chan int, done *chan bool) *Producer {
	return &Producer{msg: msg, done: done}
}
func (p *Producer) produce(max int) {
	fmt.Println("produce: Started")
	for i := 0; i < max; i++ {
		fmt.Println("produce: Sending ", i)
		*p.msg <- i
	}
	*p.done <- true
	fmt.Println("produce: Done")
}
