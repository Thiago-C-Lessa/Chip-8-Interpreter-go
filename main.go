/*
  Chip 8 interpreter

  For the clock ill use the Golang Ticker



   where a got the roms:  https://github.com/kripod/chip8-roms
  some documentation: http://devernay.free.fr/hacks/chip8/C8TECH10.HTM,
  i found the documentation in this video:  https://www.youtube.com/watch?v=MBWyVwyBMhk&t=579s



  test roms: https://github.com/corax89/chip8-test-rom
  found on this video: https://www.youtube.com/watch?v=YHkBgR6yvbY



  Memory map
  The Chip-8 language is capable of accessing up to 4KB (4,096 bytes) of RAM, from location 0x000 (0) to 0xFFF (4095). The first 512 bytes, from 0x000 to 0x1FF, are where the original interpreter was located, and should not be used by programs.

  Most Chip-8 programs start at location 0x200 (512), but some begin at 0x600 (1536). Programs beginning at 0x600 are intended for the ETI 660 computer.

  Memory Map:
  +---------------+= 0xFFF (4095) End of Chip-8 RAM
  |               |
  |               |
  |               |
  |               |
  |               |
  | 0x200 to 0xFFF|
  |     Chip-8    |
  | Program / Data|
  |     Space     |
  |               |
  |               |
  |               |
  +- - - - - - - -+= 0x600 (1536) Start of ETI 660 Chip-8 programs
  |               |
  |               |
  |               |
  +---------------+= 0x200 (512) Start of most Chip-8 programs
  | 0x000 to 0x1FF|
  | Reserved for  |
  |  interpreter  |
  +---------------+= 0x000 (0) Start of Chip-8 RAM

*/

package main

import (
	"Chip-8-Interpreter-go/chip8"
)

const (
	ROM = "chip8-test-rom-master/test_opcode.ch8"
	//ROM = "roms/Pong (1 player).ch8"
)

func main() {

	//frequencyHz := 500.0
	//interval := time.Duration(float64(time.Second) / frequencyHz)
	//
	//ticker := time.NewTicker(interval)
	//defer ticker.Stop()
	//
	//for range ticker.C {
	//  ciclo do emulador
	//}
	var chip8 = chip8.CHIP8{}
	chip8.Init(ROM)
}
