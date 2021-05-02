package main
import (
    "encoding/json"
    "fmt"
    "reflect"
)
type Domicile struct {
    Country  	string  `json:"country"`
    IsRemote 	bool	`json:"is_remote"`
}
type Employee struct {
    Name 			string 		`json:"name"`
    Entity 			string 		`json:"entity"`
    EmployeeNumber 	int 		`json:"employee_number"`
    Salary 			float32 	`json:"salary"`
	Domicile		Domicile	`json:"domicile"`
}
func main() {
    // TODO: task #1 - give me a skeleton!
    data := string(`
    {
        "name": "Golang",
        "entity": "Xendit",
        "employee_number": 10,
        "salary": 1.5,
        "domicile": {
            "country": "ID",
            "is_remote": true
        }
    }
    `)
    var employee Employee
    if err := json.Unmarshal([]byte(data), &employee); err != nil {
		fmt.Println(err)
        fmt.Println("Task #1 failed!")
        return
    }
    // TODO: task #2 - I am a legal employee, include me into your database!
    var database map[string]Employee
    database = make(map[string]Employee)

	// fmt.Println(reflect.TypeOf(database["Golang"])) // main.Employee
	// fmt.Println(reflect.TypeOf(employee)) // main.Employee

	// Database is empty.
	fmt.Println(database) // map[]
	fmt.Println(employee) // {Golang Xendit 10 1.5 {ID false}}

	// Can create like this or assign the key "Golang" to the employee struct
	// database["Golang"] = Employee{
    //     Name: "Golang",
    //     Entity: "Xendit",
    //     EmployeeNumber: 10,
    //     Salary: 1.5,
    //     Domicile: Domicile{
    //         Country: "ID",
    //         IsRemote: true,
    //     },
    // }
	database["Golang"] = employee
	
	fmt.Println(database) // map[Golang:{Golang Xendit 10 1.5 {ID true}}]

    if !reflect.DeepEqual(database["Golang"], employee) {
        fmt.Println("Task #2 failed!")
        return
    }
    return
}