package main
import "fmt"

//define the count() function
func count(num int, arr []int, ch chan int){
    total := 0
    for _, x := range arr {
        if x == num {
            total += 1
        }
    }
    ch <- total
}


func main() {
    data := []int{12, 45, 88, 42, 0, 98, 102, 42, 77, 42, 1, 8, 7, 55, 4, 12, 87, 90, 42, 42, 11, 2, 6, 53, 90, 100, 4, 32, 8}
    
    var num int
    fmt.Scanln(&num)

    ch1 := make(chan int)
    ch2 := make(chan int)

    go count(num, data[:len(data)/2], ch1)
    go count(num, data[len(data)/2:], ch2)

    //output the final result here
    fmt.Println(<-ch1 + <-ch2)
}
