package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ctcpip/notifize"
	"github.com/kardianos/osext"
	"github.com/nsf/termbox-go"
)

type app struct{}

func (a *app) init() {

	var k keyboard

	booContinue := true

	if len(os.Args) > 1 {

		if a, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
			duration = time.Millisecond * time.Duration(a*60000)
		} else {
			printHelp(os.Args[1])
			booContinue = false
		}

	}

	if booContinue {

		if duration == 0 {
			duration = time.Minute * 25
		}

		duration += time.Second * 1

		timer = time.NewTimer(duration)
		abort := make(chan bool, 1)

		scr.init()

		go countdown(duration, abort)

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

	var toggle bool
	t := time.NewTicker(time.Millisecond * 500)
	endTime := time.Now().Add(d)

	drawTime(endTime)

	for {

		select {

		case <-t.C:

			if !timerPaused {
				clearTime()
				drawTime(endTime)
			} else {

				if toggle {
					clearPauseDisplay()
				} else {
					scr.drawColoredText("PAUSED", 3, 5, termbox.ColorWhite, termbox.ColorRed)
					termbox.Flush()
				}

				toggle = !toggle

			}

		case <-pause:

			timerPaused = !timerPaused

			if timerPaused {
				timer.Stop()
				d = endTime.Sub(time.Now())
			} else {
				timer.Reset(d)
				endTime = time.Now().Add(d)
				clearPauseDisplay()
				toggle = false
			}

		case <-abort:
			return
		}

	}

}

func clearPauseDisplay() {
	scr.drawText("      ", 3, 5)
	termbox.Flush()
}

func clearTime() {
	for x := 0; x < 80; x++ {
		termbox.SetCell(x, 3, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}
}

func drawTime(endTime time.Time) {
	remainingTime := endTime.Sub(time.Now())
	scr.drawText(getTimeString(remainingTime), 3, 3)
	termbox.Flush()
}

func alertNotification() {

	appPath, err := osext.ExecutableFolder()
	if err != nil {
		panic(err)
	}

	notifize.Display("timezilla", "time is up!", true, appPath+"/clock.png")

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

func (a *app) close() { scr.close() }
