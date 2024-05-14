package main

import "fmt"

func mergeSortedArrays(arr1 []int, arr2 []int) []int {
	merged := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0

	for i &lt; len(arr1) &amp;&amp; j &lt; len(arr2) {
		if arr1[i] &lt; arr2[j] {
			merged[k] = arr1[i]
			i++
		} else {
		   merged[k] = arr2[j]
		   j++
		}
		k++
	}

	for i &lt; len(arr1) {
		merged[k] = arr1[i]
		i++
		k++
	}

	for j &lt; len(arr2) {
		merged[k] = arr2[j]
		j++
		k++
	}

	return merged
}

func main() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8, 10}

	mergedArray := mergeSortedArrays(arr1, arr2)

	fmt.Println("Слияние двух отсортированных массивов:")
	fmt.Println(mergedArray)
}
