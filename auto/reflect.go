package auto

import (
	"reflect"
)

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

/*FillStruct 自动填充Struct*/
func FillStruct(container interface{}, dict map[string]interface{}, tag string) {
	cv := reflect.ValueOf(container)

	fillStruct(cv, dict, tag)
}

func fillBool(v reflect.Value, x interface{}) {
	v.SetBool(x.(bool))
}

func fillInt(v reflect.Value, x interface{}) {
	xv := reflect.ValueOf(x)

	if v.Type() != xv.Type() {
		xv = xv.Convert(v.Type())
	}

	v.SetInt(xv.Int())
}

func fillUint(v reflect.Value, x interface{}) {
	xv := reflect.ValueOf(x)

	if v.Type() != xv.Type() {
		xv = xv.Convert(v.Type())
	}

	v.SetUint(xv.Uint())
}

func fillString(v reflect.Value, x interface{}) {
	v.SetString(x.(string))
}

func fillSlice(v reflect.Value, s []interface{}, tag string) {
	nv := reflect.MakeSlice(v.Type(), len(s), cap(s))
	kind := nv.Type().Elem().Kind()

	for i := 0; i < len(s); i++ {
		fv := nv.Index(i)
		sx := s[i]
		shunt(kind, fv, sx, tag)
	}

	v.Set(nv)
}

func fillMap(v reflect.Value, d map[string]interface{}, tag string) {
	nv := reflect.MakeMapWithSize(v.Type(), len(d))

	kt := nv.Type().Key()
	ktk := kt.Kind()
	vt := nv.Type().Elem()
	vtk := vt.Kind()

	for dk, dx := range d {
		kv := reflect.Indirect(reflect.New(kt))
		shunt(ktk, kv, dk, tag)

		xv := reflect.Indirect(reflect.New(vt))
		shunt(vtk, xv, dx, tag)

		nv.SetMapIndex(kv, xv)
	}

	v.Set(nv)
}

func fillStruct(v reflect.Value, d map[string]interface{}, tag string) {
	if reflect.Ptr == v.Kind() {
		v = v.Elem()
	}
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		ft := t.Field(i)

		dx, ok := d[ft.Tag.Get(tag)]
		if ok {
			shunt(fv.Kind(), fv, dx, tag)
		}
	}
}

func shunt(kind reflect.Kind, v reflect.Value, x interface{}, tag string) {
	switch kind {
	case reflect.Bool:
		fillBool(v, x)
	case reflect.Int:
		fillInt(v, x)
	case reflect.Int8:
		fillInt(v, x)
	case reflect.Int16:
		fillInt(v, x)
	case reflect.Int32:
		fillInt(v, x)
	case reflect.Int64:
		fillInt(v, x)
	case reflect.Uint:
		fillUint(v, x)
	case reflect.Uint8:
		fillUint(v, x)
	case reflect.Uint16:
		fillUint(v, x)
	case reflect.Uint32:
		fillUint(v, x)
	case reflect.Uint64:
		fillUint(v, x)
	case reflect.Map:
		fillMap(v, x.(map[string]interface{}), tag)
	case reflect.Ptr:
		fillStruct(v, x.(map[string]interface{}), tag)
	case reflect.Slice:
		fillSlice(v, x.([]interface{}), tag)
	case reflect.String:
		fillString(v, x)
	case reflect.Struct:
		fillStruct(v, x.(map[string]interface{}), tag)
	}
}
