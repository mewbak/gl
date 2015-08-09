# gl  
[![GoDoc](https://godoc.org/github.com/luxengine/gl?status.svg)](https://godoc.org/github.com/luxengine/gl)  
This package provides Go idiomatic types to represent OpenGL native types and associates OpenGL function to the proper Go type. The goal is to prevent errors from being made by using the type safety or the compiler without affecting performance.  
This package SHOULD be compiled with quadruple -l flag for optimisation.  
`go install -gcflags="-l -l -l -l" luxengine.net/lux`

Examples:  
From
```Go
var diffuseTex uint32
gl.GenTextures(&diffuseTex, 1)
gl.Bind(gl.TEXTURE_2D, diffuseTex)
gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)
```
To
```Go
diffuseTex := gl.GenTexture2D()
diffuseTex.Bind()
diffuseTex.Parameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
diffuseTex.Parameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
diffuseTex.Parameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
diffuseTex.Parameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
diffuseTex.Image2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)
```
That's because `GenTexture2D` returns the lux gl type `Texture2D`. It has the function
```Go
func (t Texture2D) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, uint32(t))
}
```
The inliner will take care of simply forwarding the argument to the original gl package and you end up with a code thats much more go-idiomatic for no cost.

Note that you can still use the original gl package with Lux and lux gl, but you may have to cast your uint32 to the lux gl version.

If you have a keen eye you will notice that this code 
```Go
diffuseTex := gl.GenTexture2D()
diffuseTex.Bind()
diffuseTex.Parameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
```
and this one
```Go
diffuseTex := gl.GenTexture2D()
diffuseTex.Bind()
normalTex.Parameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
```
do EXACTLY the same thing. To warn you about these we added a safety mode to lux gl
essentially every function in the package looks like this:
```Go
func funcname(args){
    if safetyflag{
        //pre operation checks and log
    }
    gl.funcname(switched args)
    if safetyflag{
        //post operation checks and log
    }
}```

you can build lux gl with `-tags safety` wich will enable all these checks. So you can verify that you aren't doing anything weird. or at least that you know about it. The great thing with that model is that these pieces of code are completelly removed when building non-safe, meaning you get the full speed.