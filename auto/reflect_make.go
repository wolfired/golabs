package auto

import "reflect"

/*MakeStruct 自动构造Struct*/
func MakeStruct(blueprint interface{}) interface{} {
	bv := reflect.ValueOf(blueprint)

	if reflect.Ptr != bv.Kind() {
		bv = reflect.New(bv.Type())
	}

	makeStruct(bv)

	return bv.Interface()
}

func makeStruct(bv reflect.Value) {
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
			makeStruct(fv)
		case reflect.Ptr:
			ft := fv.Type().Elem()

			if bt != ft {
				nv := reflect.New(ft)
				makeStruct(nv)
				fv.Set(nv)
			}
		}
	}
}
