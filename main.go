package main

import (
	"fmt"

	inv "github.com/nmluci/KissatenService/libs/Inventory"
	mem "github.com/nmluci/KissatenService/libs/Membership"
)

func main() {
	// PlaceHolder
	fmt.Println(inv.ItemCount(), mem.MemberCount())

	for opt := 0; opt != -1; {
		fmt.Println("Welcome to SchneeKatze's Cafe!")
		fmt.Println("[1] Make an Order")
		fmt.Println("[2] Cancel an Order")
		fmt.Println("[3] Check Shopping Cart")
		fmt.Println("[4] Membership Management")
		fmt.Println("[5] Checkout")
		fmt.Println("[6] Exit")
		fmt.Println("Enter your choices")
		fmt.Scanf("%d", &opt)

		switch opt {
		case 1:
			// Make an Order
		case 2:
			// Remove an Order
		case 3:
			// Show Q
		case 4:
			// Show Member Interface
		case 5:
			// Proceed
		case 6:
			opt = -1
		default:
			if opt > 6 {
				fmt.Println("Be serious please... _=")
			} else {
				fmt.Println("Placeholder desu~")
			}
		}
	}
}
