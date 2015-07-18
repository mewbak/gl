package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

const (
	DEPTH_TEST = gl.DEPTH_TEST
)

func Enable(param uint32) {
	gl.Enable(param)
	if safety {
		CheckErr()
	}
}
