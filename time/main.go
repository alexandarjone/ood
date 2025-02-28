package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	duration := 1*time.Hour + 45*time.Minute + 30*time.Second
	end := start.Add(duration)
	diff := end.Sub(start)
	fmt.Println(int(diff/(15 * time.Minute)))
}