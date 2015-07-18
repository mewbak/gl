package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//TODO:
//add function for status and log text (all the glGet)

//Program is a high level representation of OpenGL program object.
type Program uint32

//CreateProgram is an alias to glCreateProgram.
func CreateProgram() Program {
	p := Program(gl.CreateProgram())
	if safety {
		CheckErr()
	}
	return p
}

//AttachShader is an alias to glAttachShader.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glAttachShader.xml
func (p Program) AttachShader(shader uint32) {
	gl.AttachShader(uint32(p), shader)
	if safety {
		CheckErr()
	}
}

//Link is an alias to glLinkProgram.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glLinkProgram.xml
func (p Program) Link() {
	gl.LinkProgram(uint32(p))
	if safety {
		CheckErr()
	}
}

//Delete is an alias to glDeleteProgram. The progrma should not be used after calling this.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteProgram.xml
func (p Program) Delete() {
	gl.DeleteProgram(uint32(p))
	if safety {
		CheckErr()
	}
}

//Use is an alias to glUseProgram(p).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glUseProgram.xml
func (p Program) Use() {
	gl.UseProgram(uint32(p))
	if safety {
		CheckErr()
	}
}

//Unuse is an alias to glUseProgram(0).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glUseProgram.xml
func (Program) Unuse() {
	gl.UseProgram(0)
	if safety {
		CheckErr()
	}
}

//GetUniformLocation is an alias to glGetUniformLocation(p, name).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetUniformLocation.xml
func (p Program) GetUniformLocation(name string) UniformLocation {
	return UniformLocation(gl.GetUniformLocation(uint32(p), gl.Str(name+"\x00")))
}

//BindFragDataLocation is an alias to glBindFragDataLocation(p, color, name).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindFragDataLocation.xml
func (p Program) BindFragDataLocation(color uint32, name string) {
	gl.BindFragDataLocation(uint32(p), color, gl.Str(name+"\x00"))
	if safety {
		CheckErr()
	}
}

//GetLinkStatus is an alias to glGetProgramiv(p, gl.LINK_STATUS, &status).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetProgram.xml
func (p Program) GetLinkStatus() bool {
	var status int32
	gl.GetProgramiv(uint32(p), gl.LINK_STATUS, &status)
	if safety {
		CheckErr()
	}
	return status == gl.TRUE
}
