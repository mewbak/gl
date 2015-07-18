package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//All Shader types
const (
	VertexShader   = gl.VERTEX_SHADER
	FragmentShader = gl.FRAGMENT_SHADER
	GeometryShader = gl.GEOMETRY_SHADER
)

//lux standard Active texture layout.
const (
	TextureUnitDiffuse      = gl.TEXTURE0
	TextureUnitNormalMap    = gl.TEXTURE1
	TextureUnitDisplacement = gl.TEXTURE2
)

//lux standard active texture layout.
const (
	TextureUniformDiffuse      = 0
	TextureUniformNormalMap    = 1
	TextureUniformDisplacement = 2
)

//All possible framebuffer target.
const (
	DrawFramebuffer     = gl.DRAW_FRAMEBUFFER
	ReadFramebuffer     = gl.READ_FRAMEBUFFER
	ReadDrawFramebuffer = gl.FRAMEBUFFER
)

//all possible framebuffer attachement.
const (
	ColorAttachement0       = gl.COLOR_ATTACHMENT0
	ColorAttachement1       = gl.COLOR_ATTACHMENT1
	ColorAttachement2       = gl.COLOR_ATTACHMENT2
	ColorAttachement3       = gl.COLOR_ATTACHMENT3
	ColorAttachement4       = gl.COLOR_ATTACHMENT4
	ColorAttachement5       = gl.COLOR_ATTACHMENT5
	ColorAttachement6       = gl.COLOR_ATTACHMENT6
	ColorAttachement7       = gl.COLOR_ATTACHMENT7
	ColorAttachement8       = gl.COLOR_ATTACHMENT8
	ColorAttachement9       = gl.COLOR_ATTACHMENT9
	ColorAttachement10      = gl.COLOR_ATTACHMENT10
	ColorAttachement11      = gl.COLOR_ATTACHMENT11
	ColorAttachement12      = gl.COLOR_ATTACHMENT12
	ColorAttachement13      = gl.COLOR_ATTACHMENT13
	ColorAttachement14      = gl.COLOR_ATTACHMENT14
	ColorAttachement15      = gl.COLOR_ATTACHMENT15
	DepthAttachement        = gl.DEPTH_ATTACHMENT
	StencilAttachement      = gl.STENCIL_ATTACHMENT
	DepthStencilAttachement = gl.DEPTH_STENCIL_ATTACHMENT
	None                    = gl.NONE
)

//Possible DrawModes
const (
	POINTS                   = gl.POINTS
	LINE_STRIP               = gl.LINE_STRIP
	LINE_LOOP                = gl.LINE_LOOP
	LINES                    = gl.LINES
	LINE_STRIP_ADJACENCY     = gl.LINE_STRIP_ADJACENCY
	LINES_ADJACENCY          = gl.LINES_ADJACENCY
	TRIANGLE_STRIP           = gl.TRIANGLE_STRIP
	TRIANGLE_FAN             = gl.TRIANGLE_FAN
	TRIANGLES                = gl.TRIANGLES
	TRIANGLE_STRIP_ADJACENCY = gl.TRIANGLE_STRIP_ADJACENCY
	TRIANGLES_ADJACENCY      = gl.TRIANGLES_ADJACENCY
	PATCHES                  = gl.PATCHES
)

//Possible Buffer binding targets.
const (
	ARRAY_BUFFER              = gl.ARRAY_BUFFER              //OpenGL 2+
	ELEMENT_ARRAY_BUFFER      = gl.ELEMENT_ARRAY_BUFFER      //OpenGL 2+
	PIXEL_PACK_BUFFER         = gl.PIXEL_PACK_BUFFER         //OpenGL 2+
	PIXEL_UNPACK_BUFFER       = gl.PIXEL_UNPACK_BUFFER       //OpenGL 2+
	COPY_READ_BUFFER          = gl.COPY_READ_BUFFER          //OpenGL 3+
	COPY_WRITE_BUFFER         = gl.COPY_WRITE_BUFFER         //OpenGL 3+
	TEXTURE_BUFFER            = gl.TEXTURE_BUFFER            //OpenGL 3+
	TRANSFORM_FEEDBACK_BUFFER = gl.TRANSFORM_FEEDBACK_BUFFER //OpenGL 3+
	UNIFORM_BUFFER            = gl.UNIFORM_BUFFER            //OpenGL 3+
	ATOMIC_COUNTER_BUFFER     = gl.ATOMIC_COUNTER_BUFFER     //OpenGL 4+
	DRAW_INDIRECT_BUFFER      = gl.DRAW_INDIRECT_BUFFER      //OpenGL 4+
	DISPATCH_INDIRECT_BUFFER  = gl.DISPATCH_INDIRECT_BUFFER  //OpenGL 4+
	QUERY_BUFFER              = gl.QUERY_BUFFER              //OpenGL 4+
	SHADER_STORAGE_BUFFER     = gl.SHADER_STORAGE_BUFFER     //OpenGL 4+
)

//List of all texture parameter names
const (
	DEPTH_STENCIL_TEXTURE_MODE = gl.DEPTH_STENCIL_TEXTURE_MODE
	TEXTURE_BASE_LEVEL         = gl.TEXTURE_BASE_LEVEL
	TEXTURE_COMPARE_FUNC       = gl.TEXTURE_COMPARE_FUNC
	TEXTURE_COMPARE_MODE       = gl.TEXTURE_COMPARE_MODE
	TEXTURE_LOD_BIAS           = gl.TEXTURE_LOD_BIAS
	TEXTURE_MIN_FILTER         = gl.TEXTURE_MIN_FILTER
	TEXTURE_MAG_FILTER         = gl.TEXTURE_MAG_FILTER
	TEXTURE_MIN_LOD            = gl.TEXTURE_MIN_LOD
	TEXTURE_MAX_LOD            = gl.TEXTURE_MAX_LOD
	TEXTURE_MAX_LEVEL          = gl.TEXTURE_MAX_LEVEL
	TEXTURE_SWIZZLE_R          = gl.TEXTURE_SWIZZLE_R
	TEXTURE_SWIZZLE_G          = gl.TEXTURE_SWIZZLE_G
	TEXTURE_SWIZZLE_B          = gl.TEXTURE_SWIZZLE_B
	TEXTURE_SWIZZLE_A          = gl.TEXTURE_SWIZZLE_A
	TEXTURE_WRAP_S             = gl.TEXTURE_WRAP_S
	TEXTURE_WRAP_T             = gl.TEXTURE_WRAP_T
	TEXTURE_WRAP_R             = gl.TEXTURE_WRAP_R
	TEXTURE_BORDER_COLOR       = gl.TEXTURE_BORDER_COLOR
	TEXTURE_SWIZZLE_RGBA       = gl.TEXTURE_SWIZZLE_RGBA
)

//All possible compare functions
const (
	Never    = gl.NEVER
	Less     = gl.LESS
	Lequal   = gl.LEQUAL
	Greater  = gl.GREATER
	Gequal   = gl.GEQUAL
	Equal    = gl.EQUAL
	NotEqual = gl.NOTEQUAL
	Always   = gl.ALWAYS
)

//List of all possible texture filter
const (
	NEAREST                = gl.NEAREST
	LINEAR                 = gl.LINEAR
	NEAREST_MIPMAP_LINEAR  = gl.NEAREST_MIPMAP_LINEAR
	NEAREST_MIPMAP_NEAREST = gl.NEAREST_MIPMAP_NEAREST
	LINEAR_MIPMAP_LINEAR   = gl.LINEAR_MIPMAP_LINEAR
	LINEAR_MIPMAP_NEAREST  = gl.LINEAR_MIPMAP_NEAREST
)

//WrapMode is an enum to represent all the possible texture wrap modes
type WrapMode int32

//List of all possible texture wrap modes
const (
	CLAMP_TO_EDGE   = gl.CLAMP_TO_EDGE
	CLAMP_TO_BORDER = gl.CLAMP_TO_BORDER
	MIRRORED_REPEAT = gl.MIRRORED_REPEAT
	REPEAT          = gl.REPEAT
)

//List of all opengl data types
const (
	UNSIGNED_BYTE               = gl.UNSIGNED_BYTE
	BYTE                        = gl.BYTE
	UNSIGNED_SHORT              = gl.UNSIGNED_SHORT
	SHORT                       = gl.SHORT
	UNSIGNED_INT                = gl.UNSIGNED_INT
	INT                         = gl.INT
	FLOAT                       = gl.FLOAT
	UNSIGNED_BYTE_3_3_2         = gl.UNSIGNED_BYTE_3_3_2
	UNSIGNED_BYTE_2_3_3_REV     = gl.UNSIGNED_BYTE_2_3_3_REV
	UNSIGNED_SHORT_5_6_5        = gl.UNSIGNED_SHORT_5_6_5
	UNSIGNED_SHORT_5_6_5_REV    = gl.UNSIGNED_SHORT_5_6_5_REV
	UNSIGNED_SHORT_4_4_4_4      = gl.UNSIGNED_SHORT_4_4_4_4
	UNSIGNED_SHORT_4_4_4_4_REV  = gl.UNSIGNED_SHORT_4_4_4_4_REV
	UNSIGNED_SHORT_5_5_5_1      = gl.UNSIGNED_SHORT_5_5_5_1
	UNSIGNED_SHORT_1_5_5_5_REV  = gl.UNSIGNED_SHORT_1_5_5_5_REV
	UNSIGNED_INT_8_8_8_8        = gl.UNSIGNED_INT_8_8_8_8
	UNSIGNED_INT_8_8_8_8_REV    = gl.UNSIGNED_INT_8_8_8_8_REV
	UNSIGNED_INT_10_10_10_2     = gl.UNSIGNED_INT_10_10_10_2
	UNSIGNED_INT_2_10_10_10_REV = gl.UNSIGNED_INT_2_10_10_10_REV
)

const (
	DYNAMIC_DRAW = gl.DYNAMIC_DRAW
)
