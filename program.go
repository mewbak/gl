package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//TODO:
//add function for status and log text (all the glGet)

//Program is a high level representation of OpenGL program object.
type Program uint32

func CreateProgram() Program {
	return Program(gl.CreateProgram())
}

//AttachShader is an alias to glAttachShader.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glAttachShader.xml
func (p Program) AttachShader(shader uint32) {
	gl.AttachShader(uint32(p), shader)
}

//Link is an alias to glLinkProgram.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glLinkProgram.xml
func (p Program) Link() {
	gl.LinkProgram(uint32(p))
}

//Delete is an alias to glDeleteProgram. The progrma should not be used after calling this.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteProgram.xml
func (p Program) Delete() {
	gl.DeleteProgram(uint32(p))
}

//Use is an alias to glUseProgram(p).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glUseProgram.xml
func (p Program) Use() {
	gl.UseProgram(uint32(p))
}

//Unuse is an alias to glUseProgram(0).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glUseProgram.xml
func (Program) Unuse() {
	gl.UseProgram(0)
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
}
