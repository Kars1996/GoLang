package main;

// this file won't run btw lolz

import "fmt";

// avaliable types:
// int, string, bool, float64 and sum other like bs like int64 and shit

func main () {
	var a int = 4 // define varaible -> name -> type -> default value
	var b, c int // define variable -> names -> same type for em both

	var d string = "Strings must use double thingies to let shit work i guesss";

	var e bool = true // kinda like typescript (absolute peak language) in terms of true/false

	f := 9 // same as var f int = 9
	g := "uwu" // same as var g string = "uwu"
	// and so on

	fmt.Println("I wasted a lot of time on this")
	
	_ = a +b + c + d +e + f 

	fmt.Printf("I THINK the syntax is `%d` I belive idr", g) // assuming i got it rights its kinda like an f string in python or a ${} thing in typescript
}
