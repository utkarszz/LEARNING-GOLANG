package main

import "fmt"

func main(){

	//  Basic constants in GoLang
	const pi = 3.14

	fmt.Println(pi)

	//  Explicitly typed constants
	const age  int =20
	const name string = "Utkarsh Singh"
	const isStudent bool = true

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(isStudent)

	//  Multiple Constants Declaration const :You can declare multiple constants inside a single const block.
const (
	x=10
	y=20
	z=30
)
fmt.Println(x,y,z)

// Constants with iota constant
const (
	a = iota
	b
	c
	d
)
fmt.Println(a,b,c,d)
}