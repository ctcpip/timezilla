# timezilla

a simple timer for the console/terminal

![Build Status](https://img.shields.io/travis/ctcpip/timezilla.svg)
![Development Status](https://img.shields.io/badge/development%20status-pre--alpha-red.svg)
[![Code Climate](https://img.shields.io/codeclimate/github/ctcpip/timezilla.svg)](https://codeclimate.com/github/ctcpip/timezilla)
[![Issue Count](https://img.shields.io/codeclimate/issues/github/ctcpip/timezilla.svg)](https://codeclimate.com/github/ctcpip/timezilla/issues)
[![License](https://img.shields.io/badge/license-GNU%20GPLv3-blue.svg)](./LICENSE)

## how it works

* the timer will count down from the specified time
* when the time is up, three things happen:
    1. the terminal screen will flash red on/off
    1. the terminal bell will ring _(every 30 seconds)_
    1. a desktop notification is sent _(only for the following operating systems*)_
        * __GNU/Linux__ using _libnotify_ `notify-send`
        * __OS X 10.9+__ using _AppleScript_ `display notification`
* use CTRL+C to exit

\*see [notifize](http://github.com/ctcpip/notifize) package for more information on desktop notification support for other operating systems

## usage instructions

__timezilla__ [_minutes_]  
_minutes_ specified in fractional minutes  
if no _minutes_ specified, timer will default to 25 minutes ([pomodoro](http://en.wikipedia.org/wiki/Pomodoro_Technique) standard)

~~~ sh
timezilla       # default 25 minute timer

timezilla .5    # 30 second timer
timezilla 1     # 1 minute timer
timezilla 2.5   # 2 minute, 30 second timer
timezilla 5     # 5 minute timer
timezilla 10.1  # 10 minute, six second timer
~~~

### inspiration

I started this project to provide a simple timer for use with the [pomodoro technique](http://en.wikipedia.org/wiki/Pomodoro_Technique), a time management method.

### stuff nobody cares about

this project uses...

* [termbox-go](http://github.com/nsf/termbox-go)
* [notifize](http://github.com/ctcpip/notifize)
* [osext](http://github.com/kardianos/osext)
* [Go](http://golang.org)
* [semantic versioning](http://semver.org/)

### license

[GNU GPLv3](http://www.gnu.org/licenses/gpl-3.0.en.html)
