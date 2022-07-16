package main

import (
	"fmt"

	"github.com/guimpo/documentgen"
	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("Hello"))
	fmt.Println(stringutil.ToUpper("Hello"))
	cpnj, err := documentgen.CNPJ(0)
	if err == nil {
		fmt.Println(cpnj)
	}
}
