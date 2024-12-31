package utilstest

import "testing"

func AssertErrorIsNil(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Failed to execute getFeed gotted a error, %v", err)
		t.FailNow()
	}
}
