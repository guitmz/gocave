package main

import (
	"debug/elf"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func byteInSlice(a byte, list []byte) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func findCave(sectionName string, sectionBody []byte, caveSize int, sectionOffset uint64, sectionAddr uint64, sectionFlags string) {
	// For this example, we are looking for 0x00 bytes only, so we create a slice with one element (0x00 in this case).
	caveBytes := make([]byte, 1)
	caveBytes[0] = 0x00

	caveCount := 0
	for currentOffset := 0; currentOffset < len(sectionBody); currentOffset++ {
		currentByte := sectionBody[currentOffset]
		if byteInSlice(currentByte, caveBytes) {
			caveCount++
		} else {
			if caveCount >= caveSize {
				fmt.Println("\n[+] CAVE DETECTED!")
				fmt.Printf("[!] Section Name: %s\n", sectionName)
				fmt.Printf("[!] Flags: %s\n", sectionFlags)
				fmt.Printf("[!] Virtual Address: %#x\n", int(sectionAddr)+currentOffset-caveCount)
				fmt.Printf("[!] Size: %#x\n", caveCount)
				fmt.Printf("[!] Begin: %#x\n", int(sectionOffset)+currentOffset-caveCount)
				fmt.Printf("[!] End: %#x\n", int(sectionOffset)+currentOffset)
			}
			caveCount = 0
		}
	}
}

func main() {

	lenArgs := len(os.Args)
	if lenArgs < 3 || lenArgs > 3 {
		fmt.Println("Usage: gocave elf_file cave_size")
		os.Exit(1)
	}

	elfFile, err := elf.Open(os.Args[1])
	check(err)

	caveSize, err := strconv.Atoi(os.Args[2])
	check(err)

	for i := 0; i < len(elfFile.Sections); i++ {
		data, _ := elfFile.Sections[i].Data()
		findCave(elfFile.Sections[i].Name,
			data,
			caveSize,
			elfFile.Sections[i].Offset,
			elfFile.Sections[i].Addr,
			elfFile.Sections[i].Flags.String())
	}
}
