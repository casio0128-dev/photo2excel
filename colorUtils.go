package main

import "fmt"

func getRGBString(r, g, b uint32) string {
	rs := fmt.Sprintf("%04x", r)[2:]
	gs := fmt.Sprintf("%04x", g)[2:]
	bs := fmt.Sprintf("%04x", b)[2:]
	return "#" + rs + gs + bs
}
