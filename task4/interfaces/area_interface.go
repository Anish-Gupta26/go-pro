package area

import "math"

type Shape interface{
	Area() float64
}

type Rectangle struct{
	length float64
	width float64
}

func (r Rectangle) Area() float64{
	return r.length * r.width
}

type Circle struct{
	radius float64
}

func (c Circle)Area() float64{
	return math.Pi* c.radius *c.radius
}