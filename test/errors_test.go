package test

import (
	"errors"
	"testing"

	"github.com/ochanoco/torima/utils"
	"github.com/stretchr/testify/assert"
)

// test for splitErrorTagfunc TestSplitErrorTag() {
func TestSplitErrorTag(t *testing.T) {
	err := errors.New("test error: this is test error")
	tag, err := utils.SplitErrorTag(err)

	assert.NoError(t, err)
	assert.Equal(t, "test error", tag)
}

// test for findStatusCodeByErr
func TestFindStatusCodeByErr(t *testing.T) {
	err := errors.New("")
	unauthorizedErr := utils.MakeError(err, utils.UnauthorizedErrorTag)
	unexpectedErr := utils.MakeError(err, "unexpected error")

	unauthorizedErrStatusCode := utils.FindStatusCodeByErr(&unauthorizedErr)
	unexpectedError := utils.FindStatusCodeByErr(&unexpectedErr)

	assert.Equal(t, 401, unauthorizedErrStatusCode)
	assert.Equal(t, 500, unexpectedError)
}

// test for abordGin
func TestAbordGin(t *testing.T) {
	// todo: implement
}
