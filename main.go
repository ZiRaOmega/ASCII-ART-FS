package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		Fonts := os.Args[2]
		inputtext := os.Args[1]
		fmt.Println(Show_ascii(Get_ascii_char(inputtext, Fonts)))
	}
}

func Get_ascii_char(caractere string, Font string) []string {
	output := []string{}
	for _, element := range caractere {
		char := rune(element - 32)
		file, err := os.Open(Font + ".txt")
		if err != nil {
			log.Fatal(err)
			//fmt.Println(err)
		}
		scanner := bufio.NewScanner(file)
		ascii := []string{}
		for scanner.Scan() {
			ascii = append(ascii, scanner.Text())
		}
		for i := char * 9; i < (char*9)+9; i++ {
			output = append(output, ascii[i])
		}
	}
	return output
}
func Show_ascii(ascii_char []string) string {
	nbr_lettre := len(ascii_char) / 9
	var word_array []string
	var result string
	for y := 0; y < 9; y++ {
		for i := y; i < nbr_lettre*9; i += 9 {
			word_array = append(word_array, ascii_char[i])
			if ascii_char[i] != "\n" {
				result += ascii_char[i]
			}
		}
		result += "\n"
	}
	return result
}
