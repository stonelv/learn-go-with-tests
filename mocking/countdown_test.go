package main

import (
	"time"
	"reflect"
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})

	t.Run("prints 3 to go", func(t *testing.T){
		buffer := &bytes.Buffer{}
		CountDown(buffer, &CountdownOperationsSpy{})
	
		got := buffer.String()
		want := `3
2
1
go!`
	
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}