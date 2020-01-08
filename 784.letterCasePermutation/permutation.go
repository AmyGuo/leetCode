package _84_letterCasePermutation

//给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。
//
//示例:
//输入: S = "a1b2"
//输出: ["a1b2", "a1B2", "A1b2", "A1B2"]
//
//输入: S = "3z4"
//输出: ["3z4", "3Z4"]
//
//输入: S = "12345"
//输出: ["12345"]

//注意：
//
//S 的长度不超过12。
//S 仅由数字和字母组成

func LetterCasePermutation(S string) []string {
	if len(S) == 0 || len(S) > 12 {
		return nil
	}

	res := []string{""}
	for _, v := range []byte(S) {
		l := len(res)
		if v >= '0' && v <= '9' {
			for i := range res {
				res[i] = res[i] + string(v)
			}
		} else {
			for i := 0; i < l; i++ {
				res = append(res, res[i])
				if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
					res[i] = res[i] + string(lower(v))
					res[l+i] = res[i+l] + string(upper(v))
				}
			}
		}
	}

	return res

}

func lower(a byte) byte {
	if a >= 'A' && a <= 'Z' {
		return a + 32
	}
	return a
}

func upper(a byte) byte {
	if a >= 'a' && a <= 'z' {
		return a - 32
	}
	return a
}

func LetterCasePermutation2(S string) []string {
	if S == "" {
		return []string{""}
	}

	ret := make([]string, 0)
	p := LetterCasePermutation2(S[1:])
	if (S[0] >= 'a' && S[0] <= 'z') || (S[0] >= 'A' && S[0] <= 'Z') {
		for j := 0; j < len(p); j++ {
			ret = append(ret, string(lower(S[0]))+p[j])
			ret = append(ret, string(upper(S[0]))+p[j])
		}
	} else {
		for j := 0; j < len(p); j++ {
			ret = append(ret, string(S[0])+p[j])
		}
	}

	return ret
}
