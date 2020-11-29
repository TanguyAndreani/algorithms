package decorator_pattern

type Beverage interface {
	cost() int
}

type Coffee struct{}

func (c Coffee) cost() int {
	return 2
}

type Tea struct{}

func (t Tea) cost() int {
	return 1
}

type BeverageWithSugar struct {
	b Beverage
}

func (b BeverageWithSugar) cost() int {
	return 1 + b.b.cost()
}

type BeverageWithChocolate struct {
	b Beverage
}

func (b BeverageWithChocolate) cost() int {
	return 3 + b.b.cost()
}
