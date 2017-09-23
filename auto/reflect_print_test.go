package auto

import "testing"

func Test_PrintInstance(t *testing.T) {
	sun := struct{ Name string }{}
	PrintInstance(&sun)
}
