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
