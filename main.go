package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {

	tickIntervalStr := os.Getenv("TICK_INTERVAL")
	tickInterval, err := strconv.Atoi(tickIntervalStr)
	if err != nil {
		fmt.Println("interval not set, using 2")
		tickInterval = 2
	}
	targetHost := os.Getenv("TARGET_HOST")
	timer := time.NewTimer(time.Duration(tickInterval) * time.Second)
	for {
		_ = <-timer.C
		_, err := net.LookupHost(targetHost)
		if err != nil {
			fmt.Printf("error looking up TARGET_HOST: %v \n %v\n", targetHost, err)
		} else {
			fmt.Printf("Successful net.LookupHost : %v\n", targetHost)
		}
		timer = time.NewTimer(time.Duration(tickInterval) * time.Second)
	}
}
