package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
)

func main() {
	answer := newt(2, 20)
	file, _ := os.Create("./sqrt.txt")
	defer file.Close()
	bigstr := fmt.Sprint(answer)
	file.WriteString(bigstr)
}

func newt(rootOf float64, steps float64) *big.Float {
	// Since Newton's Method doubles the number of correct digits at each
	// iteration, we need at least log_2(prec) steps.
	// steps := int(math.Log2(prec))
	// but since we actually want to specify the number of steps,
	// we need the inverse of that equation.
	// log2(prec)=steps == prec=2^steps
	// finally, we actually want digits rather than bits,
	// so we multiply the bits by log2(10)
	var prec = uint(math.Log2(10) * math.Exp2(steps))
    // not all digits are going to be correct, but that's ok

	// Compute the square root of 2 using Newton's Method. We start with
	// an initial estimate for sqrt(2), and then iterate:
	//     x_{n+1} = 1/2 * ( x_n + (2.0 / x_n) )

	// Initialize values we need for the computation.
	ro := new(big.Float).SetPrec(prec).SetFloat64(rootOf)
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)

	// Use 1 as the initial estimate.
	guess := new(big.Float).SetPrec(prec).SetInt64(1)

	// We use t as a temporary variable. There's no need to set its precision
	// since big.Float values with unset (== 0) precision automatically assume
	// the largest precision of the arguments when used as the result (receiver)
	// of a big.Float operation.
	t := new(big.Float)

	// Iterate.
	for i := 0.0; i <= steps; i++ {
		fmt.Println("calculating log2(10) * exp2(", i, ") (", int(math.Log2(10)*math.Exp2(i)), ") bits...")
		t.Quo(ro, guess)   // t = 2.0 / guess_n
		t.Add(guess, t)    // t = guess_n + (2.0 / guess_n)
		guess.Mul(half, t) // guess_{n+1} = 0.5 * t
	}
	return guess
}
