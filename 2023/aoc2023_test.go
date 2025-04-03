package aoc

import (
	"log"
	"testing"
)

func TestDay1a(t *testing.T) {
	result, err := day1a()
	if err != nil {
		t.Errorf("day1a internal error")
	}
	log.Println("Result:", result)
}

func TestDay1b(t *testing.T) {
	result, err := day1b()
	if err != nil {
		t.Errorf("day1b internal error")
	}
	log.Println("Result:", result)
}
