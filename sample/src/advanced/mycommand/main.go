package main

import (
	"flag"
	"fmt"
)

func main() {
	// 프로그램은 문제가 없지만, 메모리에 할당과 인스턴스로 할당 (그냥 테스트)
	// 1.
	// var c Configuration = Configuration{}
	// 2.
	// var c *Configuration   // 구조체 선언
	// c = new(Configuration) // 구조체 포인터에 메모리 할당
	// 3.
	// c := new(Configuration)
	// 4.
	c := Configuration{}

	// 해당 커멘드의 플래그 값을 전달된 플래그 값으로 초기화 합니다.
	c.FlagEnvironmentSetting()

	// 이 함수는 메인에서 사용되며, 들어온 인자값을 파서 하는 것 같음.
	flag.Parse()

	// 그리고 처리된 결과를 터미널에 출력합니다.
	fmt.Println(c.GetResultMessage())
}
