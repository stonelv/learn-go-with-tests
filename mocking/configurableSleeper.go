package main

import (
	"time"
)

//ConfigurableSleeper used to make configurable sleeper
type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

//Sleep used to Sleep based on the config
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

//SpyTime used to mock Time
type SpyTime struct {
	durationSlept time.Duration
}

//Sleep implements Sleep method
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}