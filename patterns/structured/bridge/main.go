package main

func main() {
	xerox := Xerox{}
	epson := Epson{}

	epsonMac := Mac{printer: &epson}
	epsonMac.Print()

	xeroxWin := Win{}
	xeroxWin.SetPrinter(&xerox)
	xeroxWin.Print()
}
