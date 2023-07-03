/**
 * Find the smallest integer in the array
 * Source: https://www.codewars.com/kata/55a2d7ebe362935a210000b2/train/go
 */
package exercises

import "math"

func SmallestItemInArray(arr []int) int {
	smallest := math.MaxInt64

	for _, element := range arr {
		if element <= smallest {
			smallest = element
		}
	}

	return smallest
}
