package main
import (
    "fmt"
)
type helloWorld struct {
    foo    string
    bar    int
    xendit Xendit
}
type Xendit struct {
    Location  string
    Valuation float32
}
func main() {
    world := helloWorld{
        foo: "foo",
        bar: 10,
        xendit: Xendit{
            Location:  "Worldwide",
            Valuation: 1000,
        },
    }
    fmt.Println(world) // {foo 10 {Worldwide 1000}}
}