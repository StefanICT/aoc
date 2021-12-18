package main

import (
	. "aoc/util"
	"fmt"
)

type Scanner struct {
	Ints     []int
	i        int
	j        int
	Consumed int
}

func (scanner Scanner) Scan() bool {
	return len(scanner.Ints) != 0 &&
		scanner.i < len(scanner.Ints) &&
		scanner.j < 4
}

func (scanner *Scanner) Int(bits int) int {
	buffer := 0

	for i := 0; i < bits; i++ {
		int := (*scanner).Ints[(*scanner).i]
		mask := 1 << (3 - (*scanner).j)
		value := (int & mask >> (3 - (*scanner).j)) & 1

		buffer = (buffer << 1) | value

		(*scanner).j++
		(*scanner).Consumed++
		if (*scanner).j >= 4 {
			(*scanner).j = 0
			(*scanner).i++
		}
	}

	return buffer
}

const literal = 4
const sum = 0
const product = 1
const minimum = 2
const maximum = 3
const greater_than = 5
const less_than = 6
const equal = 7

func main() {
	transmission := ReadLines("2021/16/input")[0]

	fmt.Println(Part1(transmission))
	fmt.Println(Part2(transmission))
}

func Part1(transmission string) int {
	ints := HexToBinary(transmission)

	scanner := Scanner{Ints: ints}
	packet := Parse(&scanner)

	return packet.SumVersion()
}

func Part2(transmission string) int {
	ints := HexToBinary(transmission)

	scanner := Scanner{Ints: ints}
	packet := Parse(&scanner)

	return packet.Evaluate()
}

type Packet interface {
	SumVersion() int
	Evaluate() int
}

type Literal struct {
	Version int
	Id      int
	Value   int
}

func (literal Literal) SumVersion() int {
	return literal.Version
}

func (literal Literal) Evaluate() int {
	return literal.Value
}

type Operator struct {
	Version int
	Id      int
	Packets []*Packet
}

func (operator Operator) SumVersion() int {
	version := operator.Version
	for i := 0; i < len(operator.Packets); i++ {
		version += (*operator.Packets[i]).SumVersion()
	}
	return version
}

func Sum(packets []*Packet) int {
	result := 0
	for _, packet := range packets {
		result += (*packet).Evaluate()
	}
	return result
}

func Product(packets []*Packet) int {
	var result int
	if len(packets) == 0 {
		result = 0
	} else {
		result = 1
	}

	for _, packet := range packets {
		result *= (*packet).Evaluate()
	}
	return result
}

func Minimum(packets []*Packet) int {
	result := MaxInt
	for _, packet := range packets {
		result = Min((*packet).Evaluate(), result)
	}
	return result
}

func Maximum(packets []*Packet) int {
	result := MinInt
	for _, packet := range packets {
		result = Max((*packet).Evaluate(), result)
	}
	return result
}

func GreaterThan(packets []*Packet) int {
	if (*(packets[0])).Evaluate() > (*(packets[1])).Evaluate() {
		return 1
	}
	return 0
}

func LessThan(packets []*Packet) int {
	if (*(packets[0])).Evaluate() < (*(packets[1])).Evaluate() {
		return 1
	}
	return 0
}

func Equal(packets []*Packet) int {
	if (*(packets[0])).Evaluate() == (*(packets[1])).Evaluate() {
		return 1
	}
	return 0
}

func (operator Operator) Evaluate() int {
	var fn func([]*Packet) int
	switch operator.Id {
	case sum:
		fn = Sum
    case product:
		fn = Product
	case minimum:
		fn = Minimum
	case maximum:
		fn = Maximum
	case greater_than:
		fn = GreaterThan
	case less_than:
		fn = LessThan
	case equal:
		fn = Equal
	}
	return fn(operator.Packets)
}

func Parse(scanner *Scanner) Packet {
	version := scanner.Int(3)
	id := scanner.Int(3)

	if id == literal {
		return ParseLiteral(version, id, scanner)
	} else {
		return ParseOperator(version, id, scanner)
	}
}

func ParseLiteral(version int, id int, scanner *Scanner) Packet {
	ints := []int{}
	for scanner.Int(1) == 1 {
		ints = append(ints, scanner.Int(4))
	}
	ints = append(ints, scanner.Int(4))

	value := Combine(ints)

	return Literal{
		Version: version,
		Id:      id,
		Value:   value,
	}
}

func Combine(buffers []int) int {
	buffer := 0
	for i := 0; i < len(buffers); i++ {
		buffer <<= 4
		buffer |= int(buffers[i])
	}
	return buffer
}

func ParseOperator(version int, id int, scanner *Scanner) Packet {
	lengthTypeId := scanner.Int(1)

	if lengthTypeId == 0 {
		return ParseOperatorNumberOfBits(version, id, scanner)
	} else {
		return ParseOperatorNumberSubPackets(version, id, scanner)
	}
}

func ParseOperatorNumberOfBits(version int, id int, scanner *Scanner) Packet {
	length := scanner.Int(15)

	mark := scanner.Consumed

	packets := []*Packet{}
	for length != 0 {
		packet := Parse(scanner)
		packets = append(packets, &packet)

		length -= scanner.Consumed - mark
		mark = scanner.Consumed
	}

	return Operator{
		Version: version,
		Id:      id,
		Packets: packets,
	}
}

func ParseOperatorNumberSubPackets(version int, id int, scanner *Scanner) Packet {
	count := scanner.Int(11)

	packets := make([]*Packet, count)
	for i := 0; i < count; i++ {
		packet := Parse(scanner)
		packets[i] = &packet
	}

	return Operator{
		Version: version,
		Id:      id,
		Packets: packets,
	}
}

var mapping = map[byte]int{
	'0': 0b0000,
	'1': 0b0001,
	'2': 0b0010,
	'3': 0b0011,
	'4': 0b0100,
	'5': 0b0101,
	'6': 0b0110,
	'7': 0b0111,
	'8': 0b1000,
	'9': 0b1001,
	'A': 0b1010,
	'B': 0b1011,
	'C': 0b1100,
	'D': 0b1101,
	'E': 0b1110,
	'F': 0b1111,
}

func HexToBinary(hex string) []int {
	ints := make([]int, len(hex))
	for i := 0; i < len(hex); i++ {
		ints[i] = mapping[hex[i]]
	}
	return ints
}
