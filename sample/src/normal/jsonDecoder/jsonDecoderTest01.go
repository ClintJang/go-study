package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	responseResult struct {
		// 퍼블릭으로, 매핑 가능하게 대문자로 시작
		Ip string `json:"ip"`
	}
)

func main() {
	// 이 url을 요청하면 테스트용 json 데이터를 응답 줍니다.
	url := "http://ip.jsontest.com"

	response, err := http.Get(url)

	if err != nil {
		log.Println("error:", err)
		return
	}

	// 함수가 끝나기 직전에는 response 를 닫는다.
	defer response.Body.Close()

	// request를 해서 response 데이터(json형태로 와요)를 받아옵니다.
	var result responseResult
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Println("error:", err)
		return
	}

	// 받은 데이터를 출력합니다.
	fmt.Println(result)
	fmt.Println("") // 필요하면 메시지를 추가로 적어서 구분

	// 4칸 space
	goodJsonViewing, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println("err:", err)
		return
	}

	// Json 형태의 결과 출력
	fmt.Println(string(goodJsonViewing))
}
