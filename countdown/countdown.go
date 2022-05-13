package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(sleeper Sleeper, w io.Writer) {
	for i := countdownStart; i > 0; i-- {
		_, _ = fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	_, _ = fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(sleeper, os.Stdout)
}
