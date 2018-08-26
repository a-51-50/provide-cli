package main

import "github.com/provideservices/provide-cli/cmd"

// VERSION set during build
var VERSION = "0.0.0"

func main() {
	cmd.Execute(VERSION)
}
