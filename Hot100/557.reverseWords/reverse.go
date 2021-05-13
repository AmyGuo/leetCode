package _57_reverseWords

func reverseWords2(s string) string {
	b := []byte(s)
	length := len(b)

	if length <= 1 {
		return s
	}

	for i := 0; i < length; {
		start := i
		for i < length && b[i] != ' ' {
			i++
		}
		left, right := start, i-1
		for left < right {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}

		for i < length && b[i] == ' ' {
			i++
		}
	}
	return string(b)
}

//reverseWords("Let's go to the party!")
//s'teL og ot eht !ytrap
func reverseWords(s string) string {
	length := len(s)

	if length <= 1 {
		return s
	}

	index := 0
	res := make([]byte, 0)
	for index < length {
		start := index
		for index < length && s[index] != ' ' {
			index++
		}

		for j := start; j < index; j++ {
			res = append(res, s[start+index-j-1])
		}

		if index < length && s[index] == ' ' {
			index++
			res = append(res, ' ')
		}
	}
	return string(res)
}

//reverseWords("Let's go to the party!")
//!ytrap eht ot og s'teL
func reverseWords3(s string) string {
	b := []byte(s)
	left, right := 0, len(b)-1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
	return string(b)
}

// I am a student
//student a an I
// 先整个句子翻转，然后再翻转每个单词
func reverseWords4(s string) string {
	b := []byte(s)
	left, right := 0, len(b)-1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}

	for i := 0; i < len(b); {
		start := i
		for i < len(b) && b[i] != ' ' {
			i++
		}
		left, right := start, i-1
		for left < right {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}

		for i < len(b) && b[i] == ' ' {
			i++
		}
	}
	return string(b)
}
