package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	duration := flag.Int64("duration", 0, "duration `value` in second")
	title := flag.String("title", "no title", "title of message")

	flag.Parse()

	time.Sleep(time.Duration(*duration) * time.Second)
	err := beeep.Notify(*title, msg(flag.Args()), "")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var defaultmsg = func() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("hello from %s", user.Name)
}

func msg(s []string) string {
	if len(s) == 0 {
		return defaultmsg()
	}
	return strings.Join(s, "\n")
}
