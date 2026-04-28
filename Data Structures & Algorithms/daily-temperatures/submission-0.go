type Temp struct {
	value int
	index int
}

func dailyTemperatures(temperatures []int) []int {
	// Naive Approach: 
	// For each number, loop until we get a number larger
	// T: O(n^2)
	// S: O(1) extra space

	// Utilize Stack
	// Add numbers to a stack, which is kept monotonically decreasing
	// keep track of indices
	// if encounter a number that breaks the decreasing, calculate "distance"
	// and add to a result array
	// T: O(n)
	// T: O(n)
	//                  0  1  2  3  4  5  6
	// temperatures = [30,38,30,36,35,40,28] 
	//                                   t 
	// stack [40,28]
	// result [1,4,1,2,1,0,0]
	// dist 
	result := make([]int, len(temperatures)) // creates a list of len(temp) 0's
	stack := []Temp{}

	for i, t := range temperatures {
		curr := Temp{t,i}

		if len(stack) == 0 {
			stack = append(stack, curr)
			continue
		} 

		top := stack[len(stack) - 1]

		if curr.value <= top.value {
			stack = append(stack, curr)
			continue
		}

		for curr.value > top.value {
			distance := curr.index - top.index
			result[top.index] = distance
			stack = stack[:len(stack) - 1] // pop
			if len(stack) > 0 {
			  top = stack[len(stack) - 1] // reset top
			} else {
				break
			}
		}
		// add curr onto the stack
		stack = append(stack, curr)
	}

	// zero out the rest of the numbers
	// edit: we don't need to do this if we initialize the slice
	// with zeros
	//for len(stack) > 0 {
	//	top = stack[len(stack) - 1]
	//	result[top.index] = 0
	//	stack = stack[:len(stack - 1)]
	//}
	return result
}
