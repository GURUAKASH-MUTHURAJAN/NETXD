package main

import (
	"fmt"
)

func main() {
	var str string
	var temp byte

	j := 0
	fmt.Scanf("%s", &str)
	new_str := []byte(str)
	for i := 0; i < len(new_str)-1; i++ {

		if j == 2 {
			temp = new_str[i]
			new_str[i] = new_str[i+1]
			new_str[i+1] = temp
			j = -2
		}
		j++
	}
	str = string(new_str)
	fmt.Println(str)
}
