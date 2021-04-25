/*
Example in http requests

*/

package main

import "context"

func main() {
	rootContext := context.Background()                             // empty context
	childctx := context.WithValue(rootContext, "requestId", "1234") // extension to that context

	child2ctx, cancelFunc := context.WithCancel(childctx)

	childr3ctx := context.WithValue(rootContext, "userId", "5678")

}
