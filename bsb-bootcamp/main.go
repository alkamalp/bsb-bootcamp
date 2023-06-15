package main

import "github.com/alkamalp/bsb-bootcamp/utils/api"

func main() {
	app := api.Default()
	err := app.Start()
	if err != nil {
		panic(err)
	}

}
