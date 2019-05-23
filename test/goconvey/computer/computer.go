package computer

import (
"errors"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Division(a, b int) (int, error) {
	aa()
	if b == 0 {
		return 0, errors.New("被除数不能为 0")
	}
	return a / b, nil

}

func aa()  {
	bb()
}

func bb()  {
	for i:=10;i<10;i++{
		cc(i)
	}

	// time.Sleep(time.Microsecond)
}

func cc(i int)  {

}