package builderpattern

import "testing"

func TestBuilderPattern(t *testing.T) {
	mDirector := ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	mDirector.SetBuilder(carBuilder)
	mDirector.Construct()

	car := carBuilder.GetVehicle()
	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}
	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}
	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

}
