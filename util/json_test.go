package util

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestConvertObjectToJSON(t *testing.T) {
	Convey("Deve-se converter um objeto para JSON e vice-versa", t, func() {
		type T struct {
			Field string
		}
		obj := new(T)
		obj.Field = "A"
		obj2 := new(T)
		err := FromJSON(ToJSON(obj), obj2)
		So(err, ShouldEqual, nil)
		So(obj2.Field, ShouldEqual, obj.Field)
	})

}
