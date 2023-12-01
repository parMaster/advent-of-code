package main

import "testing"

func Test_MaxCalories_PrimitiveInput(t *testing.T) {
	maxCalories, err := maxcalories("input0.txt")
	if err != nil {
		t.Errorf("maxcalories returned error: %e", err)
	}

	if maxCalories != 24000 {
		t.Errorf("maxCalories: %d, expected: %d", maxCalories, 24000)
	}
}

func Test_MaxCalories(t *testing.T) {
	maxCalories, err := maxcalories("input.txt")
	if err != nil {
		t.Errorf("maxcalories returned error: %e", err)
	}

	if maxCalories != 67016 {
		t.Errorf("maxCalories: %d, expected: %d", maxCalories, 67016)
	}
}

func Test_Top3Calories(t *testing.T) {
	top3Calories, err := top3calories("input.txt")
	if err != nil {
		t.Errorf("top3calories returned error: %e", err)
	}

	if top3Calories != 200116 {
		t.Errorf("top3Calories: %d, expected: %d", top3Calories, 200116)
	}
}
