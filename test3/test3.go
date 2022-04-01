package main

func rand5() int8 {
	var ch = make(chan int8, 1)
	select {
	case ch <- 1:
	case ch <- 2:
	case ch <- 3:
	case ch <- 4:
	case ch <- 5:
	}
	return <-ch
}

func rand5ToRand13() int8 {
	result := (rand5()-1)*5 + rand5()
	if result <= 13 {
		return result
	}
	return rand5ToRand13()
}

func rand13ToRand5() int8 {
	result := rand5ToRand13()
	if result <= 10 {
		return result%5 + 1
	}
	return rand13ToRand5()
}
