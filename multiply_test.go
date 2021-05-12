package multiply

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	testCases := []struct {
		A        float32
		B        float32
		Expected float32
	}{
		{A: 2, B: 3, Expected: 6},
		{A: 100, B: 0.5, Expected: 50},
		{A: -2, B: 3, Expected: -6},
		{A: -2, B: -90, Expected: 180},
		{A: -2.99, B: -2.7, Expected: 8.073},
	}
	for _, tcase := range testCases {
		result, _ := Multiply(&tcase.A, &tcase.B)
		assert.NotNil(t, result)
		assert.Equal(t, *result, tcase.Expected, "")
	}
}

func TestMultiplyLongNumbers(t *testing.T) {
	var err error

	var a float32 = 0.000000008 // 9 decimal digits
	var b float32 = 0.000000002 // 9 decimal digit
	var m float32 = 2           // 9 decimal digitss

	c := a * b
	t.Logf("\n%v\n%.20f\n%g\n%T\n", c, c, c, c)
	var n float32 = 0.0000006 // 7 decimal digits
	d, err := Multiply(&n, &m)
	if err != nil {
		t.Fatalf("test case threw an unexpected error %v\n", err)
	}
	expectedD := float32(0.0000012)
	t.Logf("d variable:\n%[1]v\n%.7[1]f\n%[1]g\n%[1]T\n", *d)
	if *d != expectedD {
		t.Fatalf("%v != %v \n", d, expectedD)
	}

	p := float32(math.MaxFloat32 / 1.9)
	q := float32(2)
	// var r *float32
	r, err := Multiply(&p, &q)
	if err == nil {
		t.Fatalf("r variable is NOT infinite: %v\n", r)
	}
	assert.Equal(t, r, (*float32)(nil), "")
	t.Logf("r variable is infinite: %v\n", r)

}

func TestValidationNumber(t *testing.T) {
	// test nil value
	var c *float32 = nil
	err := validateNumber(c)
	fmt.Printf("error: %v\n", err)
	assert.Nil(t, c)
	assert.EqualError(t, err, "Number is nil", "")

	// test on infinite value
	infValue := float32(math.Inf(0))
	err = validateNumber(&infValue)
	fmt.Printf("error: %v\n", err)
	assert.Nil(t, c)
	assert.Contains(t, err.Error(), "Number is really big")

	// test on correct value
	value := float32(10)
	err = validateNumber(&value)
	fmt.Printf("error: %v\n", err)
	assert.NoError(t, err)
}
