package _0_isValid

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	fmt.Println(IsValid("()"))
	fmt.Println(IsValid("()[]{}"))
	fmt.Println(IsValid("(]"))
	fmt.Println(IsValid("([)]"))
	fmt.Println(IsValid("{[]}"))
	fmt.Println(IsValid("["))
	fmt.Println(IsValid("(("))
	fmt.Println(IsValid("){"))
}

//示例 1:
//输入: "()"
//输出: true

//示例 2:
//输入: "()[]{}"
//输出: true

//示例 3:
//输入: "(]"
//输出: false

//示例 4:
//输入: "([)]"
//输出: false

//示例 5:
//输入: "{[]}"
//输出: true
