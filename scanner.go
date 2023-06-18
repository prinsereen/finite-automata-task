package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getToken(teks string, j *int) string {
	k := len(teks)
	kata := ""

	for *j < k && (teks[*j] == ' ' || teks[*j] == '\r' || teks[*j] == '\n') {
		*j++
	}

	for *j < k && teks[*j] != ' ' && teks[*j] != '\r' && teks[*j] != '\n' {
		if teks[*j] == '>' {
			if kata != "" {
				return kata
			} else {
				*j++
				if *j < k && teks[*j] == '=' {
					*j++
					return ">="
				} else {
					return ">"
				}
			}
		} else if teks[*j] == '<' {
			if kata != "" {
				return kata
			} else {
				*j++
				if *j < k && teks[*j] == '=' {
					*j++
					return "<="
				} else {
					return "<"
				}
			}
		} else if teks[*j] == '=' {
			if kata != "" {
				return kata
			} else {
				*j++
				return "="
			}
		} else if teks[*j] == '+' {
			if kata != "" {
				return kata
			} else {
				*j++
				return "+"
			}
		} else if teks[*j] == '-' {
			if kata != "" {
				return kata
			} else {
				*j++
				return "-"
			}
		}
		kata += string(teks[*j])
		*j++
	}
	return kata
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the text input:\n")
	var input string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		input += line + "\n"
	}

	i := 0
	hsl := ""
	k := len(input)

	for i < k {
		token := getToken(input, &i)
		hsl += strings.TrimSpace(token) + "\n"
	}

	fmt.Println("Hasil:")
	fmt.Println(hsl)
}
