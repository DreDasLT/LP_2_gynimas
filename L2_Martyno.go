package main

import (
	"fmt"
)

func sender(send chan<- int, done <-chan bool, startAt int) {
	Loop:
	for {
		select {
		case send <- startAt:
			startAt++
		case <- done:
			break Loop
		}
	}
}

func receiver(receive <-chan int, sendOdd chan<- int, sendEven chan<- int, done chan<- bool) {
	sent := 0
	Loop:
	for {
		received := <- receive
		if received % 2 == 0 {
			sendEven <- received
		} else {
			sendOdd <- received
		}
		sent++
		if sent == 20 {
			done <- true
			done <- true
			done <- true
			done <- true
			break Loop
		}
	}
}

func printer(receive <-chan int, done <-chan bool, donePrinting chan<- bool)  {
	var array []int
	Loop:
	for {
		select {
		case received := <- receive:
			array = append(array, received)
		case <- done:
			fmt.Println(array)
			donePrinting <- true
			break Loop
		}
	}
}

func main() {
	sendToReceive := make(chan int)
	receiveToOdd := make(chan int)
	receiveToEven := make(chan int)
	done := make(chan bool)
	donePrintingOdd := make(chan bool)
	donePrintingEven := make(chan bool)
	go receiver(sendToReceive, receiveToOdd, receiveToEven, done)
	go printer(receiveToOdd, done, donePrintingOdd)
	go printer(receiveToEven, done, donePrintingEven)
	go sender(sendToReceive, done, 0)
	go sender(sendToReceive, done, 11)
	<- donePrintingOdd
	<- donePrintingEven
}
