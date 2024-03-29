package main

import (
	"fmt"
	"reflect"
)

func produce(ch chan<- string, i int) {
	for j := 0; j < 5; j++ {
		ch <- fmt.Sprint(i*10 + j)
	}
	close(ch)
}

func main() {
	numChans := 4
	chans := []chan string{}
	for i := 0; i < numChans; i++ {
		ch := make(chan string)
		chans = append(chans, ch)
		go produce(ch, i+1)
	}

	cases := make([]reflect.SelectCase, len(chans))
	for i, ch := range chans {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}
	remaining := len(cases)
	for remaining > 0 {
		chosen, value, ok := reflect.Select(cases)
		if !ok {
			cases[chosen].Chan = reflect.ValueOf(nil)
			remaining--
			continue
		}
		fmt.Printf("Read from channel %#v and received %s\n", chans[chosen], value.String())
	}
}
