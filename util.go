package main

import (
	"fmt"
	"strings"
	"time"
)

func getTimeString(d time.Duration) string {

	var strTime string

	if d < time.Second*1 {
		strTime = "0"
	} else {

		s := d.String()
		a := strings.Split(s, ".")

		if len(a) > 1 {
			strTime = a[0] + "s" // discard fractional seconds
		} else {
			strTime = s
		}

	}

	return strTime

}

func printHelp(arg string) {

	if arg != "-h" {
		fmt.Print("\ntimezilla: invalid argument -- '" + arg + "'\n")
	}

	fmt.Print(strHelp)

}
