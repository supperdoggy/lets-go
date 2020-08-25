package main

import (
	"fmt"
	"time"
)

func main()  {
	time1 := time.Now().Add(time.Minute*1).Unix()
	sub := time1 - time.Now().Unix()

	fmt.Print(sub/60 < 5)
}
