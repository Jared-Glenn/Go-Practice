package main
import "fmt"


func mars_age(age int) int{
    days := age*365
    mars_age := days/687
    return mars_age
}

func main() {
    var age int
    fmt.Scanln(&age)

    mars := mars_age(age)
    fmt.Println(mars)
}
