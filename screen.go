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

type screen struct{}

func (s *screen) init() {

	err := termbox.Init()

	if err != nil {
		panic(err)
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	s.writeText("timezilla 0.1.0 (C) 2016  Chris de Almeida     \"timezilla -h\" for more info", 0, 0)
	termbox.Flush()

}

func (s *screen) writeText(text string, startX, y int) {

	currX := startX

	for i := 0; i < len(text); i++ {
		termbox.SetCell(currX, y, rune(text[i]), termbox.ColorDefault, termbox.ColorDefault)
		currX++
	}

}

func (s *screen) close() { defer termbox.Close() }
