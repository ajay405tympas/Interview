package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// scheduleDeliveries distributes the orders among couriers and returns the assignment.
func scheduleDeliveries(orderId <-chan int, capacity int, couriers []chan<- int) {
	// TODO: Implement the function logic

	for i := range couriers {
		couriers[i] = make(chan int)
		for j := 0; j < capacity; j++ {
			couriers[i] <- <-orderId
		}
	}
}

type Delivery struct {
	courierId int
	orders    []int
}

func courier(in <-chan int, courierId int, capacity int, delivery chan<- Delivery) {
	// TODO: Implement the function logic
	fmt.Println("in", <-in)

	for i := 0; i < capacity; i++ {
		delivery <- Delivery{courierId, []int{<-in}}
	}
}

func main() {
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		panic(err)
	}
	defer stdout.Close()
	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	defer writer.Flush()

	reader := bufio.NewReader(os.Stdin)

	nStr, _ := reader.ReadString('\n')
	nCouriers, _ := strconv.Atoi(strings.TrimSpace(nStr))
	nStr, _ = reader.ReadString('\n')
	capacity, _ := strconv.Atoi(strings.TrimSpace(nStr))

	ordersStr, _ := reader.ReadString('\n')
	ordersStr = strings.TrimSpace(ordersStr)
	orderIDs := strings.Split(ordersStr, " ")

	orders := make([]int, 0)
	for _, idStr := range orderIDs {
		if len(idStr) > 0 {
			x, _ := strconv.Atoi(idStr)
			orders = append(orders, x)
		}
	}
	orderId := make(chan int)
	couriers := make([]chan<- int, 0)
	output := make(chan Delivery)
	wg := sync.WaitGroup{}
	for i := 0; i < nCouriers; i++ {
		in := make(chan int)
		couriers = append(couriers, in)
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			courier(in, id, capacity, output)
		}(i)
	}
	go func() {
		wg.Wait()
		close(output)
	}()

	go scheduleDeliveries(orderId, capacity, couriers)
	result := make(map[int][]int)
	done := make(chan struct{})
	go func() {
		defer close(done)

		for out := range output {
			if len(out.orders) > capacity {
				fmt.Fprintf(writer, "Courier capacity exceeded\n")
			}
			result[out.courierId] = append(result[out.courierId], out.orders...)
		}
	}()

	for _, i := range orders {
		orderId <- i
	}
	close(orderId)
	<-done

	for i := 0; i < nCouriers; i++ {
		r := result[i]
		for j, x := range r {
			if j > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, x)
		}
		fmt.Fprint(writer, "\n")
	}
}
