package web_racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, e error) {
	return ConfigurableRacer(a, b, 10*time.Second)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, e error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, _ = http.Get(url)
		close(ch)
	}()

	return ch
}
