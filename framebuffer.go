package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//Framebuffer is a high level representation of OpenGL framebuffer object.
type Framebuffer uint32

//GenFramebuffer is an alias to glGenFramebuffers(1, &fbo).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGenFramebuffers.xml
func GenFramebuffer() Framebuffer {
	var fbo uint32
	gl.GenFramebuffers(1, &fbo)
	return Framebuffer(fbo)
}

//GenFramebuffer is an alias to glGenFramebuffers(n, &fbo).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glGenFramebuffers.xml
func GenFramebuffers(n int32) []Framebuffer {
	fbos := make([]Framebuffer, n)
	gl.GenFramebuffers(n, (*uint32)(&fbos[0]))
	return fbos
}

//Bind is an alias to glBindFramebuffer(gl.FRAMEBUFFER, fbo).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindFramebuffer.xml
func (fbo Framebuffer) Bind(target FramebufferTarget) {
	gl.BindFramebuffer(uint32(target), uint32(fbo))
}

//Unbind is an alias to glBindFramebuffer(gl.FRAMEBUFFER, 0)
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glBindFramebuffer.xml
func (Framebuffer) Unbind(target FramebufferTarget) {
	gl.BindFramebuffer(uint32(target), 0)
}

//RenderBuffer is an alias to glFramebufferRenderbuffer(target, attachement, gl.RENDERBUFFER, renderbuffer).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferRenderbuffer.xml
func (Framebuffer) RenderBuffer(target FramebufferTarget, attachement FramebufferAttachement, renderbuffer RenderBuffer) {
	gl.FramebufferRenderbuffer(uint32(target), uint32(attachement), gl.RENDERBUFFER, uint32(renderbuffer))
}

//Texture is an alias to glFramebufferTexture(target, attachement, texture, level).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferTexture.xml
func (Framebuffer) Texture(target FramebufferTarget, attachement FramebufferAttachement, texture Texture2D, level int32) {
	gl.FramebufferTexture(uint32(target), uint32(attachement), uint32(texture), level)
}

//DrawBuffers is an alias to glDrawBuffers(len(attachements), &attachements[0]).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDrawBuffers.xml
func (Framebuffer) DrawBuffers(attachements ...FramebufferAttachement) {
	gl.DrawBuffers(int32(len(attachements)), (*uint32)(&attachements[0]))
}

//DrawBuffer is an alias to glDrawBuffer(attachement).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDrawBuffer.xml
func (Framebuffer) DrawBuffer(attachement FramebufferAttachement) {
	gl.DrawBuffer(uint32(attachement))
}

//ReadBuffer is as alias to glReadBuffer(attachement).
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glReadBuffer.xml
func (Framebuffer) ReadBuffer(attachement FramebufferAttachement) {
	gl.ReadBuffer(uint32(attachement))
}

//Delete is an alias to glDeleteFramebuffers(1, &fbo). The FBO should not be used after calling this.
//
//Documentation reference: https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteFramebuffers.xml
func (fbo Framebuffer) Delete() {
	gl.DeleteFramebuffers(1, (*uint32)(&fbo))
}
