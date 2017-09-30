package auto

import "reflect"

/*MakeInstance 自动构造Instance*/
func MakeInstance(blueprint interface{}) interface{} {
	bv := reflect.ValueOf(blueprint)

	if reflect.Ptr != bv.Kind() {
		bv = reflect.New(bv.Type())
	}

	makeInstance(bv)

	return bv.Interface()
}

func shuntMake() {

}

func makeInstance(bv reflect.Value) {
	if reflect.Ptr == bv.Kind() {
		bv = bv.Elem()
	}
	bt := bv.Type()

	for i := 0; i < bv.NumField(); i++ {
		fv := bv.Field(i)

		switch fv.Kind() {
		case reflect.Map:
			// fv.Set(reflect.MakeMap(fv.Type()))
		case reflect.Slice:
			// fv.Set(reflect.MakeSlice(fv.Type(), 0, 0))
		case reflect.Chan:
			// fv.Set(reflect.MakeChan(fv.Type(), 0))
		case reflect.Struct:
			makeInstance(fv)
		case reflect.Ptr:
			ft := fv.Type().Elem()

			if bt != ft {
				nv := reflect.New(ft)
				makeInstance(nv)
				fv.Set(nv)
			}
		}
	}
}
