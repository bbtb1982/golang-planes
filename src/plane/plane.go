package plane

import (
	"errors"
  "fmt"
)

var ErrPassengerBayEmpty = errors.New("PassengerBays is empty.")
var ErrPassengerBayOccupancyExceeded = errors.New("Occupancy Full cannot add more passengers.");
var ErrCargoBayEmpty = errors.New("CargoBay is already emtpy.");
var ErrCargoBayOverFlow = errors.New("CargoBay is over weight.");

var ErrFuelTankOverflow = errors.New("FuelTanks has been filled beyond capacity.")
var ErrFuelTankEmpty = errors.New("FuelTank is empty.")

type FuelTank struct {
	Level    int
	Capacity int
	UnitType string
}

func NewFuelTank() (*FuelTank, error) {
	f := FuelTank{
		UnitType: "Gallons",
		Capacity: 1000,
	}

	return &f, nil
}

func (f *FuelTank) GallonsNeeded() int {
	return f.Capacity - f.Level
}

func (f *FuelTank) GallonsRemaining() int {
	return f.Level
}

func (f *FuelTank) AddFuel(gallons int) error {
	t := f.Level + gallons
	if t > f.Capacity {
		return ErrFuelTankOverflow
	}
	f.Level = t

	return nil
}

func (f *FuelTank) RemoveFuel(gallons int) error {
	t := f.Level - gallons
	if t < 0 {
		return ErrFuelTankEmpty
	}
	f.Level = t
	return nil
}

func (f *FuelTank) IsFull() bool {
	return (f.Capacity == f.Level)
}

type CargoBay struct {
	UnitType      string
	MaxWeight     int
	CurrentWeight int
}

func NewCargoBay() (*CargoBay, error) {
	c := CargoBay{
		UnitType:  "Lbs",
		MaxWeight: 1000,
	}

	return &c, nil
}

func (c *CargoBay) AddWeight(lbs int) error {
	t := c.CurrentWeight + lbs
	if t > c.MaxWeight {
		return ErrCargoBayOverFlow
	}
	c.CurrentWeight = t
	return nil
}

func (c *CargoBay) RemoveWeight(lbs int) error {
	t := c.CurrentWeight - lbs
	if t < 0 {
		return ErrCargoBayEmpty
	}

	c.CurrentWeight = t
	return nil
}

func (c *CargoBay) RemainingWeight() int {
	return c.MaxWeight - c.CurrentWeight
}

func (c *CargoBay) IsFull() bool {
	return (c.MaxWeight == c.CurrentWeight)
}

type PassengerBay struct {
  TotalSeats int
  Occupancy int
}

func NewPassengerBay() (*PassengerBay, error) {
  pb := PassengerBay{}
  pb.TotalSeats = 1000;
  return &pb, nil;
}

func (pb *PassengerBay) AddPaasengers(p int) error {
  t := pb.Occupancy + p;
  if pb.TotalSeats < t {
    return ErrPassengerBayOccupancyExceeded;
  }
  pb.Occupancy = t;
  return nil
}
func (pb *PassengerBay) RemovePassengers(p int) error {
  t := pb.Occupancy - p;
  if t < 0 {
    return ErrPassengerBayEmpty;
  }
  pb.Occupancy = t;
  return nil
}

func (pb *PassengerBay) RemainingSeats() int {
  return pb.TotalSeats - pb.Occupancy;
}

type Plane struct {
	*FuelTank
	Altitude int
	Speed   int
}

func NewPlane() (*Plane, error) {
	ft, err := NewFuelTank()
	if err != nil {
		panic("Error creating plane FuelTank!")
	}
	p := Plane{
		FuelTank: ft,
	}
	return &p, nil
}

type CargoPlane struct {
	*Plane
	*CargoBay
}

func NewCargoPlane() (*CargoPlane, error) {
	cb, err := NewCargoBay()
	if err != nil {
		panic("error creating CargoBay!")
	}

	p, err := NewPlane()
	if err != nil {
		panic("Error creating base plane struct!")
	}

	c := CargoPlane{
		Plane:    p,
		CargoBay: cb,
	}
	return &c, nil
}

func (p *CargoPlane) IsFull() bool {
	return false
}

func (p *CargoPlane) IsFueled() bool {
	return p.FuelTank.IsFull()
}

func (p *CargoPlane) RemainingFuel() int {
	return p.FuelTank.Level
}

func (p *CargoPlane) GallonsTillFull() int {
	return p.FuelTank.Capacity
}

func (p *CargoPlane) FuelNeededToFillFuelTank() int {
	return p.FuelTank.GallonsNeeded()
}

func (p *CargoPlane) RemainingCargoWeight() (int, error) {
	if p.CargoBay.MaxWeight < p.CargoBay.CurrentWeight {
		return 0, errors.New("Plane Over Capacity.")
	}
	return (p.CargoBay.MaxWeight - p.CargoBay.CurrentWeight), nil
}

func (p *CargoPlane) Status() string {
  return fmt.Sprintf("Plane State:: Kind: CargoPlane, Altitude: %d, Speed: %d, FuelLevel: %d, CargoWeight: %d, RemainingWeight: %d",
  p.Altitude,
  p.Speed,
  p.FuelTank.Level,
  p.CargoBay.CurrentWeight,
  p.CargoBay.RemainingWeight(),
);
}

type PassengerPlane struct {
  *Plane
  *PassengerBay
}

func NewPassengerPlane() (*PassengerPlane, error) {
  passengerBay, err := NewPassengerBay()
  if err != nil {
    panic("Error creating PassengerPlane.")
  }

  plane, err := NewPlane()
  if err != nil {
    panic("Error creating PassengerPlane.")
  }
  p := PassengerPlane{
    Plane: plane,
    PassengerBay: passengerBay,
  }
  return &p, nil
}

func (p *PassengerPlane) Status() string {
  return fmt.Sprintf("Plane State:: Kind: PassengerPlane, Altitude: %d, Speed: %d, FuelLevel: %d, TotalSeats: %d, Occupants: %d, RemainingSeats: %d",
    p.Altitude,
    p.Speed,
    p.FuelTank.Level,
    p.PassengerBay.TotalSeats,
    p.PassengerBay.Occupancy,
    p.PassengerBay.RemainingSeats(),
  );

}
