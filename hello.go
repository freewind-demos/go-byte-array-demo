package main;

import "fmt"

var (
	array = [...]byte{0, 1, 2, 3, 4, 10, 11}
	slice = []byte{0, 1, 2, 3, 4, 10, 11}
)

func main() {
	arrayToHex()
	sliceToHex()
	arrayToSlice()
	sliceToArray()
}

func arrayToSlice() {
	fmt.Println("----------- arrayToSlice --------------")
	var newSlice []byte
	newSlice = array[:]
	fmt.Printf("% 02x\n", newSlice)
}

func sliceToArray() {
	fmt.Println("----------- sliceToArray --------------")
	var newArray [3]byte
	copy(newArray[:], slice)
	fmt.Printf("% 02x\n", newArray)
}

func arrayToHex() {
	fmt.Println("----------- arrayToHex --------------")
	fmt.Printf("The array is: %02x\n", array)
	fmt.Printf("The array is: % 02x\n", array)

	str := fmt.Sprintf("% 02x", array)
	fmt.Println(str)
}

func sliceToHex() {
	fmt.Println("----------- sliceToHex --------------")
	fmt.Printf("The array is: %02x\n", slice)
	fmt.Printf("The array is: % 02x\n", slice)

	str := fmt.Sprintf("% 02x", slice)
	fmt.Println(str)
}
