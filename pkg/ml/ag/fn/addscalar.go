// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fn

import (
	"brillion.io/spago/pkg/mat"
)

// Element-wise addition over two values.
type AddScalar struct {
	x1 Operand
	x2 Operand // scalar
}

func NewAddScalar(x1, x2 Operand) *AddScalar {
	return &AddScalar{x1: x1, x2: x2}
}

// Forward computes the output of the function.
// It doesn't backward on the scalar value x2.
func (r *AddScalar) Forward() mat.Matrix {
	return r.x1.Value().AddScalar(r.x2.Value().Scalar())
}

func (r *AddScalar) Backward(gy mat.Matrix) {
	if r.x1.RequiresGrad() {
		r.x1.PropagateGrad(gy)
	}
	if r.x2.RequiresGrad() {
		gx := mat.NewEmptyDense(r.x2.Value().Dims())
		for i := 0; i < gy.Rows(); i++ {
			for j := 0; j < gy.Columns(); j++ {
				gx.Set(gx.Scalar()+gy.At(i, j), 0, 0)
			}
		}
		r.x2.PropagateGrad(gx)
	}
}