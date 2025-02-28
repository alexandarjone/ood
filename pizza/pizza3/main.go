package main

type Topping interface {
	GetName() string
	GetPrice() float64
}

type Size interface {
	GetSize() string
	GetPrice() float64
}

type Pizza interface {
	GetID() int64
	AddTopping(Topping)
	RemoveTopping(Topping)
	GetPrice() float64
}

type PizzaOrder interface {
	GetID() int64
	AddPizza(Size)
	// RemovePizza takes pizza id and remove it
	RemovePizza(int64)
	// AddTopping takes pizza id and topping name, and add the topping to the pizza
	AddTopping(int64, Topping)
	// RemoveTopping takes pizza id and topping name, and remove the topping to the pizza
	RemoveTopping(int64, Topping)
	GetPrice() float64
}

type OrderCenter interface {
	NewOrder() int64
	// AddPizza takes order id, pizza size name, and add a new pizza with the given size to order
	AddPizza(int64, string)
	// RemovePizza takes order id, pizza id and remove it
	RemovePizza(int64, int64)
	// AddTopping takes order id, pizza id and topping name, and add the topping to the pizza
	AddTopping(int64, int64, string)
	// RemoveTopping takes order id, pizza id and topping name, and remove the topping to the pizza
	RemoveTopping(int64, int64, string)
	GetPrice(int64) float64
	AddToppingType(string, float64)
	AddSizeType(string, float64)
}

// pizza implementation
var pizzaID int64
type pizza struct {
	id int64
	size Size
	// topping => count
	toppings map[Topping]int
}


func (p *pizza) GetID() int64 {
	return p.id
}

func (p *pizza) AddTopping(newTopping Topping) {
	p.toppings[newTopping]++
}

func (p *pizza) RemoveTopping(targetTopping Topping) {
	p.toppings[targetTopping]--
	if p.toppings[targetTopping] <= 0 {
		delete(p.toppings, targetTopping)
	}
}

func (p *pizza) GetPrice() float64 {
	totalPrice := p.size.GetPrice()
	for topping, count := range p.toppings {
		totalPrice += topping.GetPrice() * float64(count)
	}
	return totalPrice
}

func NewPizza(size Size) Pizza {
	pizzaID++
	return &pizza{
		id: pizzaID,
		size: size,
		toppings: make(map[Topping]int),
	}
}


// pizza order implementation
var pizzaOrderID int64
type pizzaOrder struct {
	id int64
	// pizza id => Pizza
	pizzas map[int64]Pizza
}

func (p *pizzaOrder) GetID() int64 {
	return p.id
}

func (p *pizzaOrder) AddPizza(size Size) {
	newPizza := NewPizza(size)
	p.pizzas[newPizza.GetID()] = newPizza
}

// RemovePizza takes pizza id and remove it
func (p *pizzaOrder) RemovePizza(id int64) {
	delete(p.pizzas, id)
}

// AddTopping takes pizza id and topping name, and add the topping to the pizza
func (p *pizzaOrder) AddTopping(pizzaID int64, newTopping Topping) {
	p.pizzas[pizzaID].AddTopping(newTopping)
}

// RemoveTopping takes pizza id and topping name, and remove the topping to the pizza
func (p *pizzaOrder) RemoveTopping(pizzaID int64, targetTopping Topping) {
	p.pizzas[pizzaID].RemoveTopping(targetTopping)
}

func (p *pizzaOrder) GetPrice() float64 {
	totalPrice := 0.0
	for _, pizza := range p.pizzas {
		totalPrice += pizza.GetPrice()
	}
	return totalPrice
}

func NewPizzaOrder() PizzaOrder {
	pizzaOrderID++
	return &pizzaOrder{
		id: pizzaOrderID,
		pizzas: make(map[int64]Pizza),
	}
}


// OrderCenter implementation
type orderCenter struct {
	// topping name => Topping
	toppings map[string]Topping
	// size => Size
	sizes map[string]Size
	// order id => PizzaOrder
	orders map[int64]PizzaOrder
}

func (o *orderCenter) NewOrder() int64 {
	newOrder := NewPizzaOrder()
	o.orders[newOrder.GetID()] = newOrder
	return newOrder.GetID()
}

// AddPizza takes order id, pizza size name, and add a new pizza with the given size to order
func (o *orderCenter) AddPizza(orderID int64, size string) {
	o.orders[orderID].AddPizza(o.sizes[size])
}

// RemovePizza takes order id, pizza id and remove it
func (o *orderCenter) RemovePizza(orderID int64, pizzaID int64) {
	o.orders[orderID].RemovePizza(pizzaID)
}

// AddTopping takes order id, pizza id and topping name, and add the topping to the pizza
func (o *orderCenter) AddTopping(orderID int64, pizzaID int64, newTopping string) {
	o.orders[orderID].AddTopping(pizzaID, o.toppings[newTopping])
}

// RemoveTopping takes order id, pizza id and topping name, and remove the topping to the pizza
func (o *orderCenter) RemoveTopping(orderID int64, pizzaID int64, targetTopping string) {
	o.orders[orderID].RemoveTopping(pizzaID, o.toppings[targetTopping])
}

func (o *orderCenter) GetPrice(orderID int64) float64 {
	return o.orders[orderID].GetPrice()
}

func (o *orderCenter) AddToppingType(toppingName string, price float64) {
	// TODO
}
func (o *orderCenter) AddSizeType(string, float64) {
	// TODO
}

type OrderCenterConfig struct {
	// topping name => Topping
	toppings map[string]Topping
	// size => Size
	sizes map[string]Size
}

func NewOrderCenter(config OrderCenterConfig) OrderCenter {
	return &orderCenter{
		toppings: config.toppings,
		sizes: config.sizes,
		orders: make(map[int64]PizzaOrder),
	}
}