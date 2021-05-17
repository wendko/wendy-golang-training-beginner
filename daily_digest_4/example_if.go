package main
import (
    "fmt"
)
func main() {
    workHour := 40
    if workHour >= 35 && workHour <= 45 {
        fmt.Println("I am working normally")
    } else if workHour <= 35 {
        fmt.Println("I am working effectively!... Or maybe I'm too lazy (ehe)")
    } else {
        fmt.Println("Can somebody describe my working hour?")
    }
}