package chip8

import (
	"fmt"
	"log"
	"os"
)

type CHIP8 struct {
	rom                              string
	ram                              [4096]uint8
	sound_registrer, video_registrer uint8
	general_use_registrer            [16]uint8 // arr[0x0] até arr[0xF]
	pc                               uint16    // currently executing address
	sp                               uint8     //topmost level of the stack
	screen                           [32][64]uint8
}

// Init initialize chip8 struct
func (chip8 *CHIP8) Init(rom string) {
	chip8.rom = rom
	chip8.pc = 0x200 //0x200 (512) Start of most Chip-8 programs

	chip8.loadFont()
	chip8.loadRom()
}

// loadFont Loads the needed font in chip8 memory
func (chip8 *CHIP8) loadFont() {
	//"The data should be stored in the interpreter area of Chip-8 memory (0x000 to 0x1FF)."

	// 0
	chip8.ram[0x000] = 0xF0
	chip8.ram[0x001] = 0x90
	chip8.ram[0x002] = 0x90
	chip8.ram[0x003] = 0x90
	chip8.ram[0x004] = 0xF0

	// 1
	chip8.ram[0x005] = 0x20
	chip8.ram[0x006] = 0x60
	chip8.ram[0x007] = 0x20
	chip8.ram[0x008] = 0x20
	chip8.ram[0x009] = 0x70

	// 2
	chip8.ram[0x00A] = 0xF0
	chip8.ram[0x00B] = 0x10
	chip8.ram[0x00C] = 0xF0
	chip8.ram[0x00D] = 0x80
	chip8.ram[0x00E] = 0xF0

	// 3
	chip8.ram[0x00F] = 0xF0
	chip8.ram[0x010] = 0x10
	chip8.ram[0x011] = 0xF0
	chip8.ram[0x012] = 0x10
	chip8.ram[0x013] = 0xF0

	// 4
	chip8.ram[0x014] = 0x90
	chip8.ram[0x015] = 0x90
	chip8.ram[0x016] = 0xF0
	chip8.ram[0x017] = 0x10
	chip8.ram[0x018] = 0x10

	// 5
	chip8.ram[0x019] = 0xF0
	chip8.ram[0x01A] = 0x80
	chip8.ram[0x01B] = 0xF0
	chip8.ram[0x01C] = 0x10
	chip8.ram[0x01D] = 0xF0

	// 6
	chip8.ram[0x01E] = 0xF0
	chip8.ram[0x01F] = 0x80
	chip8.ram[0x020] = 0xF0
	chip8.ram[0x021] = 0x90
	chip8.ram[0x022] = 0xF0

	// 7
	chip8.ram[0x023] = 0xF0
	chip8.ram[0x024] = 0x10
	chip8.ram[0x025] = 0x20
	chip8.ram[0x026] = 0x40
	chip8.ram[0x027] = 0x40

	// 8
	chip8.ram[0x028] = 0xF0
	chip8.ram[0x029] = 0x90
	chip8.ram[0x02A] = 0xF0
	chip8.ram[0x02B] = 0x90
	chip8.ram[0x02C] = 0xF0

	// 9
	chip8.ram[0x02D] = 0xF0
	chip8.ram[0x02E] = 0x90
	chip8.ram[0x02F] = 0xF0
	chip8.ram[0x030] = 0x10
	chip8.ram[0x031] = 0xF0

	// A
	chip8.ram[0x032] = 0xF0
	chip8.ram[0x033] = 0x90
	chip8.ram[0x034] = 0xF0
	chip8.ram[0x035] = 0x90
	chip8.ram[0x036] = 0x90

	// B
	chip8.ram[0x037] = 0xE0
	chip8.ram[0x038] = 0x90
	chip8.ram[0x039] = 0xE0
	chip8.ram[0x03A] = 0x90
	chip8.ram[0x03B] = 0xE0

	// C
	chip8.ram[0x03C] = 0xF0
	chip8.ram[0x03D] = 0x80
	chip8.ram[0x03E] = 0x80
	chip8.ram[0x03F] = 0x80
	chip8.ram[0x040] = 0xF0

	// D
	chip8.ram[0x041] = 0xE0
	chip8.ram[0x042] = 0x90
	chip8.ram[0x043] = 0x90
	chip8.ram[0x044] = 0x90
	chip8.ram[0x045] = 0xE0

	// E
	chip8.ram[0x046] = 0xF0
	chip8.ram[0x047] = 0x80
	chip8.ram[0x048] = 0xF0
	chip8.ram[0x049] = 0x80
	chip8.ram[0x04A] = 0xF0

	// F
	chip8.ram[0x04B] = 0xF0
	chip8.ram[0x04C] = 0x80
	chip8.ram[0x04D] = 0xF0
	chip8.ram[0x04E] = 0x80
	chip8.ram[0x04F] = 0x80
}

// loadRom Loads the rom in the struct
func (chip8 *CHIP8) loadRom() {
	fmt.Println("Loading ROM")
	fmt.Println("\t", chip8.rom)
	//fmt.Println(chip8.readRom())
	byteArr := chip8.readRom()
	for i, b := range byteArr {
		chip8.ram[0x200+i] = b
	}
}

// readRom Reads the rom from the disk
func (chip8 *CHIP8) readRom() []byte {

	fmt.Println("\tReading ROM")
	_rom, err := os.Stat(chip8.rom)
	if err != nil {
		log.Fatal(fmt.Errorf("Error opening ROM file: %v", err))
	}
	fmt.Println("\tSize of rom", _rom.Size())

	buffer, err := os.ReadFile(chip8.rom)
	if err != nil {
		log.Fatal(fmt.Errorf("Error reading ROM: %v", err))
	}

	fmt.Println("\tRead file size", len(buffer))

	return buffer
}
