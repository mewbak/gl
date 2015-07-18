package gl

import (
	"errors"
	"github.com/go-gl/gl/v3.3-core/gl"
)

func GetError() error {
	if err := gl.GetError(); err != gl.NO_ERROR {
		return glErrorToerror(err)
	}
	return nil
}

var (
	ErrNoError                     = errors.New("GL_NO_ERROR")
	ErrInvalidEnum                 = errors.New("GL_INVALID_ENUM")
	ErrInvalidValue                = errors.New("GL_INVALID_VALUE")
	ErrInvalidOperation            = errors.New("GL_INVALID_OPERATION")
	ErrInvalidFramebufferOperation = errors.New("GL_INVALID_FRAMEBUFFER_OPERATION")
	ErrOutOfMemory                 = errors.New("GL_OUT_OF_MEMORY")
	ErrStackUnderflow              = errors.New("GL_STACK_UNDERFLOW")
	ErrStackOverflow               = errors.New("GL_STACK_OVERFLOW")
	ErrUnknown                     = errors.New("Unknown error code")
)

func glErrorToerror(err uint32) error {
	switch err {
	case gl.NO_ERROR:
		return ErrNoError
	case gl.INVALID_ENUM:
		return ErrInvalidEnum
	case gl.INVALID_VALUE:
		return ErrInvalidValue
	case gl.INVALID_OPERATION:
		return ErrInvalidOperation
	case gl.INVALID_FRAMEBUFFER_OPERATION:
		return ErrInvalidFramebufferOperation
	case gl.OUT_OF_MEMORY:
		return ErrOutOfMemory
	case gl.STACK_UNDERFLOW:
		return ErrStackUnderflow
	case gl.STACK_OVERFLOW:
		return ErrStackOverflow
	default:
		return ErrUnknown
	}
}
