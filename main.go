package main

import (
  "log/slog"
  "github.com/bbtb1982/golang-planes/src/plane"
)

func main() {
  slog.Info("Creating Planes.")

  cp, err := plane.NewCargoPlane();
  cp.Altitude = 50;
  cp.FuelTank.AddFuel(50);
  cp.CargoBay.AddWeight(50);
  if err != nil {
    panic("Error creating new Cargo Plane.")
  }

  slog.Info("CargoPlane",  "Status", cp.Status())

  pb, err := plane.NewPassengerPlane();
  pb.Altitude = 100;
  pb.FuelTank.AddFuel(100);
  pb.PassengerBay.AddPaasengers(100);
  if err != nil {
    panic("Error creating new Passenger Plane.")
  }

  slog.Info("CargoPlane",  "Status", pb.Status())


}

