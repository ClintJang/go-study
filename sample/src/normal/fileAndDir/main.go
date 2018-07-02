package main

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
)

func initFileAndFolder() error {
	if err := os.RemoveAll("makeDirFolder01"); err != nil {
		return err
	}

	if err := os.RemoveAll("makeDirFolder02"); err != nil {
		return err
	}

	if err := os.Remove("testFileMake02.txt"); err != nil {
		return err
	}
	return nil
}

func testFolderAndFile() error {

	// 폴더 만들기 1
	if err := os.Mkdir("makeDirFolder01", os.FileMode(0755)); err != nil {
		return err
	}

	// 폴더 만들기 2
	if err := os.Mkdir("makeDirFolder02", os.FileMode(0755)); err != nil {
		return err
	}

	// 첫번째 만들어진 폴더로 들어가기
	if err := os.Chdir("makeDirFolder01"); err != nil {
		return err
	}

	// 파일 만들고, 파일 매니저 열기
	f, err := os.Create("testTXT01.txt")
	if err != nil {
		return err
	}

	// "hello world 01" 문자열 파일에 쓰기
	value := []byte("hello world 01\n")
	count, err := f.Write(value)
	if err != nil {
		return err
	}

	if count != len(value) {
		return errors.New("incorrect length returned from write")
	}

	// 파일 매니저 닫기
	if err := f.Close(); err != nil {
		return err
	}

	// 만들었던 파일 열어보기
	f, err = os.Open("testTXT01.txt")
	if err != nil {
		return err
	}

	// 출력
	io.Copy(os.Stdout, f)

	// 파일 매니저 닫기
	if err := f.Close(); err != nil {
		return err
	}

	// 상위 폴더로 이동
	if err := os.Chdir(".."); err != nil {
		return err
	}

	// 폴더 지우기 (만들어진 것을 보려고 지우지 않는다.)
	// if err := os.RemoveAll("makeDirFolder01"); err != nil {
	// 	return err
	// }

	return nil
}

func copyAndToUpper(f1 *os.File, f2 *os.File) error {
	if _, err := f1.Seek(0, 0); err != nil {
		return err
	}

	var temp = new(bytes.Buffer)

	// tmp에 첫번째 파일 내용을 복사합니다.
	if _, err := io.Copy(temp, f1); err != nil {
		return err
	}

	// 대문자로 변경합니다.
	tempToUpperString := strings.ToUpper(temp.String())

	// 대문자로 변경한 문자열을 두번째에 복사합니다.
	if _, err := io.Copy(f2, strings.NewReader(tempToUpperString)); err != nil {
		return err
	}
	return nil
}

func testFile() error {
	// testFileMake01.txt 파일을 만듭니다.
	f1, err := os.Create("testFileMake01.txt")
	if err != nil {
		return err
	}

	// 여러줄의 테스트 문장을 추가합니다.
	if _, err := f1.Write([]byte(`
    test line1
    test line2
    test line3`)); err != nil {
		return err
	}

	// 두번째 testFileMake02.txt 파일을 만듭니다.
	f2, err := os.Create("testFileMake02.txt")
	if err != nil {
		return err
	}

	// 복사합니다. (복사한 문자열을 중간에 대문자로 문자를 변경하였습니다.)
	if err := copyAndToUpper(f1, f2); err != nil {
		return err
	}

	// 복사한 첫번째 파일을 지웁니다.
	if err := os.Remove("testFileMake01.txt"); err != nil {
		return err
	}

	// 두번째 파일은 지우지 않습니다.(확인해야죠)
	// if err := os.Remove("testFileMake02.txt"); err != nil {
	// 	return err
	// }

	return nil
}

func main() {
	print("폴더 : makeDirFolder01, 02 는 지우세요. \n파일 : testFileMake02.txt 도 지우세요.\n\n")
	// 테스트를 위해만들어진 폴더를 지우고 시작한다.
	if err := initFileAndFolder(); err != nil {
		println("에러가 나도 종료하지 않습니다.\n폴더가 없을 수 있으니깐요.. \n이 프로그램은 공부용 테스트 프로그램이에요.\n처음만 발생합니다.\nerr:%s", err)
		// panic(err)
	}

	if err := testFolderAndFile(); err != nil {
		// panic()함수는 현재 함수를 즉시 멈추고 현재 함수에 defer 함수들을 모두 실행한 후 즉시 리턴
		// panic 모드 실행 방식은 다시 상위함수에도 똑같이 적용되고, 계속 콜스택을 타고 올라가며 적용
		// 그리고 프로그램의 에러를 내고 종료함
		panic(err)
	}

	if err := testFile(); err != nil {
		// panic()함수는 현재 함수를 즉시 멈추고 현재 함수에 defer 함수들을 모두 실행한 후 즉시 리턴
		// panic 모드 실행 방식은 다시 상위함수에도 똑같이 적용되고, 계속 콜스택을 타고 올라가며 적용
		// 그리고 프로그램의 에러를 내고 종료함
		panic(err)
	}

	println("만들어진 폴더와 파일을 확인해보세요~ ")
}
