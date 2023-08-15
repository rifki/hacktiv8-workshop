// package main

// import "fmt"

// type ICalculate interface {
// 	add() int
// 	subtract() int
// }

// type calculate struct {
// 	a int
// 	b int
// }

// func (calc calculate) add() int {
// 	return calc.a + calc.b
// }

// func (calc calculate) subtract() int {
// 	return calc.a - calc.b
// }

// func NewCalculate(a, b int) ICalculate {
// 	return &calculate{
// 		a: a,
// 		b: b,
// 	}
// }

// func main() {
// 	calc := NewCalculate(5, 5)
// 	calcAdd := calc.add()
// 	calcSubt := calc.subtract()
// 	// print
// 	fmt.Println(calcAdd)  // Result: 10
// 	fmt.Println(calcSubt) // Result: 0
// }

// /*
// Step by step
// $ cd $GOPATH/src/
// $ mkdir -p github.com/rifki/hacktiv8-workshop
// $ go mod init
// # go mod tidy
// $ go run helloworld.go
// */
