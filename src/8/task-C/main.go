package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type SmartList struct {
	words []string
}

func New(newWords []string) *SmartList {
	return &SmartList{newWords}
}

//Sort words and printed it
func (sl *SmartList) GetAnswer() {
	result := []string{}
	for i := 0; i < 100; i++ {
		for j := 0; j < len(sl.words); j++ {
			if len(sl.words[j]) == i {
				result = append(result, sl.words[j])
			}
		}
	}
	sl.words = result
}

func (sl *SmartList) Pr() {
	for i := 0; i < len(sl.words); i++ {
		fmt.Println(sl.words[i])
	}
	//println(sl.words[0])
	//println(len(sl.words))
}

func GetParam() string {
	input := bufio.NewScanner(os.Stdin)
	var str string
	if input.Scan() {
		str = input.Text()
	}
	return str
}

func GetIntParam() (int, error) {
	input := bufio.NewScanner(os.Stdin)
	var str string
	if input.Scan() {
		str = input.Text()
	}
	return strconv.Atoi(str)
}

func main() {
	s := []string{}
	count, _ := GetIntParam()
	for i := 0; i < count; i++ {
		s = append(s, GetParam())
	}
	sort.Strings(s[:])
	sl := New(s[:])
	sl.GetAnswer()
	sl.Pr()
	print()
}
