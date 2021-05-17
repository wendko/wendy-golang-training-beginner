package main
import (
    "errors"
    "fmt"
)
func takeLeave(duration int) (string, error) {
    switch duration {
    case 1:
        return "Afik", nil
    case 2:
        return "Iman", nil
    case 3:
        return "Bo", nil
    default:
        return "", errors.New("back to work hooman")
    }
}
func main() {
    duration := 2
    hooman, err := takeLeave(duration)
    if err != nil {
        fmt.Println(err.Error())
        panic(1)
    }
    fmt.Println(hooman, "can replace you. Say thanks to this person!")
}