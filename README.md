# GO! 공부하자!
Go, 범용프로그래밍 언어인 Go에 대해 공부를 시작하며 만듭니다. 공부하며 내용을 업데이트 합니다. 

# 히스토리
- 2007년 최초 설계 : 캔 톰슨, .. 
- 2008년 본격 개발 진행

## [캔 톰슨](https://ko.wikipedia.org/wiki/%EC%BC%84_%ED%86%B0%ED%94%84%EC%8A%A8)의 참여
- AT&T 벨 연구소에서 유닉스와 플랜 9 운영 체제의 개발을 주도
-  C언어의 모체가 된 B언어를 개발
- 2006년부터 구글에서 근무
- 고 언어의 개발에 참여

# 언어의 만들어진 목적
> 범용프로그래밍 언어, 빠른 성능, 깔끔함, 간결함, 안정성, 편의성, 생산성, 쉬운 프로그래밍
- 처음에는 C 언어를 능가하는 시스템 프로그래밍 언어를 목표로 디자인된 개발언어 이지만, 더 이상 언급되지 않고 범용 언어로 소개 되고 있습니다. 
- 저수준(low-level) 프로그래밍 구조에 대한 접근을 허용하면서도 생산성을 높일 수 있는 언어

# 특징
1. 컴파일 언어
    - 인터프리터 언어가 아니기에 기계어 형태로 실행파일을 만들고 바로 실행
    - 빠름!
2. 가비지 컬렉션 (Garbage Collection) 제공
    - 실행파일 안에 가비지 컬렉터가 메로리를 관리함. 
3. 논리적 동시 처리 지원
    - 고루틴(Gorutine)
   		- 프로그램의 진입점(Entry Point) 함수를 비롯하여 다른 고루틴과 함께 동시에 실행되는 함수.
   		- 스레드 보다 적은 메모리 사용 (하나의 스레드에서 여러개의 고루틴이 사용)
    - CSP (Communicating Sequenital Processes)
    - 멀티코어 환경 지원변
4. 패키지 시스템 지원
    - 쉬운 의존성 관리
    - go get, go install 등으로 쉽게 소스코드를 가져올 수 있음
5. 편리한 모듈화
6. 규모가 크고 복잡한 프로그램 개발을 쉽게 가능하게 해줍니다.
    - 유지보수에 용이
7. 강타입 언어
	- 강타입 언어입니다. 형변환이 가능하면 약타입 언어, 가능하지 않으면 강타입 언어
```
// error
var intValue int = 10
var floatValue float32 = 1.5
var intAddfloatValue float32 = intValue + floatValue // 형변환 가능하지 않음, 컴파일 에러 발생 

invalid operation: intValue + floatValue (mismatched types int and float32)

.. 

// error
fmt.Println("string" + true)

cannot convert "string" (type untyped string) to type int
```

# 바로 웹에서 헬로우 월드 찍기!
- https://play.golang.org/
- https://golang.org/

# 설치
- https://golang.org/doc/install

## 설치 진행 (저의 기준, Mac)
1. 다운로드
	1. https://golang.org/doc/install 에서 전 맥용 설치를 찾아서..
	2. 다운로드 : https://golang.org/doc/install?download=go1.10.2.darwin-amd64.pkg
2. 실행해서 쭉쭉 진행
<img width="621" height="438" src="/Image/install_progress01.png">

3. 터미널 실행해서 설치되었는 지 확인
```
$ go version

go version go1.10.2 darwin/amd64
```

# 링크
- 깃허브
  - https://github.com/golang
  - https://github.com/golang/go

# 개발 Tool
- Visual Studio Code 가 좋은 것 같습니다. 

# 헬로우 월드 찍어보기
<img width="1014" height="607" src="/Image/helloworld00.png">

- 헬로우 월드 출력하고 시작은 반이다.. 50% 달성!
```
$ go run main.go

Hello world
```

or 

```
$ go build main.go

(main file이 만들어집니다. 윈도우라면 main.exe가 만들어지겠죠.)

$ ./main

Hello world
```



# Study Sample Code
> 헬로우 월드를 찍었으면..
1. helloworld/[main.go](https://github.com/ClintJang/go-study/blob/master/sample/helloworld/main.go) : helloworld 출력 테스트, 환경 설정 확인이죠.

## Essential Sample Code
> 기본부터 ..
1. 출력부터 해볼까나..
	- source : essental/[fmtPrint.go](https://github.com/ClintJang/go-study/blob/master/sample/essental/fmtPrint.go)
	- result
```
$ go run fmtPrint.go

printprintfprintln
print-1-printf
println
한글
6
english
日本語(にほんご)
true
true
false
false
false
10
9
10
100
한글english日本語(にほんご)
한글,english,日本語(にほんご)
경고음실행되나?
한칸만 띄워지나?
띄웅~
탭      탭
```
2. 뭐하지..

## Normal Sample Code
> 응용을 하면?
1. http 통신을 해서 json 데이터를 가져와 볼까?
- Test URL : http://ip.jsontest.com
- source : normal/[jsonDecoderTest01.go](https://github.com/ClintJang/go-study/blob/master/sample/normal/jsonDecoder/jsonDecoderTest01.go)
- result
```
{175.223.30.146}

{
    "ip": "175.223.30.146"
}
```

## Advanced Sample Code
> 언제하지?
