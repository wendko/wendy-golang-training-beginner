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
    var world2 map[string]helloWorld

	// initialize the map before adding elements
	world2 = make(map[string]helloWorld)

    world2["Earth"] = helloWorld{
        foo: "foo",
        bar: 10,
        xendit: Xendit{
            Location:  "Worldwide",
            Valuation: 1000,
        },
    }
    world2["Mars"] = helloWorld{
        foo: "foo",
        bar: 10,
        xendit: Xendit{
            Location:  "Worldwide",
            Valuation: 1000,
        },
    }
	fmt.Println(world2); // map[Earth:{foo 10 {Worldwide 1000}} Mars:{foo 10 {Worldwide 1000}}]
}