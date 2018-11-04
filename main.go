package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var yamlFile string

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		yamlFile += line
		// fmt.Printf(yamlFile)
	}

	yamlMap := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(yamlFile), &yamlMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	// fmt.Printf("--- yamlMap:\n%v\n\n", yamlMap)

	dump, err := yaml.Marshal(&yamlMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%s", string(dump))

	// for _, arg := range os.Args {
	// 	fmt.Printf("args： %s\n", arg)
	// }
}
