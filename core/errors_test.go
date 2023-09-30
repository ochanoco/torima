package core

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test for splitErrorTagfunc TestSplitErrorTag() {
func TestSplitErrorTag(t *testing.T) {
	err := errors.New("test error: this is test error")
	tag, err := splitErrorTag(err)

	assert.NoError(t, err)
	assert.Equal(t, "test error", tag)
}

// test for findStatusCodeByErr
func TestFindStatusCodeByErr(t *testing.T) {
	err := errors.New("")
	unauthorizedErr := makeError(err, unauthorizedErrorTag)
	unexpectedErr := makeError(err, "unexpected error")

	unauthorizedErrStatusCode := findStatusCodeByErr(&unauthorizedErr)
	unexpectedError := findStatusCodeByErr(&unexpectedErr)

	assert.Equal(t, 401, unauthorizedErrStatusCode)
	assert.Equal(t, 500, unexpectedError)
}

// test for abordGin
func TestAbordGin(t *testing.T) {
	// todo: implement
}
