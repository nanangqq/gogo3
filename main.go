package main

import (
	"fmt"

	"github.com/nanangqq/gogo3/dict"
)

func main() {
	dictionary := dict.Dictionary{}
	// dictionary["hello"] = "hello"
	fmt.Println(dictionary)

	def, err := dictionary.Search("hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(def)
	}

	err = dictionary.Add("hello", "hi")
	if err != nil {
		fmt.Println(err)
	}

	def, err = dictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(def)
	}
	fmt.Println(dictionary)

	baseWord := "test"
	dictionary.Add(baseWord, "first")
	err = dictionary.Update(baseWord, "second")
	if err != nil {
		fmt.Println(err)
	}
	value, _ := dictionary.Search(baseWord)
	fmt.Println(value)

	dictionary.Delete(baseWord)
	fmt.Println(dictionary)
}
