package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (app *application) getMiner() {

	response, err := http.Get("https://api.ethermine.org/miner/0x4462819ab1f460b08087ec20778a6ad94f953f3a/dashboard")
	//response, err := http.Get("https://httpbin.org/get")

	if err != nil {
		app.errorLog.Fatal("Endpoint unreachable:", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		app.errorLog.Fatal("unable to read body:", err)
	}

	fmt.Println("data from api:", string(body))
}
