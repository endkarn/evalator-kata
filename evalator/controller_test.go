package evalator

import (
	"testing"
)

func Test_checkLoadLimit_Input_LiftWeight_1300_Should_Be_False(t *testing.T) {
	expected := false
	lift := Lift{
		Weight: 1300,
	}

	actual := lift.checkLoadLimit()

	if expected != actual {
		t.Errorf("Expected is %v but get %v", expected, actual)
	}
}

func Test_checkLoadLimit_Input_LiftWeight_1000_Should_Be_True(t *testing.T) {
	expected := true
	lift := Lift{
		Weight: 1000,
	}

	actual := lift.checkLoadLimit()

	if expected != actual {
		t.Errorf("Expected is %v but get %v", expected, actual)
	}
}

func Test_checkLoadLimit_Input_LiftWeight_900_Should_Be_True(t *testing.T) {
	expected := true
	lift := Lift{
		Weight: 900,
	}

	actual := lift.checkLoadLimit()

	if expected != actual {
		t.Errorf("Expected is %v but get %v", expected, actual)
	}
}

func Test_CheckIndicator_Input_LiftCurrentFloor_1_DestinationFloor_10_Should_Be_Up(t *testing.T) {
	expected := "Up"
	lift := Lift{
		CurrentFloor: 1,
	}
	destinationFloor := 10

	actual := lift.checkIndicator(destinationFloor)

	if expected != actual {
		t.Errorf("Expected is %v but get %v", expected, actual)
	}
}
