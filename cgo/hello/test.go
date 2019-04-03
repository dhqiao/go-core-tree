package main

import "fmt"

type Slice []int
func NewSlice() Slice {
	return make(Slice, 0)
}
func (s* Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}
func main() {
	s := NewSlice()
	defer s.Add(11).Add(2)
	s.Add(3)
}


func deferDemo() error {
	err := createResource1()
	if err != nil {
		return ERR_CREATE_RESOURCE1_FAILED
	}
	defer func() {
		if err != nil {
			destroyResource1()
		}
	}()

	err = createResource2()
	if err != nil {
		return ERR_CREATE_RESOURCE2_FAILED
	}
	defer func() {
		if err != nil {
			destroyResource2()
		}
	}()

	err = createResource3()
	if err != nil {
		return ERR_CREATE_RESOURCE3_FAILED
	}
	return nil
}