package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type StencilObj struct{}

var Stencil = StencilObj{}

//Enable is an alias to glEnable(gl.STENCIL_TEST)
func (StencilObj) Enable() {
	gl.Enable(gl.STENCIL_TEST)
}

//Disable is an alias to glDisable(gl.STENCIL_TEST)
func (StencilObj) Disable() {
	gl.Disable(gl.STENCIL_TEST)
}

//Func is an alias to glStencilFunc(f, ref, mask)
func (StencilObj) Func(f StencilFunc, ref int32, mask uint32) {
	gl.StencilFunc(uint32(f), ref, mask)
}

//Op is an alias to glStencilOp(sfail, zfail, zpass)
func (StencilObj) Op(sfail, zfail, zpass StencilOp) {
	gl.StencilOp(uint32(sfail), uint32(zfail), uint32(zpass))
}

//Mask is an alias to glStencilMask(mask)
func (StencilObj) Mask(mask uint32) {
	gl.StencilMask(mask)
}

//Enum to represent all possible stencil func, prevents bad arguments
type StencilFunc uint32

//Stencil func possible values
const (
	Never    StencilFunc = gl.NEVER
	Less                 = gl.LESS
	Lequal               = gl.LEQUAL
	Greater              = gl.GREATER
	Gequal               = gl.GEQUAL
	Equal                = gl.EQUAL
	NotEqual             = gl.NOTEQUAL
	Always               = gl.ALWAYS
)

//Enum to represent all possible stencil op, prevent bad arguments
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
