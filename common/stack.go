package common

import "log"

type Stack struct {
	Items []any
}

func (st *Stack) Push(data any) {
	st.Items = append(st.Items, data)
}

func (st *Stack) Pop() any {
	if st.IsEmpty() {
		log.Panic("Stack is empty")
	}
	returnVal := st.Items[len(st.Items)-1]
	st.Items = st.Items[:len(st.Items)-1]
	return returnVal
}

func (st *Stack) IsEmpty() bool {
	return len(st.Items) == 0
}
