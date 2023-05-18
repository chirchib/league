package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"league/configs"
	"league/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	envFile := flag.String("env", "", "env for use")
	flag.Parse()

	if *envFile != "" {
		fmt.Printf("env is used %s \n", *envFile)

		err := godotenv.Load(*envFile)
		if err != nil {
			fmt.Println(fmt.Errorf("env parse err %s", err.Error()).Error())
			os.Exit(-1)
		}
	}

	fmt.Printf("version: %s \n", os.Getenv("VERSION"))

	if _, err := app.NewApp(configs.NewConfig()).Start(); err != nil {
		log.Fatal(err)
	}
}
