package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"unsafe"
)

//texture interface maybe ?
/*
type Texture interface{
	Bind()
	Unbind()
	Delete()
	//think about parameters
	//Parameteri()
}
*/

//Texture is a high-level representation of the OpenGL texture object, can be any type (TEXTURE_2D, TEXTURE_1D, etc).
type Texture uint32

//Texture2D is the high level representation of OpenGL TEXTURE_2D object. It restrict availlable functions and automatically fills the GL_TEXTURE_2D target.
type Texture2D Texture

//GenTexture2D is an alias to glGenTextures(1,&tex).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGenTextures.xml
func GenTexture2D() Texture2D {
	var tex uint32
	gl.GenTextures(1, &tex)
	return Texture2D(tex)
}

//GenTexture2D is an alias to glGenTextures(n,&tex).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGenTextures.xml
func GenTextures2D(n int32) []Texture2D {
	tex := make([]Texture2D, n)
	gl.GenTextures(1, (*uint32)(&tex[0]))
	return tex
}

//Bind (Texture2D) is an alias to glBindTexture(gl.TEXTURE_2D, t).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindTexture.xml
func (t Texture2D) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, uint32(t))
}

//Unbind is an alias to glBindTexture(0).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindTexture.xml
func (Texture2D) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

//Delete is an alias to glDeleteTextures. This texture should not be used after calling this.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteTextures.xml
func (t Texture2D) Delete() {
	gl.DeleteTextures(1, (*uint32)(&t))
}

//TexImage2D is an alias to glTexImage2D.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexImage2D.xml
func (t Texture2D) TexImage2D(level, internalformat, width, height, border int32, format, xtype uint32, pixels unsafe.Pointer) {
	gl.TexImage2D(gl.TEXTURE_2D, level, internalformat, width, height, border, format, xtype, pixels)
}

//TexParameteriv is an alias to glTexParameteri(target, pname, param)
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glTexParameter.xml
func (t Texture2D) TexParameteriv(target, pname uint32, param int32) {
	gl.TexParameteri(target, pname, param)
}

//GetTexParameteriv is an alias to glGetTexParameteriv(target, pname, params)
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexParameter.xml
func (t Texture2D) GetTexParameteriv(target, pname uint32, params *int32) {
	gl.GetTexParameteriv(target, pname, params)
}

//GetTexLevelParameteriv is an alias to glGetTexLevelParameteriv(target, level, pname, params)
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexLevelParameter.xml
func (t Texture2D) GetTexLevelParameteriv(target uint32, level int32, pname uint32, params *int32) {
	gl.GetTexLevelParameteriv(target, level, pname, params)
}

//GetTexImage is an alias to glGetTexImage(gl.TEXTURE_2D, level, format, xtype, pixels)
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexImage.xml
func (t Texture2D) GetTexImage(level int32, format, xtype uint32, pixels unsafe.Pointer) {
	gl.GetTexImage(gl.TEXTURE_2D, level, format, xtype, pixels)
}

//ReadPixels is an alias to glReadPixels(x, y, width, height, format, xtype, pixels)
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glReadPixels.xml
func (t Texture2D) ReadPixels(x, y, width, height int32, format, xtype uint32, pixels unsafe.Pointer) {
	gl.ReadPixels(x, y, width, height, format, xtype, pixels)
}

//Width is an alias to glGetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_WIDTH, &w)
//
//Documentation reference:
func (t Texture2D) Width(miplevel int32) int32 {
	var w int32
	gl.GetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_WIDTH, &w)
	return w
}

//Height is an alias to glGetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_HEIGHT, &h)
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexLevelParameter.xml
func (t Texture2D) Height(miplevel int32) int32 {
	var h int32
	gl.GetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_HEIGHT, &h)
	return h
}

//InternalFormat is an alias to glGetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_INTERNAL_FORMAT, &x)
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGetTexLevelParameter.xml
func (t Texture2D) InternalFormat(miplevel int32) uint32 {
	var x int32
	gl.GetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_INTERNAL_FORMAT, &x)
	return uint32(x)
}
