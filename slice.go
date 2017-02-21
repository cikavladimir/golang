package main 
import "fmt"

func main() {

	// A slice is a descriptor of an array segment. 
	// It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).
	x := []int{1, 2, 3, 4}
	fmt.Println(x)

	x = append(x, 5)
	fmt.Println(x)

	// This will create an "empty" slice with a pre-allocate underlying array. 
	// We specify both the length (of the array) and the capacity (of the slice):
	y := make([]int, 0, 4)

	// Print the length of array:
	fmt.Println("Len of array y is:", len(y))
	y = append(y, 5)
	fmt.Println("Len of array y is:", len(y))
	fmt.Println(y)

	original := []int{1,2,3,4} //a slice with a len and cap of 4
	other := original
	other[2] = 8
	fmt.Println(original)
	fmt.Println(other)

	// Create a new location in memory of Array and put original array there:
	other = append(original, 5)
	other[2] = 9
	fmt.Println(original)
	fmt.Println(other)
	// names := []string{"leto", "paul", "teg"}
	// process(names[0:])


	z := [3]string{"Лайка", "Белка", "Стрелка"}
	fmt.Println(len(z))
	s := z[:] 
	fmt.Println(len(s))
	fmt.Println(s)

	a := []string{"John", "Paul"}
	fmt.Println("This is a:", a)
	b := []string{"George", "Ringo", "Pete"}
	fmt.Println("This is b:", b)
	a = append(a, b...)
	fmt.Println("This is a:", a)

	fmt.Println(string(65))


}