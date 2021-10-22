package main

import (
	"fmt"
	"pools/pool"
	"time"
)

func main(){
	wp := pool.New(10)
	wp.SetTimeOut(1* time.Second)
	for i := 0; i < 20; i++ { // Open 20 requests
		ii := i
		wp.Do(func() error {
			for j := 0; j < 10; j++ {
				fmt.Println(fmt.Sprintf("%v->\t%v", ii, j))
				time.Sleep(1 * time.Second)
			}
			return nil
		})
	}

	err := wp.Wait()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("down")
}
