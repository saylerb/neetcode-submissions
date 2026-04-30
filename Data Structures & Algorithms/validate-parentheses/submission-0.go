func isValid(s string) bool {
	// in this problem
	// I don't it's possible 
	// to simply "count" up if you have opening
	// bracket, and count down if you don't
	// Mainly because, [(]) would be valid if you simply did counts


	// Naive Approach
	// walk the input string, one character at a time
	// push run/char on to a stack (if empty)

	// if the current character "matches" the top of the stack
	// pop off the top of the stack and continue

	// if it doesn't match, add it to the stack

	// right matches a left

	// s = "([{}])"
	// r = '('
	// stack []

	matching := map[rune]rune {
		'}':'{',
		')':'(',
		']':'[',
	}
	
	stack := []rune{}

	for _, r := range s {
		if len(stack) > 0 && matching[r] == stack[len(stack) - 1] {
			stack = stack[:len(stack) - 1] // pop
		} else {
			stack = append(stack, r)
		}
	}
	return len(stack) == 0
}
