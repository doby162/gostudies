package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
    "github.com/meirf/gopart" // go get github.com/meirf/gopart
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
	bigstr = TextEncode(bigstr)
	return bigstr
}

func TextEncode(x string) string {
    var charColl = [...]string {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
    ret := ""
    for idxRange := range gopart.Partition(len(x), 2) {
        n, _ := strconv.Atoi(x[idxRange.Low:idxRange.High])
        n = n % len(charColl)
        ret = ret + charColl[n]
    }
	return ret
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
