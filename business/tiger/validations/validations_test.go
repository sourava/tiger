package validations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenDateIsInvalid_ThenValidateDateStringShouldReturnFalse(t *testing.T) {
	valid := ValidateDateString("invalid date")

	assert.False(t, valid)
}

func TestWhenDateIsValid_ThenValidateDateStringShouldReturnTrue(t *testing.T) {
	valid := ValidateDateString("2020-01-01")

	assert.True(t, valid)
}

func TestWhenLatitudeIsLessThanNegative90_ThenValidateLatitudeShouldReturnFalse(t *testing.T) {
	valid := ValidateLatitude(-91)

	assert.False(t, valid)
}

func TestWhenLatitudeIsMoreThan90_ThenValidateLatitudeShouldReturnFalse(t *testing.T) {
	valid := ValidateLatitude(91)

	assert.False(t, valid)
}

func TestWhenLatitudeIsBetweenNegative90AndPositive90_ThenValidateLatitudeShouldReturnTrue(t *testing.T) {
	valid := ValidateLatitude(89)

	assert.True(t, valid)
}

func TestWhenLongitudeIsLessThanNegative180_ThenValidateLongitudeShouldReturnFalse(t *testing.T) {
	valid := ValidateLongitude(-181)

	assert.False(t, valid)
}

func TestWhenLongitudeIsMoreThan180_ThenValidateLongitudeShouldReturnFalse(t *testing.T) {
	valid := ValidateLongitude(181)

	assert.False(t, valid)
}

func TestWhenLongitudeIsBetweenNegative180AndPositive180_ThenValidateLongitudeShouldReturnTrue(t *testing.T) {
	valid := ValidateLongitude(179)

	assert.True(t, valid)
}
