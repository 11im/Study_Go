package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// func Test(t *testing.T) {
// 	rst := square(9)
// 	if rst != 81 {
// 		t.Errorf("%d\n", rst)
// 	}
// }

// func Test2(t *testing.T) {
// 	rst := square(3)
// 	if rst != 9 {
// 		t.Errorf("%d\n", rst)
// 	}
// }

func Test3(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(81, square(9), "Error")
}

func Test4(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, square(3), "Error")
}
