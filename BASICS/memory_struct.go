package main

import (
	"fmt"
	"unsafe"
)

// Before optimization
type Post struct {
	published bool  // 1 byte
	title     int32 // 16 Bytes
	content   bool  // 16 Bytes
	// likes     int16  // 2 bytes
	// 35 bytes total, 13 bytes padding
}

// After optimization
// type Post struct {
// 	title     string // 16 Bytes
// 	content   string // 16 Bytes
// 	likes     int16  // 2 bytes
// 	published bool   // 1 byte

// 	// 35 bytes total, 13 bytes padding
// }

// Because the padding is optimized, when the struct tries to allocate bytes
// It first checks for the space required for that
func main() {
	v := Post{}
	fmt.Printf("Sizeof Post: %d\n", unsafe.Sizeof(v))

	fmt.Printf("Sizeof Post: %d\n", unsafe.Offsetof(v.title))
	fmt.Printf("Sizeof Post: %d\n", unsafe.Offsetof(v.content))
	fmt.Printf("Sizeof Post: %d\n", unsafe.Offsetof(v.published))
}
