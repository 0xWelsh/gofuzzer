package main

import "fmt"

func ParseMagic(data []byte) error {
	if len(data) < 4 {
		return nil
	}
	if data[0] == 'M' && data[1] == 'Z' && data[2] == 0x90 {
		// our bug: only check 3 bytes but later use 4th as size
		size := data[3]
		if size > 200 && len(data) < int(size) {
			return fmt.Error("buffer overflow risk")
		}
	}
	return nil
}
