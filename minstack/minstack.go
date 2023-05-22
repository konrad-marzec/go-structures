package minstack

type tuple[T interface{}] struct {
	value T
	idx   int
}

type MinStack[T interface{}] struct {
	stack []tuple[T]
}

func New[T interface{}]() *MinStack[T] {
	return &MinStack[T]{stack: make([]tuple[T], 0)}
}

func (ms *MinStack[T]) Push(val T) {

}

func (ms *MinStack[T]) Pop() T {

}

func (ms *MinStack[T]) Top() T {

}

func (ms *MinStack[T]) GetMin() T {

}
