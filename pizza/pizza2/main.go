package main

type PizzaTopping interface {
	GetName() string
	GetPrice() float64
}

type PizzaSizeType int
const (
	Small PizzaSizeType = iota
	Medium
	Large
)
type PizzaSize interface {
	GetSizeType() PizzaSizeType
	GetPrice() float64
}

type Pizza interface {
	AddTopping(PizzaTopping)
	GetToppings() []PizzaTopping
	RemoveTopping(PizzaTopping) // should it be a name?
	GetSize() PizzaSize
	GetPrice() float64
}

type PizzaOrder interface {
	GetPizzas() []Pizza
	GetPrice() float64
	AddPizza(PizzaSizeType)
}

// PizzaTopping implementation
type generalPizzaTopping struct {
	name string
	price float64
}

func (g *generalPizzaTopping) GetName() string {
	return g.name
}

func (g *generalPizzaTopping) GetPrice() float64 {
	return g.price
}

func NewGeneralPizzaTopping(name string, price float64) PizzaTopping {
	return &generalPizzaTopping{
		name: name,
		price: price,
	}
}

// PizzaSize implementation
type generalPizzaSize struct {
	sizeType PizzaSizeType
	price float64
}

func (g *generalPizzaSize) GetSizeType() PizzaSizeType {
	return g.sizeType
}

func (g *generalPizzaSize) GetPrice() float64 {
	return g.price
}

func NewGeneralPizzaSize(sizeType PizzaSizeType, price float64) PizzaSize {
	return &generalPizzaSize{
		sizeType: sizeType,
		price: price,
	}
}

// Pizza implementation
type pizza struct {
	toppings []PizzaTopping
	size PizzaSize
	price float64
}
func (p *pizza) AddTopping(topping PizzaTopping) {
	p.toppings = append(p.toppings, topping)
	p.price += topping.GetPrice()
}
func (p *pizza) GetToppings() []PizzaTopping {
	return p.toppings
}
func (p *pizza) RemoveTopping(removeTopping PizzaTopping) {
	for i, topping := range p.toppings {
		if topping == removeTopping {
			p.toppings = append(p.toppings[:i], p.toppings[i+1:]...)
			return
		}
	}
}
func (p *pizza) GetSize() PizzaSize {
	return p.size
}
func (p *pizza) GetPrice() float64 {
	return p.price
}
func NewPizza(size PizzaSize) Pizza {
	return &pizza{
		size: size,
		price: size.GetPrice(),
	}
}

// PizzaOrder implementation
type pizzaOrder struct {
	pizzas []Pizza
}

func (p *pizzaOrder) GetPizzas() []Pizza {
	return p.pizzas
}
func (p *pizzaOrder) GetPrice() float64 {
	totalPrice := 0.0
	for _, pizza := range p.pizzas {
		totalPrice += pizza.GetPrice()
	}
	return totalPrice
}
func (p *pizzaOrder) AddPizza(pizza Pizza) {
	p.pizzas = append(p.pizzas, pizza)
}
