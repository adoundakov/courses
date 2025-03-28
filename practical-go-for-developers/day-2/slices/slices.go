package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int // s is a slice of ints
	fmt.Println("len", len(s))
	if s == nil {
		fmt.Println("nil slice")
	}

	// can create slices directly in code
	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 = %#v\n", s2)

	s3 := s2[1:4] // slicing operation, half-open range (takes from first not including end)
	fmt.Printf("s3 = %#v\n", s3)

	// Can omit the start or end index to slice from beginning or to end
	fmt.Println(s2[2:])
	fmt.Println(s2[:3])

	// Slicing out of the slices range will cause a panic
	// fmt.Println(s2[:100])

	// To add need to use built-in append method
	s3 = append(s3, 100)
	fmt.Printf("s3 append: %#v\n", s3)

	// However, sometimes strange things can happen:
	fmt.Printf("s2 w/o append: %#v\n", s2) // prints []int{1, 2, 3, 4, 100, 6, 7}

	// Can also make slices using the "make" built-in
	// s4 is a struct, where the underlying array points to the start of the array in memory
	s4 := make([]int, 10) // makes a slice with 10 elements, all initialized to 0
	// When we make a sub-slice, the returned value is another struct, but with the underlying
	// array pointing to the start of the sub-slice, a new array is not created
	s5 := s4[3:7] // creates a slice with 4 length, and 7 capacity (start idx to end of underlying array)
	fmt.Printf("s5: len=%d, cap=%d\n", len(s5), cap(s5))

	var s6 []int
	for i := 0; i < 1_000; i++ {
		s6 = appendInt(s6, i)
	}

	fmt.Println("s6", len(s6), cap(s6))

	fmt.Println(concat([]string{"a", "b"}, []string{"c", "d", "e"}))

	vs := []float64{2, 1, 3}
	fmt.Println(median(vs))
	vs = []float64{2, 1, 3, 4}
	fmt.Println(median(vs))
	// However, printing the values shows that they got mutated!
	fmt.Println(vs)

	// No way around it, need to copy array to prevent mutation
	vs = []float64{2, 1, 3, 4}
	fmt.Println(safeMedian(vs))
	// However, printing the values shows that they got mutated!
	fmt.Println(vs)

	// Error checking
	fmt.Println(safeMedian(nil))
}

// Let's dive into how the append function works
func appendInt(s []int, v int) []int {
	// Want to insert at the end of the slice
	i := len(s)

	// We can use the underlying array as-is if there's leftover capacity
	if len(s) < cap(s) {
		s = s[:len(s)+1]
	} else {
		// If there is no capacity, need to create new one and copy old
		// But instead of +1, we need to double the size of the underlying
		// array, it is more efficient long-term (or if we're appending)
		// a lot of elements
		s2 := make([]int, 2*len(s)+1)
		copy(s2, s) // shallow copy
		// Because we doubled the underlying array, we want to slice up to the
		// new length to prevent extra elements from leaking out
		s = s2[:len(s)+1]
	}

	// Now we can safely assign the value at the end of the array
	s[i] = v
	return s
}

func concat(s1, s2 []string) []string {
	// Restriction: no "for" loops
	s3 := make([]string, len(s1)+len(s2))

	copy(s3, s1)
	copy(s3[len(s1):], s2)

	return s3
}

func median(values []float64) float64 {
	// Median is defined as midpoint of sorted array
	// If midpoint straddles two values, median is average of the two
	sort.Float64s(values)
	i := len(values) / 2
	if len(values)%2 == 1 {
		return values[i]
	}

	return (values[i-1] + values[i]) / 2
}

func safeMedian(values []float64) (float64, error) {
	// Check for errors!
	if len(values) == 0 {
		return 0, fmt.Errorf("tried to get median of empty slice")
	}

	// Median is defined as midpoint of sorted array
	// If midpoint straddles two values, median is average of the two
	nums := make([]float64, len(values))
	copy(nums, values)

	sort.Float64s(nums)
	i := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[i], nil
	}

	return (nums[i-1] + nums[i]) / 2, nil
}
