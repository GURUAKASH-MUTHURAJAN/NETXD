package main

import (
	"fmt"

)

func main() {
	var h int
	fmt.Scan(&h)
	arr := make([][]int, h)
	for i := range arr {
		arr[i] = make([]int, i+1)
	}
	n:=(h/2)+(h+1)
	for i := 0; i < h; i++ {
		for k:=0;k<n;k++{
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			if i == j || j == 0 {
				arr[i][j] = 1
				fmt.Print(1," ")
			} else {
				arr[i][j]=arr[i-1][j-1]+arr[i-1][j]
				fmt.Print((arr[i-1][j-1]+arr[i-1][j])," ")
			}
		}
		n--
		fmt.Println()

	}
}
