// package main

// import(
// 	"fmt"
// 	"golang.org/x/exp/constraints"
// )
// func add [T constraints.Number] (a,b T) T{
// 	return a+b
// }
// func main(){
// 	res:=add(1.5,3.7)
// 	fmt.Println("Result = ",res)
// }

//we will create a return length function that will return its len

// package main

// import(
// 	"fmt"
// )

// func Length_counter[T []int | string ] (chill T) int{
// 	return len(chill)
// }


// func main(){
// 	test := "string"
// 	fmt.Println(Length_counter(test))
// }

package main

import "fmt"

func Filter[T any](items []T, predicate func(T) bool) []T {
    var result []T
    for _, v := range items {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    even := Filter(nums, func(n int) bool {
        return n%2 == 0
    })

    fmt.Println(even) // Output: [2 4]
}
