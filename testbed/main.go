package main

import (
	"fmt"
	"reflect"

	"github.com/wolfired/golabs/auto"
	"github.com/wolfired/golabs/bitorent"
)

//User user
type User struct {
	// Ages []string `torrent:"ages"` //slice
	// Ages map[string]string `torrent:"ages"` //map
	Ages map[string]Age `torrent:"ages"` //map
}

//Age age
type Age struct {
	Age uint64 `torrent:"age"`
}

func main() {
	// bs, _ := ioutil.ReadFile("./mmm.torrent")
	// mi := auto.MakeStruct(&bitorent.MetaInfo{}).(*bitorent.MetaInfo)
	// bitorent.Decode(bs, mi)
	// fmt.Println(mi)
	// fmt.Println(mi.Info)

	// var t interface{} = int(1)

	// switch t.(type) {
	// case int:
	// 	fmt.Println("int")
	// case int8:
	// 	fmt.Println("int8")
	// case int16:
	// 	fmt.Println("int16")
	// case int32:
	// 	fmt.Println("int32")
	// case int64:
	// 	fmt.Println("int64")
	// default:
	// 	fmt.Printf("unexpected type %T\n", t)
	// }

	u := auto.MakeStruct(User{}).(*User)
	// bitorent.Decode([]byte("d4:agesl1:a1:bee"), u) //slice
	// bitorent.Decode([]byte("d4:agesd1:a1:aee"), u)    //map
	bitorent.Decode([]byte("d4:agesd1:ad3:agei3eeee"), u) //map

	fmt.Println(u)
}

func doit(u interface{}) interface{} {
	v := reflect.ValueOf(u)

	if reflect.Ptr != v.Kind() {
		v = reflect.New(v.Type())
		u = v.Interface()
	}

	v = v.Elem()

	fmt.Println(v.Field(1).CanAddr())
	fmt.Println(v.Field(1).CanSet())

	v.Field(1).Set(reflect.ValueOf(Age{1}))
	fmt.Println(u)

	fmt.Println(v.Field(1).Field(0).CanAddr())
	fmt.Println(v.Field(1).Field(0).CanSet())

	v.Field(1).Field(0).SetInt(2)
	fmt.Println(u)

	v.Field(1).Field(0).Set(reflect.ValueOf(3))
	fmt.Println(u)

	return u
}
