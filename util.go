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

import (
	"fmt"
	"strings"
	"time"
)

func getTimeString(d time.Duration) string {

	var strTime string

	if d < time.Second*1 {
		strTime = "0"
	}

	s := d.String()
	a := strings.Split(s, ".")

	if len(a) > 1 {
		strTime = a[0] + "s" // discard fractional seconds
	} else {
		strTime = s
	}

	return strTime

}

func printHelp(arg string) {

	if arg != "-h" {
		fmt.Print("\ntimezilla: invalid argument -- '" + arg + "'\n")
	}

	fmt.Print(strHelp)

}
