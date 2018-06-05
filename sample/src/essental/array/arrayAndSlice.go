package main

import (
	"fmt"
)

func printSpacing(title string) {
	fmt.Println("")
	fmt.Println("-----------------------------")
	fmt.Printf("--- %s\n", title)
	fmt.Println("-----------------------------")
}

func printArray(list [2]string) {
	for index, person := range list {
		fmt.Printf("%d, %s\n", index, person)
	}
	println()
}

func printSlice(list []string) {
	for index, person := range list {
		fmt.Printf("%d, %s\n", index, person)
	}
	fmt.Printf(">> 배열 길이는 : %d, 각각의 가능한 최대 크기는 : %d", len(list), cap(list))
	println()
}

// 길이가 고정된 배열을 "배열"이라 합니다. 만들어 보자
func makeArray() {
	// 01
	var people01 [2]string = [2]string{"홍길동", "임꺽정"}
	printArray(people01) // 출력

	// 02
	var people02 [3]string = [3]string{"홍길동", "앤 해서웨이", "톰 크루즈"}
	// printArray(people02) // error, cannot use people02 (type [4]string) as type [3]string in argument to printArray
	// 출력
	for index, person := range people02 {
		fmt.Printf("%d, %s\n", index, person)
	}
	println()

	// 03, 축약
	people03 := [2]string{"임꺽정", "톰 크루즈"}
	printArray(people03) // 출력

	// 04, 축약과 간격
	people04 := [2]string{
		"임꺽정",
		"스머프", // 여러줄일 때는 마지막에 콤마
	}
	printArray(people04) // 출력

	// 05, 좀더 편의 생성
	people05 := [...]string{"임꺽정", "앤 해서웨이"}
	printArray(people05) // 출력

	// 06, 다차원 배열
	people06 := [2][2]string{
		{"임꺽정", "앤 해서웨이"},
		{"홍길동", "톰 크루즈"},
	}
	// 출력
	for rowIndex, people := range people06 {
		for columnIndex, person := range people {
			fmt.Printf("[%d][%d], %s\n", rowIndex, columnIndex, person)
		}
	}
	println()
}

// 길이가 고정되지 않고 동적으로 늘어날 수 있는 배열을 "슬라이스"라고 합니다. 만들어 보자
func makeSlice() {
	// 01
	var people01 []string = make([]string, 2, 3)
	people01[0] = "홍길동"
	people01[1] = "임꺽정"
	printSlice(people01) // 출력

	// 02
	var people02 []string = make([]string, 4, 5)
	// var people02 []string = make([]string, 4, 2) // error : len larger than cap in make([]string)
	people02[0] = "홍길동"
	people02[1] = "앤 해서웨이"
	people02[2] = "톰 크루즈"
	printSlice(people02) // 출력

	// 03
	people03 := make([]string, 2, 3)
	people03[0] = "임꺽정"
	people03[1] = "톰 크루즈"
	printSlice(people03) // 출력
}

func main() {
	printSpacing("배열 만들어보기")
	makeArray()

	printSpacing("슬라이스 만들어보기")
	makeSlice()
}
