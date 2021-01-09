package assert

import (
	"fmt"
	"reflect"
)

func Equal(want interface{}, got interface{}) {
	if !reflect.DeepEqual(got, want) {
		template := `--- Failed equality assertion ...
want type:
%T
want value:
%v
got type:
%T
got value:
%v
`
		msg := fmt.Sprintf(template, got, got, want, want)
		panic(msg)
	}
}

func True(condition bool, msg string) {
	if !condition {
		panic("Assertion failure: " + msg)
	}
}
