package main
import "fmt"

type BankAccount struct {
  holder string
  balance int
}

func (ba *BankAccount) withdraw(amount int){
    
    if (ba.balance - amount) < 0 {
        fmt.Println("Insufficient Funds")
    } else {
        ba.balance -= amount
    }
}

func main() {
  acc := BankAccount{"James Smith", 100000}
  
  var amount int
  fmt.Scanln(&amount)
  
  acc.withdraw(amount)
  fmt.Println(acc)
}
