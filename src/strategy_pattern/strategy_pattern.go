package strategy_pattern

import (
	"fmt"
)

type Duck struct {
	quacker QuackStrategy
	flyer   FlyStrategy
	name    string
}

/* Passing down calls to duck's quacker and flyer */

func (d Duck) Quack() string {
	return d.quacker.Quack(d)
}

func (d Duck) Fly() string {
	return d.flyer.Fly(d)
}

/* Allows us to define quacking and flying strategies */

type QuackStrategy interface {
	Quack(Duck) string
}

type FlyStrategy interface {
	Fly(Duck) string
}

/* Quack strategies */

type CityQuacker struct{}

func (q CityQuacker) Quack(d Duck) string {
	return fmt.Sprintf("%s quacks like a city duck!\n", d.name)
}

type PondQuacker struct{}

func (q PondQuacker) Quack(d Duck) string {
	return fmt.Sprintf("%s quacks like a pond duck!\n", d.name)
}

/* Fly strategies */

type CityFlyer struct{}

func (q CityFlyer) Fly(d Duck) string {
	return fmt.Sprintf("%s flies like a city duck!\n", d.name)
}

type PondFlyer struct{}

func (q PondFlyer) Quack(d Duck) string {
	return fmt.Sprintf("%s flies like a pond duck!\n", d.name)
}
