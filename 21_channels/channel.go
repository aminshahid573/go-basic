package main

import (
	"fmt"
)

//channel is to send data from one goroutine to another

//sending
/*
func processNum(numChan chan int) {

	for num := range numChan {

		fmt.Println("Processing number", num)
		time.Sleep(time.Second)
	}
	// fmt.Println("Processing number", <-numChan)

}
*/
//receiving
/*
func sum(result chan int, num1, num2 int) {
	numResult := num1 + num2
	result <- numResult
}
*/

//goroutine syncronization : used like waitgroup
/*
func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("processing....")
}
*/

// buffered channel part
// you can just do like this either
// emailSender(emailChan chan string, done chan bool)
// but below case can be good bcoz it improve type safety
/*
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second)
	}
}
*/
func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Received data from chan2", chan2Val)
		}
	}

	//buffered channel
	/*
		emailChan := make(chan string, 100) //non blocking
		done := make(chan bool)             //blocking

		go emailSender(emailChan, done)
		for i := 0; i < 100; i++ {
			emailChan <- fmt.Sprintf("%d@gmail.com", i)
		}

		fmt.Println("done Sending")

		//closing channel is important in buffered channel
		close(emailChan)
		<-done

	*/
	// emailChan <- "1@gmail.com"
	// emailChan <- "2@gmail.com"

	//this dont fall in deadlock coz the emailChan has 100 size and can handle upto 100 but if we exceed the size then we fall in the deadlock
	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	//can be used like waitgroup
	/*
		done := make(chan bool)
		go task(done)
		<-done //proram is block unless and until task func send dara to done channel
	*/
	// receiving
	/*
		result := make(chan int)
		go sum(result, 4, 5)

		res := <-result
		fmt.Println(res)
	*/
	// sending data from main goroutine to processNow
	/*
		numChan := make(chan int)

		go processNum(numChan)

		numChan <- 5

		for {
			numChan <- rand.Intn(100)
		}
	*/

	//simulating deadlock
	/*
		messageChan := make(chan string)

		 messageChan <- "ping" //channel is blocking

		msg := <-messageChan

		fmt.Println(msg)
	*/

}
