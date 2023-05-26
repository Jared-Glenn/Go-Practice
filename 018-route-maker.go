package main
import "fmt"

//create the route() function
func route(cities ...string) {
    for _, city := range cities {
        fmt.Print(city + "->")
    }
}

func main() {
    var n int
    fmt.Scanln(&n)

    cities := make([]string, 0)
    //take n strings as input and append them to the slice
    for i:=0; i<n; i++{
        var str string
        fmt.Scanln(&str)
        cities = append(cities, str)
    }
    //
    route(cities...)
}
