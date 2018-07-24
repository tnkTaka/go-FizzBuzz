package fizzbuzz

import "fmt"

func PrintFizzBuzz(n int) {
	result := Pattern1(n)
	for _, v := range result {
		fmt.Println(v)
	}

	result2 := Pattern2(n)
	for _, v := range result2 {
		fmt.Println(v)
	}
}

// if
func Pattern1(n int) []interface{} {
	fmt.Println("******** Pattern1 ********")

	result := make([]interface{}, n)
	for i := 1; i <= n; i++ {
		var v interface{} = i
		if i%15 == 0 {
			v = "FizzBuzz"
		} else if i%5 == 0 {
			v = "Buzz"
		} else if i%3 == 0 {
			v = "Fizz"
		}
		result[i-1] = v
	}

	return result
}

// switch
func Pattern2(n int) []interface{} {
	fmt.Println("******** Pattern2 ********")

	result := make([]interface{}, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result[i-1] = "FizzBuzz"
		case i%5 == 0:
			result[i-1] = "Buzz"
		case i%3 == 0:
			result[i-1] = "Fizz"
		default:
			result[i-1] = i
		}
	}

	return result
}