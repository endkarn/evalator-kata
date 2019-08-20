package liftcontroller_test

import "testing"

func Test_CheckLiftControl_Input_Goto_5_Current_1_Weight_100_Should_Be_Up(t *testing.T){
	expected := "Up"
	evaletor := evalator.Evalator{
		GoToFloor:5,
		Current:1,
		Weight:100,
	}

	actual := checkLiftControl(evalator)

	if expected != actual {
		t.Errorf("Expected in %s but get %s",expected,actual)
	}
}