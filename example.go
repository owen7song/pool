package example

import (
	"fmt"
	"pools/pool"
	"time"
)

//Poll 简单示列
func Poll()  {
	wp := pool.New(10)
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

//TimeoutPoll 可设置超时工作池
func TimeoutPoll()  {
	wp := pool.New(10)
	wp.SetTimeOut(time.Second * 2)
	for i := 0; i < 20; i++ { // Open 20 requests
		wp.Do(func() error {
			fmt.Println("wait")
			return nil
		})
	}

	err := wp.Wait()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("down")
}