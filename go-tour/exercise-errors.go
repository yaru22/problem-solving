package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	var z float64 = 1
	for i := 0; i < 10; i++ {
		z = z - ((z*z)-x)/(2*z)
	}
	return z, nil
}

func main() {
	var result float64
	var err error
	if result, err = Sqrt(2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	if result, err = Sqrt(-2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
