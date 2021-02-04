package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

//Racer used to return the url which is more fast
func Racer(url1, url2 string) (winner string, err error) {
	return ConfigurableRacer(url1, url2, tenSecondTimeout)
}

//ConfigurableRacer return the faster url with select
func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
