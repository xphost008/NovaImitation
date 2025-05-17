package main

import "testing"

func TestLaunchGameParam(t *testing.T) {
	if GetWindowsVersion() {
		t.Log("Yes")
	} else {
		t.Error("No")
	}
}
