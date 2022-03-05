// Code generated by "go generate github.com/gopherd/gonum/unit/constant”; DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import "github.com/gopherd/gonum/unit"

var definedTypes = []struct {
	unit unit.Uniter
	name string
}{
	{unit: unit.AbsorbedRadioactiveDose(1), name: "unit.AbsorbedRadioactiveDose"},
	{unit: unit.Acceleration(1), name: "unit.Acceleration"},
	{unit: unit.Angle(1), name: "unit.Angle"},
	{unit: unit.Area(1), name: "unit.Area"},
	{unit: unit.Capacitance(1), name: "unit.Capacitance"},
	{unit: unit.Charge(1), name: "unit.Charge"},
	{unit: unit.Conductance(1), name: "unit.Conductance"},
	{unit: unit.Current(1), name: "unit.Current"},
	{unit: unit.Dimless(1), name: "unit.Dimless"},
	{unit: unit.Energy(1), name: "unit.Energy"},
	{unit: unit.EquivalentRadioactiveDose(1), name: "unit.EquivalentRadioactiveDose"},
	{unit: unit.Force(1), name: "unit.Force"},
	{unit: unit.Frequency(1), name: "unit.Frequency"},
	{unit: unit.Inductance(1), name: "unit.Inductance"},
	{unit: unit.Length(1), name: "unit.Length"},
	{unit: unit.LuminousIntensity(1), name: "unit.LuminousIntensity"},
	{unit: unit.MagneticFlux(1), name: "unit.MagneticFlux"},
	{unit: unit.MagneticFluxDensity(1), name: "unit.MagneticFluxDensity"},
	{unit: unit.Mass(1), name: "unit.Mass"},
	{unit: unit.Mole(1), name: "unit.Mole"},
	{unit: unit.Power(1), name: "unit.Power"},
	{unit: unit.Pressure(1), name: "unit.Pressure"},
	{unit: unit.Radioactivity(1), name: "unit.Radioactivity"},
	{unit: unit.Resistance(1), name: "unit.Resistance"},
	{unit: unit.Temperature(1), name: "unit.Temperature"},
	{unit: unit.Time(1), name: "unit.Time"},
	{unit: unit.Torque(1), name: "unit.Torque"},
	{unit: unit.Velocity(1), name: "unit.Velocity"},
	{unit: unit.Voltage(1), name: "unit.Voltage"},
	{unit: unit.Volume(1), name: "unit.Volume"},
}

func definedEquivalentOf(q unit.Uniter) string {
	for _, u := range definedTypes {
		if unit.DimensionsMatch(q, u.unit) {
			return u.name
		}
	}
	return ""
}
