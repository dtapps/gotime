package gotime

import "testing"

func TestVerification(t *testing.T) {
	t.Log(Verification("2022-02-05 00:00:00", DateTimeFormat))
}
