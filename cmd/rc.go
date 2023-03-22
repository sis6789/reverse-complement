package main

import (
	"bytes"
	"reflect"
)

const asciiSize = 128

var complementArray = bytes.Repeat([]byte{'*'}, asciiSize)

func reverseComplement(v []byte) (result []byte) {
	n := len(v)
	result = make([]byte, n)
	copy(result, v)
	swap := reflect.Swapper(result)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
		//v[i], v[j] = v[j], v[i]
	}
	for i := 0; i < n; i++ {
		result[i] = complementArray[result[i]]
	}
	return
}

func complement(v []byte) (result []byte) {
	n := len(v)
	result = make([]byte, n)
	copy(result, v)
	for i := 0; i < n; i++ {
		result[i] = complementArray[result[i]]
	}
	return
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
