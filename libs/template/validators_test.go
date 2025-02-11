package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidatorString(t *testing.T) {
	err := validateString("abc")
	assert.NoError(t, err)

	err = validateString(1)
	assert.ErrorContains(t, err, "expected type string, but value is 1")

	err = validateString(true)
	assert.ErrorContains(t, err, "expected type string, but value is true")

	err = validateString("false")
	assert.NoError(t, err)
}

func TestValidatorBoolean(t *testing.T) {
	err := validateBoolean(true)
	assert.NoError(t, err)

	err = validateBoolean(1)
	assert.ErrorContains(t, err, "expected type boolean, but value is 1")

	err = validateBoolean("abc")
	assert.ErrorContains(t, err, "expected type boolean, but value is \"abc\"")

	err = validateBoolean("false")
	assert.ErrorContains(t, err, "expected type boolean, but value is \"false\"")
}

func TestValidatorNumber(t *testing.T) {
	err := validateNumber(true)
	assert.ErrorContains(t, err, "expected type float, but value is true")

	err = validateNumber(int32(1))
	require.NoError(t, err)

	err = validateNumber(int64(1))
	require.NoError(t, err)

	err = validateNumber(float32(1))
	assert.NoError(t, err)

	err = validateNumber(float64(1))
	assert.NoError(t, err)

	err = validateNumber("abc")
	assert.ErrorContains(t, err, "expected type float, but value is \"abc\"")
}

func TestValidatorInt(t *testing.T) {
	err := validateInteger(true)
	assert.ErrorContains(t, err, "expected type integer, but value is true")

	err = validateInteger(int32(1))
	assert.NoError(t, err)

	err = validateInteger(int64(1))
	assert.NoError(t, err)

	err = validateInteger(float32(1))
	assert.ErrorContains(t, err, "expected type integer, but value is 1")

	err = validateInteger(float64(1))
	assert.ErrorContains(t, err, "expected type integer, but value is 1")

	err = validateInteger("abc")
	assert.ErrorContains(t, err, "expected type integer, but value is \"abc\"")
}
