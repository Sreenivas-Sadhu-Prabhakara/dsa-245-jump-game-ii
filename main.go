package main

import "fmt"

// MinJumps returns the fewest hops from the first to the last index
// (-1 if unreachable). Implicit BFS: each "level" is the stretch of indexes
// reachable in the same number of jumps; crossing its boundary costs one.
func MinJumps(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	jumps, levelEnd, farthest := 0, 0, 0
	for i := 0; i < n-1; i++ {
		if i > farthest {
			return -1 // gap: this index was never reachable
		}
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if i == levelEnd { // leaving the current jump's reach: must hop
			jumps++
			levelEnd = farthest
		}
	}
	if levelEnd >= n-1 {
		return jumps
	}
	return -1
}

func main() {
	fmt.Println(MinJumps([]int{2, 3, 1, 1, 4}))
	fmt.Println(MinJumps([]int{3, 2, 1, 0, 4}))
}
