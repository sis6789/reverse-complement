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
		reverseComplement(w)
		fmt.Printf("           %v\n", string(w))
	}
}
