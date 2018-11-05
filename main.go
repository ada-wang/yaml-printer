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
	}

	if yamlFile == "" {
		log.Fatalln("STDIN is nil, exit")
	}

	// FIX BUG - the last line is not an empty line
	if yamlFile[len(yamlFile)-1] == '\n' {
		yamlFile += "\n"
	}
	// log.Println(yamlFile[len(yamlFile)-1])
	// log.Println('\n')

	yamlMap := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(yamlFile), &yamlMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	dump, err := yaml.Marshal(&yamlMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%s", string(dump))

	// for _, arg := range os.Args {
	// 	fmt.Printf("args： %s\n", arg)
	// }
}
