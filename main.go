package main

import (
	"fmt"
)

func main() {
	water := 400
	milk := 540
	coffee := 120
	cups := 9
	money := 550

	run := true

	for run {
		action := getAction()

		run = onAction(action, &water, &milk, &coffee, &cups, &money)
	}
}

func getAction() string {
	var action string
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	fmt.Scan(&action)

	return action
}

func onAction(action string, water, milk, coffee, cups, money *int) bool {
	toContinue := true
	switch action {
	case "remaining":
		display(*water, *milk, *coffee, *cups, *money)
	case "buy":
		buy(water, milk, coffee, cups, money)
	case "fill":
		fill(water, milk, coffee, cups)
	case "take":
		take(money)
	case "exit":
		toContinue = false
	}

	return toContinue
}

func buy(water, milk, coffee, cups, money *int) {
	var sort string
	fmt.Println("What do you want to buy?\n1 - espresso,\n2 - latte,\n3 - cappuccino,\n4 - own flavor\nback - to main menu:")
	fmt.Scan(&sort)

	switch sort {
	case "1":
		makeEspresso(water, coffee, cups, money)
	case "2":
		makeLatte(water, milk, coffee, cups, money)
	case "3":
		makeCappuccino(water, milk, coffee, cups, money)
	case "4":
		ownFlavor(water, milk, coffee, cups, money)
	}
}

func ownFlavor(water, milk, coffee, cups, money *int) {
	const milkPriceCentsPerMl = 3
	const coffeePriceCentsPerGr = 25

	var waterAmount, milkAmount, coffeeAmount int

	fmt.Println("Write how many ml of water you want:")
	fmt.Scan(&waterAmount)

	fmt.Println("Write how many ml of milk you want:")
	fmt.Scan(&milkAmount)

	fmt.Println("Write how many grams of coffee beans you want:")
	fmt.Scan(&coffeeAmount)

	if checkIngradients(*water < waterAmount, *milk < milkAmount, *coffee < coffeeAmount, *cups < 1) {
		*water -= waterAmount
		*milk -= milkAmount
		*coffee -= coffeeAmount
		*money += int((milkAmount*milkPriceCentsPerMl + coffeeAmount*coffeePriceCentsPerGr) / 100)
		*cups -= 1

		fmt.Println("I have enough resources, making you a coffee!")
	}

}

func makeEspresso(water, coffee, cups, money *int) {
	if checkIngradients(*water < 250, false, *coffee < 16, *cups < 1) {
		*water -= 250
		*coffee -= 16
		*money += 4
		*cups -= 1

		fmt.Println("I have enough resources, making you a coffee!")
	}
}

func makeLatte(water, milk, coffee, cups, money *int) {
	if checkIngradients(*water < 350, *milk < 75, *coffee < 20, *cups < 1) {
		*water -= 350
		*milk -= 75
		*coffee -= 20
		*money += 7
		*cups -= 1

		fmt.Println("I have enough resources, making you a coffee!")
	}
}

func makeCappuccino(water, milk, coffee, cups, money *int) {
	if checkIngradients(*water < 200, *milk < 100, *coffee < 12, *cups < 1) {
		*water -= 200
		*milk -= 100
		*coffee -= 12
		*money += 6
		*cups -= 1

		fmt.Println("I have enough resources, making you a coffee!")
	}
}

func checkIngradients(lackOfWater, lackOfMilk, lackOfCoffee, lackOfCups bool) bool {
	if lackOfWater {
		fmt.Println("Sorry, not enough water!")
	} else if lackOfMilk {
		fmt.Println("Sorry, not enough milk!")
	} else if lackOfCoffee {
		fmt.Println("Sorry, not enough coffee!")
	} else if lackOfCups {
		fmt.Println("Sorry, not enough cups!")
	}

	return !lackOfWater && !lackOfMilk && !lackOfCoffee && !lackOfCups
}

func fill(water, milk, coffee, cups *int) {
	var addWater, addMilk, addCoffee, addCups int

	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&addWater)
	*water += addWater

	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&addMilk)
	*milk += addMilk

	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&addCoffee)
	*coffee += addCoffee

	fmt.Println("Write how many disposable cups you want to add:")
	fmt.Scan(&addCups)
	*cups += addCups
}

func take(money *int) {
	fmt.Printf("I gave you $%d\n", *money)
	*money = 0
}

func display(water, milk, coffee, cups, money int) {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", water)
	fmt.Printf("%d ml of milk\n", milk)
	fmt.Printf("%d g of coffee beans\n", coffee)
	fmt.Printf("%d disposable cups\n", cups)
	fmt.Printf("$%d of money\n", money)
}
