// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type subscriber struct {
	id int
	ch chan string
}

type publisher struct {
	subscribers []subscriber
	lock        sync.RWMutex
}

func subscribe_1(sub subscriber) {
	for i := range sub.ch {
		fmt.Println("subscriber #1", i)
	}
}

func subscribe_2(sub subscriber) {
	for i := range sub.ch {
		fmt.Println("subscriber #2", i)
	}
}

func (p *publisher) publish(msg string) {
	p.lock.Lock()
	defer p.lock.Unlock()

	for _, sub := range p.subscribers {
		sub.ch <- msg
	}
}

func (p *publisher) subscribe(id int) subscriber {
	sub := subscriber{id: id, ch: make(chan string)}
	p.lock.Lock()
	p.subscribers = append(p.subscribers, sub)
	p.lock.Unlock()
	return sub
}

func main() {
	p := &publisher{}

	sub1 := p.subscribe(1)
	sub2 := p.subscribe(2)

	go subscribe_1(sub1)
	go subscribe_2(sub2)

	for i := 1; i < 3; i++ {
		p.publish(strconv.Itoa(i))
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Hello, 世界")
}
