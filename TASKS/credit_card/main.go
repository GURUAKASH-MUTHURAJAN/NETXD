package main

import (
	"fmt"
)

func main() {
	var num, temp, count,safe, sum, oddsum int
	fmt.Scanf("%d", &num)
	temp = num
	for temp != 0 {
		temp /= 10
		count++
	}
	if count>=10{
	temp = num
	arr := make([]int, count)

	for i := 0; i <count; i++ {
		arr[i] = (temp % 10)
		temp /= 10
		if  i%2 != 0  {
			arr[i] = arr[i] * 2 
			if arr[i] > 9 {
				safe = arr[i] % 10
				arr[i] /= 10
				arr[i] = safe + arr[i]
				
			}

		}
		
		if i%2 != 0 {
			oddsum += arr[i]
		}else{
			sum += arr[i]
		}
		
	}
	sum=sum+oddsum
	
	
	if(sum%10==0){
		fmt.Println("Valid")
	}else{
		fmt.Println("Invalid")
	}

}else{
	fmt.Println("Invalid_credit card")
}
}
