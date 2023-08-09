package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	j := 0
	read := bufio.NewReader(os.Stdin)
	str,_ = read.ReadString('\n')
	for i := 0; i < len(str); i++ {
		if str[i] == 'g' || str[i] == 'G' {
			if str[i+1] == 'o' || str[i+1] == 'O' {
				j++
			}
		}
	}
	fmt.Println(j)
	fmt.Println(str)
}
