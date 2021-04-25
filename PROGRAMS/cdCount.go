package main

import (
	"fmt"
)

func dirCounter(str string) {
	nodes := make([]string, 200)
	nodes = stringSplit(str, "/")
	// nodes := strings.Split(str, "/")
	path := make([]string, len(nodes))
	dirCounter := make(map[string]int)

	for _, val := range nodes {
		if val != ".." {
			path = append(path, val)
		} else {
			path = path[:len(path)-1]
		}
		if path[len(path)-1] == "" {
			dirCounter["Root"] += 1

		} else {
			dirCounter[path[len(path)-1]] += 1

		}
	}
	fmt.Println(dirCounter)
}

func main() {

	str := "a/b/../c/d/e/../../../../"
	dirCounter(str)

}
func stringSplit(s, sep string) []string {

	r := []rune(s)
	finalStr := make([]string, 0, len(r))
	fmt.Println(cap(finalStr))
	for i := 0; i < len(r); i++ {
		if sep == string(r[i]) {
			finalStr = append(finalStr, string(r[:i]))
			r = r[i+1:]
			fmt.Println(finalStr, i)
			i = 0
		}
	}

	return finalStr
}
