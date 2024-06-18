// https://www.youtube.com/watch?v=LGVRPFZr548&list=PL0xRBLFXXsP7-0IVCmoo2FEWBrQzfH2l8&index=2

package foohandler

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	respch := make(chan any, 2)
	userName := fetchUser()
	likes := 0
	match := ""

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go fetchUserLikes(userName, respch, wg)
	go fetchUserMatch(userName, respch, wg)

	wg.Wait()
	close(respch)

	for resp := range respch {
		i, ok := resp.(int)
		if ok {
			likes = i
		}

		s, ok := resp.(string)
		if ok {
			match = s
		}

		fmt.Println("resp:", resp)
	}

	fmt.Println("UserName: ", userName)
	fmt.Println("Likes: ", likes)
	fmt.Println("Match: ", match)
	fmt.Println("Elapsed: ", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "Bob"
}

func fetchUserLikes(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)
	respch <- 11
	wg.Done()
}

func fetchUserMatch(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respch <- "Anna"
	wg.Done()
}
