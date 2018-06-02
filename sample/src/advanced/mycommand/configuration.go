package main

import (
	"flag"
	"fmt"
)

// Configuration 구조체에 플래그 값을 저장
type Configuration struct {
	// 스트링 입력값 테스트, 결과 타이틀 설정
	title string
	// 불 변수 테스트, 개발자 여부
	isDeveloper bool
	// 정수 테스트, 나이값을 표현, 기본 값 설정도 확인
	age int
	// 값을 "1,2,3" 으로 입력 받아서 모두를 더할 수 있는 기능을 제공
	sum AddValuesOfAll
}

// 해당 커멘드의 플래그 값을 전달된 플래그 값으로 초기화
func (c *Configuration) FlagEnvironmentSetting() {
	// 서술형 플래그, 기본은 빈값
	flag.StringVar(&c.title, "title", "", "title is a string, it default to empty")
	// 축약형 플래그, 기본은 빈값
	flag.StringVar(&c.title, "t", "", "title is a string, it defaults to empty (shorthand)")

	// 서술형 플래그, 기본은 false
	flag.BoolVar(&c.isDeveloper, "isD", false, "bool Set test")
	// 축약형 플래그, 기본은 false
	flag.BoolVar(&c.isDeveloper, "isDeveloper", false, "bool Set test (shorthand)")

	// 플래그, 기본값은 -1
	flag.IntVar(&c.age, "age", -1, "It is my age value.")

	// 플래그, 값을 파싱해서 배열에 넣고 모두 더한다.
	flag.Var(&c.sum, "sum", "comma separated list of integers, add the values of all")
}

// 처리된 결과 메시지
func (c *Configuration) GetResultMessage() string {
	var returnValue string

	// 타이틀
	title := c.title
	if len(title) == 0 {
		title = "No title is set."
	}

	// 개발자인지?
	var isDeveloperString string = ""
	if c.isDeveloper {
		isDeveloperString += "Are you a developer?\n Yes. I am\n"
	} else {
		isDeveloperString += "Are you a developer?\n Yes. I'm not\n"
	}

	// 나이
	var ageString = "What is your age?"

	// 결과 메시지 종합
	returnValue = fmt.Sprintf("%s\n%s\n%s\n%d\n, age is -1 means no value is set.\n%s\n",
		title, isDeveloperString, ageString, c.age, c.sum.String())
	return returnValue
}
