package main

import (
	"fmt"
	"github.com/genekuo/hello/tempconv"
)

func main() {
	//fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	//fmt.Printf("Brrrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}