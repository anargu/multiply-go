package multiply

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetErrorMessage(t *testing.T) {
	var err error
	assert.Nil(t, GetErrorMessage(err))
	err = errors.New("message")
	assert.NotNil(t, GetErrorMessage(err))
	expected := "message"
	assert.Equal(t, *GetErrorMessage(err), expected)
}
