package main

type Warewhare struct {
}

func (w Warewhare) Robots() {
	var r Robot = Spot()
	x := [1]Robot{r}
	return x
}
