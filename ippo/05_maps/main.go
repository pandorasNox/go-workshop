package main

import "fmt"

func main() {
	createMaps()
}

func createMaps() {
	//in order to create a map we have to use the `make` function
	// otherwise we just get a nil map
	//
	//    map[key-type]value-type
	var m map[string]string
	fmt.Println("nil map:", m)
	// the key is custom
	//
	// ↓↓↓ can't be outcommended, bec map is nil, we can't assignment entry to nil map ↓↓↓
	// m["foo"] = "bar"
	fmt.Println()

	//the map key-type can be any type which implements or has
	// the functions for comparising:
	// `==` `!=`
	// var intKeyMap map[int]string
	var intKeyMap = make(map[int]string)
	fmt.Println(intKeyMap)
	// the key is custom
	intKeyMap[6] = "foo"
	fmt.Println("intKeyMap:", intKeyMap)
	fmt.Println("get value from map using the key (intKeyMap[6]):", intKeyMap[6])
	intKeyMap[576] = "bar"
	intKeyMap[2] = "bazz"
	fmt.Println("intKeyMap:", intKeyMap)
	fmt.Println()

	//range over map
	fmt.Println("range over map")
	for key, value := range intKeyMap {
		fmt.Printf("key:'%d' value:'%s'\n", key, value)
	}
}
