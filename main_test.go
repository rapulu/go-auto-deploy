package main

import (
	"testing"
) 

func TestAdd(t *testing.T){
	got := Add(5, 6)

	if got != 11{
		t.Errorf("Wants 11 got %v", got)
	}

}