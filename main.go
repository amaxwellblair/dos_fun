package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := "http://universe.rejs.io/api/v1/data"

	request := func() {
		req, _ := http.NewRequest("GET", url, nil)
		resp, _ := http.DefaultClient.Do(req)
		if resp.Status == "200" {

		}
	}

	for index := 0; index < 1000; index++ {
		go request()
		fmt.Println(index)
	}

	time.Sleep(20 * time.Second)
	// fmt.Println(ioutil.ReadAll(resp.Body))
}
