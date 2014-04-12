package go4java

type Programmer struct {
	caffeineLevel int
}

func (p *Programmer) DrinkCoffee() {
	p.caffeineLevel++
}
func (p *Programmer) isAwake() bool {
	return p.caffeineLevel > 42
}
