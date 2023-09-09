package stack

type Stack []float32
 
func (s *Stack) ExportPush(val float32) {
	s.push(val)
}

func (s *Stack) ExportPop() (float32, bool){
	return s.pop()
}


func (s *Stack) push(val float32) {
	// push values to the end of the lsit 
	*s = append(*s, val)
}

func (s *Stack) pop() (float32, bool){
	// check if stack is empty 
	if len(*s) == 0 {
		return -1, false 
	}

	firstVal := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s)-1]

	return firstVal, true
}