package main
import (
    "fmt"
    "strconv"
)
func xenDit(i int) string {
	if i % 10 == 0 {
		return "Xendit"
	}

	if i % 3 == 0 {
		if i % 2 == 0 {
			return "Dit"
		}
		return "Xen"
	}

    return strconv.Itoa(i)
}

func main() {
    tests := []struct {
        name string
        in   int
        want string
    }{
        {
            name: "Should return number 1",
            in:   1,
            want: "1",
        }, {
            name: "Should return Xen",
            in:   9,
            want: "Xen",
        }, {
            name: "Should return Dit",
            in:   24,
            want: "Dit",
        }, {
            name: "Should return Xendit",
            in:   30,
            want: "Xendit",
        },
    }
    for _, tt := range tests {
        result := xenDit(tt.in)
        if result != tt.want {
            fmt.Println("Task #1 - Failed! ", tt.name)
            panic(1)
        }
    }
    fmt.Println("Task #1 - Passed!")
}
