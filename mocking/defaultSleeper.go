package main

import (
	"time"
)

//DefaultSleeper the default implementation of Sleeper
type DefaultSleeper struct {}

//Sleep do default things
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}