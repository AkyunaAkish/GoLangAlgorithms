// The two oldest ages function/method needs to be completed. It should take an array of numbers as its argument and return the two highest numbers within the array. The returned value should be an array in the format [second oldest age, oldest age].

// The order of the numbers passed in could be any order. The following are some examples of what the method should return (shown in different languages but the logic will be the same between all three).

// TwoOldestAges([]int{1,5,87,45,8,8}) // should return [2]int{45,87}

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{93, 35, 53, 67, 17, 23, 89, 75, 15, 53}
	fmt.Println(twoOldestAges(arr))
}

func twoOldestAges(ages []int) [2]int {
	sort.Ints(ages)
	return [2]int{ages[len(ages)-2], ages[len(ages)-1]}
}
