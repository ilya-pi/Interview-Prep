package main

import "fmt"

func drawLine(screen []byte, width, x1, x2, y byte) {
	sz := byte(8)
	if x1 > width || x2 > width || y >= byte(len(screen))/(width/sz) {
		return
	}
	/*
		Find offset from left as y * width / 8 + x1 / 8
			The next byte is the first we need to start writing to at offset x1 % 8
			Then for next (x2 - x1 - x % 8 ) / 8 elements we need to set them to 255
			And fill in (x2 - x1 - x%8) %8 with 1 from left, thus offsetting it
			from right on 8 - that amount of bits
	*/
	offset := y * width / sz
	start := offset + x1/sz //(the next element after that)
	end := offset + x2/sz
	screen[start] = screen[start] | (byte(255) >> x1 % sz)
	for i := start; i < end; i++ {
		screen[i] = byte(255)
	}
	screen[end] = screen[end] | (byte(255) << (sz - x2%sz))
	// todo(ilya): use pointer to screen
}

func main() {
	screen := make([]byte, 9)
	width := byte(24)
	x1 := byte(9)
	x2 := byte(18)
	y := byte(1)
	drawLine(screen, width, x1, x2, y)
	fmt.Printf("%v\n", screen)
}
