package main

import (
	"fmt"
	"strconv"
	"strings"
)

type AddValuesOfAll []int

// ","로 들어온 값을 짤라내서 값을 모두 더하고, 그 결과를 리턴합니다.
func (c *AddValuesOfAll) String() string {
	result := ""
	result += fmt.Sprintf("The result of adding all(")
	sum := 0
	for _, v := range *c {
		result += fmt.Sprintf(" %d", v)
		sum += v

		// For Debug
		// fmt.Println(v)
		// fmt.Println(sum)
	}
	result += fmt.Sprintf(" ) is %d.", sum)
	return result
}

// flag 패키지에서 구조체에 값을 저장시킬때 사용
func (c *AddValuesOfAll) Set(value string) error {
	// , 로 잘라서
	values := strings.Split(value, ",")

	for _, v := range values {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*c = append(*c, i)
	}

	return nil
}
