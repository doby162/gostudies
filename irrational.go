package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println(square(2, 100000))
	guess := big.NewFloat(1.0)
	fmt.Println(guess)
	newton(2, guess)
	fmt.Println(guess)
	newton(2, guess)
	fmt.Println(guess)
	newton(2, guess)
	fmt.Println(guess)
	newton(2, guess)
	fmt.Println(guess)
	newton(2, guess)
	fmt.Println(guess)
	newton(2, guess)
	fmt.Println(guess)
	newton(2, guess)
	fmt.Println(guess)
	newton(2, guess)
	fmt.Println(guess)
	newton(2, guess)
	fmt.Println(guess)
	newton(2, guess)
	fmt.Println(guess.String())
}

func Sqrt(x float64, i int) float64 {
	z := 1.0 //our initital guess
	for ; i > 0; i-- {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func Digit(x float64) (int, float64) {
	i := 10
	z := 1.0 //our initital guess
	for ; i > 0; i-- {
		z = z - (z*z-x)/(2*z)
	}
	num := int(z)
	rem := z
	return num, rem
}

func square(n int64, precision int64) string {
	ans_int := strconv.Itoa(int(math.Sqrt(float64(n))))

	limit := new(big.Int).Exp(big.NewInt(10), big.NewInt(precision+1), nil)
	a := big.NewInt(5 * n)
	b := big.NewInt(5)
	five := big.NewInt(5)
	ten := big.NewInt(10)
	hundred := big.NewInt(100)

	for b.Cmp(limit) < 0 {
		if a.Cmp(b) < 0 {
			a.Mul(a, hundred)
			tmp := new(big.Int).Div(b, ten)
			tmp.Mul(tmp, hundred)
			b.Add(tmp, five)
		} else {
			a.Sub(a, b)
			b.Add(b, ten)
		}
	}
	b.Div(b, hundred)

	ans_dec := b.String()

	return ans_dec[:len(ans_int)] + "." + ans_dec[len(ans_int):]
}

func newton(n float64, guess *big.Float) {
	// *guess = *big.NewFloat(3.0) // replace a big float
	// one := big.NewFloat(1.0)
	// guess.Add(guess, one) // add to a big float
	//
	// one := big.NewFloat(1.0)
	// guess.Add(guess, one) // add to a big float

	rootOf := big.NewFloat(n)

	// right parens are 2  * guess
	rightParens := big.NewFloat(0.0)
	rightParens.Copy(guess)
	rightParens.Mul(rightParens, big.NewFloat(2.0))

	// left parens are guess * guess minus rootOf
	leftParens := big.NewFloat(0.0)
	leftParens.Copy(guess)
	leftParens.Mul(leftParens, guess)
	leftParens.Sub(leftParens, rootOf)

	// combine is left parens devided by right parens
	combine := big.NewFloat(0.0)
	combine.Quo(leftParens, rightParens)

	//finally we subtract the right side of the equasion from our guess

	guess.Sub(guess, combine)
}
