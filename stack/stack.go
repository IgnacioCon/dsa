package stack

type Stack struct {
	data []int
}

func (s *Stack) IsEmpty() bool {
	if len(s.data) > 0 {
		return false
	}
	return true
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	}

	popped := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return popped
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) Peek() int {
	if !s.IsEmpty() {
		return s.data[len(s.data)-1]
	}
	return 0
}

func (s *Stack) Length() int {
	return len(s.data)
}
