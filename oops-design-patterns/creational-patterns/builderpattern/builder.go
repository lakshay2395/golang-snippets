package builderpattern

//VehicleProduct - base struct
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

//BuildProcess - base interface
type BuildProcess interface{
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

//ManufacturingDirector - base director
type ManufacturingDirector struct{
	builder BuildProcess
}

//Construct - base construct
func (m *ManufacturingDirector) Construct() {
	m.builder.SetSeats().SetWheels().SetStructure()
}

//SetBuilder - set builder
func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}