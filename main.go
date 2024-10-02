package main

import (
	carts "OpnCart/src"
	"fmt"
)

func main() {
	cart := carts.NewCart()

	cart.AddProduct("A001", 2000, 1)
	cart.AddProduct("A002", 500, 2)

	fmt.Println("Cart Items:", cart.ListItems())

	cart.UpdateProduct("A001", 10)
	fmt.Println("Updated Cart Items:", cart.ListItems())

	cart.ApplyDiscount(carts.Discount{Name: "9/9 sale", Percentage: 10, MaxAmount: 100})
	totalAfterDiscount := cart.ApplyDiscounts()
	fmt.Printf("Total Amount after discount: %.2f\n", totalAfterDiscount)

	cart.ApplyFreebie("1", "FreebieProduct")
	fmt.Println("Cart Items after applying freebie:", cart.ListItems())

}
