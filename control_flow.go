package main
import "fmt"
import "time"

func main() {


	sum := 1
	fmt.Println("sum:", sum)

	for ; sum < 1000; {
	    sum += sum
	}

	fmt.Println("sum:", sum)
	fmt.Println(time.Now().Unix() % 2)
}