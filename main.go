package main

import (
	// "encoding/gob"
	"encoding/gob"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Couple struct {
	Index     int
	Caractere string
}

type CoupleByte struct{
	Index byte
	Caractere []byte
}

func CheckIfCharInTab(char string, chartab []string) bool {
	return slices.Contains(chartab, char)
}

func GetTheIndexe(char string, charTab []string) int {
	if len(char) > 1 {
		for idx := range charTab {
			if char[:len(char)-1] == charTab[idx] {
				return idx
			}
		}
	}
	return 0
}

func GetChar(char string) string {
	content := strings.Split(char, "")
	if len(content) > 1 {
		return content[len(content)-1]
	}
	return char
}

func Encode(content string) []Couple {
	var caractere []string
	var encodage []Couple

	contentTab := strings.Split(content, "")
	var char string

	for i := range len(contentTab) {
		if CheckIfCharInTab(char, caractere) == false {
			caractere = append(caractere, char)
			codage := Couple{
				Index:     GetTheIndexe(char, caractere),
				Caractere: GetChar(char),
			}
			encodage = append(encodage, codage)
			char = contentTab[i]
		} else {
			char += contentTab[i]
		}
	}
	codage := Couple{
		Index:     GetTheIndexe(contentTab[len(contentTab)-1], caractere),
		Caractere: contentTab[len(contentTab)-1],
	}
	encodage = append(encodage, codage)
	// fmt.Println(encodage)
	return encodage
}

func Decode(content []Couple) []string {
	var res []string
	for _, idx := range content[0:] {
		if idx.Index > 0 {
			res = append(res, res[idx.Index ]+idx.Caractere)
			// fmt.Print(res[idx.Index]+idx.Caractere)
		} else {
			res = append(res, idx.Caractere)
			// fmt.Print(idx.Caractere)
		}

	}
	
	return res
}

func ConvertIntoByte(encodage []Couple) []CoupleByte {
	var res []CoupleByte
	for _, i := range encodage {
		res = append(res, CoupleByte{
			Index: byte(uint(i.Index)),
			Caractere: []byte(i.Caractere),
		})
	}
	return nil
}

func main() {
	fmt.Println("LZ 78")
	data, _ := os.ReadFile("texte.txt")  
	encodage := Encode(string(data))
	encodedFile, _ := os.OpenFile("encoded.ed", os.O_CREATE|os.O_RDONLY|os.O_RDWR, 0644)
	encoder := gob.NewEncoder(encodedFile)
	fmt.Println(encodage)
	encoder.Encode(encodage)
}
