package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	HEIGHT int = 14
)

func main() {
	fmt.Print("\033[?25l") // hide the cursor
	for {
		print_time()

		time.Sleep(time.Second)
	}
}

func time_string() []string {
	t := time.Now()
	h := t.Hour()
	m := t.Minute()
	s := t.Second()
	out := []string{}
	vals := []int{h / 10, h % 10, 10, m / 10, m % 10, 10, s / 10, s % 10}

	for _, i := range vals {
		out = append(out, digits[i])
	}
	return out
}

func morph(strs ...string) string {
	// assuming all strings will have equal height
	if len(strs) == 0 {
		return ""
	}

	out := strings.Split(strs[0], "\n")

	for _, str := range strs[1:] {
		slc := strings.Split(str, "\n")
		for i, s := range slc {
			out[i] += "  " + s
		}
	}

	return strings.Join(out, "\n")
}

func print_time() {
	fmt.Println(morph(time_string()...))

	fmt.Printf("\033[%dA", HEIGHT+2)
}
