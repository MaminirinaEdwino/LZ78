package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type Couple struct {
	Index     int
	Caractere string
}

type CoupleByte struct {
	Index     byte
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

func CheckByteInTab(ByteTab [][]byte, curr []byte) bool {
	for i := range ByteTab {
		if bytes.Equal(ByteTab[i], curr) {
			return true
		}
	}

	return false
}
func GetByteIndexe(ByteTab [][]byte, cuur []byte) int {
	for i := range len(ByteTab) {
		if bytes.Equal(ByteTab[i], cuur) {
			return i
		}
	}
	return 0
}

func EncodeLZ78B(data []byte) ([]byte, []int) {
	var ByteTab [][]byte
	var Index []int
	var EncodeByte []byte

	var curr []byte

	for i := range len(data) {
		if !CheckByteInTab(ByteTab, curr) && len(curr) > 0 {
			ByteTab = append(ByteTab, curr)
			Index = append(Index, GetByteIndexe(ByteTab, curr[:len(curr)-1]))
			EncodeByte = append(EncodeByte, curr[len(curr)-1])
			curr = []byte{data[i]}
		} else {
			curr = append(curr, data[i])
		}
	}
	Index = append(Index, GetByteIndexe(ByteTab, []byte{data[len(data)-1]}))
	EncodeByte = append(EncodeByte, data[len(data)-1])

	fmt.Println(string(EncodeByte))
	fmt.Println(Index)
	return EncodeByte, Index
}

func DecodeLZ78B(data []byte, index []int) ([]byte) {
	var res [][]byte
	for i := range data {
		if index[i] > 0 {
			tmp := []byte{}
			if i < len(data)-1 {
				t := res[index[i]]

				// tmp = []byte{t..., data[i]}
				for i := range t {
					tmp = append(tmp, t[i])
				}
				tmp = append(tmp, data[i])
			} else {
				tmp = []byte{data[index[i]]}
			}
			res = append(res, tmp)
		} else {
			res = append(res, []byte{data[i]})
		}
	}
	fmt.Println(string(res[len(res)-1]))
	final := []byte{}
	for i := range res {
		fmt.Println(res[i])
		final = append(final, res[i]...)
	}
	
	fmt.Println(string(final), "Binary")
	return final
}

func Decode(content []Couple) []string {
	var res []string
	for _, idx := range content[0:] {
		if idx.Index > 0 {
			res = append(res, res[idx.Index]+idx.Caractere)
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
			Index:     byte(uint(i.Index)),
			Caractere: []byte(i.Caractere),
		})
	}
	return nil
}

func main() {
	fmt.Println("LZ 78")
	// fmt.Println(append([]byte("A"), []byte("A")...))

	// data, _ := os.ReadFile("1k.txt")
	// encodedData, intData := EncodeLZ78B(data)
	// // Decode(Encode(data))
	// file, _ := os.Create("output.combyte")
	// w := bufio.NewWriter(file)
	// defer w.Flush()
	

	// binary.Write(w, binary.LittleEndian, int32(len(intData)))
	// binary.Write(w, binary.LittleEndian, int32(len(encodedData)))
	// for i := range intData {
	// 	binary.Write(w, binary.LittleEndian, int32(intData[i]))
	// }
	// w.Write(encodedData)

	file1, _ := os.Open("output.combyte")
	r := bufio.NewReader(file1)

	var lenIntData, lenEncodedData int32
	var encodedInt []int
	err := binary.Read(r, binary.LittleEndian, &lenIntData)
	if err == io.EOF {
		fmt.Println("Fin de Fichier")
	}
	err = binary.Read(r, binary.LittleEndian, &lenEncodedData)
	if err == io.EOF {
		panic(err)
	}
	encodedDataContent := make([]byte, lenEncodedData)


	for range lenIntData{
		var tmp int32
		err = binary.Read(r, binary.LittleEndian, &tmp)
		encodedInt = append(encodedInt, int(tmp))
		if err == io.EOF {
			break
		}
	}

	_, err = io.ReadFull(r, encodedDataContent)
	if err == io.EOF {
		panic(err)
	}
	fmt.Println(string(DecodeLZ78B(encodedDataContent, encodedInt)))
}
