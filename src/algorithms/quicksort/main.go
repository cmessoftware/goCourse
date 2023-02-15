// In this code, the quickSort function implements the quicksort algorithm, which chooses a pivot element and partitions the array into two subarrays: one with elements smaller than the pivot and another with elements greater than the pivot. The function then recursively sorts the subarrays.

// The partition function is used by quickSort to partition the array by moving all elements smaller than the pivot to the left of it and all elements greater than the pivot to the right of it. The function then returns the index of the pivot element.

// The main function creates an example array, calls the quickSort function to sort it, and then prints the sorted array to the console.

// Note that there are many other ways to implement quicksort, and this is just one example in Go.

package main

import "fmt"

func quickSort(arr []int, low, high int) {
	if low < high {
		// Partition the array
		pivot := partition(arr, low, high)

		// Recursively sort the subarrays
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// Choose the pivot element
	pivot := arr[high]

	// Initialize the index of the smaller element
	i := low - 1

	for j := low; j < high; j++ {
		// If the current element is smaller than or equal to the pivot
		if arr[j] <= pivot {
			// Increment the index of the smaller element
			i++

			// Swap arr[i] and arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Swap arr[i+1] and arr[high], where the pivot belongs
	arr[i+1], arr[high] = arr[high], arr[i+1]

	// Return the index of the pivot element
	return i + 1
}

func main() {
	arr := []int{5, 2, 6, 3, 1, 4}

	fmt.Println("Unsorted array:", arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}
