// Area of circle and rectangle
package shape_area

import "math"

type Rectangle struct{
	length float64
	width float64
}
//function
func Area(r Rectangle) float64{
	return r.length * r.width
}

type Circle struct{
	radius float64
}
//method
func (c Circle)Area() float64{
	return math.Pi * c.radius * c.radius
}
