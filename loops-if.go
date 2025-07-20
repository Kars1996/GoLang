package main;

import "fmt";

func main() {
	name := "Kars";

	if name == "Kars" {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d. Hello %s\n", i + 1, name); // use %s to expeect a string apparently
		}
	} else {
		fmt.Println("Fuck off twat");
	}

	coolBeans := [5]string {"Me", "Myself", "I", "MeSelf", "ISelf"}

	for index, value := range coolBeans {
		fmt.Printf("%s is the %dst coolest bean\n", value, index)
	}
}
