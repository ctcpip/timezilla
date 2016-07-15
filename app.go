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
	"time"

	"github.com/nsf/termbox-go"
)

type app struct{}

func (a *app) init() {

	var k keyboard
	var b bool
	var c termbox.Attribute
	var remainingTime time.Duration

	// TODO: parse command line args

	duration := time.Second * 10
	timer := time.NewTimer(duration)
	ticker := time.NewTicker(time.Second * 1)
	countdown := time.NewTicker(time.Second * 1)
	endTime := time.Now().Add(duration)

	s.init()

	go func() {

		for _ = range countdown.C {

			for x := 0; x < 80; x++ {
				termbox.SetCell(x, 3, ' ', termbox.ColorDefault, termbox.ColorDefault)
			}

			remainingTime = endTime.Sub(time.Now())
			s.writeText(remainingTime.String(), 3, 3)

			termbox.Flush()

		}

	}()

	go func() {

		<-timer.C

		countdown.Stop()

		fmt.Print("\a")

		for _ = range ticker.C {

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

	}()

	k.init()

}

func (a *app) close() { s.close() }