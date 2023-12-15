package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/neorel/test-go/model"
)

func main() {
	res, err := http.Get("https://ergast.com/api/f1/2011/5/laps/1.json")
	if err != nil {
        fmt.Println(err)
        return
    }

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
        fmt.Println(err)
        return
    }

	fmt.Println(string(data));

	var response model.ErgastResponse
	json.Unmarshal(data, &response)

	fmt.Println(response);
}