package main

import (
	"fmt"
	"strings"
	"os"

	"github.com/alexflint/go-arg"

)

var args struct {
	I bool
	V bool
	C bool
	W bool
	Pattern string `arg:"positional"`
	Input string `arg:"positional"`
}


func IndexOfSubstring(str, subStr string) []int {
	ind := []int{} 
	fmt.Println(subStr)
	fmt.Println(str)
	fmt.Println(len(subStr))
	for i := 0; i < (len(str)-len(subStr)); i++ {
	   if str[i:i+len(subStr)] == subStr {
		ind = append(ind, i)
	}
	}
	return ind
}



func main() {
	arg.MustParse(&args)
	file:=args.Input
	dat, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	pattern := args.Pattern
	if(args.I){
		index := IndexOfSubstring(strings.ToLower(string(dat)), strings.ToLower(pattern))
		fmt.Println(index)
		fmt.Println("Match found at index number", index)
	}
	if(args.V){
		str:=string(dat)
		res:=strings.ReplaceAll(str,pattern,"")
		fmt.Println(res)
	}
	if(args.C){
		index := IndexOfSubstring(strings.ToLower(string(dat)), strings.ToLower(pattern))
		fmt.Println(len(index))
	}
	if(args.W){
		ind := []int{} 
		arr:=strings.Split(string(dat), " ")
		for i := 0; i < len(arr); i++{
			if(arr[i] == pattern){
				ind = append(ind,i)
			}
		}
		fmt.Println("Exact word found at index",ind)
	}
}