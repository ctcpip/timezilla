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

const strHeader = "timezilla 0.1.0 (C) 2016  Chris de Almeida     \"timezilla -h\" for more info"

const strHelp = `
timezilla 0.1.0 (C) 2016  Chris de Almeida    http://github.com/ctcpip/timezilla

a simple timer for the console/terminal

usage: timezilla [minutes]
   minutes specified in fractional minutes
   if no minutes specified, timer will default to 25 minutes (pomodoro standard)

timezilla       # default 25 minute timer
timezilla .5    # 30 second timer
timezilla 2.5   # 2 minute, 30 second timer
timezilla 5     # 5 minute timer
timezilla 10.1  # 10 minute, six second timer

-the timer will count down from the specified time
-when the time is up, three things happen:
   1. an OS notification is sent (using libnotify or growl)
   2. the terminal screen will flash red on/off
   3. the terminal bell will ring
-use CTRL+C to exit

`
