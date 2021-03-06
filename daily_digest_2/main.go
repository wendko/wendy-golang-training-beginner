package main
import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	// TODO: task #1 - Why is it not working?
	// constants must be known at compile time. os.Getenv evaluates during run time (src: https://stackoverflow.com/a/46352442)
	var task1 = os.Getenv("task_1")
	if !reflect.DeepEqual("", task1) {
		fmt.Println("Task #1 failed!")
	}
	// TODO: Challenge #1 - How to convert from any type into string
	var challenge1 interface{} // DO NOT CHANGE THE DATATYPE
	challenge1 = strconv.Itoa(1)
	if !reflect.DeepEqual("1", challenge1) {
		fmt.Println("Challenge #1 failed!")
	}
	fmt.Println("All passed!")
}