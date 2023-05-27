package main
import "fmt"

func main() {
    results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
    
    var added string
    fmt.Scanln(&added)
    results = append(results, added)
    
    total := 0
    
    for _, res := range results{
        if res == "w"{
            total += 3
        } else if res == "d"{
            total += 1
        }
    }
    
    fmt.Println(total)
    
}
