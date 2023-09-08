package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(i int) string {

	var fizz = "Fizz"
	var	buzz = "Buzz"
	var fizzbizz = "FizzBuzz"


	if (i % 3) == 0 && (i % 5 ) == 0 {
		return (fizzbizz)

	}else if (i % 3) == 0 {
		return fizz
	}else if (i % 5 ) == 0{
		return buzz
	}

	return strconv.FormatInt(int64(i), 10)
}
