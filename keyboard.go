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

import "github.com/nsf/termbox-go"

type keyboard struct{}

func (k *keyboard) init() { k.read() }

func (k *keyboard) read() {

loopyMcLoopface:
	for {

		switch e := termbox.PollEvent(); e.Type {

		case termbox.EventKey:

			switch {
			case e.Key == termbox.KeyCtrlC:
				break loopyMcLoopface
			}

		case termbox.EventError:
			panic(e.Err)
		}

	}

}
