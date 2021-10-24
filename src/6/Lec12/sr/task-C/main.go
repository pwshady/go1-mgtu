package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	var (
		s          string
		result_str string
	)

	data := []string{}
	out := []string{}
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	height, _ := strconv.Atoi(data[0])
	length, _ := strconv.Atoi(data[1])
	s = data[2]

	for i := 0; i < length; i++ {
		result_str += s
	}
	out = append(out, result_str)
	result_str = "\n"
	for i := 1; i < height-1; i++ {
		result_str += s
		for j := 1; j < length-1; j++ {
			result_str += " "
		}
		result_str += s
		out = append(out, result_str)
		result_str = "\n"
	}
	for i := 0; i < length; i++ {
		result_str += s
	}
	out = append(out, result_str)
	result_str = "\n"

	//Запись в файл
	out_file, _ := os.Create("output.txt")
	defer out_file.Close()
	for i := 0; i < len(out); i++ {
		out_file.WriteString(out[i])
	}
}
