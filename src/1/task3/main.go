package main //package name

import "fmt"

func main() {
	var (
		word1 string = "имя"
		word2 string = "твое"
		word3 string = "мне"
		word4 string = "знакомо"
	)
	fmt.Println(word4, word3, word2, word1)
	fmt.Println(word3, word1, word4, word2)
	fmt.Println(word2, word4, word1, word3)
}
