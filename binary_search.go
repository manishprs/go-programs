package main
import "fmt"
import "sort"
 
func binarySearch(arr []int, key int) bool {	
 
    low := 0
    high := len(arr) - 1
 
    for low <= high{
        median := (low + high) / 2
 
        if arr[median] < key {
            low = median + 1
        }else{
            high = median - 1
        }
    }
 
    if low == len(arr) || arr[low] != key {
        return false
    }
 
    return true
}
 
 
func main(){
	arr := []int{11,22,33,55,44,12,65,23,1}
	sort.Ints(arr)
    fmt.Println(binarySearch(arr, 13))
}