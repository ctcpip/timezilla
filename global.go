package main

import "time"

var scr screen
var duration time.Duration
var timer *time.Timer
var timerPaused bool
var pause = make(chan bool, 1)
