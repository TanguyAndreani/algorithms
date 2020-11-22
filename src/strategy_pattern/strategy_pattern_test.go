package strategy_pattern

import (
	"regexp"
	"testing"
)

func TestDuckStrategyPattern(t *testing.T) {
	donnie := Duck{PondQuacker{}, CityFlyer{}, "Donald Duck"}

	matched, err := regexp.MatchString(`pond`, donnie.Quack())
	if !matched {
		t.Error("Expected a pond quacker!", err)
	}

	matched, err = regexp.MatchString(`city`, donnie.Fly())
	if !matched {
		t.Error("Expected a city flyer!", err)
	}

	donnie.quacker = CityQuacker{}

	matched, err = regexp.MatchString(`city`, donnie.Quack())
	if !matched {
		t.Error("Expected a city quacker!", err)
	}
}
