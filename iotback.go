// iotback project iotback.go
package main

import "fmt"

// Initialize the application and file configuration so they are globally accessible
var c aConfig
var f []string

func init() {
	// Set values to defaults then get or set
	c = aConfig{username: "demo", password: "demo", machine: RandStringBytesMaskImprSrc(8)}
	c = aConfig.getaConfig(c)
}

func main() {

	fmt.Println(c)
}
