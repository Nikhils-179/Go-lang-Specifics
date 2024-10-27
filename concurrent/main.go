package main

import (
	"fmt"
	"time"

)

func main() {
	withoutConcurrency()
	start := time.Now()
	withConcurrency()
	wg.Wait()
	fmt.Println(time.Since(start))
}
