package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println //p now points to fmt.Println

	t := time.Now() // Get current local time (nanos precision)
	p("1.", t)      //println uses default String() method of time, shows monotonic reading

	p("3.", t.Unix())             //seconds since Jan 1, 1970 (aka Unix Time)
	p("4.", t.UnixNano())         //nanosecs since Jan 1, 1970
	p("5.", t.UnixNano()/1000000) //millis since Jan 1, 1970 (java System.currentTimeMillis)

}
