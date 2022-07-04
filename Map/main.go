package main

import "fmt"

func main() {
	var mymap map[string]string
	mymap = make(map[string]string, 1)
	mymap["522316"] = "Guntur"
	mymap["560086"] = "yeshvantpur Bangalore"
	mymap["560096"] = "Mahalakshmi Layout, Bangalore"
	mymap["560031"] = "Rajaji Nagar,Bangalore"
	fmt.Println("Map", mymap)
	for key, value := range mymap {
		fmt.Println("Key", key, "Value", value)
	}
	fmt.Println("Length of Map", len(mymap))
	delete(mymap, "522316")
	fmt.Println("After deleting the key")
	fmt.Println("\nLength of the myMap", len(mymap))
	fmt.Println("Map", mymap)
	for key1, value1 := range mymap {
		fmt.Println("key", key1, "value", value1)
	}
	val, ok := mymap["1111"]
	if ok {
		fmt.Println("Can update an element to the map, becasue there is a key ")
	} else {
		fmt.Println("Cannot update an element but can insert an element to the map, becase ther is not key")
	}
	fmt.Println("key", ok, "value", val)
}
