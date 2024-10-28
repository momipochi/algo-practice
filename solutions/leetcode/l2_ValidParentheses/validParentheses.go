package l2validparentheses

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}
	dict := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	stack := []rune{}

	for _, v := range s {
		if _, ok := dict[v]; ok {
			stack = append(stack, v)
		} else if len(stack) == 0 || dict[stack[len(stack)-1]] != v {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
