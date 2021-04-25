/*
A map is a data structure that provides you with an unordered collection
of key/value pairs (maps are also sometimes called associative arrays in
Php, hash tables in Java, or dictionaries in Python). Maps are used to
look up a value by its associated key. You store values into the map
based on a key.

A map works internally by using a hash table and which provides faster
lookup due to which we can retrive the values in a faster way by provinding key.

Map types are reference types, like pointers or slices, and so the value of m
above is nil; it doesn't point to an initialized map.
A nil map behaves like an empty map when reading, but attempts
to write to a nil map will cause a runtime panic; don't do that

*/
package main

import "fmt"

type People struct {
	name  string
	likes []string
}
type Path struct {
	path    string
	country string
}

func main() {

	var mp map[string]int // where KeyType may be any type that is comparable
	//  (more on this later), and ValueType may be any type at all,
	// including another map!
	mp = make(map[string]int) // it initializes a hashmap datastructure
	//alternative is m = map[string]int{}

	mp["India"] = 100
	j := mp["root"] // Value types zero value

	fmt.Println("Default Map value:: ", mp, "\nJ:: ", j)

	mp = map[string]int{ // Initialize multiple values
		"Australia": 50,
		"pakistan":  20,
		"Switz":     100,
	}
	// Built in map functions

	//1) Len
	fmt.Println("Length:: ", len(mp))
	//2) Delete Function
	delete(mp, "pakistan")
	// 3)Two value assignment to check presence of a key
	i, ok := mp["Australia"]
	fmt.Println("Value:: ", i, "Isthere?", ok)

	//4) Iterate over a map
	for k, v := range mp {
		fmt.Println("Key: ", k, "Value: ", v)
	}
	// Exploiting zero vlaues
	// Problem statement find people who likes icecream

	people := []*People{
		{name: "Susheel", likes: []string{"icecream", "coffee"}},
		{name: "A", likes: []string{"tea", "cream"}},
		{name: "B", likes: []string{"icecream", "tea"}},
	}

	likesmap := make(map[string][]*People)

	for _, v := range people {
		for _, s := range v.likes {
			likesmap[s] = append(likesmap[s], v)
		}
	}
	fmt.Println(likesmap["icecream"][0].name)

	// Keys in maps
	/*
	 map keys may be of any type that is comparable.
	 The language spec defines this precisely, but in short,
	 comparable types are boolean, numeric, string, pointer, channel, and
	 interface types, and structs or arrays that contain only those types.
	 Notably absent from the list are slices, maps, and functions; these types
	 cannot be compared using ==, and may not be used as map keys.
	*/

	// How struct can be used as a key
	// A use case for finding number of hits coming from various regions

	// Before a map[string]map[string]int
	// Drawback here is there was a need to check each internal path and
	// append if necessary

	// To overcome this problem a struct is used
	mapOfCountry := make(map[Path]int, 100)
	path := Path{
		path:    "/a/a/a",
		country: "IND",
	}
	mapOfCountry[path] = 20
	mapOfCountry[Path{
		path:    "/b/a/a",
		country: "Pak",
	}] = 100
	fmt.Println("NUmber of hits", mapOfCountry, "\nCapacity: ")

	// Concurrency later
}
