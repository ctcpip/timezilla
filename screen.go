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
	s.drawText(strHeader, 0, 0)
	drawHelp()
	termbox.Flush()

}

func drawHelp() {

	scr.drawColoredText(" ^C ", 0, 23, termbox.ColorBlack, termbox.ColorWhite)
	scr.drawText("Exit", 5, 23)
	scr.drawColoredText(" P ", 13, 23, termbox.ColorBlack, termbox.ColorWhite)
	scr.drawText("Pause Timer", 17, 23)

}

func (s *screen) drawColoredText(text string, startX, y int, fg, bg termbox.Attribute) {

	currX := startX

	for i := 0; i < len(text); i++ {
		termbox.SetCell(currX, y, rune(text[i]), fg, bg)
		currX++
	}

}

func (s *screen) drawText(text string, startX, y int) {
	scr.drawColoredText(text, startX, y, termbox.ColorDefault, termbox.ColorDefault)
}

func (s *screen) close() {

	if termbox.IsInit {
		defer termbox.Close()
	}

}
