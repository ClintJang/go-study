[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/ugorji/go/codec)
[![rcard](https://goreportcard.com/badge/github.com/ugorji/go/codec?v=2)](https://goreportcard.com/report/github.com/ugorji/go/codec)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/clintjang/go-study/blob/master/LICENSE)

# GO! 공부하자!
Go, 범용프로그래밍 언어인 Go에 대해 공부를 시작하며 만듭니다. 공부하며 내용을 업데이트 합니다. 

# golang 공부해야되는 이유? (개인적인 생각)
1. 2018 스택오버플러우 설문 결과 중 "Most Loved, Dreaded, and Wanted Language" 항목에서 Loved에서는 65.6%(5위), Wanted는 16.2% (3위) .. 이때 좀 알아봐야지 .. 라는 생각이 있었던 것 같네요.... (연봉순위 설문에서도 상위권 😀)
    - 링크 : https://insights.stackoverflow.com/survey/2018/
    - 2016년 부터 유명해 지기 시작..
2. 블록체인에 대해 알아보다보면, 하이퍼레져(Hyperledger)에 대해 알게 되고.. 알아가다보면.. 체인코드(Chaincode)에 대해 접하게 되고.. 이때 golang을 해야 알아가는 데 유리하겠구나 하는 생각이 들었습니다.
    - 이더리움의 메인 클라이언트인 Geth -> Go로 작성
3. 그 유명한 도커(Docker)의 개발 언어 랍니다.
4. 구글이 만든 언어 입니다.
5. 주위에서도 좋다고 합니다. 더 궁금해지는.. 공부해본 개발자 분들은 더욱 칭찬을.. 왜?
6. 심지어.. 모바일(android, iOS)도 개발이 가능하게 하려는 계획이 있다고 한다.. 덜덜 
    - 링크 : https://github.com/golang/go/wiki/Mobile

# 히스토리
- 2007년 최초 설계 : **캔 톰슨!!(Kenneth Thompson)**, 로버트 그리즈머(V8 자바스크립트 엔진 개발에 참여했던 Robert Griesemer), 롭 파이크(UTF-8을 만든 Rob Pike) .. 
- 2008년 본격 개발 진행
- 2009년 11월에 처음으로 선보임

## [캔 톰슨](https://ko.wikipedia.org/wiki/%EC%BC%84_%ED%86%B0%ED%94%84%EC%8A%A8)의 참여
- AT&T 벨 연구소에서 유닉스와 플랜 9 운영 체제의 개발을 주도
-  C언어의 모체가 된 B언어를 개발
- 2006년부터 구글에서 근무
- 고 언어의 개발에 참여

# 언어의 만들어진 목적
> **범용프로그래밍 언어**, 빠른 성능, 깔끔함, 간결함, 안정성, 편의성, 생산성, 쉬운 프로그래밍
- 범용 언어
    - 처음에는 C 언어를 능가하는 시스템 프로그래밍 언어를 목표로 디자인된 개발언어 이지만, 더 이상 언급되지 않고 범용 언어로 소개
- 저수준(low-level) 프로그래밍 구조에 대한 접근을 허용하면서도 생산성을 높일 수 있는 언어

# 특징
1. 컴파일 언어
    - 인터프리터 언어가 아니기에 기계어 형태로 실행파일을 만들고 바로 실행
    - 빠름!
        - Go 컴파일러는 직접적으로 참조하는 라이브러리의 의존성만 해석해서 컴파일 (다른 컴파일러들은 전체 라이브러리의 의존성을 탐색)
2. 가비지 컬렉션 (Garbage Collection) 제공
    - 실행파일 안에 가비지 컬렉터가 메로리를 관리함. 
    - 약간의 오버헤드가 생길 수 있겠지만, 개발자 노력?ㅋ의 감소를 얻을 수 있음.
3. 논리적 동시 처리 지원
    - 고루틴(Gorutine)
   		- 프로그램의 진입점(Entry Point) 함수를 비롯하여 다른 고루틴과 함께 동시에 실행되는 함수.
   		- 스레드와 유사하지만, 스레드 보다 적은 메모리의 사용 (하나의 스레드에서 여러개의 고루틴이 사용)
        - 많은 수의 고루틴을 시스템 부담을 최소화 하면서 만들 수 있음 (경량화된 쓰레드의 느낌)
        - 비동기 메커니즘인? [Erlang](https://namu.wiki/w/Erlang)에서 영향을 받음 : 함수형 병행성 프로그래밍 언어
    - CSP (Communicating Sequenital Processes)
    - 멀티코어 환경 지원
    - CPU 코어를 효과적으로 활용할 수 있는 손쉬운 방법을 제공
    - 채널(channel)은 내장된 동기화 기능을 이용해 고루틴 간에 형식화된(typed) 메시지를 공유할 수 있는 데이터 구조 라는 데, 간단히 고루틴간에 메시지를 주고받게 해주는 역활을 쉽게 해주는 것
4. 패키지 시스템 지원
    - 쉬운 의존성 관리
    - go get, go install 등으로 쉽게 소스코드를 가져올 수 있음
    - 쉬운 패키지를 임포트, 클라우드와 같은 분산된 환경에 유리함.
    - 아래는 go get 등으로 가져와서 사용하는 경우.
        - 터미널에서 쉘명령어로 
        ```
        $ go get github.com/clintjang/go-study/
        ```
        - 그러면 아래 루트에서 필요한 패키지가 있는 경로를 지정해서 사용할 수 있음
        ```go
        import (
            "fmt"
            // 실제는 $GOPATH/src/github.com/clintjang/원하는패키지
            "github.com/clintjang/원하는패키지" // 실경로는 아님
        )
        ```
    - $GOPATH 에 프로젝트 설정으로 다양하게 사용 가능
    - 현재의 개발 트랜드는 공유된 라이브러리를 선호하죠. (물론 검증된 것을 선호하죠..)
5. 편리한 모듈화
    - 코드 재사용성, 작은 단위의 컴퍼넌트, 패키지를 활용
6. 규모가 크고 복잡한 프로그램 개발을 쉽게 가능함
    - 유지보수에 용이
    - 최신 아키텍쳐 개발에 용이
        - 아직은 잘은 모르겠지만, MSA(Microservice Architecture)와 REST(Representational State Transfer) 모델에도 유리하다고 함
7. 강타입 언어
	- 강타입 언어입니다. 형변환이 가능하면 약타입 언어, 가능하지 않으면 강타입 언어
    - 정적 언어(static language)로서 타입의 안정성을 제공함
        - 반대로 동적 언어(dynamic language) 라면 실행 시에 잘못된 타입 때문에 발생할 수 있는 버그 들이 있을 수있음의 단점과 그에 따른 테스트 코드의 진행이 필요.
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

8. 단순한 언어 (개인적인..)
    - 최신트랜드를 따른 것 같으면서, 단순함이 있음.
    - 키워드 사용이 적고, 코드가 간결함.

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
- [vim-go](https://github.com/fatih/vim-go) 라는 도구도 있습니다.
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

# 시작시 알아두면 좋은 팁
1. go env
    1. 프로젝트에 기본이 되는 $GOPATH 에 대해 알고 시작하면 좋습니다. 
    ```
    $ go env
    >> 아래와 같은 로그를 볼 수 있고, GOPATH의 설정 위치를 알 수 있습니다. 

    GOARCH="amd64"
    GOBIN=""
    GOCACHE="/Users/clintjang/Library/Caches/go-build"
    GOEXE=""
    GOHOSTARCH="amd64"
    GOHOSTOS="darwin"
    GOOS="darwin"
    GOPATH="/Users/clintjang/go"        << 여기!
    GORACE=""
    GOROOT="/usr/local/go"
    GOTMPDIR=""
    GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
    GCCGO="gccgo"
    CC="clang"
    CXX="clang++"
    CGO_ENABLED="1"
    CGO_CFLAGS="-g -O2"
    CGO_CPPFLAGS=""
    CGO_CXXFLAGS="-g -O2"
    CGO_FFLAGS="-g -O2"
    CGO_LDFLAGS="-g -O2"
    PKG_CONFIG="pkg-config"
    GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/zs/8q7f_x6d4zb6jsy_dpvqc5m80000gn/T/go-build886547065=/tmp/go-build -gno-record-gcc-switches -fno-common"
    ```
    2. 개발 소스는 "$GOPATH/src/작업패키지폴더들/작업소스" 위치에서 작업하면 좋은 것 같습니다.
    3. GOPATH를 변경하려면.. 아래를 실행
    ```
    $ export GOPATH=/원하는경로1/원하는경로2/원하는경로3/../프로젝트명
    ```
2. go build
    1. 해당폴더에 있는 main.go 를 찾아서 컴파일 합니다.
    2. 해당 폴더명으로 실행 가능한 파일을 만듭니다.
    3. 그 파일은 컴파일한 해당 OS에서 실행가능한 파일입니다.
3. go에서 사용되는 메모리 표현중에.. ( 공부중이라 틀릴수도 있어요. )
    - 구조체라면 메모리에 할당하는 방식과 인스턴스 형태의 방식으로 할당할 수 있습니다. 
    - 메모리에 할당하는 방식은 포인터를 이용하여 참조 방식(by referance)으로 사용하고
    ```go
    // 예로 이런 형태의 선언
    var temp *TempStruct = new(TempStruct)
    temp2 := new(TempStruct2)
    ```
    - 인스턴스 라고 표현한 것은 복사 방식(by value)으로 이해하고 있고, 참조 방식으로 처리하려면 "&"를 이용해서 직접 메모리의 참조 데이터를 넘길 수도 있습니다. 초기값 설정도 가능합니다. 이 용어도 틀릴 수 있습니다.
    ```go
    // 예로 이런 형태의 선언
    var temp TempStruct = TempStruct{}
    var temp2 TempStruct2 = TempStruct2{20, 100}
    var temp3 TempStruct3 = TempStruct2{x: 20, y: 100}
    temp4 := TempStruct4{}

    // 참조 형태로 사용할때는 
    modify(&temp2, 30, 110)
    fmt.Println(temp2)  // 30, 110

    // 복사가 됨
    modify(temp3, 30, 110)
    fmt.Println(temp3) // 20, 100
    ```

# Study Sample Code
> 헬로우 월드를 찍었으면..
1. source : src/helloworld/[helloworld.go](https://github.com/ClintJang/go-study/blob/master/sample/src/helloworld/helloworld.go) : helloworld 출력 테스트, 환경 설정 확인이죠.
2. Essential Sample Code : 기본 문법을 이해하기 위한 셈플
3. Normal Sample Code : 기본 문법을 이해하기 위해 약간 응용한 셈플
4. Advanced Sample Code : 실제 사용될 간단한 수준의 셈플 프로젝트

## Essential Sample Code
> 기본부터 ..
1. 출력부터!!
	1. source : essental/print/[fmtPrint.go](https://github.com/ClintJang/go-study/blob/master/sample/src/essental/print/fmtPrint.go)
	2. result
    ```
    $ go run fmtPrint.go

    -----------------------------
    --- Print, Prinff, Println
    -----------------------------
    printprintfprintln
    print-1-printf
    println

    -----------------------------
    --- string
    -----------------------------
    한글
    6
    english
    日本語(にほんご)

    -----------------------------
    --- bool
    -----------------------------
    true
    true
    false
    false
    false

    -----------------------------
    --- int
    -----------------------------
    10
    9
    10
    100

    -----------------------------
    --- 문자열 붙이기
    -----------------------------
    한글english日本語(にほんご)
    한글,english,日本語(にほんご)

    -----------------------------
    --- 제어문자
    -----------------------------
    경고음실행되나?
    한칸만 띄워지나?
    띄웅~
    탭      탭
    ```

2. 배열을 만들어보기, 추가로 슬라이스의 개념은?
    > 배열 : 길이가 고정된 정적 배열<br />슬라이스 : 길이가 고정되지 않고 동적으로 늘어날 수 있는 배열
    1. source : essental/array/[arrayAndSlice.go](https://github.com/ClintJang/go-study/blob/master/sample/src/essental/array/arrayAndSlice.go)
	2. result
    ```
    -----------------------------
    --- 배열 만들어보기
    -----------------------------
    0, 홍길동
    1, 임꺽정

    0, 홍길동
    1, 앤 해서웨이
    2, 톰 크루즈

    0, 임꺽정
    1, 톰 크루즈

    0, 임꺽정
    1, 스머프

    0, 임꺽정
    1, 앤 해서웨이

    [0][0], 임꺽정
    [0][1], 앤 해서웨이
    [1][0], 홍길동
    [1][1], 톰 크루즈


    -----------------------------
    --- 슬라이스 만들어보기
    -----------------------------
    0, 홍길동
    1, 임꺽정
    >> 배열 길이는 : 2, 각각의 가능한 최대 크기는 : 3
    0, 홍길동
    1, 앤 해서웨이
    2, 톰 크루즈
    3,
    >> 배열 길이는 : 4, 각각의 가능한 최대 크기는 : 5
    0, 임꺽정
    1, 톰 크루즈
    >> 배열 길이는 : 2, 각각의 가능한 최대 크기는 : 3
    ```

3. 언제하지? 시간이 되면 셈플을 추가해보자~

## Normal Sample Code
> 응용을 하면?
1. http 통신을 해서 json 데이터를 가져와 볼까?
    1. Test URL : http://ip.jsontest.com
    2. source : normal/jsonDecoder/[jsonDecoderTest01.go](https://github.com/ClintJang/go-study/blob/master/sample/src/normal/jsonDecoder/jsonDecoderTest01.go)
    3. result
    ```
    {175.223.30.146}

    {
        "ip": "175.223.30.146"
    }
    ```
2. 파일 입출력 
	> 파일과 폴더를 만들고 지워보자, 만들어진 파일에 문자열을 쓰거나 파일 내용을 복사해보자.
	
	1. source : normal/fileAndDir/[main.go](https://github.com/ClintJang/go-study/blob/master/sample/src/normal/fileAndDir/main.go)
	2. result : fileAndDir 폴더에서 main.go를 run합니다.
	```
    fileAndDir $ go run main.go
    
	폴더 : makeDirFolder01, 02 는 실행전에 지우세요.
    파일 : testFileMake02.txt 도 실행전에 지우세요.

    hello world 01

    만들어진 폴더와 파일을 확인해보세요~
    ```
    그리고 폴더와 파일을 지우고 다시 실행하면..
    ```
    fileAndDir $ go run main.go
    
    폴더 : makeDirFolder01, 02 는 실행전에 지우세요.
    파일 : testFileMake02.txt 도 실행전에 지우세요.

    에러가 나도 종료하지 않습니다.
    폴더가 없을 수 있으니깐요..
    이 프로그램은 공부용 테스트 프로그램이에요.
    처음만 발생합니다.
    err:%s (0x109fbe0,0xc42008e1e0)
    hello world 01

    만들어진 폴더와 파일을 확인해보세요~
	```

	... if 문을 많이 써야 되는 부분이 아쉽네요. 
 
3. 초간단 HTTP 서버
	1. source : normal/httpServer/[main.go](https://github.com/ClintJang/go-study/blob/master/sample/src/normal/httpServer/main.go)
	2. result : httpServer 폴더에서 main.go를 run합니다.
	```
    httpServer $ go run main.go
    _
    (커서 대기중, 이제 웹 브라우저에 "localhost:9999" 를 쳐보세요)
    ```
	3. 간단한 서버가 실행됬습니다. (Http Server를 닫으려면 control + c)
	4. 서버 실행 : http://localhost:9999
	<img width="444" height="166" src="/Image/normalHttpServer00.png">
4. 언제하지? 시간이 되면 셈플을 추가해보자~

## Advanced Sample Code
> 실제 사용할만 한것은?

1. 콘솔에서 사용할 수 있는 간단한 커멘드라인 도구를 만들어보자
    1. flag 패키지 : 커맨드라인 플래그 인수를 쉽게 추가 할 수 있음. 활용해보자
    2. source
        1. advanced/mycommand/[main.go](https://github.com/ClintJang/go-study/blob/master/sample/src/advanced/mycommand/main.go) : 개발 된 프로그램의 시작 설정
        2. advanced/mycommand/[configuration.go](https://github.com/ClintJang/go-study/blob/master/sample/src/advanced/mycommand/configuration.go) : flag 패키지의 환경설정 및 출력 설정
        3. advanced/mycommand/[addvaluesofall.go](https://github.com/ClintJang/go-study/blob/master/sample/src/advanced/mycommand/addvaluesofall.go) : string값을 파싱해서 더한 결과 처리

    3. result
    
    ```
    > 1. 빌드해서 실행파일을 만듭니다. (폴더에 mycommand 생성)
    $ go build


    > 2. 사용법 확인, "-h" 인자는 환경 설정된 값을 표현해주는 기본 지원되는 기능입니다.
    $ ./mycommand -h
    Usage of ./mycommand:
    -age int
            It is my age value. (default -1)
    -isD
            bool Set test
    -isDeveloper
            bool Set test (shorthand)
    -sum value
            comma separated list of integers, add the values of all
    -t string
            title is a string, it defaults to empty (shorthand)
    -title string
            title is a string, it default to empty


    > 3. 인자값 없이 실행 
    $ ./mycommand

    No title is set.
    Are you a developer?
    Yes. I'm not

    What is your age?
    -1
    , age is -1 means no value is set.
    The result of adding all( ) is 0.


    > 4. 인자값을 주고 실행
    $ ./mycommand -t GoodTitle -isD -age 19 -sum 3,4,5

    GoodTitle
    Are you a developer?
    Yes. I am

    What is your age?
    19
    , age is -1 means no value is set.
    The result of adding all( 3 4 5 ) is 12.

    ```

2. 언제하지? 시간이 되면 셈플을 추가해보자~

## 기타
> Go 언어 사용자들을 고퍼(Gopher)라고 부르며, 고퍼들을 위한 연례행사인 고퍼콘(Gophercon)이 세계 각국에서 매년 열림.

## 좋은 링크 
- http://golang.site/
