package main

import (
	"fmt"
	"time"
)

// customer struct
type customer struct {
	name  string
	phone string
}

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosec precession
	customer            //we get info of customer struct
}

// constructor
func newOrder(id string, amount float32, status string) *order {
	//initial setup goes here
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o *order) getAmount() float32 {
	return o.amount
}

func main() {
	// myOrder1 := order{
	// 	id:     "1",
	// 	amount: 100.00,
	// 	status: "pending",
	// }

	// myOrder2 := order{
	// 	id:        "2",
	// 	amount:    200.00,
	// 	status:    "pending",
	// 	createdAt: time.Now(),
	// }

	// myOrder1.createdAt = time.Now()

	// fmt.Println("Order1 struct", myOrder1)
	// // fmt.Println("Order2 struct", myOrder2)

	// fmt.Println("Order id:", myOrder1.id)

	// //using the changeStatus method
	// myOrder1.changeStatus("processing")
	// fmt.Println("Order1 struct", myOrder1)

	//utilizing constructor
	// myOrder := newOrder("1", 100.00, "pending")

	// fmt.Println(myOrder)

	//intializing a struct without blueprint
	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{"Go", true}

	// fmt.Println(language)

	//using embedded struct
	newOrder := order{
		id:        "1",
		amount:    100.00,
		status:    "pending",
		createdAt: time.Now(),
		customer: customer{
			name:  "John Doe",
			phone: "123-456-7890",
		},
	}

	fmt.Println(newOrder)

}
