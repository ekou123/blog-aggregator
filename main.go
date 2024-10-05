package main

import (
	"example.com/sql/internal/config"
	"fmt"
)

func main() {

	configFile, err := config.Read()
	if err != nil {
		panic(err)
	}

	err = configFile.SetUser("Ethan")
	if err != nil {
		panic(err)
	}

	newConfigFile, err := config.Read()
	if err != nil {
		panic(err)
	}

	fmt.Println(newConfigFile.DbUrl)

}
