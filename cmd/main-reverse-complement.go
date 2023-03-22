package main

import "fmt"

func main() {
	for {
		var seq string
		fmt.Print("enter nucs:")
		readLen, err := fmt.Scanln(&seq)
		if err != nil || readLen == 0 {
			break
		}
		w := []byte(seq)
		vRC := reverseComplement(w)
		vC := complement(w)
		fmt.Printf("         c:%v\n", string(vC))
		fmt.Printf("        rc:%v\n", string(vRC))
	}
}
