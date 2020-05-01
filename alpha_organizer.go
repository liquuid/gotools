package main

import (
	"fmt"
	"os"
)

func CreateAlphaDirs(path string) {
	want := []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x", "c", "v", "b", "n", "m"}
	for _, name := range want {
		err := os.MkdirAll(path+"/"+name, 0700)
		if err != nil {
			fmt.Println(err)
		}
	}
}
