package dbuser_test

import (
	"testing"

	"github.com/camiloperezv/samples/users_db/pkg/dbuser"
	"github.com/stretchr/testify/assert"
)

func TestCheckName(t *testing.T) {
	var responses = []struct {
		name     string // input
		expected bool   // expected result
		msg      string // message if test fails
	}{
		{"", false, "empty did not return false"},
		{"a", false, "one character 'a' did not return false"},
		{"b", false, "one character 'a' did not return false"},
		{"ba", true, "two characters 'ba' did not return true"},
		{"abc", true, "three characters 'ba' did not return true"},
		{"abcd", true, "four characters 'ba' did not return true"},
	}
	for _, v := range responses {
		assert.Equal(t, v.expected, dbuser.CheckName(v.name), v.msg)
	}
}
