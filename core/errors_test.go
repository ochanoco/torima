package core

import (
	"errors"
	"testing"
)

// test for splitErrorTagfunc TestSplitErrorTag() {
func TestSplitErrorTag(t *testing.T) {
	err := errors.New("test error: this is test error")
	tag, err := splitErrorTag(err)

	if err != nil {
		t.Errorf("splitErrorTag() is failed: %v", err)
	}
	if tag != "test error" {
		t.Errorf("splitErrorTag() is failed: %v", err)
	}
}

// test for findStatusCodeByErr
func TestFindStatusCodeByErr(t *testing.T) {
	err := errors.New("")
	unauthorizedErr := makeError(err, unauthorizedErrorTag)
	unexpectedErr := makeError(err, "unexpected error")

	unauthorizedErrStatusCode := findStatusCodeByErr(&unauthorizedErr)
	unexpectedError := findStatusCodeByErr(&unexpectedErr)

	if unauthorizedErrStatusCode != 401 {
		t.Errorf("findStatusCodeByErr() is failed: %v", err)
	}

	if unexpectedError != 500 {
		t.Errorf("findStatusCodeByErr() is failed: %v", err)
	}
}

// test for abordGin
func TestAbordGin(t *testing.T) {
	// todo: implement
}
