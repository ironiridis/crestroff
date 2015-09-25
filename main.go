/*
This go package reads a .usp or .umc file on stdin and attemps to output a
token that may be useful for recovering the file. Some of these files have been
observed exhibiting some form of corruption that causes each byte in the stream
to be doubled. This corruption does appear to be reversible.



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
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var d bytes.Buffer
	var err error
	var c byte
	var f bool

	for in := bufio.NewReader(os.Stdin); err == nil; c, err = in.ReadByte() {
		f = (c >= 0x80 || f)
		if f {
			d.WriteByte(c / 2)
		}
	}

	b := d.Bytes()
	s := 4 + bytes.Index(b, []byte("*PP*"))
	if s < 4 {
		return
	}
	e := bytes.Index(b[s:], []byte("*EPP*"))
	if e < 0 {
		return
	}
	fmt.Printf("%s\n", string(b[s:s+e]))

}
