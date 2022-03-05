// Code generated by "go generate github.com/gopherd/gonum/unit”; DO NOT EDIT.

// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"errors"
	"fmt"
	"math"
	"unicode/utf8"
)

// Mole represents an amount in moles.
type Mole float64

const Mol Mole = 1

// Unit converts the Mole to a *Unit.
func (n Mole) Unit() *Unit {
	return New(float64(n), Dimensions{
		MoleDim: 1,
	})
}

// Mole allows Mole to implement a Moleer interface.
func (n Mole) Mole() Mole {
	return n
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension.
func (n *Mole) From(u Uniter) error {
	if !DimensionsMatch(u, Mol) {
		*n = Mole(math.NaN())
		return errors.New("unit: dimension mismatch")
	}
	*n = Mole(u.Unit().Value())
	return nil
}

func (n Mole) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", n, float64(n))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " mol"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(n))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(n))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(n))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(n))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g mol)", c, n, float64(n))
	}
}
