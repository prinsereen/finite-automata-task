package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack struct {
	n   int
	isi []string
}

func Push(stack *Stack, x string) {
	stack.n++
	stack.isi = append(stack.isi, x)
}

func Pop(stack *Stack) string {
	x := stack.isi[stack.n-1]
	stack.n--
	stack.isi = stack.isi[:stack.n]
	return x
}

func IsEmpty(stack *Stack) bool {
	return stack.n == 0
}

func ReadTop(stack *Stack) string {
	return stack.isi[stack.n-1]
}

func inputHandler(input string) string {
	_for := "f"
	_kali := "x"
	_bagi := "v"
	_minus := "m"
	_ut := "s"
	_uet := "t"
	_ld := "u"
	_let := "g"
	_bb := "y"
	_bt := "z"
	_et := "i"
	_eet := "h"
	_plus := "c"

	words := strings.Split(input, " ")
	var result [][]string

	for _, word := range words {
		result = append(result, []string{word})
	}

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			if result[i][j] == "for" {
				result[i][j] = _for
			} else if result[i][j] == "*" {
				result[i][j] = _kali
			} else if result[i][j] == "/" {
				result[i][j] = _bagi
			} else if result[i][j] == "-" {
				result[i][j] = _minus
			} else if result[i][j] == ">" {
				result[i][j] = _ut
			} else if result[i][j] == ">=" {
				result[i][j] = _uet
			} else if result[i][j] == "<" {
				result[i][j] = _ld
			} else if result[i][j] == "<=" {
				result[i][j] = _let
			} else if result[i][j] == "{" {
				result[i][j] = _bb
			} else if result[i][j] == "}" {
				result[i][j] = _bt
			} else if result[i][j] == "==" {
				result[i][j] = _eet
			} else if result[i][j] == "=" {
				result[i][j] = _et
			} else if result[i][j] == "+" {
				result[i][j] = _plus
			} else if result[i][j] == "1" || result[i][j] == "2" || result[i][j] == "3" || result[i][j] == "4" || result[i][j] == "5" {
				result[i][j] = "a"
			}
		}
	}

	var sb strings.Builder
	for i := 0; i < len(result); i++ {
		sb.WriteString(result[i][0])
	}

	str := sb.String()
	return str
}
func main() {
	stack := Stack{
		n:   0,
		isi: make([]string, 0),
	}

	tabelParsing := map[string]map[string]string{
		"L": {
			"f": "fCBAT",
			"x": "-",
			"v": "-",
			"m": "-",
			"a": "-",
			"s": "-",
			"t": "-",
			"u": "-",
			"g": "-",
			"y": "-",
			"z": "-",
			"h": "-",
			"i": "-",
			"c": "-",
			"#": "-",
		},
		"A": {
			"f": "-",
			"x": "-",
			"v": "-",
			"m": "-",
			"a": "VEVOV",
			"s": "-",
			"t": "-",
			"u": "-",
			"g": "-",
			"y": "-",
			"z": "-",
			"h": "-",
			"i": "-",
			"c": "-",
			"#": "-",
		},
		"V": {
			"f": "-",
			"x": "-",
			"v": "-",
			"m": "-",
			"a": "a",
			"s": "-",
			"t": "-",
			"u": "-",
			"g": "-",
			"y": "-",
			"z": "-",
			"h": "-",
			"i": "-",
			"c": "-",
			"#": "-",
		},
		"C": {
			"f": "-",
			"x": "-",
			"v": "-",
			"m": "-",
			"a": "VPV",
			"s": "-",
			"t": "-",
			"u": "-",
			"g": "-",
			"y": "-",
			"z": "-",
			"h": "-",
			"i": "-",
			"c": "-",
			"#": "-",
		},
		"O": {
			"f": "-",
			"x": "x",
			"v": "v",
			"m": "m",
			"a": "-",
			"s": "-",
			"t": "-",
			"u": "-",
			"g": "-",
			"y": "-",
			"z": "-",
			"h": "-",
			"i": "-",
			"c": "c",
			"#": "-",
		},
		"P": {
			"f": "-",
			"x": "-",
			"v": "-",
			"m": "-",
			"a": "-",
			"s": "s",
			"t": "t",
			"u": "u",
			"g": "g",
			"y": "-",
			"z": "-",
			"h": "h",
			"i": "-",
			"c": "-",
			"#": "-",
		},
		"B": {
			"f": "-",
			"x": "-",
			"v": "-",
			"m": "-",
			"a": "-",
			"s": "-",
			"t": "-",
			"u": "-",
			"g": "-",
			"y": "y",
			"z": "-",
			"h": "-",
			"i": "-",
			"c": "-",
			"#": "-",
		},
		"T": {
			"f": "-",
			"x": "-",
			"v": "-",
			"m": "-",
			"a": "-",
			"s": "-",
			"t": "-",
			"u": "-",
			"g": "-",
			"y": "-",
			"z": "z",
			"h": "-",
			"i": "-",
			"c": "-",
			"#": "-",
		},
		"E": {
			"f": "-",
			"x": "-",
			"v": "-",
			"m": "-",
			"a": "-",
			"s": "-",
			"t": "-",
			"u": "-",
			"g": "-",
			"y": "-",
			"z": "-",
			"h": "-",
			"i": "i",
			"c": "-",
			"#": "-",
		},
	}
	simbolAwal := "L"

	fmt.Println("Aturan Produksi :")
	fmt.Println("LL(1) DENGAN NOTASI SEDERHANA")
	fmt.Println("S -> AB | CD")
	fmt.Println("A -> xA | y")
	fmt.Println("B -> yB | x")
	fmt.Println("C -> aD | b")
	fmt.Println("D -> bC | a")
	fmt.Println()

	fmt.Println("Parse Table")
	printTable(tabelParsing)
	fmt.Println()

	var input string
	fmt.Print("Input: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	input = inputHandler(input)

	fmt.Println("Hasil:")

	Push(&stack, simbolAwal)
	inputIndex := 0

	for !IsEmpty(&stack) {
		top := ReadTop(&stack)
		simbol := ""
		if inputIndex < len(input) {
			simbol = string(input[inputIndex])
		} else {
			simbol = "#"
		}

		if top >= "a" {
			if top == simbol {
				Pop(&stack)
				inputIndex++
			} else {
				fmt.Println("Error/Ditolak")
				break
			}
		} else if top <= "Z" {
			sel := tabelParsing[top][simbol]
			if sel != "-" {
				Pop(&stack)
				for i := len(sel) - 1; i >= 0; i-- {
					Push(&stack, string(sel[i]))
				}
			} else {
				fmt.Println("Error/Ditolak")
				break
			}
		}
	}

	if IsEmpty(&stack) && inputIndex == len(input) {
		fmt.Println("DITERIMA")
	} else {
		fmt.Println("Erorr/Ditolak")
	}
}

func printTable(tbl map[string]map[string]string) {
	fmt.Println("  \tf\tx\tv\tm\ta\ts\tt\tu\tg\ty\tz\th\ti\tc\t#")
	for jdlBaris, dt := range tbl {
		fmt.Print(jdlBaris + "\t")
		for _, isi := range dt {
			fmt.Print(isi + "\t")
		}
		fmt.Println()
	}
}
