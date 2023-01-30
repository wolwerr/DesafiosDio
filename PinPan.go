package main

import "fmt"

func main() {
	for x:= 0; x <= 100; x++ {
		if x%3 == 0 && x%5 == 0 {
		fmt.Println("PingPang")
	}else if x%5 == 0 {	
		fmt.Println("Pang")
	}else if x%3 == 0{ 	
		fmt.Println("Ping")
	}else{
	fmt.Println(x)	
		}
	}
}