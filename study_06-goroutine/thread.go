/*
goroutine example

go routine can be used as the TASK of freeRTOS.

For channel: before read operation the write operation is blocked;
			 before write operation the read operation is also blocked.

This is similar as the lock for shared memory.
*/
package main

import (
	"fmt"
	"time"
)

func add(a int, b int, resChan chan int) {

	var cnt int = 0

	for{
		time.Sleep(time.Second * 5)
		resChan <- a + b
		cnt++
		a += cnt
	}


}

/*
func printResChan(ch chan int)  {

	for{

		fmt.Println("get int value from chan:", <-ch)
	}
	
}
*/

func main() {

	resCh := make([]chan int, 3) //create channel array

	/*Create 3 goroutine, each routine is assigned a separate chanel*/
	for i := 0; i < 3; i++ {
		resCh[i] = make(chan int) //Create a instance of the channel array
		go add(5, 6, resCh[i])
	}

	/*After all routines are started, the try to get the value from all channels.
	And we use a dead loop to do this again and again.

	Note: If we do not use any block method here, the goroutine cannot get chance to run.
	Because when we create the go routine in main(), main() will end its execution immediately.*/
	for {

		for chanIdx, ch := range resCh {
			fmt.Println("get value", <-ch,"from chs:",chanIdx )
		}
	}


}
