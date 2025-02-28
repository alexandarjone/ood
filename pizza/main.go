package main

import "fmt"

// Topping struct
type Topping struct {
	name string
	price float64
}

func NewTopping(name string, price float64) *Topping {
	return &Topping{name, price}
}

func (t *Topping) GetName() string {
	return t.name
}

func (t *Topping) GetPrice() float64 {
	return t.price
}

// Size struct
type Size struct {
	name string
	price float64
}

func NewSize(name string, price float64) *Size {
	return &Size{name, price}
}

func (s *Size) GetName() string {
	return s.name
}

func (s *Size) GetPrice() float64 {
	return s.price
}

// PizzaMenu struct
type PizzaMenu struct {
	toppings map[string]*Topping
	sizes map[string]*Size
}

func NewPizzaMenu() *PizzaMenu {
	return &PizzaMenu{
		toppings: make(map[string]*Topping),
		sizes: make(map[string]*Size),
	}
}

func (m *PizzaMenu) AddTopping(topping *Topping) {
	m.toppings[topping.name] = topping
}

func (m *PizzaMenu) AddSize(size *Size) {
	m.sizes[size.name] = size
}

func (m *PizzaMenu) GetTopping(name string) *Topping {
	return m.toppings[name]
}

func (m *PizzaMenu) GetSize(name string) *Size {
	return m.sizes[name]
}

func (m *PizzaMenu) GetAllToppings() []*Topping {
	var result []*Topping
	for _, topping := range m.toppings {
		result = append(result, topping)
	}
	return result
}

func (m *PizzaMenu) GetAllSizes() []*Size {
	var result []*Size
	for _, size := range m.sizes {
		result = append(result, size)
	}
	return result
}

// Pizza struct
type Pizza struct {
	size *Size
	toppings []*Topping
}

func NewPizza(size *Size) *Pizza {
	return &Pizza{
		size: size,
		toppings: []*Topping{},
	}
}

func (p *Pizza) AddTopping(topping *Topping) {
	p.toppings = append(p.toppings, topping)
}

func (p *Pizza) RemoveTopping(toppingName string) {
	var newToppings []*Topping
	for _, topping := range p.toppings {
		if topping.GetName() != toppingName {
			newToppings = append(newToppings, topping)
		}
	}
	p.toppings = newToppings
}

func (p *Pizza) CalculateCost() float64 {
	total := p.size.GetPrice()
	for _, topping := range p.toppings {
		total += topping.price
	}
	return total
}

func (p *Pizza) GetDescription() string {
	description := fmt.Sprintf("Pizza Size: %s ($%.2f)", p.size.GetName(), p.size.GetPrice())
	if len(p.toppings) > 0 {
		description += " with toppings: "
		for _, topping := range p.toppings {
			description += fmt.Sprintf("%s ($%.2f), ", topping.GetName(), topping.GetPrice())
		}
		description = description[:len(description)-2] // Remove trailing comma
	} else {
		description += " (No toppings)"
	}
	return description
}

// Oder struct
type Order struct {
	pizzas []*Pizza
}

func NewOrder() *Order {
	return &Order{
		pizzas: []*Pizza{},
	}
}

func (o *Order) AddPizza(pizza *Pizza) {
	o.pizzas = append(o.pizzas, pizza)
}

func (o *Order) RemovePizza(index int) {
	if index >= 0 && index < len(o.pizzas) {
		o.pizzas = append(o.pizzas[:index], o.pizzas[index+1:]...)
	}
}

func (o *Order) CalculateTotal() float64 {
	total := 0.0
	for _, pizza := range o.pizzas {
		total += pizza.CalculateCost()
	}
	return total
}

func (o *Order) GetOrderSummary() string {
	summary := "Order Summary:\n"
	for i, pizza := range o.pizzas {
		summary += fmt.Sprintf("Pizza %d: %s | Cost: $%.2f\n", i+1, pizza.GetDescription(), pizza.CalculateCost())
	}
	summary += fmt.Sprintf("Total Order Cost: $%.2f", o.CalculateTotal())
	return summary
}

// Main function to test implementation
func main() {
	menu := NewPizzaMenu()

	// Add Sizes
	menu.AddSize(NewSize("Small", 5.00))
	menu.AddSize(NewSize("Medium", 7.50))
	menu.AddSize(NewSize("Large", 10.00))

	// Add Toppings
	menu.AddTopping(NewTopping("Cheese", 1.50))
	menu.AddTopping(NewTopping("Pepperoni", 2.00))
	menu.AddTopping(NewTopping("Olives", 0.75))
	menu.AddTopping(NewTopping("Mushrooms", 1.25))

	// Create Pizzas
	pizza1 := NewPizza(menu.GetSize("Medium"))
	pizza1.AddTopping(menu.GetTopping("Cheese"))
	pizza1.AddTopping(menu.GetTopping("Pepperoni"))

	pizza2 := NewPizza(menu.GetSize("Large"))
	pizza2.AddTopping(menu.GetTopping("Olives"))
	pizza2.AddTopping(menu.GetTopping("Mushrooms"))

	// Create an Order
	order := NewOrder()
	order.AddPizza(pizza1)
	order.AddPizza(pizza2)

	// Print Order Summary
	fmt.Println(order.GetOrderSummary())
}