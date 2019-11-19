package main

import "sync"

var pirmasSiuntejas = make(chan int)

var pirmasSpausdintojas = make(chan int)
var antrasSpausdintojas = make(chan int)

var wg = sync.WaitGroup{}
var pabaigaGavejo = make(chan bool)
var pabaigaDarbo = make(chan bool)

func main() {

	var siunciamas = 0
	wg.Add(2)
	go pirmasSiuntimas(siunciamas)
	var siunciamas2 = 11
	go pirmasSiuntimas(siunciamas2)
	wg.Add(2)
	go spausdintojas1()
	go spausdintojas2()
	var gauta = 0
	for {
		x := <-pirmasSiuntejas
		gauta++
		if x%2 == 0 {
			pirmasSpausdintojas <- x

		} else {
			antrasSpausdintojas <- x
		}
		if gauta >= 20 {
			pabaigaGavejo <- false
			pabaigaGavejo <- false
			break
		}

	}
	wg.Wait()
	print("\n")
	print("viskas")

}
func spausdintojas1() {
	defer wg.Done()
	privateItems := make([]int, 20)
	var kiekis = 0
	for {
		select {
		case x := <-pirmasSpausdintojas:
			privateItems[kiekis] = x
			kiekis++
		case <-pabaigaDarbo:
			print("\n")
			print("spausdinu masyva pirmo \n")
			for i := 0; i < kiekis; i++ {
				print(privateItems[i])
				print("\n")
			}
			return

		}

	}
}
func spausdintojas2() {
	defer wg.Done()
	privateItems := make([]int, 20)
	var kiekis = 0
	for {
		select {
		case x := <-antrasSpausdintojas:
			privateItems[kiekis] = x
			kiekis++
		case <-pabaigaDarbo:
			print("\n")
			print("spausdinu masyva  antro \n")
			for i := 0; i < kiekis; i++ {
				print(privateItems[i])
				print("\n")
			}
			return

		}

	}
}

func pirmasSiuntimas(siuncimas int) {
	defer wg.Done()
	for {
		select {
		case x := <-pabaigaGavejo:
			if x == false {
				pabaigaDarbo <- true
				return
			}
		case pirmasSiuntejas <- siuncimas:
			siuncimas++
		}
	}
}
