package main
import (
    "encoding/json"
    "fmt"
    "reflect"
)
type Employee struct {
}
func main() {
    // TODO: task #1 - give me a skeleton!
    data := string(`
    {
        "name": "Golang",
        "entity": "Xendit",
        "employee_number": 10,
        "salary": 1.5
        "domicile": {
            "country": "ID",
            "is_remote": true
        }
    }
    `)
    var employee Employee
    if err := json.Unmarshal([]byte(data), &employee); err != nil {
        fmt.Println("Task #1 failed!")
        return
    }
    // TODO: task #2 - I am a legal employee, include me into your database!
    var database map[string]Employee
    if !reflect.DeepEqual(database["Golang"], employee) {
        fmt.Println("Task #2 failed!")
        return
    }
    return
}