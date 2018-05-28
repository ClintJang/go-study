package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	responseResult struct {
		Ip string `json:"ip"`
	}
)

func main() {
	url := "http://ip.jsontest.com"

	resp, err := http.Get(url)

	if err != nil {
		log.Println("error:", err)
		return
	}

	defer resp.Body.Close()

	var rs responseResult
	err = json.NewDecoder(resp.Body).Decode(&rs)
	if err != nil {
		log.Println("error:", err)
		return
	}

	// 받은 데이터
	fmt.Println(rs)
	fmt.Println("")

	// 4칸 space
	goodJsonViewing, err := json.MarshalIndent(rs, "", "    ")
	if err != nil {
		log.Println("err:", err)
		return
	}

	fmt.Println(string(goodJsonViewing))
}
