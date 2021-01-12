package collections_test

import (
	"testing"

	c "github.com/pvillela/go-light-collections/pkg/collections"
	"github.com/stretchr/testify/assert"
)

////
// Maps used as inputs to functions below.

var mEmpty = c.MapAnyAny{}

var mBase = c.MapAnyAny{1: "w1", 22: "w22", 333: "w333", 4444: "w4444"}

////
// Tests

func TestMapCopy(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.MapAnyAny
	}{
		{"Copy: non-empty map", mBase},
		{"Copy: empty map", mEmpty},
		{"Copy: map", nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Copy()
		assert.Equal(t, cs.receiver, got, cs.msg)
		assert.True(t, &cs.receiver != &got, cs.msg)
	}
}
