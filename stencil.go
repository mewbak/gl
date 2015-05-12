package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type stencilObj struct{}

//Stencil is a proxy for all the stencil functions. It should make your code cleaner.
var Stencil = stencilObj{}

//Enable is an alias to glEnable(gl.STENCIL_TEST)
func (stencilObj) Enable() {
	gl.Enable(gl.STENCIL_TEST)
}

//Disable is an alias to glDisable(gl.STENCIL_TEST)
func (stencilObj) Disable() {
	gl.Disable(gl.STENCIL_TEST)
}

//Func is an alias to glStencilFunc(f, ref, mask)
func (stencilObj) Func(f CompareFunc, ref int32, mask uint32) {
	gl.StencilFunc(uint32(f), ref, mask)
}

//Op is an alias to glStencilOp(sfail, zfail, zpass)
func (stencilObj) Op(sfail, zfail, zpass StencilOp) {
	gl.StencilOp(uint32(sfail), uint32(zfail), uint32(zpass))
}

//Mask is an alias to glStencilMask(mask)
func (stencilObj) Mask(mask uint32) {
	gl.StencilMask(mask)
}

//StencilOp is an enum to represent all possible stencil op, prevent bad arguments
type StencilOp uint32

//Stencil op possible values
const (
	Keep     StencilOp = gl.KEEP
	Zero               = gl.ZERO
	Replace            = gl.REPLACE
	Incr               = gl.INCR
	IncrWrap           = gl.INCR_WRAP
	Decr               = gl.DECR
	DecrWrap           = gl.DECR_WRAP
	Invert             = gl.INVERT
)
