package main

import "fmt"

func main() {
	// n := [0,1]
	n := []int{1, 4}
	fmt.Println(n, n[0], n[1])

	type Person struct {
		name string
		id   int
	}

	s := Person{name: "Uma"}
	fmt.Println(s, s.name)

	s2 := struct {
		name string
	}{
		name: "Kha",
	}
	fmt.Println(s2, s2.name)

	m1 := map[string]int{}
	fmt.Println(m1)
	m1 = map[string]int{
		"Foo": 1,
		"Aoo": 2,
		"Coo": 4,
		"zoo": 0,
	}
	fmt.Println(m1)

	// s2.foo := sdfdsjhfg  <= not possible in struct, can't create new fields

	m1["new"] = 6
	fmt.Println(m1)
	fmt.Println(m1["new"])
	v, keyExists := m1["hjsdgf"]
	fmt.Println(keyExists, v)

	// for key, val := range m1 {
	// 	fmt.Println(key, val)
	// }

	// mp := map[string]Person{
	// 	"shu": Person{name: "shu", id: 145},
	// }
	// fmt.Println(mp)

	// createMaps()
}

// {
// 	"foo": "bar",
// 	"baz": 5,
// 	"someArray": [1,2,3]
// }

// func createMaps() {
// 	//in order to create a map we have to use the `make` function
// 	// otherwise we just get a nil map
// 	//
// 	//    map[key-type]value-type
// 	var m map[string]string
// 	fmt.Println("nil map:", m)
// 	// the key is custom
// 	//
// 	// ↓↓↓ can't be outcommended, bec map is nil, we can't assignment entry to nil map ↓↓↓
// 	// m["foo"] = "bar"
// 	fmt.Println()

// 	//the map key-type can be any type which implements or has
// 	// the functions for comparising:
// 	// `==` `!=`
// 	// var intKeyMap map[int]string
// 	var intKeyMap = make(map[int]string)
// 	fmt.Println(intKeyMap)
// 	// the key is custom
// 	intKeyMap[6] = "foo"
// 	fmt.Println("intKeyMap:", intKeyMap)
// 	fmt.Println("get value from map using the key (intKeyMap[6]):", intKeyMap[6])
// 	intKeyMap[576] = "bar"
// 	intKeyMap[2] = "bazz"
// 	fmt.Println("intKeyMap:", intKeyMap)
// 	fmt.Println()

// 	//range over map
// 	fmt.Println("range over map")
// 	for key, value := range intKeyMap {
// 		fmt.Printf("key:'%d' value:'%s'\n", key, value)
// 	}
// }
