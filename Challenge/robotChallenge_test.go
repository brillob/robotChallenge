package challenge_test

import (
	challenge "challenge/robotwarehouse/Challenge"
	"testing"
)

func TestChallenge(t *testing.T) {
	if challenge.BestRobot() != "Robot" {
		t.Fatal("Not what we are expecting")
	}

}
