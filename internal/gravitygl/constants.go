package gravitygl

/*
	Partially generated from the Khronos OpenGL API specification in XML
	format, which is covered by the license:

	Copyright (c) 2013-2014 The Khronos Group Inc.
	Permission is hereby granted, free of charge, to any person obtaining a
	copy of this software and/or associated documentation files (the
	"Materials"), to deal in the Materials without restriction, including
	without limitation the rights to use, copy, modify, merge, publish,
	distribute, sublicense, and/or sell copies of the Materials, and to
	permit persons to whom the Materials are furnished to do so, subject to
	the following conditions:
	The above copyright notice and this permission notice shall be included
	in all copies or substantial portions of the Materials.
	THE MATERIALS ARE PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
	EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
	MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
	IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
	CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
	TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
	MATERIALS OR THE USE OR OTHER DEALINGS IN THE MATERIALS.
*/

import "fmt"

// Enum represents a gl type or constant
type Enum EnumType

func (v Enum) String() string {
	switch v {
	case 0x0:
		return "0"
	case 0x1:
		return "1"
	case 0x2:
		return "2"
	case 0x3:
		return "LINE_STRIP"
	case 0x4:
		return "4"
	case 0x5:
		return "TRIANGLE_STRIP"
	case 0x6:
		return "TRIANGLE_FAN"
	case 0x300:
		return "SRC_COLOR"
	case 0x301:
		return "ONE_MINUS_SRC_COLOR"
	case 0x302:
		return "SRC_ALPHA"
	case 0x303:
		return "ONE_MINUS_SRC_ALPHA"
	case 0x304:
		return "DST_ALPHA"
	case 0x305:
		return "ONE_MINUS_DST_ALPHA"
	case 0x306:
		return "DST_COLOR"
	case 0x307:
		return "ONE_MINUS_DST_COLOR"
	case 0x308:
		return "SRC_ALPHA_SATURATE"
	case 0x8006:
		return "FUNC_ADD"
	case 0x8009:
		return "32777"
	case 0x883d:
		return "BLEND_EQUATION_ALPHA"
	case 0x800a:
		return "FUNC_SUBTRACT"
	case 0x800b:
		return "FUNC_REVERSE_SUBTRACT"
	case 0x80c8:
		return "BLEND_DST_RGB"
	case 0x80c9:
		return "BLEND_SRC_RGB"
	case 0x80ca:
		return "BLEND_DST_ALPHA"
	case 0x80cb:
		return "BLEND_SRC_ALPHA"
	case 0x8001:
		return "CONSTANT_COLOR"
	case 0x8002:
		return "ONE_MINUS_CONSTANT_COLOR"
	case 0x8003:
		return "CONSTANT_ALPHA"
	case 0x8004:
		return "ONE_MINUS_CONSTANT_ALPHA"
	case 0x8005:
		return "BLEND_COLOR"
	case 0x8892:
		return "ARRAY_BUFFER"
	case 0x8893:
		return "ELEMENT_ARRAY_BUFFER"
	case 0x8894:
		return "ARRAY_BUFFER_BINDING"
	case 0x8895:
		return "ELEMENT_ARRAY_BUFFER_BINDING"
	case 0x88e0:
		return "STREAM_DRAW"
	case 0x88e4:
		return "STATIC_DRAW"
	case 0x88e8:
		return "DYNAMIC_DRAW"
	case 0x8764:
		return "BUFFER_SIZE"
	case 0x8765:
		return "BUFFER_USAGE"
	case 0x8626:
		return "CURRENT_VERTEX_ATTRIB"
	case 0x404:
		return "FRONT"
	case 0x405:
		return "BACK"
	case 0x408:
		return "FRONT_AND_BACK"
	case 0xde1:
		return "TEXTURE_2D"
	case 0xb44:
		return "CULL_FACE"
	case 0xbe2:
		return "BLEND"
	case 0xbd0:
		return "DITHER"
	case 0xb90:
		return "STENCIL_TEST"
	case 0xb71:
		return "DEPTH_TEST"
	case 0xc11:
		return "SCISSOR_TEST"
	case 0x8037:
		return "POLYGON_OFFSET_FILL"
	case 0x809e:
		return "SAMPLE_ALPHA_TO_COVERAGE"
	case 0x80a0:
		return "SAMPLE_COVERAGE"
	case 0x500:
		return "INVALID_ENUM"
	case 0x501:
		return "INVALID_VALUE"
	case 0x502:
		return "INVALID_OPERATION"
	case 0x505:
		return "OUT_OF_MEMORY"
	case 0x900:
		return "CW"
	case 0x901:
		return "CCW"
	case 0xb21:
		return "LINE_WIDTH"
	case 0x846d:
		return "ALIASED_POINT_SIZE_RANGE"
	case 0x846e:
		return "ALIASED_LINE_WIDTH_RANGE"
	case 0xb45:
		return "CULL_FACE_MODE"
	case 0xb46:
		return "FRONT_FACE"
	case 0xb70:
		return "DEPTH_RANGE"
	case 0xb72:
		return "DEPTH_WRITEMASK"
	case 0xb73:
		return "DEPTH_CLEAR_VALUE"
	case 0xb74:
		return "DEPTH_FUNC"
	case 0xb91:
		return "STENCIL_CLEAR_VALUE"
	case 0xb92:
		return "STENCIL_FUNC"
	case 0xb94:
		return "STENCIL_FAIL"
	case 0xb95:
		return "STENCIL_PASS_DEPTH_FAIL"
	case 0xb96:
		return "STENCIL_PASS_DEPTH_PASS"
	case 0xb97:
		return "STENCIL_REF"
	case 0xb93:
		return "STENCIL_VALUE_MASK"
	case 0xb98:
		return "STENCIL_WRITEMASK"
	case 0x8800:
		return "STENCIL_BACK_FUNC"
	case 0x8801:
		return "STENCIL_BACK_FAIL"
	case 0x8802:
		return "STENCIL_BACK_PASS_DEPTH_FAIL"
	case 0x8803:
		return "STENCIL_BACK_PASS_DEPTH_PASS"
	case 0x8ca3:
		return "STENCIL_BACK_REF"
	case 0x8ca4:
		return "STENCIL_BACK_VALUE_MASK"
	case 0x8ca5:
		return "STENCIL_BACK_WRITEMASK"
	case 0xba2:
		return "VIEWPORT"
	case 0xc10:
		return "SCISSOR_BOX"
	case 0xc22:
		return "COLOR_CLEAR_VALUE"
	case 0xc23:
		return "COLOR_WRITEMASK"
	case 0xcf5:
		return "UNPACK_ALIGNMENT"
	case 0xd05:
		return "PACK_ALIGNMENT"
	case 0xd33:
		return "MAX_TEXTURE_SIZE"
	case 0xd3a:
		return "MAX_VIEWPORT_DIMS"
	case 0xd50:
		return "SUBPIXEL_BITS"
	case 0xd52:
		return "RED_BITS"
	case 0xd53:
		return "GREEN_BITS"
	case 0xd54:
		return "BLUE_BITS"
	case 0xd55:
		return "ALPHA_BITS"
	case 0xd56:
		return "DEPTH_BITS"
	case 0xd57:
		return "STENCIL_BITS"
	case 0x2a00:
		return "POLYGON_OFFSET_UNITS"
	case 0x8038:
		return "POLYGON_OFFSET_FACTOR"
	case 0x8069:
		return "TEXTURE_BINDING_2D"
	case 0x80a8:
		return "SAMPLE_BUFFERS"
	case 0x80a9:
		return "SAMPLES"
	case 0x80aa:
		return "SAMPLE_COVERAGE_VALUE"
	case 0x80ab:
		return "SAMPLE_COVERAGE_INVERT"
	case 0x86a2:
		return "NUM_COMPRESSED_TEXTURE_FORMATS"
	case 0x86a3:
		return "COMPRESSED_TEXTURE_FORMATS"
	case 0x1100:
		return "DONT_CARE"
	case 0x1101:
		return "FASTEST"
	case 0x1102:
		return "NICEST"
	case 0x8192:
		return "GENERATE_MIPMAP_HINT"
	case 0x1400:
		return "BYTE"
	case 0x1401:
		return "UNSIGNED_BYTE"
	case 0x1402:
		return "SHORT"
	case 0x1403:
		return "UNSIGNED_SHORT"
	case 0x1404:
		return "INT"
	case 0x1405:
		return "UNSIGNED_INT"
	case 0x1406:
		return "FLOAT"
	case 0x140c:
		return "FIXED"
	case 0x1902:
		return "DEPTH_COMPONENT"
	case 0x1906:
		return "ALPHA"
	case 0x1907:
		return "RGB"
	case 0x1908:
		return "RGBA"
	case 0x1909:
		return "LUMINANCE"
	case 0x190a:
		return "LUMINANCE_ALPHA"
	case 0x8033:
		return "UNSIGNED_SHORT_4_4_4_4"
	case 0x8034:
		return "UNSIGNED_SHORT_5_5_5_1"
	case 0x8363:
		return "UNSIGNED_SHORT_5_6_5"
	case 0x8869:
		return "MAX_VERTEX_ATTRIBS"
	case 0x8dfb:
		return "MAX_VERTEX_UNIFORM_VECTORS"
	case 0x8dfc:
		return "MAX_VARYING_VECTORS"
	case 0x8b4d:
		return "MAX_COMBINED_TEXTURE_IMAGE_UNITS"
	case 0x8b4c:
		return "MAX_VERTEX_TEXTURE_IMAGE_UNITS"
	case 0x8872:
		return "MAX_TEXTURE_IMAGE_UNITS"
	case 0x8dfd:
		return "MAX_FRAGMENT_UNIFORM_VECTORS"
	case 0x8b4f:
		return "SHADER_TYPE"
	case 0x8b80:
		return "DELETE_STATUS"
	case 0x8b82:
		return "LINK_STATUS"
	case 0x8b83:
		return "VALIDATE_STATUS"
	case 0x8b85:
		return "ATTACHED_SHADERS"
	case 0x8b86:
		return "ACTIVE_UNIFORMS"
	case 0x8b87:
		return "ACTIVE_UNIFORM_MAX_LENGTH"
	case 0x8b89:
		return "ACTIVE_ATTRIBUTES"
	case 0x8b8a:
		return "ACTIVE_ATTRIBUTE_MAX_LENGTH"
	case 0x8b8c:
		return "SHADING_LANGUAGE_VERSION"
	case 0x8b8d:
		return "CURRENT_PROGRAM"
	case 0x200:
		return "NEVER"
	case 0x201:
		return "LESS"
	case 0x202:
		return "EQUAL"
	case 0x203:
		return "LEQUAL"
	case 0x204:
		return "GREATER"
	case 0x205:
		return "NOTEQUAL"
	case 0x206:
		return "GEQUAL"
	case 0x207:
		return "ALWAYS"
	case 0x1e00:
		return "KEEP"
	case 0x1e01:
		return "REPLACE"
	case 0x1e02:
		return "INCR"
	case 0x1e03:
		return "DECR"
	case 0x150a:
		return "INVERT"
	case 0x8507:
		return "INCR_WRAP"
	case 0x8508:
		return "DECR_WRAP"
	case 0x1f00:
		return "VENDOR"
	case 0x1f01:
		return "RENDERER"
	case 0x1f02:
		return "VERSION"
	case 0x1f03:
		return "EXTENSIONS"
	case 0x2600:
		return "NEAREST"
	case 0x2601:
		return "LINEAR"
	case 0x2700:
		return "NEAREST_MIPMAP_NEAREST"
	case 0x2701:
		return "LINEAR_MIPMAP_NEAREST"
	case 0x2702:
		return "NEAREST_MIPMAP_LINEAR"
	case 0x2703:
		return "LINEAR_MIPMAP_LINEAR"
	case 0x2800:
		return "TEXTURE_MAG_FILTER"
	case 0x2801:
		return "TEXTURE_MIN_FILTER"
	case 0x2802:
		return "TEXTURE_WRAP_S"
	case 0x2803:
		return "TEXTURE_WRAP_T"
	case 0x1702:
		return "TEXTURE"
	case 0x8513:
		return "TEXTURE_CUBE_MAP"
	case 0x8514:
		return "TEXTURE_BINDING_CUBE_MAP"
	case 0x8515:
		return "TEXTURE_CUBE_MAP_POSITIVE_X"
	case 0x8516:
		return "TEXTURE_CUBE_MAP_NEGATIVE_X"
	case 0x8517:
		return "TEXTURE_CUBE_MAP_POSITIVE_Y"
	case 0x8518:
		return "TEXTURE_CUBE_MAP_NEGATIVE_Y"
	case 0x8519:
		return "TEXTURE_CUBE_MAP_POSITIVE_Z"
	case 0x851a:
		return "TEXTURE_CUBE_MAP_NEGATIVE_Z"
	case 0x851c:
		return "MAX_CUBE_MAP_TEXTURE_SIZE"
	case 0x84c0:
		return "TEXTURE0"
	case 0x84c1:
		return "TEXTURE1"
	case 0x84c2:
		return "TEXTURE2"
	case 0x84c3:
		return "TEXTURE3"
	case 0x84c4:
		return "TEXTURE4"
	case 0x84c5:
		return "TEXTURE5"
	case 0x84c6:
		return "TEXTURE6"
	case 0x84c7:
		return "TEXTURE7"
	case 0x84c8:
		return "TEXTURE8"
	case 0x84c9:
		return "TEXTURE9"
	case 0x84ca:
		return "TEXTURE10"
	case 0x84cb:
		return "TEXTURE11"
	case 0x84cc:
		return "TEXTURE12"
	case 0x84cd:
		return "TEXTURE13"
	case 0x84ce:
		return "TEXTURE14"
	case 0x84cf:
		return "TEXTURE15"
	case 0x84d0:
		return "TEXTURE16"
	case 0x84d1:
		return "TEXTURE17"
	case 0x84d2:
		return "TEXTURE18"
	case 0x84d3:
		return "TEXTURE19"
	case 0x84d4:
		return "TEXTURE20"
	case 0x84d5:
		return "TEXTURE21"
	case 0x84d6:
		return "TEXTURE22"
	case 0x84d7:
		return "TEXTURE23"
	case 0x84d8:
		return "TEXTURE24"
	case 0x84d9:
		return "TEXTURE25"
	case 0x84da:
		return "TEXTURE26"
	case 0x84db:
		return "TEXTURE27"
	case 0x84dc:
		return "TEXTURE28"
	case 0x84dd:
		return "TEXTURE29"
	case 0x84de:
		return "TEXTURE30"
	case 0x84df:
		return "TEXTURE31"
	case 0x84e0:
		return "ACTIVE_TEXTURE"
	case 0x2901:
		return "REPEAT"
	case 0x812f:
		return "CLAMP_TO_EDGE"
	case 0x8370:
		return "MIRRORED_REPEAT"
	case 0x8622:
		return "VERTEX_ATTRIB_ARRAY_ENABLED"
	case 0x8623:
		return "VERTEX_ATTRIB_ARRAY_SIZE"
	case 0x8624:
		return "VERTEX_ATTRIB_ARRAY_STRIDE"
	case 0x8625:
		return "VERTEX_ATTRIB_ARRAY_TYPE"
	case 0x886a:
		return "VERTEX_ATTRIB_ARRAY_NORMALIZED"
	case 0x8645:
		return "VERTEX_ATTRIB_ARRAY_POINTER"
	case 0x889f:
		return "VERTEX_ATTRIB_ARRAY_BUFFER_BINDING"
	case 0x8b9a:
		return "IMPLEMENTATION_COLOR_READ_TYPE"
	case 0x8b9b:
		return "IMPLEMENTATION_COLOR_READ_FORMAT"
	case 0x8b81:
		return "COMPILE_STATUS"
	case 0x8b84:
		return "INFO_LOG_LENGTH"
	case 0x8b88:
		return "SHADER_SOURCE_LENGTH"
	case 0x8dfa:
		return "SHADER_COMPILER"
	case 0x8df8:
		return "SHADER_BINARY_FORMATS"
	case 0x8df9:
		return "NUM_SHADER_BINARY_FORMATS"
	case 0x8df0:
		return "LOW_FLOAT"
	case 0x8df1:
		return "MEDIUM_FLOAT"
	case 0x8df2:
		return "HIGH_FLOAT"
	case 0x8df3:
		return "LOW_INT"
	case 0x8df4:
		return "MEDIUM_INT"
	case 0x8df5:
		return "HIGH_INT"
	case 0x8d40:
		return "FRAMEBUFFER"
	case 0x8d41:
		return "RENDERBUFFER"
	case 0x8056:
		return "RGBA4"
	case 0x8057:
		return "RGB5_A1"
	case 0x8d62:
		return "RGB565"
	case 0x81a5:
		return "DEPTH_COMPONENT16"
	case 0x8d48:
		return "STENCIL_INDEX8"
	case 0x8d42:
		return "RENDERBUFFER_WIDTH"
	case 0x8d43:
		return "RENDERBUFFER_HEIGHT"
	case 0x8d44:
		return "RENDERBUFFER_INTERNAL_FORMAT"
	case 0x8d50:
		return "RENDERBUFFER_RED_SIZE"
	case 0x8d51:
		return "RENDERBUFFER_GREEN_SIZE"
	case 0x8d52:
		return "RENDERBUFFER_BLUE_SIZE"
	case 0x8d53:
		return "RENDERBUFFER_ALPHA_SIZE"
	case 0x8d54:
		return "RENDERBUFFER_DEPTH_SIZE"
	case 0x8d55:
		return "RENDERBUFFER_STENCIL_SIZE"
	case 0x8cd0:
		return "FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE"
	case 0x8cd1:
		return "FRAMEBUFFER_ATTACHMENT_OBJECT_NAME"
	case 0x8cd2:
		return "FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL"
	case 0x8cd3:
		return "FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE"
	case 0x8ce0:
		return "COLOR_ATTACHMENT0"
	case 0x8d00:
		return "DEPTH_ATTACHMENT"
	case 0x8d20:
		return "STENCIL_ATTACHMENT"
	case 0x8cd5:
		return "FRAMEBUFFER_COMPLETE"
	case 0x8cd6:
		return "FRAMEBUFFER_INCOMPLETE_ATTACHMENT"
	case 0x8cd7:
		return "FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT"
	case 0x8cd9:
		return "FRAMEBUFFER_INCOMPLETE_DIMENSIONS"
	case 0x8cdd:
		return "FRAMEBUFFER_UNSUPPORTED"
	case 0x8ca6:
		return "36006"
	case 0x8ca7:
		return "RENDERBUFFER_BINDING"
	case 0x84e8:
		return "MAX_RENDERBUFFER_SIZE"
	case 0x506:
		return "INVALID_FRAMEBUFFER_OPERATION"
	case 0x100:
		return "DEPTH_BUFFER_BIT"
	case 0x400:
		return "STENCIL_BUFFER_BIT"
	case 0x4000:
		return "COLOR_BUFFER_BIT"
	case 0x8b50:
		return "FLOAT_VEC2"
	case 0x8b51:
		return "FLOAT_VEC3"
	case 0x8b52:
		return "FLOAT_VEC4"
	case 0x8b53:
		return "INT_VEC2"
	case 0x8b54:
		return "INT_VEC3"
	case 0x8b55:
		return "INT_VEC4"
	case 0x8b56:
		return "BOOL"
	case 0x8b57:
		return "BOOL_VEC2"
	case 0x8b58:
		return "BOOL_VEC3"
	case 0x8b59:
		return "BOOL_VEC4"
	case 0x8b5a:
		return "FLOAT_MAT2"
	case 0x8b5b:
		return "FLOAT_MAT3"
	case 0x8b5c:
		return "FLOAT_MAT4"
	case 0x8b5e:
		return "SAMPLER_2D"
	case 0x8b60:
		return "SAMPLER_CUBE"
	case 0x8b30:
		return "FRAGMENT_SHADER"
	case 0x8b31:
		return "VERTEX_SHADER"
	case 0x8a35:
		return "ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH"
	case 0x8a36:
		return "ACTIVE_UNIFORM_BLOCKS"
	case 0x911a:
		return "ALREADY_SIGNALED"
	case 0x8c2f:
		return "ANY_SAMPLES_PASSED"
	case 0x8d6a:
		return "ANY_SAMPLES_PASSED_CONSERVATIVE"
	case 0x1905:
		return "BLUE"
	case 0x911f:
		return "BUFFER_ACCESS_FLAGS"
	case 0x9120:
		return "BUFFER_MAP_LENGTH"
	case 0x9121:
		return "BUFFER_MAP_OFFSET"
	case 0x88bc:
		return "BUFFER_MAPPED"
	case 0x88bd:
		return "BUFFER_MAP_POINTER"
	case 0x1800:
		return "COLOR"
	case 0x8cea:
		return "COLOR_ATTACHMENT10"
	case 0x8ce1:
		return "COLOR_ATTACHMENT1"
	case 0x8ceb:
		return "COLOR_ATTACHMENT11"
	case 0x8cec:
		return "COLOR_ATTACHMENT12"
	case 0x8ced:
		return "COLOR_ATTACHMENT13"
	case 0x8cee:
		return "COLOR_ATTACHMENT14"
	case 0x8cef:
		return "COLOR_ATTACHMENT15"
	case 0x8ce2:
		return "COLOR_ATTACHMENT2"
	case 0x8ce3:
		return "COLOR_ATTACHMENT3"
	case 0x8ce4:
		return "COLOR_ATTACHMENT4"
	case 0x8ce5:
		return "COLOR_ATTACHMENT5"
	case 0x8ce6:
		return "COLOR_ATTACHMENT6"
	case 0x8ce7:
		return "COLOR_ATTACHMENT7"
	case 0x8ce8:
		return "COLOR_ATTACHMENT8"
	case 0x8ce9:
		return "COLOR_ATTACHMENT9"
	case 0x884e:
		return "COMPARE_REF_TO_TEXTURE"
	case 0x9270:
		return "COMPRESSED_R11_EAC"
	case 0x9272:
		return "COMPRESSED_RG11_EAC"
	case 0x9274:
		return "COMPRESSED_RGB8_ETC2"
	case 0x9276:
		return "COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2"
	case 0x9278:
		return "COMPRESSED_RGBA8_ETC2_EAC"
	case 0x9271:
		return "COMPRESSED_SIGNED_R11_EAC"
	case 0x9273:
		return "COMPRESSED_SIGNED_RG11_EAC"
	case 0x9279:
		return "COMPRESSED_SRGB8_ALPHA8_ETC2_EAC"
	case 0x9275:
		return "COMPRESSED_SRGB8_ETC2"
	case 0x9277:
		return "COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2"
	case 0x911c:
		return "CONDITION_SATISFIED"
	case 0x8f36:
		return "36662"
	case 0x8f37:
		return "36663"
	case 0x8865:
		return "CURRENT_QUERY"
	case 0x1801:
		return "DEPTH"
	case 0x88f0:
		return "DEPTH24_STENCIL8"
	case 0x8cad:
		return "DEPTH32F_STENCIL8"
	case 0x81a6:
		return "DEPTH_COMPONENT24"
	case 0x8cac:
		return "DEPTH_COMPONENT32F"
	case 0x84f9:
		return "DEPTH_STENCIL"
	case 0x821a:
		return "DEPTH_STENCIL_ATTACHMENT"
	case 0x8825:
		return "DRAW_BUFFER0"
	case 0x882f:
		return "DRAW_BUFFER10"
	case 0x8826:
		return "DRAW_BUFFER1"
	case 0x8830:
		return "DRAW_BUFFER11"
	case 0x8831:
		return "DRAW_BUFFER12"
	case 0x8832:
		return "DRAW_BUFFER13"
	case 0x8833:
		return "DRAW_BUFFER14"
	case 0x8834:
		return "DRAW_BUFFER15"
	case 0x8827:
		return "DRAW_BUFFER2"
	case 0x8828:
		return "DRAW_BUFFER3"
	case 0x8829:
		return "DRAW_BUFFER4"
	case 0x882a:
		return "DRAW_BUFFER5"
	case 0x882b:
		return "DRAW_BUFFER6"
	case 0x882c:
		return "DRAW_BUFFER7"
	case 0x882d:
		return "DRAW_BUFFER8"
	case 0x882e:
		return "DRAW_BUFFER9"
	case 0x8ca9:
		return "DRAW_FRAMEBUFFER"
	case 0x88ea:
		return "DYNAMIC_COPY"
	case 0x88e9:
		return "DYNAMIC_READ"
	case 0x8dad:
		return "FLOAT_32_UNSIGNED_INT_24_8_REV"
	case 0x8b65:
		return "FLOAT_MAT2x3"
	case 0x8b66:
		return "FLOAT_MAT2x4"
	case 0x8b67:
		return "FLOAT_MAT3x2"
	case 0x8b68:
		return "FLOAT_MAT3x4"
	case 0x8b69:
		return "FLOAT_MAT4x2"
	case 0x8b6a:
		return "FLOAT_MAT4x3"
	case 0x8b8b:
		return "FRAGMENT_SHADER_DERIVATIVE_HINT"
	case 0x8215:
		return "FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE"
	case 0x8214:
		return "FRAMEBUFFER_ATTACHMENT_BLUE_SIZE"
	case 0x8210:
		return "FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING"
	case 0x8211:
		return "FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE"
	case 0x8216:
		return "FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE"
	case 0x8213:
		return "FRAMEBUFFER_ATTACHMENT_GREEN_SIZE"
	case 0x8212:
		return "FRAMEBUFFER_ATTACHMENT_RED_SIZE"
	case 0x8217:
		return "FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE"
	case 0x8cd4:
		return "FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER"
	case 0x8218:
		return "FRAMEBUFFER_DEFAULT"
	case 0x8d56:
		return "FRAMEBUFFER_INCOMPLETE_MULTISAMPLE"
	case 0x8219:
		return "FRAMEBUFFER_UNDEFINED"
	case 0x1904:
		return "GREEN"
	case 0x140b:
		return "HALF_FLOAT"
	case 0x8d9f:
		return "INT_2_10_10_10_REV"
	case 0x8c8c:
		return "INTERLEAVED_ATTRIBS"
	case 0x8dca:
		return "INT_SAMPLER_2D"
	case 0x8dcf:
		return "INT_SAMPLER_2D_ARRAY"
	case 0x8dcb:
		return "INT_SAMPLER_3D"
	case 0x8dcc:
		return "INT_SAMPLER_CUBE"
	//case 0xffffffff:
	//	return "INVALID_INDEX"
	case 0x821b:
		return "MAJOR_VERSION"
	case 0x10:
		return "MAP_FLUSH_EXPLICIT_BIT"
	case 0x8:
		return "MAP_INVALIDATE_BUFFER_BIT"
	case 0x20:
		return "MAP_UNSYNCHRONIZED_BIT"
	case 0x8008:
		return "MAX"
	case 0x8073:
		return "MAX_3D_TEXTURE_SIZE"
	case 0x88ff:
		return "MAX_ARRAY_TEXTURE_LAYERS"
	case 0x8cdf:
		return "MAX_COLOR_ATTACHMENTS"
	case 0x8a33:
		return "MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS"
	case 0x8a2e:
		return "MAX_COMBINED_UNIFORM_BLOCKS"
	case 0x8a31:
		return "MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS"
	case 0x8824:
		return "MAX_DRAW_BUFFERS"
	case 0x8d6b:
		return "MAX_ELEMENT_INDEX"
	case 0x80e9:
		return "MAX_ELEMENTS_INDICES"
	case 0x80e8:
		return "MAX_ELEMENTS_VERTICES"
	case 0x9125:
		return "MAX_FRAGMENT_INPUT_COMPONENTS"
	case 0x8a2d:
		return "MAX_FRAGMENT_UNIFORM_BLOCKS"
	case 0x8b49:
		return "MAX_FRAGMENT_UNIFORM_COMPONENTS"
	case 0x8905:
		return "MAX_PROGRAM_TEXEL_OFFSET"
	case 0x8d57:
		return "MAX_SAMPLES"
	case 0x9111:
		return "MAX_SERVER_WAIT_TIMEOUT"
	case 0x84fd:
		return "MAX_TEXTURE_LOD_BIAS"
	case 0x8c8a:
		return "MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS"
	case 0x8c8b:
		return "MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS"
	case 0x8c80:
		return "MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS"
	case 0x8a30:
		return "MAX_UNIFORM_BLOCK_SIZE"
	case 0x8a2f:
		return "MAX_UNIFORM_BUFFER_BINDINGS"
	case 0x8b4b:
		return "MAX_VARYING_COMPONENTS"
	case 0x9122:
		return "MAX_VERTEX_OUTPUT_COMPONENTS"
	case 0x8a2b:
		return "MAX_VERTEX_UNIFORM_BLOCKS"
	case 0x8b4a:
		return "MAX_VERTEX_UNIFORM_COMPONENTS"
	case 0x8007:
		return "MIN"
	case 0x821c:
		return "MINOR_VERSION"
	case 0x8904:
		return "MIN_PROGRAM_TEXEL_OFFSET"
	case 0x821d:
		return "NUM_EXTENSIONS"
	case 0x87fe:
		return "NUM_PROGRAM_BINARY_FORMATS"
	case 0x9380:
		return "NUM_SAMPLE_COUNTS"
	case 0x9112:
		return "OBJECT_TYPE"
	case 0xd02:
		return "PACK_ROW_LENGTH"
	case 0xd04:
		return "PACK_SKIP_PIXELS"
	case 0xd03:
		return "PACK_SKIP_ROWS"
	case 0x88eb:
		return "PIXEL_PACK_BUFFER"
	case 0x88ed:
		return "PIXEL_PACK_BUFFER_BINDING"
	case 0x88ec:
		return "PIXEL_UNPACK_BUFFER"
	case 0x88ef:
		return "PIXEL_UNPACK_BUFFER_BINDING"
	case 0x8d69:
		return "PRIMITIVE_RESTART_FIXED_INDEX"
	case 0x87ff:
		return "PROGRAM_BINARY_FORMATS"
	case 0x8741:
		return "PROGRAM_BINARY_LENGTH"
	case 0x8257:
		return "PROGRAM_BINARY_RETRIEVABLE_HINT"
	case 0x8866:
		return "QUERY_RESULT"
	case 0x8867:
		return "QUERY_RESULT_AVAILABLE"
	case 0x8c3a:
		return "R11F_G11F_B10F"
	case 0x822d:
		return "R16F"
	case 0x8233:
		return "R16I"
	case 0x8234:
		return "R16UI"
	case 0x822e:
		return "R32F"
	case 0x8235:
		return "R32I"
	case 0x8236:
		return "R32UI"
	case 0x8229:
		return "R8"
	case 0x8231:
		return "R8I"
	case 0x8f94:
		return "R8_SNORM"
	case 0x8232:
		return "R8UI"
	case 0x8c89:
		return "RASTERIZER_DISCARD"
	case 0xc02:
		return "READ_BUFFER"
	case 0x8ca8:
		return "READ_FRAMEBUFFER"
	case 0x8caa:
		return "READ_FRAMEBUFFER_BINDING"
	case 0x1903:
		return "RED"
	case 0x8d94:
		return "RED_INTEGER"
	case 0x8cab:
		return "RENDERBUFFER_SAMPLES"
	case 0x8227:
		return "RG"
	case 0x822f:
		return "RG16F"
	case 0x8239:
		return "RG16I"
	case 0x823a:
		return "RG16UI"
	case 0x8230:
		return "RG32F"
	case 0x823b:
		return "RG32I"
	case 0x823c:
		return "RG32UI"
	case 0x822b:
		return "RG8"
	case 0x8237:
		return "RG8I"
	case 0x8f95:
		return "RG8_SNORM"
	case 0x8238:
		return "RG8UI"
	case 0x8059:
		return "RGB10_A2"
	case 0x906f:
		return "RGB10_A2UI"
	case 0x881b:
		return "RGB16F"
	case 0x8d89:
		return "RGB16I"
	case 0x8d77:
		return "RGB16UI"
	case 0x8815:
		return "RGB32F"
	case 0x8d83:
		return "RGB32I"
	case 0x8d71:
		return "RGB32UI"
	case 0x8051:
		return "RGB8"
	case 0x8d8f:
		return "RGB8I"
	case 0x8f96:
		return "RGB8_SNORM"
	case 0x8d7d:
		return "RGB8UI"
	case 0x8c3d:
		return "RGB9_E5"
	case 0x881a:
		return "RGBA16F"
	case 0x8d88:
		return "RGBA16I"
	case 0x8d76:
		return "RGBA16UI"
	case 0x8814:
		return "RGBA32F"
	case 0x8d82:
		return "RGBA32I"
	case 0x8d70:
		return "RGBA32UI"
	case 0x8058:
		return "RGBA8"
	case 0x8d8e:
		return "RGBA8I"
	case 0x8f97:
		return "RGBA8_SNORM"
	case 0x8d7c:
		return "RGBA8UI"
	case 0x8d99:
		return "RGBA_INTEGER"
	case 0x8d98:
		return "RGB_INTEGER"
	case 0x8228:
		return "RG_INTEGER"
	case 0x8dc1:
		return "SAMPLER_2D_ARRAY"
	case 0x8dc4:
		return "SAMPLER_2D_ARRAY_SHADOW"
	case 0x8b62:
		return "SAMPLER_2D_SHADOW"
	case 0x8b5f:
		return "SAMPLER_3D"
	case 0x8919:
		return "SAMPLER_BINDING"
	case 0x8dc5:
		return "SAMPLER_CUBE_SHADOW"
	case 0x8c8d:
		return "SEPARATE_ATTRIBS"
	case 0x9119:
		return "SIGNALED"
	case 0x8f9c:
		return "SIGNED_NORMALIZED"
	case 0x8c40:
		return "SRGB"
	case 0x8c41:
		return "SRGB8"
	case 0x8c43:
		return "SRGB8_ALPHA8"
	case 0x88e6:
		return "STATIC_COPY"
	case 0x88e5:
		return "STATIC_READ"
	case 0x1802:
		return "STENCIL"
	case 0x88e2:
		return "STREAM_COPY"
	case 0x88e1:
		return "STREAM_READ"
	case 0x9113:
		return "SYNC_CONDITION"
	case 0x9116:
		return "SYNC_FENCE"
	case 0x9115:
		return "SYNC_FLAGS"
	case 0x9117:
		return "SYNC_GPU_COMMANDS_COMPLETE"
	case 0x9114:
		return "SYNC_STATUS"
	case 0x8c1a:
		return "TEXTURE_2D_ARRAY"
	case 0x806f:
		return "TEXTURE_3D"
	case 0x813c:
		return "TEXTURE_BASE_LEVEL"
	case 0x8c1d:
		return "TEXTURE_BINDING_2D_ARRAY"
	case 0x806a:
		return "TEXTURE_BINDING_3D"
	case 0x884d:
		return "TEXTURE_COMPARE_FUNC"
	case 0x884c:
		return "TEXTURE_COMPARE_MODE"
	case 0x912f:
		return "TEXTURE_IMMUTABLE_FORMAT"
	case 0x82df:
		return "TEXTURE_IMMUTABLE_LEVELS"
	case 0x813d:
		return "TEXTURE_MAX_LEVEL"
	case 0x813b:
		return "TEXTURE_MAX_LOD"
	case 0x813a:
		return "TEXTURE_MIN_LOD"
	case 0x8e45:
		return "TEXTURE_SWIZZLE_A"
	case 0x8e44:
		return "TEXTURE_SWIZZLE_B"
	case 0x8e43:
		return "TEXTURE_SWIZZLE_G"
	case 0x8e42:
		return "TEXTURE_SWIZZLE_R"
	case 0x8072:
		return "TEXTURE_WRAP_R"
	case 0x911b:
		return "TIMEOUT_EXPIRED"
	case 0x8e22:
		return "TRANSFORM_FEEDBACK"
	case 0x8e24:
		return "TRANSFORM_FEEDBACK_ACTIVE"
	case 0x8e25:
		return "TRANSFORM_FEEDBACK_BINDING"
	case 0x8c8e:
		return "TRANSFORM_FEEDBACK_BUFFER"
	case 0x8c8f:
		return "TRANSFORM_FEEDBACK_BUFFER_BINDING"
	case 0x8c7f:
		return "TRANSFORM_FEEDBACK_BUFFER_MODE"
	case 0x8c85:
		return "TRANSFORM_FEEDBACK_BUFFER_SIZE"
	case 0x8c84:
		return "TRANSFORM_FEEDBACK_BUFFER_START"
	case 0x8e23:
		return "TRANSFORM_FEEDBACK_PAUSED"
	case 0x8c88:
		return "TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN"
	case 0x8c76:
		return "TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH"
	case 0x8c83:
		return "TRANSFORM_FEEDBACK_VARYINGS"
	case 0x8a3c:
		return "UNIFORM_ARRAY_STRIDE"
	case 0x8a43:
		return "UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES"
	case 0x8a42:
		return "UNIFORM_BLOCK_ACTIVE_UNIFORMS"
	case 0x8a3f:
		return "UNIFORM_BLOCK_BINDING"
	case 0x8a40:
		return "UNIFORM_BLOCK_DATA_SIZE"
	case 0x8a3a:
		return "UNIFORM_BLOCK_INDEX"
	case 0x8a41:
		return "UNIFORM_BLOCK_NAME_LENGTH"
	case 0x8a46:
		return "UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER"
	case 0x8a44:
		return "UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER"
	case 0x8a11:
		return "UNIFORM_BUFFER"
	case 0x8a28:
		return "UNIFORM_BUFFER_BINDING"
	case 0x8a34:
		return "UNIFORM_BUFFER_OFFSET_ALIGNMENT"
	case 0x8a2a:
		return "UNIFORM_BUFFER_SIZE"
	case 0x8a29:
		return "UNIFORM_BUFFER_START"
	case 0x8a3e:
		return "UNIFORM_IS_ROW_MAJOR"
	case 0x8a3d:
		return "UNIFORM_MATRIX_STRIDE"
	case 0x8a39:
		return "UNIFORM_NAME_LENGTH"
	case 0x8a3b:
		return "UNIFORM_OFFSET"
	case 0x8a38:
		return "UNIFORM_SIZE"
	case 0x8a37:
		return "UNIFORM_TYPE"
	case 0x806e:
		return "UNPACK_IMAGE_HEIGHT"
	case 0xcf2:
		return "UNPACK_ROW_LENGTH"
	case 0x806d:
		return "UNPACK_SKIP_IMAGES"
	case 0xcf4:
		return "UNPACK_SKIP_PIXELS"
	case 0xcf3:
		return "UNPACK_SKIP_ROWS"
	case 0x9118:
		return "UNSIGNALED"
	case 0x8c3b:
		return "UNSIGNED_INT_10F_11F_11F_REV"
	case 0x8368:
		return "UNSIGNED_INT_2_10_10_10_REV"
	case 0x84fa:
		return "UNSIGNED_INT_24_8"
	case 0x8c3e:
		return "UNSIGNED_INT_5_9_9_9_REV"
	case 0x8dd2:
		return "UNSIGNED_INT_SAMPLER_2D"
	case 0x8dd7:
		return "UNSIGNED_INT_SAMPLER_2D_ARRAY"
	case 0x8dd3:
		return "UNSIGNED_INT_SAMPLER_3D"
	case 0x8dd4:
		return "UNSIGNED_INT_SAMPLER_CUBE"
	case 0x8dc6:
		return "UNSIGNED_INT_VEC2"
	case 0x8dc7:
		return "UNSIGNED_INT_VEC3"
	case 0x8dc8:
		return "UNSIGNED_INT_VEC4"
	case 0x8c17:
		return "UNSIGNED_NORMALIZED"
	case 0x85b5:
		return "VERTEX_ARRAY_BINDING"
	case 0x88fe:
		return "VERTEX_ATTRIB_ARRAY_DIVISOR"
	case 0x88fd:
		return "VERTEX_ATTRIB_ARRAY_INTEGER"
	case 0x911d:
		return "WAIT_FAILED"
	default:
		return fmt.Sprintf("gl.Enum(0x%x)", uint32(v))
	}
}

const (
	GL_POINT_SMOOTH_HINT   = 0x0C51
	GL_LINE_SMOOTH_HINT    = 0x0C52
	GL_POLYGON_SMOOTH_HINT = 0x0C53

	GL_POLYGON_SMOOTH = 0x0B41
	GL_LINE_SMOOTH    = 0x0B20
)

const (
	POINTS                                       = 0x0000
	LINES                                        = 0x0001
	LINE_LOOP                                    = 0x0002
	LINE_STRIP                                   = 0x0003
	TRIANGLES                                    = 0x0004
	TRIANGLE_STRIP                               = 0x0005
	TRIANGLE_FAN                                 = 0x0006
	SRC_COLOR                                    = 0x0300
	ONE_MINUS_SRC_COLOR                          = 0x0301
	SRC_ALPHA                                    = 0x0302
	ONE_MINUS_SRC_ALPHA                          = 0x0303
	DST_ALPHA                                    = 0x0304
	ONE_MINUS_DST_ALPHA                          = 0x0305
	DST_COLOR                                    = 0x0306
	ONE_MINUS_DST_COLOR                          = 0x0307
	SRC_ALPHA_SATURATE                           = 0x0308
	FUNC_ADD                                     = 0x8006
	BLEND_EQUATION                               = 0x8009
	BLEND_EQUATION_RGB                           = 0x8009
	BLEND_EQUATION_ALPHA                         = 0x883D
	FUNC_SUBTRACT                                = 0x800A
	FUNC_REVERSE_SUBTRACT                        = 0x800B
	BLEND_DST_RGB                                = 0x80C8
	BLEND_SRC_RGB                                = 0x80C9
	BLEND_DST_ALPHA                              = 0x80CA
	BLEND_SRC_ALPHA                              = 0x80CB
	CONSTANT_COLOR                               = 0x8001
	ONE_MINUS_CONSTANT_COLOR                     = 0x8002
	CONSTANT_ALPHA                               = 0x8003
	ONE_MINUS_CONSTANT_ALPHA                     = 0x8004
	BLEND_COLOR                                  = 0x8005
	ARRAY_BUFFER                                 = 0x8892
	ELEMENT_ARRAY_BUFFER                         = 0x8893
	ARRAY_BUFFER_BINDING                         = 0x8894
	ELEMENT_ARRAY_BUFFER_BINDING                 = 0x8895
	STREAM_DRAW                                  = 0x88E0
	STATIC_DRAW                                  = 0x88E4
	DYNAMIC_DRAW                                 = 0x88E8
	BUFFER_SIZE                                  = 0x8764
	BUFFER_USAGE                                 = 0x8765
	CURRENT_VERTEX_ATTRIB                        = 0x8626
	FRONT                                        = 0x0404
	BACK                                         = 0x0405
	FRONT_AND_BACK                               = 0x0408
	TEXTURE_2D                                   = 0x0DE1
	CULL_FACE                                    = 0x0B44
	BLEND                                        = 0x0BE2
	DITHER                                       = 0x0BD0
	STENCIL_TEST                                 = 0x0B90
	DEPTH_TEST                                   = 0x0B71
	SCISSOR_TEST                                 = 0x0C11
	POLYGON_OFFSET_FILL                          = 0x8037
	SAMPLE_ALPHA_TO_COVERAGE                     = 0x809E
	SAMPLE_COVERAGE                              = 0x80A0
	INVALID_ENUM                                 = 0x0500
	INVALID_VALUE                                = 0x0501
	INVALID_OPERATION                            = 0x0502
	OUT_OF_MEMORY                                = 0x0505
	CW                                           = 0x0900
	CCW                                          = 0x0901
	LINE_WIDTH                                   = 0x0B21
	ALIASED_POINT_SIZE_RANGE                     = 0x846D
	ALIASED_LINE_WIDTH_RANGE                     = 0x846E
	CULL_FACE_MODE                               = 0x0B45
	FRONT_FACE                                   = 0x0B46
	DEPTH_RANGE                                  = 0x0B70
	DEPTH_WRITEMASK                              = 0x0B72
	DEPTH_CLEAR_VALUE                            = 0x0B73
	DEPTH_FUNC                                   = 0x0B74
	STENCIL_CLEAR_VALUE                          = 0x0B91
	STENCIL_FUNC                                 = 0x0B92
	STENCIL_FAIL                                 = 0x0B94
	STENCIL_PASS_DEPTH_FAIL                      = 0x0B95
	STENCIL_PASS_DEPTH_PASS                      = 0x0B96
	STENCIL_REF                                  = 0x0B97
	STENCIL_VALUE_MASK                           = 0x0B93
	STENCIL_WRITEMASK                            = 0x0B98
	STENCIL_BACK_FUNC                            = 0x8800
	STENCIL_BACK_FAIL                            = 0x8801
	STENCIL_BACK_PASS_DEPTH_FAIL                 = 0x8802
	STENCIL_BACK_PASS_DEPTH_PASS                 = 0x8803
	STENCIL_BACK_REF                             = 0x8CA3
	STENCIL_BACK_VALUE_MASK                      = 0x8CA4
	STENCIL_BACK_WRITEMASK                       = 0x8CA5
	VIEWPORT                                     = 0x0BA2
	SCISSOR_BOX                                  = 0x0C10
	COLOR_CLEAR_VALUE                            = 0x0C22
	COLOR_WRITEMASK                              = 0x0C23
	UNPACK_ALIGNMENT                             = 0x0CF5
	PACK_ALIGNMENT                               = 0x0D05
	MAX_TEXTURE_SIZE                             = 0x0D33
	MAX_VIEWPORT_DIMS                            = 0x0D3A
	SUBPIXEL_BITS                                = 0x0D50
	RED_BITS                                     = 0x0D52
	GREEN_BITS                                   = 0x0D53
	BLUE_BITS                                    = 0x0D54
	ALPHA_BITS                                   = 0x0D55
	DEPTH_BITS                                   = 0x0D56
	STENCIL_BITS                                 = 0x0D57
	POLYGON_OFFSET_UNITS                         = 0x2A00
	POLYGON_OFFSET_FACTOR                        = 0x8038
	TEXTURE_BINDING_2D                           = 0x8069
	SAMPLE_BUFFERS                               = 0x80A8
	SAMPLES                                      = 0x80A9
	SAMPLE_COVERAGE_VALUE                        = 0x80AA
	SAMPLE_COVERAGE_INVERT                       = 0x80AB
	NUM_COMPRESSED_TEXTURE_FORMATS               = 0x86A2
	COMPRESSED_TEXTURE_FORMATS                   = 0x86A3
	DONT_CARE                                    = 0x1100
	FASTEST                                      = 0x1101
	NICEST                                       = 0x1102
	GENERATE_MIPMAP_HINT                         = 0x8192
	BYTE                                         = 0x1400
	UNSIGNED_BYTE                                = 0x1401
	SHORT                                        = 0x1402
	UNSIGNED_SHORT                               = 0x1403
	INT                                          = 0x1404
	UNSIGNED_INT                                 = 0x1405
	FLOAT                                        = 0x1406
	FIXED                                        = 0x140C
	DEPTH_COMPONENT                              = 0x1902
	ALPHA                                        = 0x1906
	RGB                                          = 0x1907
	RGBA                                         = 0x1908
	LUMINANCE                                    = 0x1909
	LUMINANCE_ALPHA                              = 0x190A
	UNSIGNED_SHORT_4_4_4_4                       = 0x8033
	UNSIGNED_SHORT_5_5_5_1                       = 0x8034
	UNSIGNED_SHORT_5_6_5                         = 0x8363
	MAX_VERTEX_ATTRIBS                           = 0x8869
	MAX_VERTEX_UNIFORM_VECTORS                   = 0x8DFB
	MAX_VARYING_VECTORS                          = 0x8DFC
	MAX_COMBINED_TEXTURE_IMAGE_UNITS             = 0x8B4D
	MAX_VERTEX_TEXTURE_IMAGE_UNITS               = 0x8B4C
	MAX_TEXTURE_IMAGE_UNITS                      = 0x8872
	MAX_FRAGMENT_UNIFORM_VECTORS                 = 0x8DFD
	SHADER_TYPE                                  = 0x8B4F
	DELETE_STATUS                                = 0x8B80
	LINK_STATUS                                  = 0x8B82
	VALIDATE_STATUS                              = 0x8B83
	ATTACHED_SHADERS                             = 0x8B85
	ACTIVE_UNIFORMS                              = 0x8B86
	ACTIVE_UNIFORM_MAX_LENGTH                    = 0x8B87
	ACTIVE_ATTRIBUTES                            = 0x8B89
	ACTIVE_ATTRIBUTE_MAX_LENGTH                  = 0x8B8A
	SHADING_LANGUAGE_VERSION                     = 0x8B8C
	CURRENT_PROGRAM                              = 0x8B8D
	NEVER                                        = 0x0200
	LESS                                         = 0x0201
	EQUAL                                        = 0x0202
	LEQUAL                                       = 0x0203
	GREATER                                      = 0x0204
	NOTEQUAL                                     = 0x0205
	GEQUAL                                       = 0x0206
	ALWAYS                                       = 0x0207
	KEEP                                         = 0x1E00
	REPLACE                                      = 0x1E01
	INCR                                         = 0x1E02
	DECR                                         = 0x1E03
	INVERT                                       = 0x150A
	INCR_WRAP                                    = 0x8507
	DECR_WRAP                                    = 0x8508
	VENDOR                                       = 0x1F00
	RENDERER                                     = 0x1F01
	VERSION                                      = 0x1F02
	EXTENSIONS                                   = 0x1F03
	NEAREST                                      = 0x2600
	LINEAR                                       = 0x2601
	NEAREST_MIPMAP_NEAREST                       = 0x2700
	LINEAR_MIPMAP_NEAREST                        = 0x2701
	NEAREST_MIPMAP_LINEAR                        = 0x2702
	LINEAR_MIPMAP_LINEAR                         = 0x2703
	TEXTURE_MAG_FILTER                           = 0x2800
	TEXTURE_MIN_FILTER                           = 0x2801
	TEXTURE_WRAP_S                               = 0x2802
	TEXTURE_WRAP_T                               = 0x2803
	TEXTURE                                      = 0x1702
	TEXTURE_CUBE_MAP                             = 0x8513
	TEXTURE_BINDING_CUBE_MAP                     = 0x8514
	TEXTURE_CUBE_MAP_POSITIVE_X                  = 0x8515
	TEXTURE_CUBE_MAP_NEGATIVE_X                  = 0x8516
	TEXTURE_CUBE_MAP_POSITIVE_Y                  = 0x8517
	TEXTURE_CUBE_MAP_NEGATIVE_Y                  = 0x8518
	TEXTURE_CUBE_MAP_POSITIVE_Z                  = 0x8519
	TEXTURE_CUBE_MAP_NEGATIVE_Z                  = 0x851A
	MAX_CUBE_MAP_TEXTURE_SIZE                    = 0x851C
	TEXTURE0                                     = 0x84C0
	TEXTURE1                                     = 0x84C1
	TEXTURE2                                     = 0x84C2
	TEXTURE3                                     = 0x84C3
	TEXTURE4                                     = 0x84C4
	TEXTURE5                                     = 0x84C5
	TEXTURE6                                     = 0x84C6
	TEXTURE7                                     = 0x84C7
	TEXTURE8                                     = 0x84C8
	TEXTURE9                                     = 0x84C9
	TEXTURE10                                    = 0x84CA
	TEXTURE11                                    = 0x84CB
	TEXTURE12                                    = 0x84CC
	TEXTURE13                                    = 0x84CD
	TEXTURE14                                    = 0x84CE
	TEXTURE15                                    = 0x84CF
	TEXTURE16                                    = 0x84D0
	TEXTURE17                                    = 0x84D1
	TEXTURE18                                    = 0x84D2
	TEXTURE19                                    = 0x84D3
	TEXTURE20                                    = 0x84D4
	TEXTURE21                                    = 0x84D5
	TEXTURE22                                    = 0x84D6
	TEXTURE23                                    = 0x84D7
	TEXTURE24                                    = 0x84D8
	TEXTURE25                                    = 0x84D9
	TEXTURE26                                    = 0x84DA
	TEXTURE27                                    = 0x84DB
	TEXTURE28                                    = 0x84DC
	TEXTURE29                                    = 0x84DD
	TEXTURE30                                    = 0x84DE
	TEXTURE31                                    = 0x84DF
	ACTIVE_TEXTURE                               = 0x84E0
	REPEAT                                       = 0x2901
	CLAMP_TO_EDGE                                = 0x812F
	MIRRORED_REPEAT                              = 0x8370
	VERTEX_ATTRIB_ARRAY_ENABLED                  = 0x8622
	VERTEX_ATTRIB_ARRAY_SIZE                     = 0x8623
	VERTEX_ATTRIB_ARRAY_STRIDE                   = 0x8624
	VERTEX_ATTRIB_ARRAY_TYPE                     = 0x8625
	VERTEX_ATTRIB_ARRAY_NORMALIZED               = 0x886A
	VERTEX_ATTRIB_ARRAY_POINTER                  = 0x8645
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING           = 0x889F
	IMPLEMENTATION_COLOR_READ_TYPE               = 0x8B9A
	IMPLEMENTATION_COLOR_READ_FORMAT             = 0x8B9B
	COMPILE_STATUS                               = 0x8B81
	INFO_LOG_LENGTH                              = 0x8B84
	SHADER_SOURCE_LENGTH                         = 0x8B88
	SHADER_COMPILER                              = 0x8DFA
	SHADER_BINARY_FORMATS                        = 0x8DF8
	NUM_SHADER_BINARY_FORMATS                    = 0x8DF9
	LOW_FLOAT                                    = 0x8DF0
	MEDIUM_FLOAT                                 = 0x8DF1
	HIGH_FLOAT                                   = 0x8DF2
	LOW_INT                                      = 0x8DF3
	MEDIUM_INT                                   = 0x8DF4
	HIGH_INT                                     = 0x8DF5
	FRAMEBUFFER                                  = 0x8D40
	RENDERBUFFER                                 = 0x8D41
	RGBA4                                        = 0x8056
	RGB5_A1                                      = 0x8057
	RGB565                                       = 0x8D62
	DEPTH_COMPONENT16                            = 0x81A5
	STENCIL_INDEX8                               = 0x8D48
	RENDERBUFFER_WIDTH                           = 0x8D42
	RENDERBUFFER_HEIGHT                          = 0x8D43
	RENDERBUFFER_INTERNAL_FORMAT                 = 0x8D44
	RENDERBUFFER_RED_SIZE                        = 0x8D50
	RENDERBUFFER_GREEN_SIZE                      = 0x8D51
	RENDERBUFFER_BLUE_SIZE                       = 0x8D52
	RENDERBUFFER_ALPHA_SIZE                      = 0x8D53
	RENDERBUFFER_DEPTH_SIZE                      = 0x8D54
	RENDERBUFFER_STENCIL_SIZE                    = 0x8D55
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE           = 0x8CD0
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME           = 0x8CD1
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL         = 0x8CD2
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = 0x8CD3
	COLOR_ATTACHMENT0                            = 0x8CE0
	DEPTH_ATTACHMENT                             = 0x8D00
	STENCIL_ATTACHMENT                           = 0x8D20
	FRAMEBUFFER_COMPLETE                         = 0x8CD5
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT            = 0x8CD6
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT    = 0x8CD7
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS            = 0x8CD9
	FRAMEBUFFER_UNSUPPORTED                      = 0x8CDD
	FRAMEBUFFER_BINDING                          = 0x8CA6
	RENDERBUFFER_BINDING                         = 0x8CA7
	MAX_RENDERBUFFER_SIZE                        = 0x84E8
	INVALID_FRAMEBUFFER_OPERATION                = 0x0506
)

const (
	DEPTH_BUFFER_BIT   = 0x00000100
	STENCIL_BUFFER_BIT = 0x00000400
	COLOR_BUFFER_BIT   = 0x00004000
)

const (
	FLOAT_VEC2   = 0x8B50
	FLOAT_VEC3   = 0x8B51
	FLOAT_VEC4   = 0x8B52
	INT_VEC2     = 0x8B53
	INT_VEC3     = 0x8B54
	INT_VEC4     = 0x8B55
	BOOL         = 0x8B56
	BOOL_VEC2    = 0x8B57
	BOOL_VEC3    = 0x8B58
	BOOL_VEC4    = 0x8B59
	FLOAT_MAT2   = 0x8B5A
	FLOAT_MAT3   = 0x8B5B
	FLOAT_MAT4   = 0x8B5C
	SAMPLER_2D   = 0x8B5E
	SAMPLER_CUBE = 0x8B60
)

const (
	FRAGMENT_SHADER = 0x8B30
	VERTEX_SHADER   = 0x8B31
)

const (
	FALSE    = 0
	TRUE     = 1
	ZERO     = 0
	ONE      = 1
	NO_ERROR = 0
	NONE     = 0
)

const (
	LINE  = 0x1B01
	FILL  = 0x1B02
	POINT = 0x1B00
)
