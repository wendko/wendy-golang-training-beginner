go
package main
import (
    "fmt"
)
func describeMyWorkingHour(workHour int) (string, error) {
	// Multiple ifs are better than else if
    if workHour >= 35 && workHour <= 45 {
        return "I am working normally", nil
    }
    if workHour <= 35 {
        return "I am working effectively!... Or maybe I'm too lazy (ehe)", nil
    }
    return "", errors.New("UNKNOWN")
}
func main() {
    workHour := 40
    explanation, err := describeMyWorkingHour(workHour)
	// Check negative case first
    if err != nil {
        fmt.Println(err.Error())
        panic(1)
    }
    fmt.Println(explanation)
}