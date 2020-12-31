package singletonpattern

import "testing"

func TestGetInstance(t *testing.T){
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("new object must have been made")
	}
	expectedCounter := counter1
	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Error("count must be one")
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("instances must be same")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Error("count must be two")
	}
}