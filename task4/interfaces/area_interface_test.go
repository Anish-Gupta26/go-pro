package area

import "testing"

func TestProperty(t *testing.T){
	findArea := func (t *testing.T,s Shape,want float64){
		t.Helper()
		got := s.Area()
		if got!=want{
			t.Errorf("Error in test with %#v: - Got %g, but want %g",s,got,want)
		}
		t.Log("Success")
	}
	t.Run("Rectangle",func(t *testing.T) {
		obj := Rectangle{2,3}
		want := 6.0
		findArea(t,obj,want)

	})
	t.Run("Circle",func(t *testing.T) {
		obj := Circle{2.0}
		want := 5.0
		findArea(t,obj,want)
	})
}