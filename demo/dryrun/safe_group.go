package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/elin-croft/go-utils/framework/safe_group"
)

func main() {
	wg := safe_group.NewGroup("test")
	t1 := time.Now().UnixMilli()
	wg.Go(func() error {
		n := time.Duration(rand.Intn(30))
		fmt.Printf("first task sleep %d seconds\n", n)
		time.Sleep(n * time.Second)
		fmt.Printf("sleep %d seconds done\n", n)
		return nil
	})

	wg.Go(func() error {
		n := time.Duration(rand.Intn(30))
		fmt.Printf("second task sleep %d seconds\n", n)
		time.Sleep(n * time.Second)
		fmt.Printf("sleep %d seconds done\n", n)
		return nil
	})
	wg.Wait()
	fmt.Printf("total time: %d milliseconds\n", (time.Now().UnixMilli() - t1))
}
