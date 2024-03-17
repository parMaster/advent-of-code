package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func hexToBin(hex string) string {
	bin := ""
	for _, s := range hex {
		i, _ := strconv.ParseUint(string(s), 16, 64)
		bin += fmt.Sprintf("%04b", i)
	}
	return bin
}

func parsePacket(packet string, maxPackets int) (string, []int, []int) {
	versions := []int{}
	values := []int{}
	for len(packet) > 6 && strings.Contains(packet, "1") && maxPackets > 0 {
		maxPackets--
		// fmt.Println("Parsing packet", packet)
		v, _ := strconv.ParseUint(packet[:3], 2, 64)
		t, _ := strconv.ParseUint(packet[3:6], 2, 64)
		// fmt.Println("\tversion", v)
		versions = append(versions, int(v))
		// fmt.Println("\ttype", t)
		packet = packet[6:]

		if t == 4 {
			// fmt.Println("Literal packet")
			literal := ""
			lastGroup := false
			for len(packet) >= 5 && !lastGroup {
				group := packet[:5]
				if group[0] == '0' {
					lastGroup = true
				}
				// fmt.Println(group)
				literal += group[1:5]
				packet = packet[5:]
			}
			literalNumber, _ := strconv.ParseUint(literal, 2, 64)
			values = append(values, int(literalNumber))
			// fmt.Println("Literal Number:", literalNumber)
		} else {
			var vals []int

			// fmt.Println("Operator packet", packet)
			lenghtTypeId := string(packet[0])
			// fmt.Println("LenghtTypeId:", lenghtTypeId)
			packet = packet[1:]
			if lenghtTypeId == "0" {
				subBits := packet[:15]
				packet = packet[15:]
				lenght, _ := strconv.ParseUint(subBits, 2, 64)
				// fmt.Println("Subpackets Lenght:", lenght)
				var vers []int
				_, vers, vals = parsePacket(packet[:int(lenght)], 999)
				packet = packet[int(lenght):]
				versions = append(versions, vers...)

			} else if lenghtTypeId == "1" {
				subBits := packet[:11]
				packet = packet[11:]
				num, _ := strconv.ParseUint(subBits, 2, 64)
				// fmt.Println("Subpackets Number:", num)
				var vers []int
				packet, vers, vals = parsePacket(packet, int(num))
				versions = append(versions, vers...)
			}

			// fmt.Println("Values:", vals)

			switch t {
			case 0:
				// fmt.Println("Addition")
				sum := 0
				for _, v := range vals {
					sum += v
				}
				vals = []int{sum}
			case 1:
				// fmt.Println("Multiplication")
				mul := 1
				for _, v := range vals {
					mul *= v
				}
				vals = []int{mul}
			case 2:
				// fmt.Println("Minimum")
				min := vals[0]
				for _, v := range vals {
					if v < min {
						min = v
					}
				}
				vals = []int{min}
			case 3:
				// fmt.Println("Maximum")
				max := vals[0]
				for _, v := range vals {
					if v > max {
						max = v
					}
				}
				vals = []int{max}
			case 5:
				// fmt.Println("Greater than")
				if vals[0] > vals[1] {
					vals = []int{1}
				} else {
					vals = []int{0}
				}
			case 6:
				// fmt.Println("Less than")
				if vals[0] < vals[1] {
					vals = []int{1}
				} else {
					vals = []int{0}
				}
			case 7:
				// fmt.Println("Equal")
				// fmt.Println(vals)
				if vals[0] == vals[1] {
					vals = []int{1}
				} else {
					vals = []int{0}
				}
			}
			values = append(values, vals...)
		}

		// fmt.Println("Packet left:", packet)
	}
	return packet, versions, values
}

func solve(transmission string) (p1 int, p2 int) {
	_, versions, values := parsePacket(hexToBin(transmission), 99999)
	vSum := 0
	for _, v := range versions {
		vSum += v
	}
	// fmt.Println("Sum of versions:", vSum)
	// fmt.Println("Values:", values)
	return vSum, values[0]
}

func main() {
	start := time.Now()
	fmt.Println("Day 16: Packet Decoder")
	p1, p2 := solve("E20D7880532D4E551A5791BD7B8C964C1548CB3EC1FCA41CC00C6D50024400C202A65C00C20257C008AF70024C00810039C00C3002D400A300258040F200D6040093002CC0084003FA52DB8134DE620EC01DECC4C8A5B55E204B6610189F87BDD3B30052C01493E2DC9F1724B3C1F8DC801E249E8D66C564715589BCCF08B23CA1A00039D35FD6AC5727801500260B8801F253D467BFF99C40182004223B4458D2600E42C82D07CC01D83F0521C180273D5C8EE802B29F7C9DA1DCACD1D802469FF57558D6A65372113005E4DB25CF8C0209B329D0D996C92605009A637D299AEF06622CE4F1D7560141A52BC6D91C73CD732153BF862F39BA49E6BA8C438C010E009AA6B75EF7EE53BBAC244933A48600B025AD7C074FEB901599A49808008398142013426BD06FA00D540010C87F0CA29880370E21D42294A6E3BCF0A080324A006824E3FCBE4A782E7F356A5006A587A56D3699CF2F4FD6DF60862600BF802F25B4E96BDD26049802333EB7DDB401795FC36BD26A860094E176006A0200FC4B8790B4001098A50A61748D2DEDDF4C6200F4B6FE1F1665BED44015ACC055802B23BD87C8EF61E600B4D6BAD5800AA4E5C8672E4E401D0CC89F802D298F6A317894C7B518BE4772013C2803710004261EC318B800084C7288509E56FD6430052482340128FB37286F9194EE3D31FA43BACAF2802B12A7B83E4017E4E755E801A2942A9FCE757093005A6D1F803561007A17C3B8EE0008442085D1E8C0109E3BC00CDE4BFED737A90DC97FDAE6F521B97B4619BE17CC01D94489E1C9623000F924A7C8C77EA61E6679F7398159DE7D84C015A0040670765D5A52D060200C92801CA8A531194E98DA3CCF8C8C017C00416703665A2141008CF34EF8019A080390962841C1007217C5587E60164F81C9A5CE0E4AA549223002E32BDCEA36B2E100A160008747D8B705C001098DB13A388803F1AE304600")
	fmt.Println("\tPart One:", p1) // 979
	fmt.Println("\tPart Two:", p2) // 277110354175
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
