package main

import "fmt"

func main() {

	fmt.Print("print")
	fmt.Printf("printf")
	fmt.Println("println")
	fmt.Print("print")
	fmt.Printf("-%d-printf\n", 1)
	fmt.Println("println")

	var kor string = "한글"
	var eng string = "english"
	var jp string = "日本語(にほんご)"
	jpSame := "日本語(にほんご)" // string
	var isBool bool = true
	isBoolSame := false
	var intValue int = 10
	intValueSame := 10*5 + 25*2

	// string
	fmt.Println(kor)
	fmt.Println(len(kor))
	fmt.Println(eng)
	fmt.Println(jp)

	// bool
	fmt.Println(jp == jpSame)
	fmt.Println(isBool)
	fmt.Println(!isBool)
	fmt.Println(isBool == isBoolSame)
	fmt.Println(intValue > intValueSame)

	// int
	fmt.Println(intValue)
	intValue--
	fmt.Println(intValue)
	intValue++
	fmt.Println(intValue)
	// fmt.Println(intValue++) // error
	fmt.Println(intValueSame)
	// --intValue	// error
	// ++intValue	// error

	// 문자열 붙이기
	fmt.Println(kor + eng + jp)
	fmt.Println(kor + "," + eng + "," + jp)

	// 제어문자
	fmt.Println("경고음\a실행되나?")
	fmt.Println("한칸만 띄워지나?\n띄웅~")
	fmt.Println("탭\t탭")
}
