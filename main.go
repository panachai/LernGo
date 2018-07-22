package main

import linenotify "github.com/utahta/go-linenotify"

func main() {
	token := "Pd8OH3yI8quQyZyc4Xpe5wu3Kr1UOlz3MM4t32HwEeK"

	c := linenotify.New()
	c.Notify(token, "hello world", "", "", nil)
}
