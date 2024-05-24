// package main

// import "fmt"

// type array struct {
// 	data []int
// }

// func (a *array) append(n int) {

// 	a.data = append(a.data, n)

// 	fmt.Println("appended item:", n)
// }

// func (a *array) printArrayData() {

// 	fmt.Println("array data:",a.data)

// }

// func (a *array) deletevalue(n int) {

// 	for i, v := range a.data {
// 		if v == n {
// 			a.data = append(a.data[:i], a.data[i+1:]...)
// 			fmt.Println("deleted item:", n)
// 		}
// 	}
// }

// func main() {

// 	myarray := array{}	
// 	myarray.append(10)
// 	myarray.append(20)
// 	myarray.append(30)
// 	myarray.append(40)
// 	myarray.append(50)
// 	myarray.append(60)
// 	myarray.append(70)
// 	myarray.append(80)
// 	myarray.append(90)
// 	myarray.append(100)
// 	myarray.append(110)
// 	myarray.append(120)
// 	myarray.append(130)
// 	myarray.append(140)
// 	myarray.append(150)
// 	myarray.append(160)
// 	myarray.append(170)
// 	myarray.append(180)
// 	myarray.append(190)
// 	myarray.append(200)

// 	myarray.deletevalue(10)
// }