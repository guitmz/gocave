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

func findCave(section *elf.Section, caveSize int) {
	// For this example, we are looking for 0x00 bytes only, so we create a slice with one element (0x00 in this case).
	caveBytes := make([]byte, 1)
	caveBytes[0] = 0x00

	sectionBody, _ := section.Data()
	caveCount := 0
	for currentOffset := 0; currentOffset < len(sectionBody); currentOffset++ {
		currentByte := sectionBody[currentOffset]
		if byteInSlice(currentByte, caveBytes) {
			caveCount++
		} else {
			if caveCount >= caveSize {
				fmt.Println("\n[+] CAVE DETECTED!")
				fmt.Printf("[!] Section Name: %s\n", section.Name)
				fmt.Printf("[!] Section Offset: %#x\n", section.Offset)
				fmt.Printf("[!] Section Size: %#x (%d bytes)\n", section.Size, int(section.Size))
				fmt.Printf("[!] Section Flags: %s\n", section.Flags.String())
				fmt.Printf("[!] Virtual Address: %#x\n", int(section.Addr)+currentOffset-caveCount)
				fmt.Printf("[!] Cave Begin: %#x\n", int(section.Offset)+currentOffset-caveCount)
				fmt.Printf("[!] Cave End: %#x\n", int(section.Offset)+currentOffset)
				fmt.Printf("[!] Cave Size: %#x (%d bytes)\n", caveCount, int(caveCount))
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

	for _, section := range elfFile.Sections {
		findCave(section, caveSize)
	}
}
