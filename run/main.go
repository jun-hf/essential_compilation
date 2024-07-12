package main


func main() {
	var a Example

	switch a.(type) {
	case Example:
		print("example")
	default:
		print("Herel")
	}
}

type Example interface {

}