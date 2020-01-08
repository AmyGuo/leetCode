package _0_isValid

import "errors"

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//
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

type stack []interface{}

func (s *stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *stack) Pop() (interface{}, error) {
	if s.Len() == 0 {
		return nil, errors.New("nil")
	}
	temp := *s
	value := temp[temp.Len()-1]
	*s = temp[:temp.Len()-1]
	return value, nil
}

func (s *stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *stack) Len() int {
	return len(*s)
}
func IsValid(s string) bool {
	if s == "" {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	stk := new(stack)
	for _, v := range s {
		if v == '{' {
			stk.Push('}')
		} else if v == '[' {
			stk.Push(']')
		} else if v == '(' {
			stk.Push(')')
		} else if stk.IsEmpty() {
			return false
		} else if c, err := stk.Pop(); err == nil && c != v {
			return false
		}
	}
	return stk.IsEmpty()
}
