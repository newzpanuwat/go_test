package main

import (
	"fmt"
)

func test(a int){
	{
		a := a + 10
		fmt.Println(a)		
	}
	fmt.Println(a)
}

func main(){
	 test(10)
}

