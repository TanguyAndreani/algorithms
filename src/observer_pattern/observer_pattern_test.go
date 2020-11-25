package observer_pattern

import (
	"testing"
)

type MyObserver struct{}

var count int

func (o *MyObserver) update(progress, percentage int) {
	count = progress
}

func TestLoadbar(t *testing.T) {
	max := 100

	bar := Loadbar{max: max}
	obs := MyObserver{}

	bar.register(&obs)

	for i := 0; i < 140; /*20*/ i++ {
		bar.increment(5)
	}

	if count != max {
		t.Error("count was not incremented!", count)
	}
}
