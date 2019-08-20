package evalator

import "testing"

func Test_CheckLoadLimit_Input_LiftWeight_1300_Should_Be_False(t *testing.T){
	expected := false
	lift := Lift{
		Weight:1300,
	}

	actual := lift.CheckLoadLimit()

	if expected != actual {
		t.Errorf("Expected is %b but get %b",expected,actual)
	}
}

func Test_CheckLoadLimit_Input_LiftWeight_1000_Should_Be_False(t *testing.T){
	expected := false
	lift := Lift{
		Weight:1000,
	}

	actual := lift.CheckLoadLimit()

	if expected != actual {
		t.Errorf("Expected is %b but get %b",expected,actual)
	}
}

func Test_CheckLoadLimit_Input_LiftWeight_900_Should_Be_True(t *testing.T){
	expected := false
	lift := Lift{
		Weight:900,
	}

	actual := lift.CheckLoadLimit()

	if expected != actual {
		t.Errorf("Expected is %b but get %b",expected,actual)
	}
}