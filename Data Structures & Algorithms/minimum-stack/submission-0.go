// if the minimum is set each time a push
// happens, and there's only a single min,
// There's no way to know what the new min should be 
// when popping a value off

// The only way I could see this working is if we
// treat the stack as a ledger, and each time we 
// push a value onto the stack, we also record the 
// minimum at that point in time

// push(1) -> [{1,1}]
// push(2) -> [{1,1},{2,1}] 
// push(0) -> [{1,1},{2,1},{0,0}] 
// getMin() -> 0
// pop() -> [{1,1},{2,1}]
// top() -> 2
// getMin() -> 1

type Data struct {
	val int
	currentMin int
}

type MinStack struct {
	data []Data
}

func Constructor() MinStack {
	return MinStack{ data: []Data{} }
}

func (this *MinStack) Push(val int) {
	var d Data
	if len(this.data) > 0 {
		t := this.topData() // bug: this.Top returns int, so we can't use that
		var min int = t.currentMin
		if val < min {
			min = val
		}
		d = Data{ val, min }
	} else {
		d = Data{val, val}
	}
	this.data = append(this.data, d)
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) topData() Data {
	return this.data[len(this.data) - 1]
}

func (this *MinStack) Top() int {
	top := this.topData()
	return top.val
}

func (this *MinStack) GetMin() int {
	top := this.topData()
	return top.currentMin
}
