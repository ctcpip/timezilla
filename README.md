# timezilla
a simple timer for the console/terminal

![Development Status](https://img.shields.io/badge/development%20status-pre--alpha-red.svg)
[![Code Climate](https://img.shields.io/codeclimate/github/ctcpip/netrisse.svg)](https://codeclimate.com/github/ctcpip/timezilla)
[![Issue Count](https://img.shields.io/codeclimate/issues/github/ctcpip/netrisse.svg)](https://codeclimate.com/github/ctcpip/timezilla)
[![License](https://img.shields.io/badge/license-GNU%20GPLv3-blue.svg)](./LICENSE)


## how it works

* the timer will count down from the specified time
* when the time is up, three things happen:
  1. an OS notification is sent (using libnotify or growl)
  1. the terminal screen will flash red on/off
  1. the terminal bell will ring
* use CTRL+C to exit

## usage instructions

__timezilla__ [_minutes_]  
_minutes_ specified in fractional minutes  
if no _minutes_ specified, timer will default to 25 minutes ([pomodoro](http://en.wikipedia.org/wiki/Pomodoro_Technique) standard)

~~~ sh
$ timezilla       # default 25 minute timer

$ timezilla .5    # 30 second timer
$ timezilla 1     # 1 minute timer
$ timezilla 2.5   # 2 minute, 30 second timer
$ timezilla 5     # 5 minute timer
$ timezilla 10.1  # 10 minute, six second timer
~~~

### inspiration

I started this project to provide a simple timer for use with the [pomodoro technique](http://en.wikipedia.org/wiki/Pomodoro_Technique), a time management method.

### stuff nobody cares about

this project uses...

* [termbox-go](http://github.com/nsf/termbox-go)
* [notificator](http://github.com/0xAX/notificator)
* [osext](http://github.com/kardianos/osext)
* [Go](http://golang.org)
* [semantic versioning](http://semver.org/)

### license

[GNU GPLv3](http://www.gnu.org/licenses/gpl-3.0.en.html)
