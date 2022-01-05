package main

import "github.com/KarlvenK/kinx/knet"

func main() {
	s := knet.NewServer("[kinx v0.1]")
	s.Serve()
}
