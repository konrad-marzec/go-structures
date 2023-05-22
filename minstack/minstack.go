package minstack

type tuple struct {
	value int
	idx   int
}

type MinStack struct {
	stack []tuple
}

func New() *MinStack {
	return &MinStack{stack: make([]tuple, 0)}
}

func (ms *MinStack) Push(val int) {
	l := len(ms.stack)

	if l <= 0 {
		ms.stack = append(ms.stack, tuple{value: val, idx: 0})
	} else {
		idx := ms.stack[l-1].idx

		if ms.stack[idx].value >= val {
			idx = len(ms.stack)
		}

		ms.stack = append(ms.stack, tuple{value: val, idx: idx})
	}
}

func (ms *MinStack) Pop() int {
	if len(ms.stack) == 0 {
		return -1
	}

	e := ms.stack[len(ms.stack)-1]

	ms.stack = ms.stack[:len(ms.stack)-1]

	return e.value
}

func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		return -1
	}

	return ms.stack[len(ms.stack)-1].value
}

func (ms *MinStack) GetMin() int {
	if len(ms.stack) == 0 {
		return -1
	}

	return ms.stack[ms.stack[len(ms.stack)-1].idx].value
}
