package main

import "fmt"

type student struct {
	name       string
	age        int8
	rollNumber string
	gender     string
}

// struct method
func (s student) printStudentDetails() {
	fmt.Printf("Name: %v, Age: %v", s.name, s.age)
}

func main() {
	fmt.Println("Arrays")
	var intArr [3]int32
	intArr[0] = 1
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intAnotherArr = [3]int32{1, 2, 3}
	fmt.Println(intAnotherArr)
	fmt.Println("## Slices")
	var intSlice []int32 = []int32{10, 20, 30}
	intSlice = append(intSlice, 50)
	fmt.Println(intSlice)

	fmt.Println("## Maps")
	var intMap map[string]int32 = map[string]int32{
		"one": 1,
		"two": 2,
	}
	fmt.Println(intMap)

	fmt.Println("### Maps using make")
	anotherMap := make(map[string]uint8)
	fmt.Println(anotherMap)
	anotherMap["Yogendra"] = 33
	anotherMap["Srijana"] = 30
	fmt.Println(anotherMap)
	fmt.Printf("Age of Yogendra is %v\n", anotherMap["Yogendra"])
	delete(anotherMap, "Yogendra")
	fmt.Println(anotherMap)

	for key, value := range anotherMap {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}

	fmt.Println("## Strings")
	myString := "Hemmstraße"
	fmt.Println(len(myString))
	for i, v := range myString {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
	fmt.Println(myString[0])
	myAnotherString := []rune(myString)
	fmt.Println("Length of myAnotherString: ", len(myAnotherString))

	fmt.Println("## Structs")

	yt := student{
		name:       "Yogendra",
		age:        33,
		rollNumber: "BEX-066-448",
		gender:     "male",
	}
	yt.printStudentDetails()

}
