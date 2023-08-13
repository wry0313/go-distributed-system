package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// func main() {
// 	c := make(chan int) // unbuffered no ability to hold any value
// 	go func() {
// 		<-c
// 	}()
// 	c <- 2 // send 2 to channel will block until this channel receives a value from another go routine. remember the purpose of channel is to communicate between go routines
// }

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "foo", "bar")
	userID := 10
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("results: ", val)
	fmt.Println("took: ", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	val := ctx.Value("foo")
	fmt.Println(val)
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	respch := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		respch <- Response{value: val, err: err}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party time out")
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 150)
	return 666, nil
}

// method call can be inconsistent
// we can control this process by using context
