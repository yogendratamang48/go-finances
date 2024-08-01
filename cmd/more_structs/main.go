package main

import "fmt"

type gasEngine struct {
	name     string
	mpg      float32
	galloons float32
	owner
}

type electricEngine struct {
	name  string
	mpkwh float32
	kwh   float32
	owner
}

type owner struct {
	name string
}

func (g gasEngine) milesLeft() float32 {
	return g.mpg * g.galloons
}

func (e electricEngine) milesLeft() float32 {
	return e.mpkwh * e.kwh
}

func canMakeit(e engine, miles float32) {
	if e.milesLeft() > miles {
		fmt.Println("You can make it")
	} else {
		fmt.Println("You can't make it")
	}
}

type engine interface {
	milesLeft() float32
}

func makeItSquare(array [5]int32) [5]int32 {
	for i, value := range array {
		array[i] = value * value
	}
	return array
}

func makeItSquareWithPointer(array *[5]int32) [5]int32 {
	for i := range array {
		array[i] = array[i] * array[i]
	}
	return *array
}

func main() {
	fmt.Println("# More structs")
	var myBmw gasEngine = gasEngine{name: "My BMW", mpg: 28, galloons: 10, owner: owner{name: "Yogendra"}}
	var myTeslaModel3 electricEngine = electricEngine{name: "My Tesla Model 3", mpkwh: 5.5, kwh: 100, owner: owner{name: "Yogendra"}}
	fmt.Println(myBmw)
	fmt.Println(myTeslaModel3)
	fmt.Println("Miles left for myBMW: ", myBmw.milesLeft())
	fmt.Println("Miles left for My Tesla: ", myTeslaModel3.milesLeft())
	fmt.Println("You can see that we milesLeft() method is needed for both Engines, but we don't have to implement it twice.")
	fmt.Println("Let's use interface to solve this problem")
	canMakeit(myBmw, 200)
	canMakeit(myTeslaModel3, 200)

	fmt.Println("# Pointers")
	var p *int32
	var anotherP *int32 = new(int32)
	var i int32 = 10
	fmt.Println("Pointer p: ", p)
	fmt.Println("Another pointer: ", anotherP)
	fmt.Printf("Value of anotherP points to: %v\n", *anotherP)
	fmt.Println("Value i: ", i)
	*anotherP = 100
	fmt.Printf("Value of anotherP points to: %v\n", *anotherP)
	fmt.Println("Let's point p to i's address")
	p = &i
	fmt.Println("Value p points to: ", *p)
	fmt.Println("Lets update p's value to 200")
	*p = 200
	fmt.Println("Value i: ", i)
	fmt.Println("Lets update i and see")
	i = 300
	fmt.Println("Value p points to: ", *p, "Value of p should also change")
	fmt.Println("Using pointers and functions")
	fmt.Println("When you call a function without a pointer, values will be copied.")
	myArray := [5]int32{1, 2, 3, 4, 5}
	fmt.Println("Input: ", myArray)
	fmt.Printf("Memory Location of myArray: %p\n", &myArray)
	anotherArray := makeItSquare(myArray)
	fmt.Println("output: ", anotherArray)
	fmt.Printf("Memory Location of AnotherArray: %p\n", &anotherArray)
	yetAnotherArray := makeItSquareWithPointer(&myArray)
	fmt.Printf("Memory Location of yetAnotherArray: %p\n", &yetAnotherArray)
	fmt.Printf("Memory Location of Original Array: %p\n", &myArray)
	fmt.Println("Output: ", yetAnotherArray)

}
