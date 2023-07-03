/**
 *  Beginner - Lost Without a Map
 *	Given an array of integers, return a new array with each value doubled.
 *	For example:
 *	[1, 2, 3] --> [2, 4, 6]

 * 	source: https://www.codewars.com/kata/57f781872e3d8ca2a000007e/train/go
 */

package exercises

func DoubleArray(arr []int) []int {
	results := []int{}

	for _, element := range arr {
		results = append(results, element*2)
	}

	return results
}

// Alternative Solution using Make instead.
func AlternateDoubleArray(arr []int) []int {
	results := make([]int, len(arr))

	for i, element := range arr {
		results[i] = element * 2
	}

	return results
}
