package main

import (
	"fmt"
	"sort"
)

func main() {
	var name string = "45"
	//var name2 string = "456"
	fmt.Printf("Name is %X", name)
	fmt.Println("Hello mr")
	fmt.Println("mr")
	fmt.Println("The Bill Amount is")
	names := []string{"mr", "aj", "bk"}
	fmt.Println(names)
	names[2] = "vny"
	fmt.Println(names)
	names = append(names, "ss")
	fmt.Println(names[1:2], len(names))
	ages := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(sort.SearchInts(ages, 23))

}
