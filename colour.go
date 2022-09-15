/*
 * Copyright (c) 2022 Mattia Cabrini <dev@mattiacabrini.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package goutility

const (
	Reset  string = "\033[0m"
	Red           = "\033[31m"
	Green         = "\033[32m"
	Yellow        = "\033[33m"
	Blue          = "\033[34m"
	Purple        = "\033[35m"
	Cyan          = "\033[36m"
	Gray          = "\033[37m"
	White         = "\033[97m"
)

var (
	reset  string = "\033[0m"
	red           = "\033[31m"
	green         = "\033[32m"
	yellow        = "\033[33m"
	blue          = "\033[34m"
	purple        = "\033[35m"
	cyan          = "\033[36m"
	gray          = "\033[37m"
	white         = "\033[97m"
)

func NotInteractive() {
	reset = ""
	red = ""
	green = ""
	yellow = ""
	blue = ""
	purple = ""
	cyan = ""
	gray = ""
	white = ""
}

func Interactive() {
	reset = Reset
	red = Red
	green = Green
	yellow = Yellow
	blue = Blue
	purple = Purple
	cyan = Cyan
	gray = Gray
	white = White
}
