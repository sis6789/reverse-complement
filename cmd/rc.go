package main

import (
	"bytes"
	"reflect"
)

const asciiSize = 128

var complementArray = bytes.Repeat([]byte{'*'}, asciiSize)

func reverseComplement(v []byte) {
	n := len(v)
	swap := reflect.Swapper(v)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
		//v[i], v[j] = v[j], v[i]
	}
	for i := 0; i < n; i++ {
		v[i] = complementArray[v[i]]
	}
}

func init() {
	complementArray['A'] = 'T'
	complementArray['C'] = 'G'
	complementArray['G'] = 'C'
	complementArray['T'] = 'A'
	complementArray['a'] = 't'
	complementArray['c'] = 'g'
	complementArray['g'] = 'c'
	complementArray['t'] = 'a'
}
