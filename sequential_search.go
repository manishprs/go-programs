package main
  
import "fmt"
 
func seqSearch(arr []int, key int) bool {
    for _, item := range arr {
        if item == key {
            return true
        }
    }
    return false
} 
  
func main() {
    arr := []int{11,22,33,55,44,12,65,23,1}
    fmt.Println(seqSearch(arr,99))
}