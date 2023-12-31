package greetings

import (
	"regexp"
	"testing"
)

func TestSayHello(t *testing.T) {
	SayHello()

}

func TestSayHello2_empty(t *testing.T) {
	expectedValue := regexp.MustCompile("name is mandatory")

	msg, err := SayHelloV2("")

	if msg != "" || err == nil {
		t.Fatalf(`Actual: msg=%q, error=%v, \n Expected: msg=%#q, error=nil`, msg, err, expectedValue)
		// t.Fatalf(`SayHelloV2("") = %q, %v, want "", error`, msg, err)
	}

	if !expectedValue.MatchString(err.Error()) {
		t.Fatalf("Actual: %v, Expected: %v", err, expectedValue.String())
	}
}
