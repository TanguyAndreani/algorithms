package observer_pattern

type ObserverLoadbar interface {
	update(int, int)
}

type ObservableLoadbar interface {
	increment(int) int
	decrement(int) int
	register(ObserverLoadbar)
	update_observers()
}

type Loadbar struct{
	observers []ObserverLoadbar
	progress int
	max int
	min int
}

func (l *Loadbar) increment(x int) int {
	if l.progress+x <= l.max {
		l.progress += x
		l.update_observers()
	} else {
		if l.progress != l.max {
			l.update_observers()
		}
		l.progress = l.max
	}
	return l.progress
}

func (l *Loadbar) decrement(x int) int {
	if l.progress-x >= l.min {
		l.progress -= x
		l.update_observers()
	} else {
		if l.progress != l.min {
			l.update_observers()
		}
		l.progress = l.min
	}
	return l.progress
}

func (l *Loadbar) register(o ObserverLoadbar) {
	l.observers = append(l.observers, o)
}

func (l *Loadbar) update_observers() {
	for _, o := range l.observers {
		o.update(l.progress, l.progress / l.max * 100)
	}
}
