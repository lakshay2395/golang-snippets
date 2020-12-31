package builderpattern

//CarBuilder - base builder
type CarBuilder struct{
	v VehicleProduct
}

//SetWheels - set wheels
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

//SetSeats - set seats
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

//SetStructure - set structure
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

//GetVehicle - get vehicle
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

//BikeBuilder - base builder
type BikeBuilder struct{
	v VehicleProduct
}

//SetWheels - set wheels
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

//SetSeats - set seats
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

//SetStructure - set structure
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}

//GetVehicle - get vehicle
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}