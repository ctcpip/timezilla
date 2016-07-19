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
	"time"

	"github.com/0xAX/notificator"
	"github.com/kardianos/osext"
	"github.com/nsf/termbox-go"
)

type app struct{}

func (a *app) init() {

	var k keyboard
	var d time.Duration

	booContinue := true

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
			go alertBell()
			go alertVisual()
			alertNotification()
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

func alertNotification() {

	var notify *notificator.Notificator

	appPath, err := osext.ExecutableFolder()
	if err != nil {
		panic(err)
	}

	// TODO - create a new notification package for use here instead
	notify = notificator.New(notificator.Options{})
	notify.Push("timezilla", "time is up!", appPath+"/clock.png", notificator.UR_CRITICAL)

}

func alertBell() {

	// ring the terminal bell

	fmt.Print("\a")

	t := time.NewTicker(time.Second * 30)

	for _ = range t.C {
		fmt.Print("\a")
	}

}

func alertVisual() {

	var b bool
	var c termbox.Attribute

	t := time.NewTicker(time.Second * 1)

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

func (a *app) close() { s.close() }
