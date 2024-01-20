package shape_area
import "testing"

func TestStruct(t *testing.T){
	obj := Rectangle{5,6}
	got:= Area(obj)
	want:= 30.0
	if got!=want{
		t.Errorf("Error: Got %v, but want %v",got,want)
	}
	t.Log("Success")
}