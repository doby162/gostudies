package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxIters := 25 // remove this check if you get your hands on a super computer
	rootOf, _ := strconv.ParseFloat(os.Args[1], 64)
	guess := new(big.Float).SetInt64(1)

	for i := 0; i <= maxIters; i++ {
		var t = time.Now()
		answer := newt(rootOf, float64(i), guess)
		bigstr := guessToText(answer)
		index := strings.Index(bigstr, os.Args[2])
		fmt.Println("took", time.Now().Sub(t))
		if index != -1 {
			fmt.Println("\nFound \"", os.Args[2], "\" at index ", index, "in the base24 aproximation of sqrt(", os.Args[1], ") after ", i, " iterations")
			i = maxIters + 1
		}
	}
	fmt.Println("finished")
}

func guessToText(answer *big.Float) string {
	bigstr := fmt.Sprint(answer)
	bigstr = strings.Replace(bigstr, ".", "", 1)
	var bignum, _ = new(big.Int).SetString(bigstr, 0)
	bigstr = TextEncode(bignum.Text(26))
	return bigstr
}

func TextEncode(x string) string {
	x = strings.Replace(x, "p", "z", -1)
	x = strings.Replace(x, "o", "y", -1)
	x = strings.Replace(x, "n", "x", -1)
	x = strings.Replace(x, "m", "w", -1)
	x = strings.Replace(x, "l", "v", -1)
	x = strings.Replace(x, "k", "u", -1)
	x = strings.Replace(x, "j", "t", -1)
	x = strings.Replace(x, "i", "s", -1)
	x = strings.Replace(x, "h", "r", -1)
	x = strings.Replace(x, "g", "q", -1)
	x = strings.Replace(x, "f", "p", -1)
	x = strings.Replace(x, "e", "o", -1)
	x = strings.Replace(x, "d", "n", -1)
	x = strings.Replace(x, "c", "m", -1)
	x = strings.Replace(x, "b", "l", -1)
	x = strings.Replace(x, "a", "k", -1)
	x = strings.Replace(x, "9", "j", -1)
	x = strings.Replace(x, "8", "i", -1)
	x = strings.Replace(x, "7", "h", -1)
	x = strings.Replace(x, "6", "g", -1)
	x = strings.Replace(x, "5", "f", -1)
	x = strings.Replace(x, "4", "e", -1)
	x = strings.Replace(x, "3", "d", -1)
	x = strings.Replace(x, "2", "c", -1)
	x = strings.Replace(x, "1", "a", -1)
	x = strings.Replace(x, "0", "a", -1)
	return x
}

func newt(rootOf float64, step float64, guess *big.Float) *big.Float {
	// Since Newton's Method doubles the number of correct digits at each
	// iteration, we need at least log_2(prec) steps.
	// steps := int(math.Log2(prec))
	// but since we actually want to specify the number of steps,
	// we need the inverse of that equation.
	// log2(prec)=steps == prec=2^steps
	// finally, we actually want digits rather than bits,
	// so we multiply the bits by log2(10)
	var prec = uint(math.Log2(10) * math.Exp2(step))
	// not all digits are going to be correct, but that's ok

	// Compute the square root of 2 using Newton's Method. We start with
	// an initial estimate for sqrt(2), and then iterate:
	//     x_{n+1} = 1/2 * ( x_n + (2.0 / x_n) )

	// Initialize values we need for the computation.
	ro := new(big.Float).SetPrec(prec).SetFloat64(rootOf)
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)

	// Use 1 as the initial estimate.
	guess.SetPrec(prec)

	// We use t as a temporary variable. There's no need to set its precision
	// since big.Float values with unset (== 0) precision automatically assume
	// the largest precision of the arguments when used as the result (receiver)
	// of a big.Float operation.
	t := new(big.Float)

	// Iterate.
	fmt.Println("calculating log2(10) * exp2(", step, ") (", prec, ") bits...")
	t.Quo(ro, guess)   // t = 2.0 / guess_n
	t.Add(guess, t)    // t = guess_n + (2.0 / guess_n)
	guess.Mul(half, t) // guess_{n+1} = 0.5 * t
	return guess
}
