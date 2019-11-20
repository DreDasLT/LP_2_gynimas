package main

import "sync"

var pirmasSiuntejas = make(chan float64)

var vidurkiss = make(chan float64)
var antrasSpausdintojas = make(chan int)

var wg = sync.WaitGroup{}
var pabaigaGavejo = make(chan bool)
var pabaigarezultat = make(chan bool)
var pabaigamain = make(chan bool)

var isviso = 4
func main() {



	wg.Add(5)
	go pirmasSiuntimas()
	go pirmasSiuntimas()
	go pirmasSiuntimas()
	go pirmasSiuntimas()
	go spausdintojas2()
	var gauta = 0;
	var skaiciai = make([]float64, 40)
	var isviso = 0;
	var vidurkis = 0.0;
	for {
		select {
			case <-pabaigamain:
			print("pabaigaa main \n")
			return

			case x := <-pirmasSiuntejas:
			skaiciai[isviso] = x
			gauta++
			isviso++

			if gauta == 2 {
				gauta = 0
				for i := 0; i < 40; i++ {
					vidurkis = vidurkis + skaiciai[i]
				}
				vidurkis = vidurkis / 40
				vidurkiss <- vidurkis
			}


		}


	}
	wg.Wait()
	print("\n")
	print("viskas")
	
}

func spausdintojas2() {
	defer wg.Done()
	for {
		select {
		case <-pabaigarezultat:
			print("pabaigaa \n")
			return
		case x := <-vidurkiss:
			print(x)
			print("\n")

		}

	}
}

func pirmasSiuntimas() {
	defer wg.Done()
	defer func() {
		isviso = isviso-1
		if isviso==0 {
			pabaigamain <-true
			pabaigarezultat <-true
		}
	}()
	var itemCount = 10
	skaiciai := make([]float64, itemCount)
	for i := 1; i <  itemCount+1; i++ {
		skaiciai[i-1] = float64(i)
	}

	for {
		if itemCount!=0 {
			pirmasSiuntejas <- skaiciai[10-itemCount]
			itemCount--
		} else{
			break
		}
	}
}






