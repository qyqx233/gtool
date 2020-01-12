package datastruct

type Stack []interface{}

func NewStack() Stack {
	return make([]interface{}, 0, 8)
}

func (stack Stack) Push(v interface{}) Stack {
	return append(stack, v)
}

func (stack Stack) Pop() (interface{}, Stack) {
	if len(stack) == 0 {
		return nil, stack
	}
	var pos = len(stack) - 1
	return stack[pos], stack[:pos]
}
