package day9

import (
	"fmt"
	"slices"
	"zsoki/aoc/common"
)

func Day9a() {
	lines := make(chan string)
	go common.ReadLines("input/day9.txt", lines)

	var disk = make([]int, 0)
	for line := range lines {
		fileId := 0
		for idx, char := range line {
			blockLength := common.ToInt(string(char))
			if idx%2 == 0 {
				// File blockStruct
				if blockLength != 0 {
					disk = append(disk, slices.Repeat([]int{fileId}, blockLength)...)
				}
				fileId++
			} else {
				if blockLength != 0 {
					disk = append(disk, slices.Repeat([]int{-1}, blockLength)...)
				}
			}
		}
	}

	// Rearrange
	reverseIdx := len(disk) - 1
	for position, fileId := range disk {
		// Finished, break.
		if position >= reverseIdx {
			break
		}
		if fileId != -1 {
			// Not empty
			continue
		}
		// Found empty space, look non-empty from the back
		for disk[reverseIdx] == -1 {
			reverseIdx--
			// Finished, break.
			if position >= reverseIdx {
				break
			}
		}
		disk[position] = disk[reverseIdx]
		disk[reverseIdx] = -1
	}

	//fmt.Println(disk)

	// Checkshum
	checkSum := 0
	for position, fileId := range disk {
		if fileId == -1 {
			// Finished
			break
		}
		checkSum += position * fileId
	}

	fmt.Printf("CheckSum: %d\n", checkSum)
}

type blockStruct struct {
	fileId int
	length int
}

const empty = -1

func Day9b() {
	lines := make(chan string)
	go common.ReadLines("input/day9.txt", lines)

	var disk = make([]blockStruct, 0)
	for line := range lines {
		fileId := 0
		for idx, char := range line {
			blockLength := common.ToInt(string(char))
			if idx%2 == 0 {
				// File block
				disk = append(disk, blockStruct{fileId, blockLength})
				fileId++
			} else {
				// Empty block
				disk = append(disk, blockStruct{empty, blockLength})
			}
		}
	}

	// Rearrange
	for rightIdx := len(disk) - 1; rightIdx > 0; rightIdx-- {
		if disk[rightIdx].fileId == empty {
			continue
		}

		// Found file block at rightIdx.
		for leftIdx := 0; leftIdx < rightIdx; leftIdx++ {
			if disk[leftIdx].fileId != empty {
				continue
			}

			// Found empty block at leftIdx
			// Need to join consequent empty blocks, in case we moved the files from there
			for disk[leftIdx+1].fileId == empty {
				newDisk := make([]blockStruct, len(disk)-1)
				copy(newDisk[:leftIdx], disk[:leftIdx])
				newDisk[leftIdx].fileId = empty
				newDisk[leftIdx].length = disk[leftIdx].length + disk[leftIdx+1].length
				copy(newDisk[leftIdx+1:], disk[leftIdx+2:])
				rightIdx-- // Array just became shorter.
				disk = newDisk
			}

			fits := disk[leftIdx].length >= disk[rightIdx].length
			if fits {
				// Move rightIdx file blockStruct into leftIdx empty blockStruct.
				// If empty blockStruct are longer, it must be split to have a remainder as empty blockStruct.
				var newDisk []blockStruct
				emptyRemaining := disk[leftIdx].length - disk[rightIdx].length
				if emptyRemaining > 0 {
					newDisk = make([]blockStruct, len(disk)+1)
					copy(newDisk[:leftIdx], disk[:leftIdx])
					newDisk[leftIdx] = disk[rightIdx]
					newDisk[leftIdx+1] = blockStruct{empty, emptyRemaining}
					copy(newDisk[leftIdx+2:rightIdx+1], disk[leftIdx+1:rightIdx])
					newDisk[rightIdx+1] = blockStruct{empty, disk[rightIdx].length}
					copy(newDisk[rightIdx+2:], disk[rightIdx+1:])
				} else {
					newDisk = make([]blockStruct, len(disk))
					copy(newDisk[:leftIdx], disk[:leftIdx])
					newDisk[leftIdx] = disk[rightIdx]
					copy(newDisk[leftIdx+1:rightIdx], disk[leftIdx+1:rightIdx])
					newDisk[rightIdx] = blockStruct{empty, disk[rightIdx].length}
					copy(newDisk[rightIdx+1:], disk[rightIdx+1:])
				}
				disk = newDisk

				// Replacement done, need to step rightIdx -= 1
				break
			}
		}
	}

	// Checksum
	checkSum := 0
	position := -1
	for _, block := range disk {
		for i := 0; i < block.length; i++ {
			position++
			if block.fileId != empty {
				checkSum += position * block.fileId
			}
		}
	}

	fmt.Printf("\n\nCheckSum: %d\n", checkSum)
}

func printDisk(disk []blockStruct) {
	fmt.Println()
	for _, blk := range disk {
		for i := 0; i < blk.length; i++ {
			switch blk.fileId {
			case empty:
				fmt.Print(".")
			default:
				fmt.Print(blk.fileId)
			}
		}
	}
}
