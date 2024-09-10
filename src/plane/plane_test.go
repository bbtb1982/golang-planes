package plane

import (
  "errors"
  "log/slog"
  "testing"
  "strings"
  //"github.com/davecgh/go-spew/spew"
)

func TestFuelTank(t *testing.T) {
  ft, err := NewFuelTank();
  if err != nil {
    t.Errorf("Failed createing new FuelTank.");
  }

  // Assert can addFuel to Tank
  ft.AddFuel(1)
  if ft.Level != 1 {
    t.Errorf("Error adding one gallon of fuel.");
  }
  ft.AddFuel(1)
  if ft.Level != 2 {
    t.Errorf("Error adding one gallon of fuel.");
  }

  // assert can remove fuel.
  ft.RemoveFuel(2)
  if ft.Level != 0 {
    t.Error("Error removing 2 gallons fuel from tank.");
  }

  // run the tank empty.
  if err = ft.RemoveFuel(3); !errors.Is(ErrFuelTankEmpty, err) {
    t.Error("Error removing 3 gallons fuel from tank.");
  }

 // overfill tank
 if err = ft.AddFuel(1001); !errors.Is(ErrFuelTankOverflow, err) {
   t.Error("Error overfilled tank.");
 }

 // fill tank full
 if err = ft.AddFuel(1000); errors.Is(ErrFuelTankOverflow, err) {
   t.Error("Error overfilled tank.");
 }

 // check if full
 if !ft.IsFull() {
   t.Error("Error tank should be full.");
 }

  // check gallons remaining
  if ft.GallonsRemaining() != 1000 {
    t.Error("Error checking remainder in tank.");
  }

  ft.RemoveFuel(1)
  if ft.GallonsRemaining() != 999 {
    t.Error("Error checking remainder in tank.");
  }

  if ft.GallonsNeeded() != 1 {
    t.Error("Error checking gallons needed to fill tank.");
  }
}

func TestCargoBay(t *testing.T) {
  // can create CargoBay
  cb, err := NewCargoBay();
  if err != nil {
    t.Error("Error checking gallons needed to fill tank.");
  }

  // add weight to cargobay
  cb.AddWeight(1)
  if cb.CurrentWeight != 1 {
    t.Error("Error adding weight to CargoBay.");
  }

  cb.RemoveWeight(1)
  if cb.CurrentWeight != 0 {
    t.Error("Error removing weight to CargoBay.");
  }

  if err = cb.RemoveWeight(1); !errors.Is(err, ErrCargoBayEmpty) {
    t.Error("Error emptying cargo bay beyond zero.")
  }

  cb.AddWeight(cb.MaxWeight)
  if !cb.IsFull() {
    t.Error("Error checking if CargoBay IsFull.")
  }
}

func TestPassengerBay(t *testing.T) {
  pb, err := NewPassengerBay();
  if err != nil {
    t.Error("Error Creating passengerbay")
  }
  // can increment passengers
  pb.AddPaasengers(10)
  if pb.Occupancy != 10 {
    t.Error("Error Adding 10 passengers to passengerBay.")
  }

  // can decrement passengers
  pb.RemovePassengers(5)
  if pb.Occupancy != 5 {
    t.Error("Error removing passengers.")
  }

  // can read remainingSeats.
  if pb.RemainingSeats() != 995 {
    t.Error("Error discovering remaining seats.")
  }
}

func TestCargoPlane(t *testing.T) {
  p, err := NewCargoPlane();
  if err != nil {
    t.Error("Error creating Plane.")
  }

  if p.Altitude != 0 {
    t.Error("Error reading initialized altitude.")
  }
}

func TestCanPrintState(t *testing.T) {
  p, err := NewCargoPlane();
  if err != nil {
    t.Error("Error creating CargoPlane.")
  }
  p.CargoBay.AddWeight(10)
  actual := p.Status();
  expected := "Plane State:: Kind: CargoPlane, Altitude: 0, Speed: 0, FuelLevel: 0, CargoWeight: 10";

  if strings.Compare(actual, expected) != 0 {
    slog.Info("", "actual ", actual)
    slog.Info("", "expected", expected)
    t.Error("status is not what is expected.")
  }
}

func TestCanPrintPassengerPlaneStatus(t *testing.T) {
  p, err := NewPassengerPlane();
  if err != nil {
    t.Error("Error creating passenger plane.")
  }
  p.Altitude = 10;
  p.Speed = 10;
  p.FuelTank.AddFuel(5)
  p.PassengerBay.AddPaasengers(10)

  actual := p.Status();
  expected := "Plane State:: Kind: PassengerPlane, Altitude: 10, Speed: 10, FuelLevel: 5, TotalSeats: 1000, Occupants: 10, RemainingSeats: 990"

  if strings.Compare(actual, expected) != 0 {
    slog.Info("", "actual", actual)
    slog.Info("", "expected", expected)
    t.Error("status is not what is expected.")
  }
}
