package tests

import "testing"

func TestMainDirector(t *testing.T) {
	testerOk := MainDirectorTester{
		ResultBody: TEST_RESP_BODY1,
		Cookie:     "ochanoco-token=test",
	}
	RunCommonTest(t, &testerOk, "main/director_ok")

}
