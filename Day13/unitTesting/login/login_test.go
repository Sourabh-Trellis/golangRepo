package login_test

import (
	"testing"
)

func TestLogin(t *testing.T) {

	// 1. arrange
	expected := "Login successfull."

	// 2. act
	actual := Login("sourabh@trellis", "sourabh123")

	// 3.assert
	if expected != actual {
		t.Errorf("Expected '%v',got '%v'", expected, actual)
	}
}
