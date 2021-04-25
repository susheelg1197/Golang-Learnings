package main

func reverse(number int) {
	println(number)
}
func main() {

	number := []int{1, 2, 3, 4, 5}
	for _, i := range number {
		defer reverse(i)
	}
}
