package shape_area
import "testing"

func TestArea(t *testing.T){
	t.Run("Rectangle area",func(t *testing.T) {
		obj := Rectangle{5.0,6.0}
		got:= Area(obj)
		want:= 30.0
		if got!=want{
			t.Errorf("Error: Got %v, but want %v",got,want)
		}
		t.Log("Success......")
	})
	t.Run("Circle area",func(t *testing.T) {
		obj := Circle{5.0}
		got := obj.Area()
		want := 78.53981633974483
		if got!=want{
			t.Errorf("Error: Got %v, but want %v",got,want)
		}
		t.Log("Success...")
	})
}