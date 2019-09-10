package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

var ntp_server string = "0.ru.pool.ntp.org"

func main() {
	time, err := ntp.Time(ntp_server)

	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

        fmt.Printf("%d:%d:%d\n", time.Hour(), time.Minute(), time.Second())
}
