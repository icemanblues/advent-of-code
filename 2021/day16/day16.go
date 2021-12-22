package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "16"
	dayTitle = "Packet Decoder"
)

var hexBinMap = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func convert(packet string) string {
	builder := strings.Builder{}
	for _, r := range packet {
		builder.WriteString(hexBinMap[r])
	}
	return builder.String()
}

func Parse(filename string) string {
	input, _ := util.ReadInput(filename)
	return convert(input[0])
}

func decode(packet string, index int) (uint64, uint64, int) {
	i := index
	sumVer, _ := strconv.ParseUint(packet[i:i+3], 2, 64)
	i += 3
	t, _ := strconv.ParseUint(packet[i:i+3], 2, 64)
	i += 3

	var value uint64 = 0
	switch t {
	case 4: // literal packet
		literalBuilder := strings.Builder{}
		for loop := "1"; loop == "1"; {
			loop = packet[i : i+1]
			i++
			literalBuilder.WriteString(packet[i : i+4])
			i += 4
		}
		value, _ = strconv.ParseUint(literalBuilder.String(), 2, 64)

	default: // operator packet
		var subValues []uint64
		lenType := packet[i : i+1]
		i++
		switch lenType {
		case "0":
			totalLength, _ := strconv.ParseUint(packet[i:i+15], 2, 64)
			i += 15
			// keep reading literal subpackets until total length
			target := i + int(totalLength)
			for i < target {
				jSumVer, jv, ji := decode(packet, i)
				subValues = append(subValues, jv)
				sumVer += jSumVer
				i = ji
			}
		case "1":
			numSubPackets, _ := strconv.ParseUint(packet[i:i+11], 2, 64)
			i += 11
			for j := 0; j < int(numSubPackets); j++ {
				jSumVer, jv, ji := decode(packet, i)
				subValues = append(subValues, jv)
				sumVer += jSumVer
				i = ji
			}
		}

		// compute the value of the operator packet
		switch t {
		case 0: // sum
			for _, e := range subValues {
				value += e
			}
		case 1: // product
			value = 1
			for _, e := range subValues {
				value *= e
			}
		case 2: // min
			value = subValues[0]
			for i := 1; i < len(subValues); i++ {
				if subValues[i] < value {
					value = subValues[i]
				}
			}
		case 3: // max
			value = subValues[0]
			for i := 1; i < len(subValues); i++ {
				if subValues[i] > value {
					value = subValues[i]
				}
			}
		case 5: // gt
			if subValues[0] > subValues[1] {
				value = 1
			} else {
				value = 0
			}
		case 6: // lt
			if subValues[0] < subValues[1] {
				value = 1
			} else {
				value = 0
			}
		case 7: // equal to
			if subValues[0] == subValues[1] {
				value = 1
			} else {
				value = 0
			}
		}
	}

	return sumVer, value, i
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	input := Parse("input.txt")
	sumVer, value, _ := decode(input, 0)
	fmt.Printf("Part 1: sum of version: %v\n", sumVer)
	fmt.Printf("Part 2: value: %v\n", value)
}
