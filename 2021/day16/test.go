package main

import "fmt"

func testSumVer(packet string, expectedSumVer uint64) {
	fmt.Printf("Testing start: %v\n", packet)
	input := convert(packet)
	sumVer, value, idx := decode(input, 0)
	fmt.Printf("idx %v sumVer %v value %v\n", idx, sumVer, value)
	if sumVer == expectedSumVer {
		fmt.Println("PASS!")
	} else {
		fmt.Println("FAIL!")
	}
	fmt.Printf("Testing end: %v\n", packet)
	fmt.Println()
}

func test1() {
	testSumVer("8A004A801A8002F478", 16)
	testSumVer("620080001611562C8802118E34", 12)
	testSumVer("C0015000016115A2E0802F182340", 23)
	testSumVer("A0016C880162017C3686B18A3D4780", 31)
}

func testOperators(packet string, expected uint64) {
	fmt.Printf("Testing start: %v\n", packet)
	bin := convert(packet)
	sumVer, value, idx := decode(bin, 0)
	fmt.Printf("idx %v sumVer %v value %v\n", idx, sumVer, value)
	if value == expected {
		fmt.Println("PASS!")
	} else {
		fmt.Println("FAIL!")
	}
	fmt.Printf("Testing end: %v\n", packet)
	fmt.Println()
}

func test2() {
	testOperators("C200B40A82", 3)
	testOperators("04005AC33890", 54)
	testOperators("880086C3E88112", 7)
	testOperators("CE00C43D881120", 9)
	testOperators("D8005AC2A8F0", 1)
	testOperators("F600BC2D8F", 0)
	testOperators("9C005AC2F8F0", 0)
	testOperators("9C0141080250320F1802104A08", 1)
}
