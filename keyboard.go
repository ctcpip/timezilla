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
      case e.Ch == 'p', e.Ch == 'P':
        pause <- true
      }

    case termbox.EventError:
      panic(e.Err)
    }

  }

}
