package assert

import (
	"fmt"
	"reflect"
)

func Equal(got interface{}, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		template := `--- Failed equality assertion ...
got type:
%T
got value:
%v
want type:
%T
want value:
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
