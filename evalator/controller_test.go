package evalator_test

import (
	"evalator-kata/evalator"
	"testing"
)

func Test_CheckLoadLimit_Input_LiftWeight_1300_Should_Be_False(t *testing.T) {
	expected := false
	lift := evalator.Lift{
		Weight: 1300,
	}

	actual := lift.CheckLoadLimit()

	if expected != actual {
		t.Errorf("Expected is %v but get %v", expected, actual)
	}
}

func Test_CheckLoadLimit_Input_LiftWeight_1000_Should_Be_True(t *testing.T) {
	expected := true
	lift := evalator.Lift{
		Weight: 1000,
	}

	actual := lift.CheckLoadLimit()

	if expected != actual {
		t.Errorf("Expected is %v but get %v", expected, actual)
	}
}

func Test_CheckLoadLimit_Input_LiftWeight_900_Should_Be_True(t *testing.T) {
	expected := true
	lift := evalator.Lift{
		Weight: 900,
	}

	actual := lift.CheckLoadLimit()

	if expected != actual {
		t.Errorf("Expected is %v but get %v", expected, actual)
	}
}

func Test_CheckIndicator_Input_LiftCurrentFloor_1_DestinationFloor_10_Should_Be_Up(t *testing.T){
	expected := "Up"
	lift := evalator.Lift{
		CurrentFloor: 1,
	}
	destinationFloor := 10

	actual := lift.CheckIndicator(destinationFloor)

	if expected != actual {
		t.Errorf("Expected is %v but get %v", expected, actual)
	}
}