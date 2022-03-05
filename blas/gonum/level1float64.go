// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonum

import (
	"math"

	"github.com/gopherd/gonum/blas"
	"github.com/gopherd/gonum/internal/asm/f64"
)

var _ blas.Float64Level1 = Implementation{}

// Dnrm2 computes the Euclidean norm of a vector,
//  sqrt(\sum_i x[i] * x[i]).
// This function returns 0 if incX is negative.
func (Implementation) Dnrm2(n int, x []float64, incX int) float64 {
	if incX < 1 {
		if incX == 0 {
			panic(zeroIncX)
		}
		return 0
	}
	if len(x) <= (n-1)*incX {
		panic(shortX)
	}
	if n < 2 {
		if n == 1 {
			return math.Abs(x[0])
		}
		if n == 0 {
			return 0
		}
		panic(nLT0)
	}
	if incX == 1 {
		return f64.L2NormUnitary(x[:n])
	}
	return f64.L2NormInc(x, uintptr(n), uintptr(incX))
}

// Dasum computes the sum of the absolute values of the elements of x.
//  \sum_i |x[i]|
// Dasum returns 0 if incX is negative.
func (Implementation) Dasum(n int, x []float64, incX int) float64 {
	var sum float64
	if n < 0 {
		panic(nLT0)
	}
	if incX < 1 {
		if incX == 0 {
			panic(zeroIncX)
		}
		return 0
	}
	if len(x) <= (n-1)*incX {
		panic(shortX)
	}
	if incX == 1 {
		x = x[:n]
		for _, v := range x {
			sum += math.Abs(v)
		}
		return sum
	}
	for i := 0; i < n; i++ {
		sum += math.Abs(x[i*incX])
	}
	return sum
}

// Idamax returns the index of an element of x with the largest absolute value.
// If there are multiple such indices the earliest is returned.
// Idamax returns -1 if n == 0.
func (Implementation) Idamax(n int, x []float64, incX int) int {
	if incX < 1 {
		if incX == 0 {
			panic(zeroIncX)
		}
		return -1
	}
	if len(x) <= (n-1)*incX {
		panic(shortX)
	}
	if n < 2 {
		if n == 1 {
			return 0
		}
		if n == 0 {
			return -1 // Netlib returns invalid index when n == 0.
		}
		panic(nLT0)
	}
	idx := 0
	max := math.Abs(x[0])
	if incX == 1 {
		for i, v := range x[:n] {
			absV := math.Abs(v)
			if absV > max {
				max = absV
				idx = i
			}
		}
		return idx
	}
	ix := incX
	for i := 1; i < n; i++ {
		v := x[ix]
		absV := math.Abs(v)
		if absV > max {
			max = absV
			idx = i
		}
		ix += incX
	}
	return idx
}

// Dswap exchanges the elements of two vectors.
//  x[i], y[i] = y[i], x[i] for all i
func (Implementation) Dswap(n int, x []float64, incX int, y []float64, incY int) {
	if incX == 0 {
		panic(zeroIncX)
	}
	if incY == 0 {
		panic(zeroIncY)
	}
	if n < 1 {
		if n == 0 {
			return
		}
		panic(nLT0)
	}
	if (incX > 0 && len(x) <= (n-1)*incX) || (incX < 0 && len(x) <= (1-n)*incX) {
		panic(shortX)
	}
	if (incY > 0 && len(y) <= (n-1)*incY) || (incY < 0 && len(y) <= (1-n)*incY) {
		panic(shortY)
	}
	if incX == 1 && incY == 1 {
		x = x[:n]
		for i, v := range x {
			x[i], y[i] = y[i], v
		}
		return
	}
	var ix, iy int
	if incX < 0 {
		ix = (-n + 1) * incX
	}
	if incY < 0 {
		iy = (-n + 1) * incY
	}
	for i := 0; i < n; i++ {
		x[ix], y[iy] = y[iy], x[ix]
		ix += incX
		iy += incY
	}
}

// Dcopy copies the elements of x into the elements of y.
//  y[i] = x[i] for all i
func (Implementation) Dcopy(n int, x []float64, incX int, y []float64, incY int) {
	if incX == 0 {
		panic(zeroIncX)
	}
	if incY == 0 {
		panic(zeroIncY)
	}
	if n < 1 {
		if n == 0 {
			return
		}
		panic(nLT0)
	}
	if (incX > 0 && len(x) <= (n-1)*incX) || (incX < 0 && len(x) <= (1-n)*incX) {
		panic(shortX)
	}
	if (incY > 0 && len(y) <= (n-1)*incY) || (incY < 0 && len(y) <= (1-n)*incY) {
		panic(shortY)
	}
	if incX == 1 && incY == 1 {
		copy(y[:n], x[:n])
		return
	}
	var ix, iy int
	if incX < 0 {
		ix = (-n + 1) * incX
	}
	if incY < 0 {
		iy = (-n + 1) * incY
	}
	for i := 0; i < n; i++ {
		y[iy] = x[ix]
		ix += incX
		iy += incY
	}
}

// Daxpy adds alpha times x to y
//  y[i] += alpha * x[i] for all i
func (Implementation) Daxpy(n int, alpha float64, x []float64, incX int, y []float64, incY int) {
	if incX == 0 {
		panic(zeroIncX)
	}
	if incY == 0 {
		panic(zeroIncY)
	}
	if n < 1 {
		if n == 0 {
			return
		}
		panic(nLT0)
	}
	if (incX > 0 && len(x) <= (n-1)*incX) || (incX < 0 && len(x) <= (1-n)*incX) {
		panic(shortX)
	}
	if (incY > 0 && len(y) <= (n-1)*incY) || (incY < 0 && len(y) <= (1-n)*incY) {
		panic(shortY)
	}
	if alpha == 0 {
		return
	}
	if incX == 1 && incY == 1 {
		f64.AxpyUnitary(alpha, x[:n], y[:n])
		return
	}
	var ix, iy int
	if incX < 0 {
		ix = (-n + 1) * incX
	}
	if incY < 0 {
		iy = (-n + 1) * incY
	}
	f64.AxpyInc(alpha, x, y, uintptr(n), uintptr(incX), uintptr(incY), uintptr(ix), uintptr(iy))
}

// Drotg computes a plane rotation
//  ⎡  c s ⎤ ⎡ a ⎤ = ⎡ r ⎤
//  ⎣ -s c ⎦ ⎣ b ⎦   ⎣ 0 ⎦
// satisfying c^2 + s^2 = 1.
//
// The computation uses the formulas
//  sigma = sgn(a)    if |a| >  |b|
//        = sgn(b)    if |b| >= |a|
//  r = sigma*sqrt(a^2 + b^2)
//  c = 1; s = 0      if r = 0
//  c = a/r; s = b/r  if r != 0
//  c >= 0            if |a| > |b|
//
// The subroutine also computes
//  z = s    if |a| > |b|,
//    = 1/c  if |b| >= |a| and c != 0
//    = 1    if c = 0
// This allows c and s to be reconstructed from z as follows:
//  If z = 1, set c = 0, s = 1.
//  If |z| < 1, set c = sqrt(1 - z^2) and s = z.
//  If |z| > 1, set c = 1/z and s = sqrt(1 - c^2).
//
// NOTE: There is a discrepancy between the reference implementation and the
// BLAS technical manual regarding the sign for r when a or b are zero. Drotg
// agrees with the definition in the manual and other common BLAS
// implementations.
func (Implementation) Drotg(a, b float64) (c, s, r, z float64) {
	// Implementation based on Supplemental Material to:
	// Edward Anderson. 2017. Algorithm 978: Safe Scaling in the Level 1 BLAS.
	// ACM Trans. Math. Softw. 44, 1, Article 12 (July 2017), 28 pages.
	// DOI: https://doi.org/10.1145/3061665
	const (
		safmin = 0x1p-1022
		safmax = 1 / safmin
	)
	anorm := math.Abs(a)
	bnorm := math.Abs(b)
	switch {
	case bnorm == 0:
		c = 1
		s = 0
		r = a
		z = 0
	case anorm == 0:
		c = 0
		s = 1
		r = b
		z = 1
	default:
		maxab := math.Max(anorm, bnorm)
		scl := math.Min(math.Max(safmin, maxab), safmax)
		var sigma float64
		if anorm > bnorm {
			sigma = math.Copysign(1, a)
		} else {
			sigma = math.Copysign(1, b)
		}
		ascl := a / scl
		bscl := b / scl
		r = sigma * (scl * math.Sqrt(ascl*ascl+bscl*bscl))
		c = a / r
		s = b / r
		switch {
		case anorm > bnorm:
			z = s
		case c != 0:
			z = 1 / c
		default:
			z = 1
		}
	}
	return c, s, r, z
}

// Drotmg computes the modified Givens rotation. See
// http://www.netlib.org/lapack/explore-html/df/deb/drotmg_8f.html
// for more details.
func (Implementation) Drotmg(d1, d2, x1, y1 float64) (p blas.DrotmParams, rd1, rd2, rx1 float64) {
	// The implementation of Drotmg used here is taken from Hopkins 1997
	// Appendix A: https://doi.org/10.1145/289251.289253
	// with the exception of the gam constants below.

	const (
		gam    = 4096.0
		gamsq  = gam * gam
		rgamsq = 1.0 / gamsq
	)

	if d1 < 0 {
		p.Flag = blas.Rescaling // Error state.
		return p, 0, 0, 0
	}

	if d2 == 0 || y1 == 0 {
		p.Flag = blas.Identity
		return p, d1, d2, x1
	}

	var h11, h12, h21, h22 float64
	if (d1 == 0 || x1 == 0) && d2 > 0 {
		p.Flag = blas.Diagonal
		h12 = 1
		h21 = -1
		x1 = y1
		d1, d2 = d2, d1
	} else {
		p2 := d2 * y1
		p1 := d1 * x1
		q2 := p2 * y1
		q1 := p1 * x1
		if math.Abs(q1) > math.Abs(q2) {
			p.Flag = blas.OffDiagonal
			h11 = 1
			h22 = 1
			h21 = -y1 / x1
			h12 = p2 / p1
			u := 1 - float64(h12*h21)
			if u <= 0 {
				p.Flag = blas.Rescaling // Error state.
				return p, 0, 0, 0
			}

			d1 /= u
			d2 /= u
			x1 *= u
		} else {
			if q2 < 0 {
				p.Flag = blas.Rescaling // Error state.
				return p, 0, 0, 0
			}

			p.Flag = blas.Diagonal
			h21 = -1
			h12 = 1
			h11 = p1 / p2
			h22 = x1 / y1
			u := 1 + float64(h11*h22)
			d1, d2 = d2/u, d1/u
			x1 = y1 * u
		}
	}

	for d1 <= rgamsq && d1 != 0 {
		p.Flag = blas.Rescaling
		d1 = (d1 * gam) * gam
		x1 /= gam
		h11 /= gam
		h12 /= gam
	}
	for d1 > gamsq {
		p.Flag = blas.Rescaling
		d1 = (d1 / gam) / gam
		x1 *= gam
		h11 *= gam
		h12 *= gam
	}

	for math.Abs(d2) <= rgamsq && d2 != 0 {
		p.Flag = blas.Rescaling
		d2 = (d2 * gam) * gam
		h21 /= gam
		h22 /= gam
	}
	for math.Abs(d2) > gamsq {
		p.Flag = blas.Rescaling
		d2 = (d2 / gam) / gam
		h21 *= gam
		h22 *= gam
	}

	switch p.Flag {
	case blas.Diagonal:
		p.H = [4]float64{0: h11, 3: h22}
	case blas.OffDiagonal:
		p.H = [4]float64{1: h21, 2: h12}
	case blas.Rescaling:
		p.H = [4]float64{h11, h21, h12, h22}
	default:
		panic(badFlag)
	}

	return p, d1, d2, x1
}

// Drot applies a plane transformation.
//  x[i] = c * x[i] + s * y[i]
//  y[i] = c * y[i] - s * x[i]
func (Implementation) Drot(n int, x []float64, incX int, y []float64, incY int, c float64, s float64) {
	if incX == 0 {
		panic(zeroIncX)
	}
	if incY == 0 {
		panic(zeroIncY)
	}
	if n < 1 {
		if n == 0 {
			return
		}
		panic(nLT0)
	}
	if (incX > 0 && len(x) <= (n-1)*incX) || (incX < 0 && len(x) <= (1-n)*incX) {
		panic(shortX)
	}
	if (incY > 0 && len(y) <= (n-1)*incY) || (incY < 0 && len(y) <= (1-n)*incY) {
		panic(shortY)
	}
	if incX == 1 && incY == 1 {
		x = x[:n]
		for i, vx := range x {
			vy := y[i]
			x[i], y[i] = c*vx+s*vy, c*vy-s*vx
		}
		return
	}
	var ix, iy int
	if incX < 0 {
		ix = (-n + 1) * incX
	}
	if incY < 0 {
		iy = (-n + 1) * incY
	}
	for i := 0; i < n; i++ {
		vx := x[ix]
		vy := y[iy]
		x[ix], y[iy] = c*vx+s*vy, c*vy-s*vx
		ix += incX
		iy += incY
	}
}

// Drotm applies the modified Givens rotation to the 2×n matrix.
func (Implementation) Drotm(n int, x []float64, incX int, y []float64, incY int, p blas.DrotmParams) {
	if incX == 0 {
		panic(zeroIncX)
	}
	if incY == 0 {
		panic(zeroIncY)
	}
	if n <= 0 {
		if n == 0 {
			return
		}
		panic(nLT0)
	}
	if (incX > 0 && len(x) <= (n-1)*incX) || (incX < 0 && len(x) <= (1-n)*incX) {
		panic(shortX)
	}
	if (incY > 0 && len(y) <= (n-1)*incY) || (incY < 0 && len(y) <= (1-n)*incY) {
		panic(shortY)
	}

	if p.Flag == blas.Identity {
		return
	}

	switch p.Flag {
	case blas.Rescaling:
		h11 := p.H[0]
		h12 := p.H[2]
		h21 := p.H[1]
		h22 := p.H[3]
		if incX == 1 && incY == 1 {
			x = x[:n]
			for i, vx := range x {
				vy := y[i]
				x[i], y[i] = float64(vx*h11)+float64(vy*h12), float64(vx*h21)+float64(vy*h22)
			}
			return
		}
		var ix, iy int
		if incX < 0 {
			ix = (-n + 1) * incX
		}
		if incY < 0 {
			iy = (-n + 1) * incY
		}
		for i := 0; i < n; i++ {
			vx := x[ix]
			vy := y[iy]
			x[ix], y[iy] = float64(vx*h11)+float64(vy*h12), float64(vx*h21)+float64(vy*h22)
			ix += incX
			iy += incY
		}
	case blas.OffDiagonal:
		h12 := p.H[2]
		h21 := p.H[1]
		if incX == 1 && incY == 1 {
			x = x[:n]
			for i, vx := range x {
				vy := y[i]
				x[i], y[i] = vx+float64(vy*h12), float64(vx*h21)+vy
			}
			return
		}
		var ix, iy int
		if incX < 0 {
			ix = (-n + 1) * incX
		}
		if incY < 0 {
			iy = (-n + 1) * incY
		}
		for i := 0; i < n; i++ {
			vx := x[ix]
			vy := y[iy]
			x[ix], y[iy] = vx+float64(vy*h12), float64(vx*h21)+vy
			ix += incX
			iy += incY
		}
	case blas.Diagonal:
		h11 := p.H[0]
		h22 := p.H[3]
		if incX == 1 && incY == 1 {
			x = x[:n]
			for i, vx := range x {
				vy := y[i]
				x[i], y[i] = float64(vx*h11)+vy, -vx+float64(vy*h22)
			}
			return
		}
		var ix, iy int
		if incX < 0 {
			ix = (-n + 1) * incX
		}
		if incY < 0 {
			iy = (-n + 1) * incY
		}
		for i := 0; i < n; i++ {
			vx := x[ix]
			vy := y[iy]
			x[ix], y[iy] = float64(vx*h11)+vy, -vx+float64(vy*h22)
			ix += incX
			iy += incY
		}
	}
}

// Dscal scales x by alpha.
//  x[i] *= alpha
// Dscal has no effect if incX < 0.
func (Implementation) Dscal(n int, alpha float64, x []float64, incX int) {
	if incX < 1 {
		if incX == 0 {
			panic(zeroIncX)
		}
		return
	}
	if n < 1 {
		if n == 0 {
			return
		}
		panic(nLT0)
	}
	if (n-1)*incX >= len(x) {
		panic(shortX)
	}
	if alpha == 0 {
		if incX == 1 {
			x = x[:n]
			for i := range x {
				x[i] = 0
			}
			return
		}
		for ix := 0; ix < n*incX; ix += incX {
			x[ix] = 0
		}
		return
	}
	if incX == 1 {
		f64.ScalUnitary(alpha, x[:n])
		return
	}
	f64.ScalInc(alpha, x, uintptr(n), uintptr(incX))
}
