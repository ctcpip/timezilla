/*

timezilla - a simple timer for the console/terminal
Copyright (C) 2016  Chris de Almeida

http://github.com/ctcpip/timezilla

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/0xAX/notificator"
	"github.com/kardianos/osext"
	"github.com/nsf/termbox-go"
)

type app struct{}

func (a *app) init() {

	var k keyboard

	booContinue := true
	var d time.Duration

	if len(os.Args) > 1 {

		if a, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
			d = time.Millisecond * time.Duration(a*60000)
		} else {
			printHelp(os.Args[1])
			booContinue = false
		}

	}

	if booContinue {

		if d == 0 {
			d = time.Minute * 25
		}

		d += time.Second * 1

		timer := time.NewTimer(d)
		abort := make(chan bool, 1)

		s.init()

		go countdown(d, abort)

		go func() {
			<-timer.C
			abort <- false
			alert()
		}()

		k.init()

	}

}

func countdown(d time.Duration, abort chan bool) {

	var remainingTime time.Duration

	t := time.NewTicker(time.Second * 1)
	endTime := time.Now().Add(d)

	remainingTime = endTime.Sub(time.Now())
	s.writeText(getTimeString(remainingTime), 3, 3)
	termbox.Flush()

	for {

		select {
		case <-t.C:

			for x := 0; x < 80; x++ {
				termbox.SetCell(x, 3, ' ', termbox.ColorDefault, termbox.ColorDefault)
			}

			remainingTime = endTime.Sub(time.Now())

			s.writeText(getTimeString(remainingTime), 3, 3)

			termbox.Flush()

		case <-abort:
			return
		}

	}

}

func getTimeString(d time.Duration) string {

	if d < time.Second*1 {
		return "0"
	}

	return strings.Split(d.String(), ".")[0] + "s"

}

func alert() {

	var b bool
	var c termbox.Attribute
	var notify *notificator.Notificator

	t := time.NewTicker(time.Second * 1)
	appPath, err := osext.ExecutableFolder()
	if err != nil {
		panic(err)
	}

	fmt.Print("\a") // ring the terminal bell

	notify = notificator.New(notificator.Options{})
	notify.Push("timezilla", "time is up!", appPath+"/clock.png", notificator.UR_CRITICAL)

	for _ = range t.C {

		if b {
			c = termbox.ColorBlack
		} else {
			c = termbox.ColorRed
		}

		for y := 2; y < 24; y++ {

			for x := 0; x < 80; x++ {
				termbox.SetCell(x, y, ' ', c, c)
			}

		}

		termbox.Flush()

		b = !b

	}

}

func printHelp(arg string) {

	if arg != "-h" {
		fmt.Print("\ntimezilla: invalid argument -- '" + arg + "'\n")
	}

	fmt.Print(strHelp)

}

func (a *app) close() { s.close() }
