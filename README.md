# GO! 공부하자!
Go, 범용프로그래밍 언어인 Go에 대해 공부를 시작하며 만듭니다. 공부하며 내용을 업데이트 합니다. 

# golang 공부해야되는 이유? (개인적인 생각)
1. 2018 스택오버플러우 설문 결과 중 "Most Loved, Dreaded, and Wanted Language" 항목에서 Loved에서는 65.6%(5위), Wanted는 16.2% (3위) .. 이때 좀 더 공부해봐야지 .. 라는 생각이 있었네요... (연봉순위 설문에서도 상위권 😀)
    - 링크 : https://insights.stackoverflow.com/survey/2018/
2. 블록체인에 대해 알아보다보면, 하이퍼레져(Hyperledger)에 대해 알게 되고.. 알아가다보면.. 체인코드(Chaincode)에 대해 접하게 되고.. 이때 golang을 해야 알아가는 데 유리하겠구나 하는 생각이 들었습니다.
3. 그 유명한 도커(Docker)의 개발 언어 랍니다.
4. 구글이 만든 언어 입니다.
5. 주위에서도 좋다고 합니다. 더 궁금해지는..

# 히스토리
- 2007년 최초 설계 : **캔 톰슨!!**, 로버트 그리즈머(V8 자바스크립트 엔진 개발에 참여했던 Robert Griesemer), 롭 파이크(UTF-8을 만든 Rob Pike) .. 
- 2008년 본격 개발 진행
- 2009년 11월에 처음으로 선보임

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
        - Go 컴파일러는 직접적으로 참조하는 라이브러리의 의존성만 해석해서 컴파일 (다른 컴파일러들은 전체 라이브러리의 의존성을 탐색)
2. 가비지 컬렉션 (Garbage Collection) 제공
    - 실행파일 안에 가비지 컬렉터가 메로리를 관리함. 
    - 약간의 오버헤드가 생길 수 있겠지만, 개발자의 노력의 감소를 얻을 수 있음.
3. 논리적 동시 처리 지원
    - 고루틴(Gorutine)
   		- 프로그램의 진입점(Entry Point) 함수를 비롯하여 다른 고루틴과 함께 동시에 실행되는 함수.
   		- 스레드와 유사하지만, 스레드 보다 적은 메모리의 사용 (하나의 스레드에서 여러개의 고루틴이 사용)
        - 많은 수의 고루틴을 시스템 부담을 최소화 하면서 만들 수 있음 (경량화된 쓰레드의 느낌)
    - CSP (Communicating Sequenital Processes)
    - 멀티코어 환경 지원
    - CPU 코어를 효과적으로 활용할 수 있는 손쉬운 방법을 제공
    - 채널(channel)은 내장된 동기화 기능을 이용해 고루틴 간에 형식화된(typed) 메시지를 공유할 수 있는 데이터 구조 라는 데, 간단히 고루틴간에 메시지를 주고받게 해주는 역활을 쉽게 해주는 것
4. 패키지 시스템 지원
    - 쉬운 의존성 관리
    - go get, go install 등으로 쉽게 소스코드를 가져올 수 있음
    - 쉬운 패키지를 임포트, 클라우드와 같은 분산된 환경에 유리함.
    ```go
    import (
        "fmt"
        "github.com/clintjang/testutil" // 실경로는 아님
    )
    ```
    - 현재의 개발 트랜드는 공유된 라이브러리를 선호하죠. (물론 검증된 것을 선호하죠..)
5. 편리한 모듈화
    - 코드 재사용성, 작은 단위의 컴퍼넌트, 패키지를 활용
6. 규모가 크고 복잡한 프로그램 개발을 쉽게 가능하게 해줍니다.
    - 유지보수에 용이
7. 강타입 언어
	- 강타입 언어입니다. 형변환이 가능하면 약타입 언어, 가능하지 않으면 강타입 언어
    - 정적 언어(static language)로서 타입의 안정성을 제공함
        - 반대로 동적 언어(dynamic language) 라면 실행 시에 잘못된 타입 때문에 발생할 수 있는 버그 들이 있을 수있음의 단점과 그에 따른 테스트 코드의 진행이 필요.
8. 단순한 언어 (개인적인..)
    - 최신트랜드를 따른 것 같으면서, 단순함이 있습니다. 
    - 키워드 사용이 적고, 코드가 간결해요.
9. 최신 아키텍쳐 개발에 용이
    - 아직은 잘은 모르겠지만, MSA(Microservice Architecture)와 REST(Representational State Transfer) 모델에도 유리하다고 합니다.

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
1. 다운로드 (현재 버전은 1.10.2 이군요..)
	1. https://golang.org/doc/install 에서 전 맥용 설치를 찾아서..
	2. 맥용 다운로드 : https://golang.org/doc/install?download=go1.10.2.darwin-amd64.pkg
    3. (만약 윈도우즈라면)윈도우즈 다운로드 : https://golang.org/doc/install?download=go1.10.2.windows-amd64.msi
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
1. source : src/helloworld/[helloworld.go](https://github.com/ClintJang/go-study/blob/master/sample/src/helloworld/helloworld.go) : helloworld 출력 테스트, 환경 설정 확인이죠.

## Essential Sample Code
> 기본부터 ..
1. 출력부터 해볼까나..
	- source : essental/[fmtPrint.go](https://github.com/ClintJang/go-study/blob/master/sample/src/essental/print/fmtPrint.go)
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
- source : normal/jsonDecoder/[jsonDecoderTest01.go](https://github.com/ClintJang/go-study/blob/master/sample/src/normal/jsonDecoder/jsonDecoderTest01.go)
- result
```
{175.223.30.146}

{
    "ip": "175.223.30.146"
}
```

## Advanced Sample Code
> 언제하지?
