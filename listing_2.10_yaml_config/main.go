package main

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	config, err := yaml.ReadFile("conf.yaml")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T\n", config)

	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))

}
