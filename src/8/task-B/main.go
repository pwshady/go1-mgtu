package main

import (
	"bufio"
	"fmt"
	"os"
)

type Sequence struct {
	message string
}

func New(newMessage string) *Sequence {
	return &Sequence{newMessage}
}

//Calc max length and return this value
func (s *Sequence) FindMax() int {
	rune_str1 := []rune(s.message)
	rune_str1 = append(rune_str1, 'p')
	var o_tec int
	var o_max int

	for i := 0; i < len(rune_str1); i++ {
		if rune_str1[i] == 'о' || rune_str1[i] == 'О' {
			o_tec++
		} else {
			if o_tec >= o_max {
				o_max = o_tec
				o_tec = 0
			}
		}
	}
	return o_max
}

func GetParam() string {
	input := bufio.NewScanner(os.Stdin)
	var str string
	if input.Scan() {
		str = input.Text()
	}
	return str
}

func main() {
	a := New(GetParam())
	fmt.Println((a.FindMax()))
}
