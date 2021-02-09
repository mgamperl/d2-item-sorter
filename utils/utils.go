package utils

import (
	"fmt"
	"log"
	"strconv"

	"golang.org/x/sys/windows/registry"
)

func GetD2SavePath() string {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Blizzard Entertainment\Diablo II`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	s, _, err := k.GetStringValue("Save Path")
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func reverseByte(b byte) byte {
	var d byte
	for i := 0; i < 8; i++ {
		d <<= 1
		d |= b & 1
		b >>= 1
	}
	return d
}

func reverseBits(b uint64, n uint) uint64 {
	var d uint64
	for i := 0; i < int(n); i++ {
		d <<= 1
		d |= b & 1
		b >>= 1
	}
	return d
}

func WriteBits(data []byte, offset uint, size uint, value int) []byte {
	valueToWrite := fmt.Sprintf("%08b", reverseBits(uint64(value), 8))
	byteStart := int(offset / 8)
	byteEnd := int((offset + size - 1) / 8)
	bytesToRead := byteEnd - byteStart + 1
	dataNew := data[:byteStart]
	//fmt.Printf("%08b\n", data)

	strIdx := 0
	for byteIdx := 0; byteIdx < bytesToRead; byteIdx++ {
		currByte := byteStart + byteIdx
		bitsRaw := reverseBits(uint64(data[currByte]), 8)

		bitString := fmt.Sprintf("%08b", bitsRaw)
		result := ""

		for bitIdx := range bitString {
			currentBitPos := uint(((currByte * 8) + bitIdx))
			if offset <= currentBitPos && currentBitPos < offset+size {
				result += string(valueToWrite[strIdx])
				strIdx++
			} else {
				result += string(bitString[bitIdx])
			}
		}

		byteResult, _ := strconv.ParseUint(result, 2, 8)
		//fmt.Printf("\t\t\tByte #%d: %08b, result: %s write: %08b\n", byteIdx, bitsRaw, result, byteResult)
		dataNew = append(dataNew, byte(reverseBits(byteResult, 8)))
	}

	dataNew = append(dataNew, data[byteEnd+1:]...)
	//fmt.Printf("%08b\n", dataNew)
	return dataNew
}
