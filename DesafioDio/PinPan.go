package main

import "fmt"

func main() {
	for x:= 0; x <= 100; x++ {
		if x%3 == 0 && x%5 == 0 {
		fmt.Println("PinPan")
	}else if x%5 == 0 {	
		fmt.Println("Pan")
	}else if x%3 == 0{ 	
		fmt.Println("Pin")
	}else{
	fmt.Println(x)	
		}
	}
}