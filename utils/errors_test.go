package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test for splitErrorTagfunc TestSplitErrorTag() {
func TestSplitErrorTag(t *testing.T) {
	err := errors.New("test error: this is test error")
	tag, err := SplitErrorTag(err)

	assert.NoError(t, err)
	assert.Equal(t, "test error", tag)
}

// test for findStatusCodeByErr
func TestFindStatusCodeByErr(t *testing.T) {
	err := errors.New("")
	unauthorizedErr := MakeError(err, UnauthorizedErrorTag)
	unexpectedErr := MakeError(err, "unexpected error")

	unauthorizedErrStatusCode := FindStatusCodeByErr(&unauthorizedErr)
	unexpectedError := FindStatusCodeByErr(&unexpectedErr)

	assert.Equal(t, 401, unauthorizedErrStatusCode)
	assert.Equal(t, 500, unexpectedError)
}

// test for abordGin
func TestAbordGin(t *testing.T) {
	// todo: implement
}
