package gl

import "fmt"

const (
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     = 0x82AF
	IMAGE_CLASS_4_X_32                                         = 0x82B9
	TEXTURE_MAX_LEVEL                                          = 0x813D
	COLOR_ATTACHMENT8                                          = 0x8CE8
	TEXTURE_1D_ARRAY                                           = 0x8C18
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        = 0x92D4
	READ_WRITE                                                 = 0x88BA
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 = 0x92CE
	UNSIGNED_SHORT_1_5_5_5_REV                                 = 0x8366
	TRANSFORM_FEEDBACK_PAUSED                                  = 0x8E23
	VIEW_CLASS_96_BITS                                         = 0x82C5
	REPLACE                                                    = 0x1E01
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      = 0x8C88
	MAX_COMPUTE_WORK_GROUP_SIZE                                = 0x91BF
	TEXTURE_ALPHA_TYPE                                         = 0x8C13
	TEXTURE_COMPARE_FUNC                                       = 0x884D
	MAX_DEPTH_TEXTURE_SAMPLES                                  = 0x910F
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         = 0x9103
	COLOR_RENDERABLE                                           = 0x8286
	FRONT_LEFT                                                 = 0x0400
	STREAM_DRAW                                                = 0x88E0
	MAX_VARYING_COMPONENTS                                     = 0x8B4B
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            = 0x8E1F
	TEXTURE20                                                  = 0x84D4
	VIEW_CLASS_8_BITS                                          = 0x82CB
	BUFFER_MAPPED                                              = 0x88BC
	ARRAY_SIZE                                                 = 0x92FB
	RG8UI                                                      = 0x8238
	TEXTURE_GREEN_TYPE                                         = 0x8C11
	FRAMEBUFFER_RENDERABLE_LAYERED                             = 0x828A
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           = 0x92D3
	UNSIGNED_INT_SAMPLER_2D_RECT                               = 0x8DD5
	MAX_IMAGE_UNITS                                            = 0x8F38
	TRUE                                                       = 1
	GREEN_INTEGER                                              = 0x8D95
	QUERY_BY_REGION_NO_WAIT                                    = 0x8E16
	INCR                                                       = 0x1E02
	COMPARE_REF_TO_TEXTURE                                     = 0x884E
	COLOR_ENCODING                                             = 0x8296
	SHADER                                                     = 0x82E1
	REPEAT                                                     = 0x2901
	TEXTURE26                                                  = 0x84DA
	SEPARATE_ATTRIBS                                           = 0x8C8D
	CLIP_DISTANCE2                                             = 0x3002
	OR_INVERTED                                                = 0x150D
	SRGB_EXT                                                   = 0x8C40
	NAND                                                       = 0x150E
	VIEWPORT                                                   = 0x0BA2
	RGB12                                                      = 0x8053
	TEXTURE_VIEW_NUM_LEVELS                                    = 0x82DC
	MAX_VERTEX_STREAMS                                         = 0x8E71
	UNIFORM_OFFSET                                             = 0x8A3B
	RENDERBUFFER_ALPHA_SIZE                                    = 0x8D53
	ANY_SAMPLES_PASSED_CONSERVATIVE                            = 0x8D6A
	TEXTURE_HEIGHT                                             = 0x1001
	TEXTURE21                                                  = 0x84D5
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      = 0x8E86
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             = 0x82B1
	UNIFORM_BUFFER                                             = 0x8A11
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               = 0x900B
	SAMPLER                                                    = 0x82E6
	VERTEX_ATTRIB_RELATIVE_OFFSET                              = 0x82D5
	TEXTURE_BUFFER_OFFSET                                      = 0x919D
	FRAGMENT_SHADER                                            = 0x8B30
	MAX_FRAMEBUFFER_WIDTH                                      = 0x9315
	FRAMEBUFFER_ATTACHMENT_LAYERED                             = 0x8DA7
	IMAGE_TEXEL_SIZE                                           = 0x82A7
	TEXTURE_SWIZZLE_RGBA                                       = 0x8E46
	IMAGE_BINDING_FORMAT                                       = 0x906E
	BUFFER_ACCESS_FLAGS                                        = 0x911F
	STENCIL_PASS_DEPTH_PASS                                    = 0x0B96
	COPY_READ_BUFFER                                           = COPY_READ_BUFFER_BINDING
	VIEWPORT_SUBPIXEL_BITS                                     = 0x825C
	TESS_GEN_MODE                                              = 0x8E76
	TEXTURE27                                                  = 0x84DB
	TESS_CONTROL_OUTPUT_VERTICES                               = 0x8E75
	TEXTURE_SWIZZLE_A                                          = 0x8E45
	PIXEL_BUFFER_BARRIER_BIT                                   = 0x00000080
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      = 0x8210
	TRANSFORM_FEEDBACK                                         = 0x8E22
	DECR                                                       = 0x1E03
	RGBA2                                                      = 0x8055
	LINE_LOOP                                                  = 0x0002
	COMPRESSED_SRGB_S3TC_DXT1_EXT                              = 0x8C4C
	FULL_SUPPORT                                               = 0x82B7
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          = 0x8C8F
	SRC_ALPHA                                                  = 0x0302
	SAMPLE_POSITION                                            = 0x8E50
	SAMPLER_1D                                                 = 0x8B5D
	COMPRESSED_SRGB8_ETC2                                      = 0x9275
	TEXTURE_BLUE_SIZE                                          = 0x805E
	VIEW_CLASS_16_BITS                                         = 0x82CA
	RG8I                                                       = 0x8237
	FLOAT_VEC3                                                 = 0x8B51
	COLOR_ATTACHMENT4                                          = 0x8CE4
	CLEAR_BUFFER                                               = 0x82B4
	TEXTURE_1D                                                 = 0x0DE0
	FIXED                                                      = 0x140C
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  = 0x8CD7
	COMPRESSED_SRGB_ALPHA_EXT                                  = 0x8C49
	BLEND_SRC_ALPHA                                            = 0x80CB
	SIGNALED                                                   = 0x9119
	DEBUG_LOGGED_MESSAGES                                      = 0x9145
	REFERENCED_BY_TESS_CONTROL_SHADER                          = 0x9307
	RGB32UI                                                    = 0x8D71
	VALIDATE_STATUS                                            = 0x8B83
	IMAGE_PIXEL_TYPE                                           = 0x82AA
	SRC_ALPHA_SATURATE                                         = 0x0308
	COMPUTE_SHADER                                             = 0x91B9
	SAMPLE_ALPHA_TO_ONE                                        = 0x809F
	DEBUG_SOURCE_APPLICATION                                   = 0x824A
	RENDERER                                                   = 0x1F01
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               = 0x92C5
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         = 0x889F
	UPPER_LEFT                                                 = 0x8CA2
	SHADER_STORAGE_BLOCK                                       = 0x92E6
	FRACTIONAL_ODD                                             = 0x8E7B
	DOUBLE_MAT4x3                                              = 0x8F4E
	VIEW_CLASS_BPTC_UNORM                                      = 0x82D2
	SAMPLER_2D_MULTISAMPLE                                     = 0x9108
	ACTIVE_UNIFORM_BLOCKS                                      = 0x8A36
	UNSIGNED_INT_ATOMIC_COUNTER                                = 0x92DB
	ONE_MINUS_SRC1_COLOR                                       = 0x88FA
	COLOR_ATTACHMENT2                                          = 0x8CE2
	TEXTURE_BINDING_2D_ARRAY                                   = 0x8C1D
	UNPACK_SWAP_BYTES                                          = 0x0CF0
	UNSIGNED_INT_SAMPLER_1D                                    = 0x8DD1
	MAX_COMPUTE_UNIFORM_BLOCKS                                 = 0x91BB
	UNPACK_SKIP_IMAGES                                         = 0x806D
	UNSIGNED_SHORT                                             = 0x1403
	MAX_COMBINED_ATOMIC_COUNTERS                               = 0x92D7
	DEBUG_OUTPUT                                               = 0x92E0
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x92C9
	FRAMEBUFFER_COMPLETE                                       = 0x8CD5
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                = 0x8518
	TEXTURE_2D                                                 = 0x0DE1
	DEPTH_WRITEMASK                                            = 0x0B72
	DRAW_INDIRECT_BUFFER                                       = 0x8F3F
	SRGB8_EXT                                                  = 0x8C41
	IMPLEMENTATION_COLOR_READ_TYPE                             = 0x8B9A
	MAX_TEXTURE_SIZE                                           = 0x0D33
	STENCIL_COMPONENTS                                         = 0x8285
	R16_SNORM                                                  = 0x8F98
	MAX_COMBINED_IMAGE_UNIFORMS                                = 0x90CF
	TEXTURE_COMPARE_MODE                                       = 0x884C
	UNSIGNED_BYTE_2_3_3_REV                                    = 0x8362
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       = 0x8CD2
	COLOR_ATTACHMENT7                                          = 0x8CE7
	EQUAL                                                      = 0x0202
	IMAGE_1D                                                   = 0x904C
	TEXTURE_MAX_ANISOTROPY_EXT                                 = 0x84FE
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             = 0x9061
	RGBA_INTEGER                                               = 0x8D99
	CLAMP_TO_BORDER                                            = 0x812D
	WRITE_ONLY                                                 = 0x88B9
	MAX_LAYERS                                                 = 0x8281
	TEXTURE_GREEN_SIZE                                         = 0x805D
	STENCIL_BACK_PASS_DEPTH_PASS                               = 0x8803
	TEXTURE_BASE_LEVEL                                         = 0x813C
	UNSIGNED_INT_24_8                                          = 0x84FA
	TEXTURE6                                                   = 0x84C6
	GREEN                                                      = 0x1904
	SLUMINANCE8_EXT                                            = 0x8C47
	TRIANGLE_FAN                                               = 0x0006
	STENCIL_INDEX1                                             = 0x8D46
	TEXTURE_CUBE_MAP                                           = 0x8513
	MAX_GEOMETRY_ATOMIC_COUNTERS                               = 0x92D5
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5C
	DEBUG_OUTPUT_SYNCHRONOUS                                   = 0x8242
	STENCIL_PASS_DEPTH_FAIL                                    = 0x0B95
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           = 0x910C
	FRAMEBUFFER_BLEND                                          = 0x828B
	BLEND_EQUATION_RGB                                         = 0x8009
	SAMPLER_CUBE                                               = 0x8B60
	AND_INVERTED                                               = 0x1504
	TEXTURE3                                                   = 0x84C3
	INTERNALFORMAT_DEPTH_SIZE                                  = 0x8275
	SRGB_READ                                                  = 0x8297
	POINTS                                                     = 0x0000
	COMPRESSED_RG_RGTC2                                        = 0x8DBD
	GEOMETRY_SHADER_BIT                                        = 0x00000004
	VERTEX_PROGRAM_POINT_SIZE                                  = 0x8642
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             = 0x8B4C
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  = 0x8A44
	ALREADY_SIGNALED                                           = 0x911A
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          = 0x886C
	TEXTURE11                                                  = 0x84CB
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            = 0x00000020
	MAX_FRAMEBUFFER_SAMPLES                                    = 0x9318
	TIMESTAMP                                                  = 0x8E28
	TEXTURE_VIEW_NUM_LAYERS                                    = 0x82DE
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       = 0x82AE
	SLUMINANCE8_ALPHA8_EXT                                     = 0x8C45
	MAX_VARYING_VECTORS                                        = 0x8DFC
	UNSIGNALED                                                 = 0x9118
	FLOAT_VEC2                                                 = 0x8B50
	TRANSFORM_FEEDBACK_VARYING                                 = 0x92F4
	BOOL_VEC3                                                  = 0x8B58
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             = 0x9107
	SAMPLE_ALPHA_TO_COVERAGE                                   = 0x809E
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               = 0x88FC
	MEDIUM_FLOAT                                               = 0x8DF1
	RENDERBUFFER_STENCIL_SIZE                                  = 0x8D55
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           = 0x8214
	VERTEX_SUBROUTINE_UNIFORM                                  = 0x92EE
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   = 0x9276
	DECR_WRAP                                                  = 0x8508
	VIEW_CLASS_S3TC_DXT1_RGBA                                  = 0x82CD
	BOOL                                                       = 0x8B56
	PROGRAM_POINT_SIZE                                         = 0x8642
	UNSIGNED_INT_SAMPLER_BUFFER                                = 0x8DD8
	VIEW_CLASS_RGTC2_RG                                        = 0x82D1
	LOW_INT                                                    = 0x8DF3
	REFERENCED_BY_VERTEX_SHADER                                = 0x9306
	ELEMENT_ARRAY_BUFFER_BINDING                               = 0x8895
	POLYGON_SMOOTH                                             = 0x0B41
	VERTEX_ATTRIB_BINDING                                      = 0x82D4
	COMPRESSED_SIGNED_RED_RGTC1                                = 0x8DBC
	UNIFORM_BLOCK                                              = 0x92E2
	FRAGMENT_TEXTURE                                           = 0x829F
	RGBA16                                                     = 0x805B
	LEFT                                                       = 0x0406
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         = 0x90C8
	BLUE                                                       = 0x1905
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               = 0x8CD3
	POLYGON_OFFSET_POINT                                       = 0x2A01
	DOUBLE_MAT2                                                = 0x8F46
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       = 0x8E81
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        = 0x8E47
	CLIP_DISTANCE0                                             = 0x3000
	COPY_WRITE_BUFFER                                          = COPY_WRITE_BUFFER_BINDING
	RGB10                                                      = 0x8052
	IMAGE_2D                                                   = 0x904D
	SAMPLER_2D_RECT_SHADOW                                     = 0x8B64
	TEXTURE_BLUE_TYPE                                          = 0x8C12
	INT_IMAGE_CUBE_MAP_ARRAY                                   = 0x905F
	SAMPLE_BUFFERS                                             = 0x80A8
	MAX_COMPUTE_WORK_GROUP_COUNT                               = 0x91BE
	POLYGON_OFFSET_LINE                                        = 0x2A02
	FRAMEBUFFER_DEFAULT                                        = 0x8218
	MAX_ARRAY_TEXTURE_LAYERS                                   = 0x88FF
	PACK_COMPRESSED_BLOCK_WIDTH                                = 0x912B
	R8UI                                                       = 0x8232
	INTERNALFORMAT_ALPHA_TYPE                                  = 0x827B
	POLYGON_OFFSET_FILL                                        = 0x8037
	TESS_GEN_SPACING                                           = 0x8E77
	FRONT_AND_BACK                                             = 0x0408
	COMPUTE_SUBROUTINE_UNIFORM                                 = 0x92F3
	UNSIGNED_BYTE                                              = 0x1401
	TEXTURE_MIN_FILTER                                         = 0x2801
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           = 0x90D6
	ACTIVE_UNIFORMS                                            = 0x8B86
	FRAMEBUFFER                                                = 0x8D40
	QUERY                                                      = 0x82E3
	SRGB_WRITE                                                 = 0x8298
	INTERNALFORMAT_BLUE_SIZE                                   = 0x8273
	UNSIGNED_INT_IMAGE_2D_ARRAY                                = 0x9069
	TEXTURE_STENCIL_SIZE                                       = 0x88F1
	MAX_NUM_COMPATIBLE_SUBROUTINES                             = 0x92F8
	SRGB_ALPHA                                                 = 0x8C42
	MAP_INVALIDATE_BUFFER_BIT                                  = 0x0008
	MAX_VERTEX_ATOMIC_COUNTERS                                 = 0x92D2
	DOUBLE_VEC2                                                = 0x8FFC
	DRAW_BUFFER9                                               = 0x882E
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              = 0x8A42
	MIN_MAP_BUFFER_ALIGNMENT                                   = 0x90BC
	INTERNALFORMAT_SHARED_SIZE                                 = 0x8277
	DEBUG_CALLBACK_USER_PARAM                                  = 0x8245
	UNIFORM_BARRIER_BIT                                        = 0x00000004
	TEXTURE24                                                  = 0x84D8
	COMPRESSED_RG11_EAC                                        = 0x9272
	SAMPLE_MASK                                                = 0x8E51
	DRAW_BUFFER11                                              = 0x8830
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               = 0x8E1E
	IMAGE_BUFFER                                               = 0x9051
	MAX_DEBUG_GROUP_STACK_DEPTH                                = 0x826C
	LOCATION_INDEX                                             = 0x930F
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       = 0x8E49
	STACK_OVERFLOW                                             = 0x0503
	COLOR_ATTACHMENT6                                          = 0x8CE6
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            = 0x90CB
	PACK_SKIP_IMAGES                                           = 0x806B
	MAX_COMBINED_DIMENSIONS                                    = 0x8282
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       = 0x8A43
	EQUIV                                                      = 0x1509
	TESS_EVALUATION_SHADER_BIT                                 = 0x00000010
	VERTEX_ATTRIB_ARRAY_TYPE                                   = 0x8625
	TEXTURE_CUBE_MAP_SEAMLESS                                  = 0x884F
	STENCIL_FAIL                                               = 0x0B94
	INTERNALFORMAT_PREFERRED                                   = 0x8270
	DOUBLE_MAT3                                                = 0x8F47
	FRONT_RIGHT                                                = 0x0401
	INTERNALFORMAT_SUPPORTED                                   = 0x826F
	RGB10_A2UI                                                 = 0x906F
	SRGB8                                                      = 0x8C41
	ONE_MINUS_DST_ALPHA                                        = 0x0305
	SAMPLER_3D                                                 = 0x8B5F
	TEXTURE_COMPRESSED_IMAGE_SIZE                              = 0x86A0
	STENCIL_INDEX4                                             = 0x8D47
	PACK_COMPRESSED_BLOCK_DEPTH                                = 0x912D
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            = 0x8E89
	VERTEX_ATTRIB_ARRAY_SIZE                                   = 0x8623
	POINT_SPRITE_COORD_ORIGIN                                  = 0x8CA0
	INT_IMAGE_BUFFER                                           = 0x905C
	COMPRESSED_RGBA_S3TC_DXT3_EXT                              = 0x83F2
	VERTEX_SHADER_BIT                                          = 0x00000001
	PROXY_TEXTURE_2D_MULTISAMPLE                               = 0x9101
	OBJECT_TYPE                                                = 0x9112
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          = 0x92CC
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        = 0x92C6
	COLOR_ATTACHMENT3                                          = 0x8CE3
	RGB5                                                       = 0x8050
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         = 0x92F0
	TEXTURE_2D_MULTISAMPLE                                     = 0x9100
	POINT_FADE_THRESHOLD_SIZE                                  = 0x8128
	MAX_HEIGHT                                                 = 0x827F
	NUM_SHADER_BINARY_FORMATS                                  = 0x8DF9
	NONE                                                       = 0
	MAX_CUBE_MAP_TEXTURE_SIZE                                  = 0x851C
	DEPTH_BUFFER_BIT                                           = 0x00000100
	FRAMEBUFFER_SRGB                                           = 0x8DB9
	SAMPLE_COVERAGE                                            = 0x80A0
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         = 0x90DC
	TEXTURE_CUBE_MAP_POSITIVE_X                                = 0x8515
	TEXTURE_FETCH_BARRIER_BIT                                  = 0x00000008
	TEXTURE_WRAP_T                                             = 0x2803
	OFFSET                                                     = 0x92FC
	PROXY_TEXTURE_RECTANGLE                                    = 0x84F7
	RG_INTEGER                                                 = 0x8228
	LINES_ADJACENCY                                            = 0x000A
	ALL_BARRIER_BITS                                           = 0xFFFFFFFF
	TEXTURE_BINDING_RECTANGLE                                  = 0x84F6
	NEAREST_MIPMAP_NEAREST                                     = 0x2700
	TEXTURE12                                                  = 0x84CC
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            = 0x8212
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              = 0x8DD6
	READ_FRAMEBUFFER_BINDING                                   = 0x8CAA
	PROGRAM                                                    = 0x82E2
	TEXTURE_INTERNAL_FORMAT                                    = 0x1003
	STENCIL_BACK_PASS_DEPTH_FAIL                               = 0x8802
	BUFFER_BINDING                                             = 0x9302
	MAX_ELEMENT_INDEX                                          = 0x8D6B
	FRAGMENT_SHADER_DERIVATIVE_HINT                            = 0x8B8B
	DRAW_BUFFER14                                              = 0x8833
	MAX_ELEMENTS_VERTICES                                      = 0x80E8
	VERTEX_ATTRIB_ARRAY_INTEGER                                = 0x88FD
	PROGRAM_INPUT                                              = 0x92E3
	R8_SNORM                                                   = 0x8F94
	DRAW_INDIRECT_BUFFER_BINDING                               = 0x8F43
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     = 0x90D8
	INT_IMAGE_2D_RECT                                          = 0x905A
	POINT                                                      = 0x1B00
	LOW_FLOAT                                                  = 0x8DF0
	IMAGE_BINDING_LEVEL                                        = 0x8F3B
	RGB16_SNORM                                                = 0x8F9A
	MAX_FRAGMENT_UNIFORM_BLOCKS                                = 0x8A2D
	ZERO                                                       = 0
	MAX_DEBUG_MESSAGE_LENGTH                                   = 0x9143
	INT_IMAGE_1D_ARRAY                                         = 0x905D
	FRAMEBUFFER_RENDERABLE                                     = 0x8289
	PROGRAM_PIPELINE_BINDING                                   = 0x825A
	TEXTURE17                                                  = 0x84D1
	DRAW_BUFFER                                                = 0x0C01
	COPY_WRITE_BUFFER_BINDING                                  = 0x8F37
	FRAMEBUFFER_DEFAULT_HEIGHT                                 = 0x9311
	GEOMETRY_VERTICES_OUT                                      = 0x8916
	MAX_VERTEX_UNIFORM_VECTORS                                 = 0x8DFB
	MAX_VERTEX_UNIFORM_COMPONENTS                              = 0x8B4A
	TEXTURE_BINDING_1D                                         = 0x8068
	DST_COLOR                                                  = 0x0306
	DEBUG_GROUP_STACK_DEPTH                                    = 0x826D
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           = 0x9279
	BLUE_INTEGER                                               = 0x8D96
	UNPACK_ALIGNMENT                                           = 0x0CF5
	R8                                                         = 0x8229
	MIN_PROGRAM_TEXEL_OFFSET                                   = 0x8904
	RGB5_A1                                                    = 0x8057
	TESS_CONTROL_TEXTURE                                       = 0x829C
	MAX_PATCH_VERTICES                                         = 0x8E7D
	BACK_LEFT                                                  = 0x0402
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                = 0x851A
	COMPRESSED_RGBA8_ETC2_EAC                                  = 0x9278
	LOWER_LEFT                                                 = 0x8CA1
	PACK_IMAGE_HEIGHT                                          = 0x806C
	TEXTURE4                                                   = 0x84C4
	INVALID_INDEX                                              = 0xFFFFFFFF
	READ_ONLY                                                  = 0x88B8
	MAX_COLOR_TEXTURE_SAMPLES                                  = 0x910E
	TEXTURE31                                                  = 0x84DF
	BLEND_SRC_RGB                                              = 0x80C9
	SRC_COLOR                                                  = 0x0300
	PACK_SKIP_ROWS                                             = 0x0D03
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             = 0x9128
	TEXTURE_SWIZZLE_G                                          = 0x8E43
	LOCATION                                                   = 0x930E
	VIEW_CLASS_RGTC1_RED                                       = 0x82D0
	INT_SAMPLER_CUBE_MAP_ARRAY                                 = 0x900E
	MAP_INVALIDATE_RANGE_BIT                                   = 0x0004
	UNIFORM                                                    = 0x92E1
	CURRENT_VERTEX_ATTRIB                                      = 0x8626
	SAMPLE_SHADING                                             = 0x8C36
	STENCIL_TEST                                               = 0x0B90
	STENCIL_BUFFER_BIT                                         = 0x00000400
	VIEW_CLASS_32_BITS                                         = 0x82C8
	ATOMIC_COUNTER_BUFFER_INDEX                                = 0x9301
	MATRIX_STRIDE                                              = 0x92FF
	COLOR_BUFFER_BIT                                           = 0x00004000
	ACTIVE_VARIABLES                                           = 0x9305
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            = 0x919F
	DEBUG_TYPE_PUSH_GROUP                                      = 0x8269
	R3_G3_B2                                                   = 0x2A10
	DRAW_FRAMEBUFFER_BINDING                                   = FRAMEBUFFER_BINDING
	PATCH_VERTICES                                             = 0x8E72
	INTERNALFORMAT_RED_SIZE                                    = 0x8271
	PROXY_TEXTURE_1D                                           = 0x8063
	SAMPLER_2D_MULTISAMPLE_ARRAY                               = 0x910B
	SMOOTH_POINT_SIZE_RANGE                                    = 0x0B12
	RGBA16I                                                    = 0x8D88
	OUT_OF_MEMORY                                              = 0x0505
	MANUAL_GENERATE_MIPMAP                                     = 0x8294
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             = 0x9124
	MAP_READ_BIT                                               = 0x0001
	BLEND_SRC                                                  = 0x0BE1
	RGBA_SNORM                                                 = 0x8F93
	ONE                                                        = 1
	MAX_GEOMETRY_SHADER_INVOCATIONS                            = 0x8E5A
	NAME_LENGTH                                                = 0x92F9
	INFO_LOG_LENGTH                                            = 0x8B84
	BLEND_DST                                                  = 0x0BE0
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        = 0x82AC
	VIEW_CLASS_S3TC_DXT5_RGBA                                  = 0x82CF
	SHADER_STORAGE_BUFFER                                      = 0x90D2
	TEXTURE_CUBE_MAP_POSITIVE_Z                                = 0x8519
	MAX_DRAW_BUFFERS                                           = 0x8824
	DRAW_BUFFER4                                               = 0x8829
	FLOAT_MAT4                                                 = 0x8B5C
	DEPTH_COMPONENTS                                           = 0x8284
	REFERENCED_BY_COMPUTE_SHADER                               = 0x930B
	MAX_GEOMETRY_IMAGE_UNIFORMS                                = 0x90CD
	ATOMIC_COUNTER_BUFFER_START                                = 0x92C2
	COLOR_ATTACHMENT1                                          = 0x8CE1
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        = 0x92DA
	TRIANGLES                                                  = 0x0004
	FILL                                                       = 0x1B02
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         = 0x8264
	DEBUG_TYPE_PERFORMANCE                                     = 0x8250
	RGB16F                                                     = 0x881B
	TEXTURE_ALPHA_SIZE                                         = 0x805F
	CONTEXT_FLAG_DEBUG_BIT                                     = 0x00000002
	DEPTH_CLEAR_VALUE                                          = 0x0B73
	TRANSFORM_FEEDBACK_VARYINGS                                = 0x8C83
	SHADER_IMAGE_LOAD                                          = 0x82A4
	MAX_WIDTH                                                  = 0x827E
	MAX_COLOR_ATTACHMENTS                                      = 0x8CDF
	ISOLINES                                                   = 0x8E7A
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        = 0x92CA
	RGB32F                                                     = 0x8815
	IMAGE_BINDING_NAME                                         = 0x8F3A
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   = 0x8A33
	DEBUG_TYPE_ERROR                                           = 0x824C
	DOUBLE_MAT3x4                                              = 0x8F4C
	AND                                                        = 0x1501
	DEBUG_SOURCE_API                                           = 0x8246
	MAX_VERTEX_UNIFORM_BLOCKS                                  = 0x8A2B
	TEXTURE13                                                  = 0x84CD
	TIMEOUT_EXPIRED                                            = 0x911B
	GEOMETRY_INPUT_TYPE                                        = 0x8917
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             = 0x824D
	UNSIGNED_SHORT_5_6_5_REV                                   = 0x8364
	NOTEQUAL                                                   = 0x0205
	IMAGE_BINDING_LAYERED                                      = 0x8F3C
	STENCIL                                                    = 0x1802
	COLOR_LOGIC_OP                                             = 0x0BF2
	INTERNALFORMAT_RED_TYPE                                    = 0x8278
	TRIANGLE_STRIP                                             = 0x0005
	TEXTURE_BORDER_COLOR                                       = 0x1004
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             = 0x8E70
	IMAGE_BINDING_ACCESS                                       = 0x8F3E
	IMAGE_2D_RECT                                              = 0x904F
	NOR                                                        = 0x1508
	SAMPLE_COVERAGE_INVERT                                     = 0x80AB
	CW                                                         = 0x0900
	ARRAY_STRIDE                                               = 0x92FE
	VIEW_CLASS_S3TC_DXT1_RGB                                   = 0x82CC
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       = 0x9105
	ALL_SHADER_BITS                                            = 0xFFFFFFFF
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT                        = 0x8C4F
	IMAGE_CLASS_4_X_8                                          = 0x82BF
	ONE_MINUS_SRC_ALPHA                                        = 0x0303
	DEBUG_SEVERITY_MEDIUM                                      = 0x9147
	DYNAMIC_COPY                                               = 0x88EA
	MAX_SAMPLE_MASK_WORDS                                      = 0x8E59
	GEQUAL                                                     = 0x0206
	IMAGE_CLASS_2_X_16                                         = 0x82BD
	NEAREST_MIPMAP_LINEAR                                      = 0x2702
	COMPUTE_TEXTURE                                            = 0x82A0
	DEBUG_SEVERITY_NOTIFICATION                                = 0x826B
	DEPTH32F_STENCIL8                                          = 0x8CAD
	UNSIGNED_SHORT_5_5_5_1                                     = 0x8034
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  = 0x910D
	SRGB8_ALPHA8                                               = 0x8C43
	R16F                                                       = 0x822D
	COMPRESSED_SIGNED_RG11_EAC                                 = 0x9273
	FIRST_VERTEX_CONVENTION                                    = 0x8E4D
	ONE_MINUS_DST_COLOR                                        = 0x0307
	ANY_SAMPLES_PASSED                                         = 0x8C2F
	UNSIGNED_INT_SAMPLER_3D                                    = 0x8DD3
	BUFFER                                                     = 0x82E0
	UNIFORM_TYPE                                               = 0x8A37
	TEXTURE_COMPRESSION_HINT                                   = 0x84EF
	UNIFORM_BUFFER_START                                       = 0x8A29
	TEXTURE14                                                  = 0x84CE
	VERTEX_SHADER                                              = 0x8B31
	BUFFER_USAGE                                               = 0x8765
	ELEMENT_ARRAY_BARRIER_BIT                                  = 0x00000002
	AUTO_GENERATE_MIPMAP                                       = 0x8295
	TEXTURE10                                                  = 0x84CA
	DOUBLE_VEC3                                                = 0x8FFD
	NUM_PROGRAM_BINARY_FORMATS                                 = 0x87FE
	UNPACK_SKIP_PIXELS                                         = 0x0CF4
	COPY                                                       = 0x1503
	TESS_CONTROL_SUBROUTINE                                    = 0x92E9
	SLUMINANCE_EXT                                             = 0x8C46
	FRONT_FACE                                                 = 0x0B46
	MAX_FRAGMENT_INPUT_COMPONENTS                              = 0x9125
	FLOAT_MAT4x2                                               = 0x8B69
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                = 0x8A46
	DEBUG_SOURCE_THIRD_PARTY                                   = 0x8249
	GET_TEXTURE_IMAGE_FORMAT                                   = 0x8291
	RGB                                                        = 0x1907
	LAST_VERTEX_CONVENTION                                     = 0x8E4E
	INTERNALFORMAT_GREEN_TYPE                                  = 0x8279
	SHADER_SOURCE_LENGTH                                       = 0x8B88
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS                      = 0x8F9F
	UNSIGNED_INT_8_8_8_8_REV                                   = 0x8367
	SAMPLER_BUFFER                                             = 0x8DC2
	RGBA16F                                                    = 0x881A
	MAX_VIEWPORT_DIMS                                          = 0x0D3A
	R32UI                                                      = 0x8236
	OR_REVERSE                                                 = 0x150B
	UNSIGNED_INT_IMAGE_3D                                      = 0x9064
	TESS_CONTROL_SHADER                                        = 0x8E88
	SAMPLER_2D_SHADOW                                          = 0x8B62
	DEBUG_SOURCE_OTHER                                         = 0x824B
	RGB565                                                     = 0x8D62
	DYNAMIC_DRAW                                               = 0x88E8
	ARRAY_BUFFER                                               = 0x8892
	RG32I                                                      = 0x823B
	UNPACK_COMPRESSED_BLOCK_SIZE                               = 0x912A
	CONDITION_SATISFIED                                        = 0x911C
	INT_SAMPLER_CUBE                                           = 0x8DCC
	INTERNALFORMAT_ALPHA_SIZE                                  = 0x8274
	IS_ROW_MAJOR                                               = 0x9300
	RIGHT                                                      = 0x0407
	DRAW_FRAMEBUFFER                                           = 0x8CA9
	RGB8UI                                                     = 0x8D7D
	TEXTURE30                                                  = 0x84DE
	OR                                                         = 0x1507
	UNPACK_COMPRESSED_BLOCK_WIDTH                              = 0x9127
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      = 0x8C76
	R16                                                        = 0x822A
	UNIFORM_IS_ROW_MAJOR                                       = 0x8A3E
	BLOCK_INDEX                                                = 0x92FD
	SHADER_TYPE                                                = 0x8B4F
	INT_VEC3                                                   = 0x8B54
	TEXTURE_BINDING_2D_MULTISAMPLE                             = 0x9104
	MIPMAP                                                     = 0x8293
	STENCIL_ATTACHMENT                                         = 0x8D20
	AND_REVERSE                                                = 0x1502
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         = 0x92DC
	RGBA16_SNORM                                               = 0x8F9B
	UNIFORM_ARRAY_STRIDE                                       = 0x8A3C
	FRAGMENT_SUBROUTINE                                        = 0x92EC
	RG16I                                                      = 0x8239
	RGBA                                                       = 0x1908
	ACTIVE_ATTRIBUTES                                          = 0x8B89
	FLOAT_MAT2                                                 = 0x8B5A
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   = 0x8E85
	HIGH_FLOAT                                                 = 0x8DF2
	PRIMITIVE_RESTART_FIXED_INDEX                              = 0x8D69
	MIN_SAMPLE_SHADING_VALUE                                   = 0x8C37
	PATCHES                                                    = 0x000E
	UNIFORM_BUFFER_BINDING                                     = 0x8A28
	TEXTURE16                                                  = 0x84D0
	MAX_COMPUTE_IMAGE_UNIFORMS                                 = 0x91BD
	RGB9_E5                                                    = 0x8C3D
	LESS                                                       = 0x0201
	COPY_INVERTED                                              = 0x150C
	UNSIGNED_INT_SAMPLER_2D                                    = 0x8DD2
	TRIANGLES_ADJACENCY                                        = 0x000C
	TEXTURE_COMPRESSED                                         = 0x86A1
	INT_IMAGE_2D_ARRAY                                         = 0x905E
	DEPTH_ATTACHMENT                                           = 0x8D00
	DOUBLE_MAT2x4                                              = 0x8F4A
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            = 0x91BC
	RG16UI                                                     = 0x823A
	DEPTH_RANGE                                                = 0x0B70
	POLYGON_SMOOTH_HINT                                        = 0x0C53
	RENDERBUFFER_RED_SIZE                                      = 0x8D50
	MAX_TESS_GEN_LEVEL                                         = 0x8E7E
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        = 0x910A
	VERTEX_SUBROUTINE                                          = 0x92E8
	PROGRAM_PIPELINE                                           = 0x82E4
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             = 0x8262
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 = 0x8C80
	COLOR_ATTACHMENT5                                          = 0x8CE5
	TESS_CONTROL_SUBROUTINE_UNIFORM                            = 0x92EF
	MAX_SAMPLES                                                = 0x8D57
	BUFFER_SIZE                                                = 0x8764
	COMPRESSED_SIGNED_RG_RGTC2                                 = 0x8DBE
	MAX_VIEWPORTS                                              = 0x825B
	STENCIL_BACK_VALUE_MASK                                    = 0x8CA4
	MAX_DEPTH                                                  = 0x8280
	FLOAT_MAT4x3                                               = 0x8B6A
	DOUBLE_MAT4                                                = 0x8F48
	DYNAMIC_READ                                               = 0x88E9
	INT_SAMPLER_2D_RECT                                        = 0x8DCD
	PACK_COMPRESSED_BLOCK_HEIGHT                               = 0x912C
	TEXTURE29                                                  = 0x84DD
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        = 0x90C9
	VERTEX_ATTRIB_ARRAY_LONG                                   = 0x874E
	BGRA                                                       = 0x80E1
	TEXTURE_2D_MULTISAMPLE_ARRAY                               = 0x9102
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       = MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS
	PRIMITIVE_RESTART_INDEX                                    = 0x8F9E
	DEPTH                                                      = 0x1801
	STENCIL_RENDERABLE                                         = 0x8288
	READ_BUFFER                                                = 0x0C02
	COLOR_ATTACHMENT12                                         = 0x8CEC
	TESS_EVALUATION_SUBROUTINE                                 = 0x92EA
	VERTEX_ATTRIB_ARRAY_ENABLED                                = 0x8622
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              = 0x900D
	RGB16UI                                                    = 0x8D77
	INTERNALFORMAT_GREEN_SIZE                                  = 0x8272
	TEXTURE_VIEW                                               = 0x82B5
	ACTIVE_TEXTURE                                             = 0x84E0
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5E
	LINEAR_MIPMAP_NEAREST                                      = 0x2701
	BGR_INTEGER                                                = 0x8D9A
	COMPUTE_LOCAL_WORK_SIZE                                    = 0x8267
	RG32UI                                                     = 0x823C
	RG8_SNORM                                                  = 0x8F95
	FALSE                                                      = 0
	GEOMETRY_OUTPUT_TYPE                                       = 0x8918
	NUM_ACTIVE_VARIABLES                                       = 0x9304
	FLOAT                                                      = 0x1406
	COLOR_ATTACHMENT0                                          = 0x8CE0
	MAP_FLUSH_EXPLICIT_BIT                                     = 0x0010
	TEXTURE_DEPTH_SIZE                                         = 0x884A
	MULTISAMPLE                                                = 0x809D
	COMPRESSED_SRGB_EXT                                        = 0x8C48
	UNSIGNED_INT_VEC4                                          = 0x8DC8
	VIEW_CLASS_64_BITS                                         = 0x82C6
	COMPRESSED_SRGB                                            = 0x8C48
	FLOAT_MAT2x4                                               = 0x8B66
	INTERNALFORMAT_STENCIL_TYPE                                = 0x827D
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             = 0x8C85
	RENDERBUFFER_WIDTH                                         = 0x8D42
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       = 0x8DA8
	KEEP                                                       = 0x1E00
	DOUBLE_MAT3x2                                              = 0x8F4B
	TEXTURE_SWIZZLE_R                                          = 0x8E42
	PROGRAM_BINARY_FORMATS                                     = 0x87FF
	STENCIL_BACK_FUNC                                          = 0x8800
	R32I                                                       = 0x8235
	UNSIGNED_INT_SAMPLER_CUBE                                  = 0x8DD4
	RG16F                                                      = 0x822F
	GEOMETRY_TEXTURE                                           = 0x829E
	UNIFORM_MATRIX_STRIDE                                      = 0x8A3D
	RENDERBUFFER                                               = 0x8D41
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    = 0x906C
	SHADER_STORAGE_BUFFER_START                                = 0x90D4
	SAMPLES                                                    = 0x80A9
	VIEW_CLASS_24_BITS                                         = 0x82C9
	TEXTURE_RED_SIZE                                           = 0x805C
	RG32F                                                      = 0x8230
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      = 0x82AD
	DEBUG_SOURCE_SHADER_COMPILER                               = 0x8248
	SRGB                                                       = 0x8C40
	SHADER_IMAGE_STORE                                         = 0x82A5
	R11F_G11F_B10F                                             = 0x8C3A
	DOUBLEBUFFER                                               = 0x0C32
	GEOMETRY_SHADER_INVOCATIONS                                = 0x887F
	FRAMEBUFFER_BARRIER_BIT                                    = 0x00000400
	INTERNALFORMAT_BLUE_TYPE                                   = 0x827A
	PACK_ROW_LENGTH                                            = 0x0D02
	MAX_COMBINED_UNIFORM_BLOCKS                                = 0x8A2E
	INVERT                                                     = 0x150A
	SYNC_FENCE                                                 = 0x9116
	SAMPLER_1D_ARRAY_SHADOW                                    = 0x8DC3
	SHORT                                                      = 0x1402
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             = 0x900A
	MAX_PROGRAM_TEXEL_OFFSET                                   = 0x8905
	POLYGON_MODE                                               = 0x0B40
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     = 0x8E80
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 = 0x9314
	SCISSOR_BOX                                                = 0x0C10
	INT_SAMPLER_1D_ARRAY                                       = 0x8DCE
	VERTEX_ATTRIB_ARRAY_DIVISOR                                = 0x88FE
	UNPACK_COMPRESSED_BLOCK_DEPTH                              = 0x9129
	DEBUG_TYPE_OTHER                                           = 0x8251
	LINE_STRIP                                                 = 0x0003
	VENDOR                                                     = 0x1F00
	UNSIGNED_INT_IMAGE_CUBE                                    = 0x9066
	RGB8_SNORM                                                 = 0x8F96
	DEPTH_STENCIL_ATTACHMENT                                   = 0x821A
	MAX_TEXTURE_BUFFER_SIZE                                    = 0x8C2B
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          = 0x8215
	INVALID_ENUM                                               = 0x0500
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        = 0x8217
	TEXTURE_3D                                                 = 0x806F
	FRAMEBUFFER_BINDING                                        = 0x8CA6
	MAX_VARYING_FLOATS                                         = 0x8B4B
	BUFFER_UPDATE_BARRIER_BIT                                  = 0x00000200
	PROXY_TEXTURE_2D_ARRAY                                     = 0x8C1B
	FIXED_ONLY                                                 = 0x891D
	UNSIGNED_INT_IMAGE_1D_ARRAY                                = 0x9068
	STENCIL_WRITEMASK                                          = 0x0B98
	VERTEX_ARRAY_BINDING                                       = 0x85B5
	MAX_FRAMEBUFFER_HEIGHT                                     = 0x9316
	BUFFER_VARIABLE                                            = 0x92E5
	IMAGE_COMPATIBILITY_CLASS                                  = 0x82A8
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            = 0x92C4
	RGBA12                                                     = 0x805A
	QUERY_WAIT                                                 = 0x8E13
	TEXTURE_BUFFER_DATA_STORE_BINDING                          = 0x8C2D
	LINE                                                       = 0x1B01
	RGB16                                                      = 0x8054
	DELETE_STATUS                                              = 0x8B80
	DOUBLE                                                     = 0x140A
	DRAW_BUFFER0                                               = 0x8825
	DEBUG_SEVERITY_HIGH                                        = 0x9146
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         = 0x90DA
	MAX_CLIP_DISTANCES                                         = 0x0D32
	COLOR_WRITEMASK                                            = 0x0C23
	CLIP_DISTANCE3                                             = 0x3003
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      = 0x8211
	LAYER_PROVOKING_VERTEX                                     = 0x825E
	UNIFORM_SIZE                                               = 0x8A38
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       = 0x8A35
	COMPRESSED_SRGB_ALPHA                                      = 0x8C49
	UNSIGNED_INT_5_9_9_9_REV                                   = 0x8C3E
	DRAW_BUFFER15                                              = 0x8834
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            = 0x8A34
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              = 0x824E
	PATCH_DEFAULT_INNER_LEVEL                                  = 0x8E73
	GEOMETRY_SHADER                                            = 0x8DD9
	IMAGE_BINDING_LAYER                                        = 0x8F3D
	TEXTURE_WRAP_S                                             = 0x2802
	UNSIGNED_INT_IMAGE_BUFFER                                  = 0x9067
	IS_PER_PATCH                                               = 0x92E7
	RGB8                                                       = 0x8051
	MAX_FRAGMENT_UNIFORM_VECTORS                               = 0x8DFD
	COMPRESSED_SIGNED_R11_EAC                                  = 0x9271
	SHADER_STORAGE_BARRIER_BIT                                 = 0x2000
	RED_INTEGER                                                = 0x8D94
	DRAW_BUFFER1                                               = 0x8826
	ELEMENT_ARRAY_BUFFER                                       = 0x8893
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              = 0x8DD7
	ACTIVE_SUBROUTINE_MAX_LENGTH                               = 0x8E48
	READ_PIXELS_TYPE                                           = 0x828E
	TEXTURE                                                    = 0x1702
	INT_SAMPLER_2D                                             = 0x8DCA
	PROGRAM_OUTPUT                                             = 0x92E4
	GEOMETRY_SUBROUTINE                                        = 0x92EB
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 = 0x90EC
	UNPACK_SKIP_ROWS                                           = 0x0CF3
	IMAGE_CLASS_11_11_10                                       = 0x82C2
	DRAW_BUFFER13                                              = 0x8832
	CONTEXT_CORE_PROFILE_BIT                                   = 0x00000001
	BOOL_VEC2                                                  = 0x8B57
	COMPILE_STATUS                                             = 0x8B81
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            = 0x8B49
	PROVOKING_VERTEX                                           = 0x8E4F
	TRANSFORM_FEEDBACK_BUFFER                                  = 0x8C8E
	CLAMP_TO_EDGE                                              = 0x812F
	COLOR_COMPONENTS                                           = 0x8283
	INT_SAMPLER_3D                                             = 0x8DCB
	STATIC_COPY                                                = 0x88E6
	BLEND                                                      = 0x0BE2
	MAX_RECTANGLE_TEXTURE_SIZE                                 = 0x84F8
	SMOOTH_LINE_WIDTH_GRANULARITY                              = 0x0B23
	INT_SAMPLER_2D_ARRAY                                       = 0x8DCF
	DEPTH_STENCIL_TEXTURE_MODE                                 = 0x90EA
	MAX_GEOMETRY_INPUT_COMPONENTS                              = 0x9123
	BUFFER_MAP_LENGTH                                          = 0x9120
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             = 0x886A
	TEXTURE28                                                  = 0x84DC
	GREATER                                                    = 0x0204
	RGBA32UI                                                   = 0x8D70
	IMPLEMENTATION_COLOR_READ_FORMAT                           = 0x8B9B
	TESS_EVALUATION_SHADER                                     = 0x8E87
	UNPACK_ROW_LENGTH                                          = 0x0CF2
	TEXTURE_MAX_LOD                                            = 0x813B
	DOUBLE_VEC4                                                = 0x8FFE
	TEXTURE_CUBE_MAP_ARRAY                                     = 0x9009
	COLOR_CLEAR_VALUE                                          = 0x0C22
	TEXTURE_BINDING_2D                                         = 0x8069
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    = 0x92CD
	NUM_COMPRESSED_TEXTURE_FORMATS                             = 0x86A2
	TEXTURE_WRAP_R                                             = 0x8072
	MAX_COMPUTE_LOCAL_INVOCATIONS                              = 0x90EB
	FLOAT_MAT3x2                                               = 0x8B67
	LINEAR                                                     = 0x2601
	IMAGE_CLASS_1_X_16                                         = 0x82BE
	FRAMEBUFFER_UNSUPPORTED                                    = 0x8CDD
	PROXY_TEXTURE_3D                                           = 0x8070
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              = 0x92D9
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         = 0x90CC
	DEPTH_CLAMP                                                = 0x864F
	RENDERBUFFER_HEIGHT                                        = 0x8D43
	NEVER                                                      = 0x0200
	TEXTURE_BINDING_CUBE_MAP                                   = 0x8514
	DOUBLE_MAT2x3                                              = 0x8F49
	TEXTURE_IMAGE_TYPE                                         = 0x8290
	DST_ALPHA                                                  = 0x0304
	TEXTURE_2D_ARRAY                                           = 0x8C1A
	DEBUG_TYPE_POP_GROUP                                       = 0x826A
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        = 0x8E7F
	COPY_READ_BUFFER_BINDING                                   = 0x8F36
	RG16                                                       = 0x822C
	MAX_VERTEX_OUTPUT_COMPONENTS                               = 0x9122
	IMAGE_CLASS_4_X_16                                         = 0x82BC
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            = 0x90C7
	UNIFORM_BLOCK_INDEX                                        = 0x8A3A
	MAX_DEBUG_LOGGED_MESSAGES                                  = 0x9144
	VERTEX_BINDING_STRIDE                                      = 0x82D8
	COLOR_ATTACHMENT11                                         = 0x8CEB
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         = 0x84F1
	DOUBLE_MAT4x2                                              = 0x8F4D
	ACTIVE_UNIFORM_MAX_LENGTH                                  = 0x8B87
	PACK_SKIP_PIXELS                                           = 0x0D04
	DEPTH_FUNC                                                 = 0x0B74
	COLOR_ATTACHMENT9                                          = 0x8CE9
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        = 0x92D0
	MAX_SUBROUTINES                                            = 0x8DE7
	DRAW_BUFFER5                                               = 0x882A
	MAX_TEXTURE_LOD_BIAS                                       = 0x84FD
	RENDERBUFFER_INTERNAL_FORMAT                               = 0x8D44
	STEREO                                                     = 0x0C33
	PROXY_TEXTURE_1D_ARRAY                                     = 0x8C19
	SAMPLER_2D                                                 = 0x8B5E
	MAX_TEXTURE_MAX_ANISOTROPY_EXT                             = 0x84FF
	NUM_SHADING_LANGUAGE_VERSIONS                              = 0x82E9
	RENDERBUFFER_DEPTH_SIZE                                    = 0x8D54
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    = 0x8266
	SRGB_ALPHA_EXT                                             = 0x8C42
	VERTEX_BINDING_DIVISOR                                     = 0x82D6
	MAX_VERTEX_IMAGE_UNIFORMS                                  = 0x90CA
	STENCIL_INDEX                                              = 0x1901
	UNIFORM_BLOCK_BINDING                                      = 0x8A3F
	TESS_EVALUATION_TEXTURE                                    = 0x829D
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         = 0x90DD
	CONTEXT_PROFILE_MASK                                       = 0x9126
	BYTE                                                       = 0x1400
	PROGRAM_SEPARABLE                                          = 0x8258
	R32F                                                       = 0x822E
	TRANSFORM_FEEDBACK_BUFFER_START                            = 0x8C84
	VIEW_CLASS_48_BITS                                         = 0x82C7
	UNSIGNED_NORMALIZED                                        = 0x8C17
	RGBA16UI                                                   = 0x8D76
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         = 0x8E5D
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         = 0x8CDC
	LINK_STATUS                                                = 0x8B82
	TEXTURE22                                                  = 0x84D6
	R8I                                                        = 0x8231
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              = 0x8C8A
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       = 0x8DE1
	MIRRORED_REPEAT                                            = 0x8370
	SAMPLER_2D_ARRAY_SHADOW                                    = 0x8DC4
	SAMPLER_BINDING                                            = 0x8919
	DEPTH_COMPONENT32                                          = 0x81A7
	VERTEX_TEXTURE                                             = 0x829B
	SAMPLE_MASK_VALUE                                          = 0x8E52
	UNSIGNED_INT_8_8_8_8                                       = 0x8035
	LINES                                                      = 0x0001
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          = 0x00000002
	PACK_SWAP_BYTES                                            = 0x0D00
	QUERY_BY_REGION_WAIT                                       = 0x8E15
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         = 0x8CDB
	SYNC_FLAGS                                                 = 0x9115
	MAX_GEOMETRY_OUTPUT_VERTICES                               = 0x8DE0
	BACK                                                       = 0x0405
	CURRENT_PROGRAM                                            = 0x8B8D
	TRANSFORM_FEEDBACK_ACTIVE                                  = 0x8E24
	DISPATCH_INDIRECT_BUFFER                                   = 0x90EE
	SHADER_COMPILER                                            = 0x8DFA
	INVALID_OPERATION                                          = 0x0502
	UNIFORM_BLOCK_DATA_SIZE                                    = 0x8A40
	UNSIGNED_SHORT_5_6_5                                       = 0x8363
	VERTEX_BINDING_OFFSET                                      = 0x82D7
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   = 0x8E4C
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          = 0x90DB
	QUERY_NO_WAIT                                              = 0x8E14
	CLIP_DISTANCE6                                             = 0x3006
	COMPRESSED_RGBA_S3TC_DXT1_EXT                              = 0x83F1
	COMPRESSED_RED                                             = 0x8225
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        = 0x92CF
	RGB_INTEGER                                                = 0x8D98
	COMPRESSED_RED_RGTC1                                       = 0x8DBB
	DEPTH_COMPONENT16                                          = 0x81A5
	BACK_RIGHT                                                 = 0x0403
	NUM_EXTENSIONS                                             = 0x821D
	SHADER_IMAGE_ATOMIC                                        = 0x82A6
	SLUMINANCE_ALPHA_EXT                                       = 0x8C44
	COMPRESSED_RGBA_S3TC_DXT5_EXT                              = 0x83F3
	TEXTURE23                                                  = 0x84D7
	SAMPLER_1D_SHADOW                                          = 0x8B61
	SUBPIXEL_BITS                                              = 0x0D50
	MAX_LABEL_LENGTH                                           = 0x82E8
	TEXTURE_BINDING_3D                                         = 0x806A
	UNSIGNED_INT_10F_11F_11F_REV                               = 0x8C3B
	POINT_SIZE_GRANULARITY                                     = 0x0B13
	STENCIL_VALUE_MASK                                         = 0x0B93
	TOP_LEVEL_ARRAY_SIZE                                       = 0x930C
	SYNC_GPU_COMMANDS_COMPLETE                                 = 0x9117
	DONT_CARE                                                  = 0x1100
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            = 0x84F0
	TEXTURE_MAG_FILTER                                         = 0x2800
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          = 0x906A
	IMAGE_CLASS_2_X_8                                          = 0x82C0
	SHADER_STORAGE_BUFFER_BINDING                              = 0x90D3
	DITHER                                                     = 0x0BD0
	INT_2_10_10_10_REV                                         = 0x8D9F
	CLIP_DISTANCE1                                             = 0x3001
	TEXTURE25                                                  = 0x84D9
	ACTIVE_RESOURCES                                           = 0x92F5
	TEXTURE_UPDATE_BARRIER_BIT                                 = 0x00000100
	MAX_UNIFORM_BLOCK_SIZE                                     = 0x8A30
	FRONT                                                      = 0x0404
	NUM_SAMPLE_COUNTS                                          = 0x9380
	TOP_LEVEL_ARRAY_STRIDE                                     = 0x930D
	PACK_COMPRESSED_BLOCK_SIZE                                 = 0x912E
	PIXEL_UNPACK_BUFFER_BINDING                                = 0x88EF
	MAX_COMPUTE_ATOMIC_COUNTERS                                = 0x8265
	UNSIGNED_INT_VEC2                                          = 0x8DC6
	POINT_SIZE_RANGE                                           = 0x0B12
	TEXTURE19                                                  = 0x84D3
	COMPRESSED_RGB                                             = 0x84ED
	COMPATIBLE_SUBROUTINES                                     = 0x8E4B
	GEOMETRY_SUBROUTINE_UNIFORM                                = 0x92F1
	TEXTURE1                                                   = 0x84C1
	SAMPLER_CUBE_MAP_ARRAY                                     = 0x900C
	TEXTURE_BINDING_1D_ARRAY                                   = 0x8C1C
	VIEWPORT_BOUNDS_RANGE                                      = 0x825D
	TEXTURE_WIDTH                                              = 0x1000
	ATTACHED_SHADERS                                           = 0x8B85
	HALF_FLOAT                                                 = 0x140B
	COMPRESSED_R11_EAC                                         = 0x9270
	SAMPLER_1D_ARRAY                                           = 0x8DC0
	TEXTURE_SHARED_SIZE                                        = 0x8C3F
	IMAGE_PIXEL_FORMAT                                         = 0x82A9
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    = 0x8E82
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           = 0x8C29
	IMAGE_2D_ARRAY                                             = 0x9053
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          = 0x8213
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    = 0x92C8
	TEXTURE8                                                   = 0x84C8
	READ_PIXELS                                                = 0x828C
	MAX_RENDERBUFFER_SIZE                                      = 0x84E8
	TEXTURE_MIN_LOD                                            = 0x813A
	COMPRESSED_SLUMINANCE_ALPHA_EXT                            = 0x8C4B
	MAX_TEXTURE_IMAGE_UNITS                                    = 0x8872
	BGRA_INTEGER                                               = 0x8D9B
	SRC1_COLOR                                                 = 0x88F9
	TEXTURE_DEPTH_TYPE                                         = 0x8C16
	INT_IMAGE_2D_MULTISAMPLE                                   = 0x9060
	UNSIGNED_INT_IMAGE_2D                                      = 0x9063
	MAX_FRAGMENT_IMAGE_UNIFORMS                                = 0x90CE
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  = 0x90D9
	INCR_WRAP                                                  = 0x8507
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           = 0x8243
	UNSIGNED_INT_VEC3                                          = 0x8DC7
	SAMPLER_2D_RECT                                            = 0x8B63
	FLOAT_VEC4                                                 = 0x8B52
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             = 0x92D8
	PATCH_DEFAULT_OUTER_LEVEL                                  = 0x8E74
	MAX_VERTEX_ATTRIB_BINDINGS                                 = 0x82DA
	RENDERBUFFER_BLUE_SIZE                                     = 0x8D52
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           = 0x8DE8
	TEXTURE_SWIZZLE_B                                          = 0x8E44
	CLAMP_READ_COLOR                                           = 0x891C
	INTERLEAVED_ATTRIBS                                        = 0x8C8C
	RG                                                         = 0x8227
	FRAGMENT_SHADER_BIT                                        = 0x00000002
	NOOP                                                       = 0x1505
	UNSIGNED_SHORT_4_4_4_4_REV                                 = 0x8365
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         = 0x90ED
	TEXTURE9                                                   = 0x84C9
	STENCIL_BACK_WRITEMASK                                     = 0x8CA5
	SET                                                        = 0x150F
	R16I                                                       = 0x8233
	TEXTURE0                                                   = 0x84C0
	RGBA8_SNORM                                                = 0x8F97
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    = 0x8C8B
	COLOR_ATTACHMENT15                                         = 0x8CEF
	MAX_VERTEX_ATTRIBS                                         = 0x8869
	DEPTH_COMPONENT32F                                         = 0x8CAC
	FRACTIONAL_EVEN                                            = 0x8E7C
	SHADING_LANGUAGE_VERSION                                   = 0x8B8C
	SHADER_BINARY_FORMATS                                      = 0x8DF8
	UNSIGNED_INT_IMAGE_2D_RECT                                 = 0x9065
	DRAW_BUFFER10                                              = 0x882F
	MEDIUM_INT                                                 = 0x8DF4
	DEPTH_STENCIL                                              = 0x84F9
	READ_FRAMEBUFFER                                           = 0x8CA8
	TEXTURE_BUFFER                                             = 0x8C2A
	CLEAR                                                      = 0x1500
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          = 0x92C7
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           = 0x8B4D
	DEBUG_TYPE_PORTABILITY                                     = 0x824F
	UNDEFINED_VERTEX                                           = 0x8260
	DRAW_BUFFER6                                               = 0x882B
	IMAGE_CLASS_10_10_10_2                                     = 0x82C3
	TESS_GEN_POINT_MODE                                        = 0x8E79
	INT_IMAGE_2D                                               = 0x9058
	FRAMEBUFFER_DEFAULT_SAMPLES                                = 0x9313
	BUFFER_ACCESS                                              = 0x88BB
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                = 0x8B8A
	INVALID_VALUE                                              = 0x0501
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5B
	ATOMIC_COUNTER_BUFFER_BINDING                              = 0x92C1
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         = 0x8D56
	PIXEL_PACK_BUFFER                                          = 0x88EB
	STENCIL_INDEX8                                             = 0x8D48
	DEPTH_RENDERABLE                                           = 0x8287
	DRAW_BUFFER3                                               = 0x8828
	DEBUG_SOURCE_WINDOW_SYSTEM                                 = 0x8247
	IMAGE_3D                                                   = 0x904E
	MAX_FRAGMENT_ATOMIC_COUNTERS                               = 0x92D6
	RED_SNORM                                                  = 0x8F90
	LINE_WIDTH                                                 = 0x0B21
	VERTEX_ATTRIB_ARRAY_POINTER                                = 0x8645
	QUERY_RESULT                                               = 0x8866
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          = 0x82D9
	VIEW_COMPATIBILITY_CLASS                                   = 0x82B6
	LINE_WIDTH_GRANULARITY                                     = 0x0B23
	UNSIGNED_INT                                               = 0x1405
	DRAW_BUFFER2                                               = 0x8827
	FLOAT_MAT2x3                                               = 0x8B65
	UNSIGNED_INT_IMAGE_1D                                      = 0x9062
	DEPTH_COMPONENT24                                          = 0x81A6
	PIXEL_PACK_BUFFER_BINDING                                  = 0x88ED
	IMAGE_CUBE_MAP_ARRAY                                       = 0x9054
	STENCIL_INDEX16                                            = 0x8D49
	DEPTH_COMPONENT                                            = 0x1902
	NO_ERROR                                                   = 0
	TEXTURE_LOD_BIAS                                           = 0x8501
	TEXTURE_SAMPLES                                            = 0x9106
	COMPRESSED_RGBA                                            = 0x84EE
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              = 0x8F39
	ALWAYS                                                     = 0x0207
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         = 0x8E83
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        = 0x92CB
	RED                                                        = 0x1903
	UNPACK_IMAGE_HEIGHT                                        = 0x806E
	UNPACK_LSB_FIRST                                           = 0x0CF1
	POLYGON_OFFSET_FACTOR                                      = 0x8038
	BGR                                                        = 0x80E0
	LINE_WIDTH_RANGE                                           = 0x0B22
	CLIP_DISTANCE4                                             = 0x3004
	POLYGON_OFFSET_UNITS                                       = 0x2A00
	PROGRAM_BINARY_LENGTH                                      = 0x8741
	RENDERBUFFER_BINDING                                       = 0x8CA7
	TEXTURE_RECTANGLE                                          = 0x84F5
	FLOAT_MAT3x4                                               = 0x8B68
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       = 0x886D
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        = 0x0001
	CLIP_DISTANCE5                                             = 0x3005
	ACTIVE_SUBROUTINES                                         = 0x8DE5
	STATIC_DRAW                                                = 0x88E4
	IMAGE_CLASS_2_X_32                                         = 0x82BA
	COMPRESSED_SLUMINANCE_EXT                                  = 0x8C4A
	BLEND_EQUATION_ALPHA                                       = 0x883D
	IMAGE_CUBE                                                 = 0x9050
	SCISSOR_TEST                                               = 0x0C11
	TEXTURE_DEPTH                                              = 0x8071
	SIGNED_NORMALIZED                                          = 0x8F9C
	DRAW_BUFFER8                                               = 0x882D
	REFERENCED_BY_TESS_EVALUATION_SHADER                       = 0x9308
	DEPTH24_STENCIL8                                           = 0x88F0
	IMAGE_2D_MULTISAMPLE                                       = 0x9055
	HIGH_INT                                                   = 0x8DF5
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        = 0x900F
	SYNC_FLUSH_COMMANDS_BIT                                    = 0x00000001
	TEXTURE_IMAGE_FORMAT                                       = 0x828F
	READ_PIXELS_FORMAT                                         = 0x828D
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     = 0x8A31
	INT_VEC2                                                   = 0x8B53
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  = 0x9277
	COLOR_ATTACHMENT13                                         = 0x8CED
	MAX_ELEMENTS_INDICES                                       = 0x80E9
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT                        = 0x8C4E
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            = 0x00000001
	TIME_ELAPSED                                               = 0x88BF
	INT_IMAGE_3D                                               = 0x9059
	ONE_MINUS_SRC1_ALPHA                                       = 0x88FB
	PIXEL_UNPACK_BUFFER                                        = 0x88EC
	MAX_SERVER_WAIT_TIMEOUT                                    = 0x9111
	COMPRESSED_RG                                              = 0x8226
	TRIANGLE_STRIP_ADJACENCY                                   = 0x000D
	BLEND_DST_ALPHA                                            = 0x80CA
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         = 0x90D7
	UNIFORM_BLOCK_NAME_LENGTH                                  = 0x8A41
	MAP_UNSYNCHRONIZED_BIT                                     = 0x0020
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            = 0x82B2
	TESS_CONTROL_SHADER_BIT                                    = 0x00000008
	INTERNALFORMAT_STENCIL_SIZE                                = 0x8276
	IMAGE_1D_ARRAY                                             = 0x9052
	LEQUAL                                                     = 0x0203
	PACK_LSB_FIRST                                             = 0x0D01
	UNSIGNED_SHORT_4_4_4_4                                     = 0x8033
	MAX_FRAMEBUFFER_LAYERS                                     = 0x9317
	TEXTURE_IMMUTABLE_FORMAT                                   = 0x912F
	INT_SAMPLER_2D_MULTISAMPLE                                 = 0x9109
	MAX_IMAGE_SAMPLES                                          = 0x906D
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           = TRANSFORM_FEEDBACK_PAUSED
	TEXTURE_CUBE_MAP_POSITIVE_Y                                = 0x8517
	MAX_3D_TEXTURE_SIZE                                        = 0x8073
	ATOMIC_COUNTER_BUFFER_SIZE                                 = 0x92C3
	TRANSFORM_FEEDBACK_BARRIER_BIT                             = 0x00000800
	STENCIL_CLEAR_VALUE                                        = 0x0B91
	TEXTURE_COMPRESSED_BLOCK_SIZE                              = 0x82B3
	ALIASED_LINE_WIDTH_RANGE                                   = 0x846E
	STATIC_READ                                                = 0x88E5
	CAVEAT_SUPPORT                                             = 0x82B8
	UNIFORM_NAME_LENGTH                                        = 0x8A39
	RGBA32F                                                    = 0x8814
	INVALID_FRAMEBUFFER_OPERATION                              = 0x0506
	MAX_INTEGER_SAMPLES                                        = 0x9110
	STREAM_READ                                                = 0x88E1
	LINE_SMOOTH                                                = 0x0B20
	RGBA4                                                      = 0x8056
	CLIP_DISTANCE7                                             = 0x3007
	RG16_SNORM                                                 = 0x8F99
	VERSION                                                    = 0x1F02
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         = 0x8CD1
	ARRAY_BUFFER_BINDING                                       = 0x8894
	FRAGMENT_SUBROUTINE_UNIFORM                                = 0x92F2
	NEAREST                                                    = 0x2600
	PRIMITIVE_RESTART                                          = 0x8F9D
	INT_SAMPLER_1D                                             = 0x8DC9
	TEXTURE_SHADOW                                             = 0x82A1
	STENCIL_BACK_REF                                           = 0x8CA3
	SRGB8_ALPHA8_EXT                                           = 0x8C43
	COLOR_ATTACHMENT10                                         = 0x8CEA
	RG8                                                        = 0x822B
	LINEAR_MIPMAP_LINEAR                                       = 0x2703
	RENDERBUFFER_GREEN_SIZE                                    = 0x8D51
	PROXY_TEXTURE_2D                                           = 0x8064
	LINE_SMOOTH_HINT                                           = 0x0C52
	MAX_NAME_LENGTH                                            = 0x92F6
	SAMPLER_2D_ARRAY                                           = 0x8DC1
	REFERENCED_BY_GEOMETRY_SHADER                              = 0x9309
	DISPLAY_LIST                                               = 0x82E7
	UNIFORM_BUFFER_SIZE                                        = 0x8A2A
	QUERY_RESULT_AVAILABLE                                     = 0x8867
	COMPRESSED_RGB_S3TC_DXT1_EXT                               = 0x83F0
	SMOOTH_LINE_WIDTH_RANGE                                    = 0x0B22
	RGB4                                                       = 0x804F
	FRAMEBUFFER_DEFAULT_LAYERS                                 = 0x9312
	IMAGE_CLASS_1_X_8                                          = 0x82C1
	FASTEST                                                    = 0x1101
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          = 0x8CD6
	TESS_GEN_VERTEX_ORDER                                      = 0x8E78
	INT_SAMPLER_BUFFER                                         = 0x8DD0
	BOOL_VEC4                                                  = 0x8B59
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       = 0x8CD4
	TEXTURE_BINDING_BUFFER                                     = 0x8C2C
	ACTIVE_PROGRAM                                             = 0x8259
	QUERY_COUNTER_BITS                                         = 0x8864
	SMOOTH_POINT_SIZE_GRANULARITY                              = 0x0B13
	MAX_NUM_ACTIVE_VARIABLES                                   = 0x92F7
	FRAMEBUFFER_DEFAULT_WIDTH                                  = 0x9310
	TEXTURE_GATHER                                             = 0x82A2
	ALPHA                                                      = 0x1906
	R16UI                                                      = 0x8234
	TEXTURE_IMMUTABLE_LEVELS                                   = 0x82DF
	FILTER                                                     = 0x829A
	UNSIGNED_INT_2_10_10_10_REV                                = 0x8368
	TEXTURE_CUBE_MAP_NEGATIVE_X                                = 0x8516
	MAX_TESS_PATCH_COMPONENTS                                  = 0x8E84
	SYNC_STATUS                                                = 0x9114
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT                        = 0x8C4D
	REFERENCED_BY_FRAGMENT_SHADER                              = 0x930A
	DRAW_BUFFER12                                              = 0x8831
	WAIT_FAILED                                                = 0x911D
	SHADER_STORAGE_BUFFER_SIZE                                 = 0x90D5
	ATOMIC_COUNTER_BUFFER                                      = 0x92C0
	TEXTURE2                                                   = 0x84C2
	VIEWPORT_INDEX_PROVOKING_VERTEX                            = 0x825F
	NICEST                                                     = 0x1102
	RGB8I                                                      = 0x8D8F
	DEBUG_TYPE_MARKER                                          = 0x8268
	PROXY_TEXTURE_CUBE_MAP                                     = 0x851B
	ONE_MINUS_SRC_COLOR                                        = 0x0301
	RGB_SNORM                                                  = 0x8F92
	POINT_SIZE                                                 = 0x0B11
	TEXTURE_RED_TYPE                                           = 0x8C10
	TEXTURE5                                                   = 0x84C5
	MAX_COMPUTE_UNIFORM_COMPONENTS                             = 0x8263
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          = 0x8216
	MAX_UNIFORM_LOCATIONS                                      = 0x826E
	CONTEXT_FLAGS                                              = 0x821E
	RG_SNORM                                                   = 0x8F91
	TEXTURE7                                                   = 0x84C7
	TEXTURE_VIEW_MIN_LAYER                                     = 0x82DD
	RGB16I                                                     = 0x8D89
	QUADS                                                      = 0x0007
	MINOR_VERSION                                              = 0x821C
	PRIMITIVES_GENERATED                                       = 0x8C87
	STREAM_COPY                                                = 0x88E2
	RGBA8                                                      = 0x8058
	TRANSFORM_FEEDBACK_BINDING                                 = 0x8E25
	PROGRAM_BINARY_RETRIEVABLE_HINT                            = 0x8257
	DEBUG_CALLBACK_FUNCTION                                    = 0x8244
	MAJOR_VERSION                                              = 0x821B
	INT_IMAGE_1D                                               = 0x9057
	TEXTURE15                                                  = 0x84CF
	COMPRESSED_TEXTURE_FORMATS                                 = 0x86A3
	PACK_ALIGNMENT                                             = 0x0D05
	RENDERBUFFER_SAMPLES                                       = 0x8CAB
	STENCIL_REF                                                = 0x0B97
	FRAMEBUFFER_UNDEFINED                                      = 0x8219
	SRC1_ALPHA                                                 = 0x8589
	LOGIC_OP_MODE                                              = 0x0BF0
	LINE_STRIP_ADJACENCY                                       = 0x000B
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        = 0x92D1
	DEPTH_TEST                                                 = 0x0B71
	UNSIGNED_INT_10_10_10_2                                    = 0x8036
	BLEND_DST_RGB                                              = 0x80C8
	DISPATCH_INDIRECT_BUFFER_BINDING                           = 0x90EF
	XOR                                                        = 0x1506
	VERTEX_ATTRIB_ARRAY_STRIDE                                 = 0x8624
	NUM_COMPATIBLE_SUBROUTINES                                 = 0x8E4A
	COMMAND_BARRIER_BIT                                        = 0x00000040
	INTERNALFORMAT_DEPTH_TYPE                                  = 0x827C
	MAX_UNIFORM_BUFFER_BINDINGS                                = 0x8A2F
	CCW                                                        = 0x0901
	CURRENT_QUERY                                              = 0x8865
	TEXTURE_GATHER_SHADOW                                      = 0x82A3
	SAMPLER_CUBE_SHADOW                                        = 0x8DC5
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           = TRANSFORM_FEEDBACK_ACTIVE
	CULL_FACE                                                  = 0x0B44
	INT_VEC4                                                   = 0x8B55
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            = 0x8DDF
	SAMPLES_PASSED                                             = 0x8914
	VIEW_CLASS_BPTC_FLOAT                                      = 0x82D3
	IMAGE_CLASS_1_X_32                                         = 0x82BB
	ATOMIC_COUNTER_BARRIER_BIT                                 = 0x00001000
	COLOR_ATTACHMENT14                                         = 0x8CEE
	INT                                                        = 0x1404
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         = 0x8CD0
	BUFFER_MAP_OFFSET                                          = 0x9121
	UNSIGNED_BYTE_3_3_2                                        = 0x8032
	STENCIL_BACK_FAIL                                          = 0x8801
	TIMEOUT_IGNORED                                            = 0xFFFFFFFFFFFFFFFF
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          = 0x906B
	TEXTURE18                                                  = 0x84D2
	RGBA8I                                                     = 0x8D8E
	INT_IMAGE_CUBE                                             = 0x905B
	CULL_FACE_MODE                                             = 0x0B45
	FLOAT_MAT3                                                 = 0x8B5B
	TEXTURE_BUFFER_SIZE                                        = 0x919E
	IMAGE_2D_MULTISAMPLE_ARRAY                                 = 0x9056
	SYNC_CONDITION                                             = 0x9113
	BUFFER_DATA_SIZE                                           = 0x9303
	EXTENSIONS                                                 = 0x1F03
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     = 0x90DF
	RGBA32I                                                    = 0x8D82
	TYPE                                                       = 0x92FA
	RASTERIZER_DISCARD                                         = 0x8C89
	TEXTURE_VIEW_MIN_LEVEL                                     = 0x82DB
	DRAW_BUFFER7                                               = 0x882C
	COMPRESSED_RGB8_ETC2                                       = 0x9274
	FLOAT_32_UNSIGNED_INT_24_8_REV                             = 0x8DAD
	STENCIL_FUNC                                               = 0x0B92
	VIEW_CLASS_S3TC_DXT3_RGBA                                  = 0x82CE
	MAP_WRITE_BIT                                              = 0x0002
	ACTIVE_SUBROUTINE_UNIFORMS                                 = 0x8DE6
	COLOR                                                      = 0x1800
	BUFFER_MAP_POINTER                                         = 0x88BD
	COMPUTE_SUBROUTINE                                         = 0x92ED
	RGB32I                                                     = 0x8D83
	TRANSFORM_FEEDBACK_BUFFER_MODE                             = 0x8C7F
	DEBUG_SEVERITY_LOW                                         = 0x9148
	MAX_SHADER_STORAGE_BLOCK_SIZE                              = 0x90DE
	STACK_UNDERFLOW                                            = 0x0504
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5F
	GET_TEXTURE_IMAGE_TYPE                                     = 0x8292
	SAMPLE_COVERAGE_VALUE                                      = 0x80AA
	RGB10_A2                                                   = 0x8059
	VIEW_CLASS_128_BITS                                        = 0x82C4
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         = 0x8E8A
	RGBA8UI                                                    = 0x8D7C
)

//	Returns the name of the specified Enum. Not recommended in a real-time loop.
func (_ GlUtil) EnumName(enum Enum) (name string) {
	switch enum {
	case STACK_OVERFLOW:
		name = "GL_STACK_OVERFLOW"
	case OUT_OF_MEMORY:
		name = "GL_OUT_OF_MEMORY"
	case VERTEX_SHADER:
		name = "GL_VERTEX_SHADER"
	case TESS_CONTROL_SHADER:
		name = "GL_TESS_CONTROL_SHADER"
	case INVALID_ENUM:
		name = "GL_INVALID_ENUM"
	case GEOMETRY_SHADER:
		name = "GL_GEOMETRY_SHADER"
	case TESS_EVALUATION_SHADER:
		name = "GL_TESS_EVALUATION_SHADER"
	case INVALID_OPERATION:
		name = "GL_INVALID_OPERATION"
	case INVALID_VALUE:
		name = "GL_INVALID_VALUE"
	case INVALID_FRAMEBUFFER_OPERATION:
		name = "GL_INVALID_FRAMEBUFFER_OPERATION"
	case STACK_UNDERFLOW:
		name = "GL_STACK_UNDERFLOW"
	case FRAGMENT_SHADER:
		name = "GL_FRAGMENT_SHADER"
	case COMPUTE_SHADER:
		name = "GL_COMPUTE_SHADER"
	case SAMPLER_2D_ARRAY:
		name = "GL_SAMPLER_2D_ARRAY"
	case REFERENCED_BY_GEOMETRY_SHADER:
		name = "GL_REFERENCED_BY_GEOMETRY_SHADER"
	case DISPLAY_LIST:
		name = "GL_DISPLAY_LIST"
	case UNIFORM_BUFFER_SIZE:
		name = "GL_UNIFORM_BUFFER_SIZE"
	case QUERY_RESULT_AVAILABLE:
		name = "GL_QUERY_RESULT_AVAILABLE"
	case COMPRESSED_RGB_S3TC_DXT1_EXT:
		name = "GL_COMPRESSED_RGB_S3TC_DXT1_EXT"
	case SMOOTH_LINE_WIDTH_RANGE:
		name = "GL_SMOOTH_LINE_WIDTH_RANGE"
	case RGB4:
		name = "GL_RGB4"
	case FRAMEBUFFER_DEFAULT_LAYERS:
		name = "GL_FRAMEBUFFER_DEFAULT_LAYERS"
	case IMAGE_CLASS_1_X_8:
		name = "GL_IMAGE_CLASS_1_X_8"
	case FASTEST:
		name = "GL_FASTEST"
	case FRAMEBUFFER_INCOMPLETE_ATTACHMENT:
		name = "GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT"
	case TESS_GEN_VERTEX_ORDER:
		name = "GL_TESS_GEN_VERTEX_ORDER"
	case INT_SAMPLER_BUFFER:
		name = "GL_INT_SAMPLER_BUFFER"
	case BOOL_VEC4:
		name = "GL_BOOL_VEC4"
	case FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER:
		name = "GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER"
	case TEXTURE_BINDING_BUFFER:
		name = "GL_TEXTURE_BINDING_BUFFER"
	case ACTIVE_PROGRAM:
		name = "GL_ACTIVE_PROGRAM"
	case QUERY_COUNTER_BITS:
		name = "GL_QUERY_COUNTER_BITS"
	case SMOOTH_POINT_SIZE_GRANULARITY:
		name = "GL_SMOOTH_POINT_SIZE_GRANULARITY"
	case MAX_NUM_ACTIVE_VARIABLES:
		name = "GL_MAX_NUM_ACTIVE_VARIABLES"
	case FRAMEBUFFER_DEFAULT_WIDTH:
		name = "GL_FRAMEBUFFER_DEFAULT_WIDTH"
	case TEXTURE_GATHER:
		name = "GL_TEXTURE_GATHER"
	case ALPHA:
		name = "GL_ALPHA"
	case R16UI:
		name = "GL_R16UI"
	case TEXTURE_IMMUTABLE_LEVELS:
		name = "GL_TEXTURE_IMMUTABLE_LEVELS"
	case FILTER:
		name = "GL_FILTER"
	case UNSIGNED_INT_2_10_10_10_REV:
		name = "GL_UNSIGNED_INT_2_10_10_10_REV"
	case TEXTURE_CUBE_MAP_NEGATIVE_X:
		name = "GL_TEXTURE_CUBE_MAP_NEGATIVE_X"
	case MAX_TESS_PATCH_COMPONENTS:
		name = "GL_MAX_TESS_PATCH_COMPONENTS"
	case SYNC_STATUS:
		name = "GL_SYNC_STATUS"
	case COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT:
		name = "GL_COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT"
	case REFERENCED_BY_FRAGMENT_SHADER:
		name = "GL_REFERENCED_BY_FRAGMENT_SHADER"
	case DRAW_BUFFER12:
		name = "GL_DRAW_BUFFER12"
	case WAIT_FAILED:
		name = "GL_WAIT_FAILED"
	case SHADER_STORAGE_BUFFER_SIZE:
		name = "GL_SHADER_STORAGE_BUFFER_SIZE"
	case ATOMIC_COUNTER_BUFFER:
		name = "GL_ATOMIC_COUNTER_BUFFER"
	case TEXTURE2:
		name = "GL_TEXTURE2"
	case VIEWPORT_INDEX_PROVOKING_VERTEX:
		name = "GL_VIEWPORT_INDEX_PROVOKING_VERTEX"
	case NICEST:
		name = "GL_NICEST"
	case RGB8I:
		name = "GL_RGB8I"
	case DEBUG_TYPE_MARKER:
		name = "GL_DEBUG_TYPE_MARKER"
	case PROXY_TEXTURE_CUBE_MAP:
		name = "GL_PROXY_TEXTURE_CUBE_MAP"
	case ONE_MINUS_SRC_COLOR:
		name = "GL_ONE_MINUS_SRC_COLOR"
	case RGB_SNORM:
		name = "GL_RGB_SNORM"
	case POINT_SIZE:
		name = "GL_POINT_SIZE"
	case TEXTURE_RED_TYPE:
		name = "GL_TEXTURE_RED_TYPE"
	case TEXTURE5:
		name = "GL_TEXTURE5"
	case MAX_COMPUTE_UNIFORM_COMPONENTS:
		name = "GL_MAX_COMPUTE_UNIFORM_COMPONENTS"
	case FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE"
	case MAX_UNIFORM_LOCATIONS:
		name = "GL_MAX_UNIFORM_LOCATIONS"
	case CONTEXT_FLAGS:
		name = "GL_CONTEXT_FLAGS"
	case RG_SNORM:
		name = "GL_RG_SNORM"
	case TEXTURE7:
		name = "GL_TEXTURE7"
	case TEXTURE_VIEW_MIN_LAYER:
		name = "GL_TEXTURE_VIEW_MIN_LAYER"
	case RGB16I:
		name = "GL_RGB16I"
	case QUADS:
		name = "GL_QUADS"
	case MINOR_VERSION:
		name = "GL_MINOR_VERSION"
	case PRIMITIVES_GENERATED:
		name = "GL_PRIMITIVES_GENERATED"
	case STREAM_COPY:
		name = "GL_STREAM_COPY"
	case RGBA8:
		name = "GL_RGBA8"
	case TRANSFORM_FEEDBACK_BINDING:
		name = "GL_TRANSFORM_FEEDBACK_BINDING"
	case PROGRAM_BINARY_RETRIEVABLE_HINT:
		name = "GL_PROGRAM_BINARY_RETRIEVABLE_HINT"
	case DEBUG_CALLBACK_FUNCTION:
		name = "GL_DEBUG_CALLBACK_FUNCTION"
	case MAJOR_VERSION:
		name = "GL_MAJOR_VERSION"
	case INT_IMAGE_1D:
		name = "GL_INT_IMAGE_1D"
	case TEXTURE15:
		name = "GL_TEXTURE15"
	case COMPRESSED_TEXTURE_FORMATS:
		name = "GL_COMPRESSED_TEXTURE_FORMATS"
	case PACK_ALIGNMENT:
		name = "GL_PACK_ALIGNMENT"
	case RENDERBUFFER_SAMPLES:
		name = "GL_RENDERBUFFER_SAMPLES"
	case STENCIL_REF:
		name = "GL_STENCIL_REF"
	case FRAMEBUFFER_UNDEFINED:
		name = "GL_FRAMEBUFFER_UNDEFINED"
	case SRC1_ALPHA:
		name = "GL_SRC1_ALPHA"
	case LOGIC_OP_MODE:
		name = "GL_LOGIC_OP_MODE"
	case LINE_STRIP_ADJACENCY:
		name = "GL_LINE_STRIP_ADJACENCY"
	case MAX_COMBINED_ATOMIC_COUNTER_BUFFERS:
		name = "GL_MAX_COMBINED_ATOMIC_COUNTER_BUFFERS"
	case DEPTH_TEST:
		name = "GL_DEPTH_TEST"
	case UNSIGNED_INT_10_10_10_2:
		name = "GL_UNSIGNED_INT_10_10_10_2"
	case BLEND_DST_RGB:
		name = "GL_BLEND_DST_RGB"
	case DISPATCH_INDIRECT_BUFFER_BINDING:
		name = "GL_DISPATCH_INDIRECT_BUFFER_BINDING"
	case XOR:
		name = "GL_XOR"
	case VERTEX_ATTRIB_ARRAY_STRIDE:
		name = "GL_VERTEX_ATTRIB_ARRAY_STRIDE"
	case NUM_COMPATIBLE_SUBROUTINES:
		name = "GL_NUM_COMPATIBLE_SUBROUTINES"
	case COMMAND_BARRIER_BIT:
		name = "GL_COMMAND_BARRIER_BIT"
	case INTERNALFORMAT_DEPTH_TYPE:
		name = "GL_INTERNALFORMAT_DEPTH_TYPE"
	case MAX_UNIFORM_BUFFER_BINDINGS:
		name = "GL_MAX_UNIFORM_BUFFER_BINDINGS"
	case CCW:
		name = "GL_CCW"
	case CURRENT_QUERY:
		name = "GL_CURRENT_QUERY"
	case TEXTURE_GATHER_SHADOW:
		name = "GL_TEXTURE_GATHER_SHADOW"
	case SAMPLER_CUBE_SHADOW:
		name = "GL_SAMPLER_CUBE_SHADOW"
	case TRANSFORM_FEEDBACK_BUFFER_ACTIVE:
		name = "GL_TRANSFORM_FEEDBACK_BUFFER_ACTIVE"
	case CULL_FACE:
		name = "GL_CULL_FACE"
	case INT_VEC4:
		name = "GL_INT_VEC4"
	case MAX_GEOMETRY_UNIFORM_COMPONENTS:
		name = "GL_MAX_GEOMETRY_UNIFORM_COMPONENTS"
	case SAMPLES_PASSED:
		name = "GL_SAMPLES_PASSED"
	case VIEW_CLASS_BPTC_FLOAT:
		name = "GL_VIEW_CLASS_BPTC_FLOAT"
	case IMAGE_CLASS_1_X_32:
		name = "GL_IMAGE_CLASS_1_X_32"
	case ATOMIC_COUNTER_BARRIER_BIT:
		name = "GL_ATOMIC_COUNTER_BARRIER_BIT"
	case COLOR_ATTACHMENT14:
		name = "GL_COLOR_ATTACHMENT14"
	case INT:
		name = "GL_INT"
	case FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE"
	case BUFFER_MAP_OFFSET:
		name = "GL_BUFFER_MAP_OFFSET"
	case UNSIGNED_BYTE_3_3_2:
		name = "GL_UNSIGNED_BYTE_3_3_2"
	case STENCIL_BACK_FAIL:
		name = "GL_STENCIL_BACK_FAIL"
	case UNSIGNED_INT_IMAGE_2D_MULTISAMPLE:
		name = "GL_UNSIGNED_INT_IMAGE_2D_MULTISAMPLE"
	case TEXTURE18:
		name = "GL_TEXTURE18"
	case RGBA8I:
		name = "GL_RGBA8I"
	case INT_IMAGE_CUBE:
		name = "GL_INT_IMAGE_CUBE"
	case CULL_FACE_MODE:
		name = "GL_CULL_FACE_MODE"
	case FLOAT_MAT3:
		name = "GL_FLOAT_MAT3"
	case TEXTURE_BUFFER_SIZE:
		name = "GL_TEXTURE_BUFFER_SIZE"
	case IMAGE_2D_MULTISAMPLE_ARRAY:
		name = "GL_IMAGE_2D_MULTISAMPLE_ARRAY"
	case SYNC_CONDITION:
		name = "GL_SYNC_CONDITION"
	case BUFFER_DATA_SIZE:
		name = "GL_BUFFER_DATA_SIZE"
	case EXTENSIONS:
		name = "GL_EXTENSIONS"
	case SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT:
		name = "GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT"
	case RGBA32I:
		name = "GL_RGBA32I"
	case TYPE:
		name = "GL_TYPE"
	case RASTERIZER_DISCARD:
		name = "GL_RASTERIZER_DISCARD"
	case TEXTURE_VIEW_MIN_LEVEL:
		name = "GL_TEXTURE_VIEW_MIN_LEVEL"
	case DRAW_BUFFER7:
		name = "GL_DRAW_BUFFER7"
	case COMPRESSED_RGB8_ETC2:
		name = "GL_COMPRESSED_RGB8_ETC2"
	case FLOAT_32_UNSIGNED_INT_24_8_REV:
		name = "GL_FLOAT_32_UNSIGNED_INT_24_8_REV"
	case STENCIL_FUNC:
		name = "GL_STENCIL_FUNC"
	case VIEW_CLASS_S3TC_DXT3_RGBA:
		name = "GL_VIEW_CLASS_S3TC_DXT3_RGBA"
	case MAP_WRITE_BIT:
		name = "GL_MAP_WRITE_BIT"
	case ACTIVE_SUBROUTINE_UNIFORMS:
		name = "GL_ACTIVE_SUBROUTINE_UNIFORMS"
	case COLOR:
		name = "GL_COLOR"
	case BUFFER_MAP_POINTER:
		name = "GL_BUFFER_MAP_POINTER"
	case COMPUTE_SUBROUTINE:
		name = "GL_COMPUTE_SUBROUTINE"
	case RGB32I:
		name = "GL_RGB32I"
	case TRANSFORM_FEEDBACK_BUFFER_MODE:
		name = "GL_TRANSFORM_FEEDBACK_BUFFER_MODE"
	case DEBUG_SEVERITY_LOW:
		name = "GL_DEBUG_SEVERITY_LOW"
	case MAX_SHADER_STORAGE_BLOCK_SIZE:
		name = "GL_MAX_SHADER_STORAGE_BLOCK_SIZE"
	case MAX_PROGRAM_TEXTURE_GATHER_OFFSET:
		name = "GL_MAX_PROGRAM_TEXTURE_GATHER_OFFSET"
	case GET_TEXTURE_IMAGE_TYPE:
		name = "GL_GET_TEXTURE_IMAGE_TYPE"
	case SAMPLE_COVERAGE_VALUE:
		name = "GL_SAMPLE_COVERAGE_VALUE"
	case RGB10_A2:
		name = "GL_RGB10_A2"
	case VIEW_CLASS_128_BITS:
		name = "GL_VIEW_CLASS_128_BITS"
	case MAX_TESS_EVALUATION_UNIFORM_BLOCKS:
		name = "GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS"
	case RGBA8UI:
		name = "GL_RGBA8UI"
	case SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE:
		name = "GL_SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE"
	case IMAGE_CLASS_4_X_32:
		name = "GL_IMAGE_CLASS_4_X_32"
	case TEXTURE_MAX_LEVEL:
		name = "GL_TEXTURE_MAX_LEVEL"
	case COLOR_ATTACHMENT8:
		name = "GL_COLOR_ATTACHMENT8"
	case TEXTURE_1D_ARRAY:
		name = "GL_TEXTURE_1D_ARRAY"
	case MAX_TESS_EVALUATION_ATOMIC_COUNTERS:
		name = "GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS"
	case READ_WRITE:
		name = "GL_READ_WRITE"
	case MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS:
		name = "GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS"
	case UNSIGNED_SHORT_1_5_5_5_REV:
		name = "GL_UNSIGNED_SHORT_1_5_5_5_REV"
	case TRANSFORM_FEEDBACK_PAUSED:
		name = "GL_TRANSFORM_FEEDBACK_PAUSED"
	case VIEW_CLASS_96_BITS:
		name = "GL_VIEW_CLASS_96_BITS"
	case REPLACE:
		name = "GL_REPLACE"
	case TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:
		name = "GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN"
	case MAX_COMPUTE_WORK_GROUP_SIZE:
		name = "GL_MAX_COMPUTE_WORK_GROUP_SIZE"
	case TEXTURE_ALPHA_TYPE:
		name = "GL_TEXTURE_ALPHA_TYPE"
	case TEXTURE_COMPARE_FUNC:
		name = "GL_TEXTURE_COMPARE_FUNC"
	case MAX_DEPTH_TEXTURE_SAMPLES:
		name = "GL_MAX_DEPTH_TEXTURE_SAMPLES"
	case PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY:
		name = "GL_PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY"
	case COLOR_RENDERABLE:
		name = "GL_COLOR_RENDERABLE"
	case FRONT_LEFT:
		name = "GL_FRONT_LEFT"
	case STREAM_DRAW:
		name = "GL_STREAM_DRAW"
	case MAX_VARYING_COMPONENTS:
		name = "GL_MAX_VARYING_COMPONENTS"
	case MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS:
		name = "GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS"
	case TEXTURE20:
		name = "GL_TEXTURE20"
	case VIEW_CLASS_8_BITS:
		name = "GL_VIEW_CLASS_8_BITS"
	case BUFFER_MAPPED:
		name = "GL_BUFFER_MAPPED"
	case ARRAY_SIZE:
		name = "GL_ARRAY_SIZE"
	case RG8UI:
		name = "GL_RG8UI"
	case TEXTURE_GREEN_TYPE:
		name = "GL_TEXTURE_GREEN_TYPE"
	case FRAMEBUFFER_RENDERABLE_LAYERED:
		name = "GL_FRAMEBUFFER_RENDERABLE_LAYERED"
	case MAX_TESS_CONTROL_ATOMIC_COUNTERS:
		name = "GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS"
	case UNSIGNED_INT_SAMPLER_2D_RECT:
		name = "GL_UNSIGNED_INT_SAMPLER_2D_RECT"
	case MAX_IMAGE_UNITS:
		name = "GL_MAX_IMAGE_UNITS"
	case TRUE:
		name = "GL_TRUE"
	case GREEN_INTEGER:
		name = "GL_GREEN_INTEGER"
	case QUERY_BY_REGION_NO_WAIT:
		name = "GL_QUERY_BY_REGION_NO_WAIT"
	case INCR:
		name = "GL_INCR"
	case COMPARE_REF_TO_TEXTURE:
		name = "GL_COMPARE_REF_TO_TEXTURE"
	case COLOR_ENCODING:
		name = "GL_COLOR_ENCODING"
	case SHADER:
		name = "GL_SHADER"
	case REPEAT:
		name = "GL_REPEAT"
	case TEXTURE26:
		name = "GL_TEXTURE26"
	case SEPARATE_ATTRIBS:
		name = "GL_SEPARATE_ATTRIBS"
	case CLIP_DISTANCE2:
		name = "GL_CLIP_DISTANCE2"
	case OR_INVERTED:
		name = "GL_OR_INVERTED"
	case SRGB_EXT:
		name = "GL_SRGB_EXT"
	case NAND:
		name = "GL_NAND"
	case VIEWPORT:
		name = "GL_VIEWPORT"
	case RGB12:
		name = "GL_RGB12"
	case TEXTURE_VIEW_NUM_LEVELS:
		name = "GL_TEXTURE_VIEW_NUM_LEVELS"
	case MAX_VERTEX_STREAMS:
		name = "GL_MAX_VERTEX_STREAMS"
	case UNIFORM_OFFSET:
		name = "GL_UNIFORM_OFFSET"
	case RENDERBUFFER_ALPHA_SIZE:
		name = "GL_RENDERBUFFER_ALPHA_SIZE"
	case ANY_SAMPLES_PASSED_CONSERVATIVE:
		name = "GL_ANY_SAMPLES_PASSED_CONSERVATIVE"
	case TEXTURE_HEIGHT:
		name = "GL_TEXTURE_HEIGHT"
	case TEXTURE21:
		name = "GL_TEXTURE21"
	case MAX_TESS_EVALUATION_OUTPUT_COMPONENTS:
		name = "GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS"
	case TEXTURE_COMPRESSED_BLOCK_WIDTH:
		name = "GL_TEXTURE_COMPRESSED_BLOCK_WIDTH"
	case UNIFORM_BUFFER:
		name = "GL_UNIFORM_BUFFER"
	case PROXY_TEXTURE_CUBE_MAP_ARRAY:
		name = "GL_PROXY_TEXTURE_CUBE_MAP_ARRAY"
	case SAMPLER:
		name = "GL_SAMPLER"
	case VERTEX_ATTRIB_RELATIVE_OFFSET:
		name = "GL_VERTEX_ATTRIB_RELATIVE_OFFSET"
	case TEXTURE_BUFFER_OFFSET:
		name = "GL_TEXTURE_BUFFER_OFFSET"
	case MAX_FRAMEBUFFER_WIDTH:
		name = "GL_MAX_FRAMEBUFFER_WIDTH"
	case FRAMEBUFFER_ATTACHMENT_LAYERED:
		name = "GL_FRAMEBUFFER_ATTACHMENT_LAYERED"
	case IMAGE_TEXEL_SIZE:
		name = "GL_IMAGE_TEXEL_SIZE"
	case TEXTURE_SWIZZLE_RGBA:
		name = "GL_TEXTURE_SWIZZLE_RGBA"
	case IMAGE_BINDING_FORMAT:
		name = "GL_IMAGE_BINDING_FORMAT"
	case BUFFER_ACCESS_FLAGS:
		name = "GL_BUFFER_ACCESS_FLAGS"
	case STENCIL_PASS_DEPTH_PASS:
		name = "GL_STENCIL_PASS_DEPTH_PASS"
	case COPY_READ_BUFFER:
		name = "GL_COPY_READ_BUFFER"
	case VIEWPORT_SUBPIXEL_BITS:
		name = "GL_VIEWPORT_SUBPIXEL_BITS"
	case TESS_GEN_MODE:
		name = "GL_TESS_GEN_MODE"
	case TEXTURE27:
		name = "GL_TEXTURE27"
	case TESS_CONTROL_OUTPUT_VERTICES:
		name = "GL_TESS_CONTROL_OUTPUT_VERTICES"
	case TEXTURE_SWIZZLE_A:
		name = "GL_TEXTURE_SWIZZLE_A"
	case PIXEL_BUFFER_BARRIER_BIT:
		name = "GL_PIXEL_BUFFER_BARRIER_BIT"
	case FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING:
		name = "GL_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING"
	case TRANSFORM_FEEDBACK:
		name = "GL_TRANSFORM_FEEDBACK"
	case DECR:
		name = "GL_DECR"
	case RGBA2:
		name = "GL_RGBA2"
	case COMPRESSED_SRGB_S3TC_DXT1_EXT:
		name = "GL_COMPRESSED_SRGB_S3TC_DXT1_EXT"
	case FULL_SUPPORT:
		name = "GL_FULL_SUPPORT"
	case TRANSFORM_FEEDBACK_BUFFER_BINDING:
		name = "GL_TRANSFORM_FEEDBACK_BUFFER_BINDING"
	case SRC_ALPHA:
		name = "GL_SRC_ALPHA"
	case SAMPLE_POSITION:
		name = "GL_SAMPLE_POSITION"
	case SAMPLER_1D:
		name = "GL_SAMPLER_1D"
	case COMPRESSED_SRGB8_ETC2:
		name = "GL_COMPRESSED_SRGB8_ETC2"
	case TEXTURE_BLUE_SIZE:
		name = "GL_TEXTURE_BLUE_SIZE"
	case VIEW_CLASS_16_BITS:
		name = "GL_VIEW_CLASS_16_BITS"
	case RG8I:
		name = "GL_RG8I"
	case FLOAT_VEC3:
		name = "GL_FLOAT_VEC3"
	case COLOR_ATTACHMENT4:
		name = "GL_COLOR_ATTACHMENT4"
	case CLEAR_BUFFER:
		name = "GL_CLEAR_BUFFER"
	case TEXTURE_1D:
		name = "GL_TEXTURE_1D"
	case FIXED:
		name = "GL_FIXED"
	case FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT:
		name = "GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT"
	case COMPRESSED_SRGB_ALPHA_EXT:
		name = "GL_COMPRESSED_SRGB_ALPHA_EXT"
	case BLEND_SRC_ALPHA:
		name = "GL_BLEND_SRC_ALPHA"
	case SIGNALED:
		name = "GL_SIGNALED"
	case DEBUG_LOGGED_MESSAGES:
		name = "GL_DEBUG_LOGGED_MESSAGES"
	case REFERENCED_BY_TESS_CONTROL_SHADER:
		name = "GL_REFERENCED_BY_TESS_CONTROL_SHADER"
	case RGB32UI:
		name = "GL_RGB32UI"
	case VALIDATE_STATUS:
		name = "GL_VALIDATE_STATUS"
	case IMAGE_PIXEL_TYPE:
		name = "GL_IMAGE_PIXEL_TYPE"
	case SRC_ALPHA_SATURATE:
		name = "GL_SRC_ALPHA_SATURATE"
	case SAMPLE_ALPHA_TO_ONE:
		name = "GL_SAMPLE_ALPHA_TO_ONE"
	case DEBUG_SOURCE_APPLICATION:
		name = "GL_DEBUG_SOURCE_APPLICATION"
	case RENDERER:
		name = "GL_RENDERER"
	case ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS:
		name = "GL_ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS"
	case VERTEX_ATTRIB_ARRAY_BUFFER_BINDING:
		name = "GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING"
	case UPPER_LEFT:
		name = "GL_UPPER_LEFT"
	case SHADER_STORAGE_BLOCK:
		name = "GL_SHADER_STORAGE_BLOCK"
	case FRACTIONAL_ODD:
		name = "GL_FRACTIONAL_ODD"
	case DOUBLE_MAT4x3:
		name = "GL_DOUBLE_MAT4x3"
	case VIEW_CLASS_BPTC_UNORM:
		name = "GL_VIEW_CLASS_BPTC_UNORM"
	case SAMPLER_2D_MULTISAMPLE:
		name = "GL_SAMPLER_2D_MULTISAMPLE"
	case ACTIVE_UNIFORM_BLOCKS:
		name = "GL_ACTIVE_UNIFORM_BLOCKS"
	case UNSIGNED_INT_ATOMIC_COUNTER:
		name = "GL_UNSIGNED_INT_ATOMIC_COUNTER"
	case ONE_MINUS_SRC1_COLOR:
		name = "GL_ONE_MINUS_SRC1_COLOR"
	case COLOR_ATTACHMENT2:
		name = "GL_COLOR_ATTACHMENT2"
	case TEXTURE_BINDING_2D_ARRAY:
		name = "GL_TEXTURE_BINDING_2D_ARRAY"
	case UNPACK_SWAP_BYTES:
		name = "GL_UNPACK_SWAP_BYTES"
	case UNSIGNED_INT_SAMPLER_1D:
		name = "GL_UNSIGNED_INT_SAMPLER_1D"
	case MAX_COMPUTE_UNIFORM_BLOCKS:
		name = "GL_MAX_COMPUTE_UNIFORM_BLOCKS"
	case UNPACK_SKIP_IMAGES:
		name = "GL_UNPACK_SKIP_IMAGES"
	case UNSIGNED_SHORT:
		name = "GL_UNSIGNED_SHORT"
	case MAX_COMBINED_ATOMIC_COUNTERS:
		name = "GL_MAX_COMBINED_ATOMIC_COUNTERS"
	case DEBUG_OUTPUT:
		name = "GL_DEBUG_OUTPUT"
	case ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER:
		name = "GL_ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER"
	case FRAMEBUFFER_COMPLETE:
		name = "GL_FRAMEBUFFER_COMPLETE"
	case TEXTURE_CUBE_MAP_NEGATIVE_Y:
		name = "GL_TEXTURE_CUBE_MAP_NEGATIVE_Y"
	case TEXTURE_2D:
		name = "GL_TEXTURE_2D"
	case DEPTH_WRITEMASK:
		name = "GL_DEPTH_WRITEMASK"
	case DRAW_INDIRECT_BUFFER:
		name = "GL_DRAW_INDIRECT_BUFFER"
	case SRGB8_EXT:
		name = "GL_SRGB8_EXT"
	case IMPLEMENTATION_COLOR_READ_TYPE:
		name = "GL_IMPLEMENTATION_COLOR_READ_TYPE"
	case MAX_TEXTURE_SIZE:
		name = "GL_MAX_TEXTURE_SIZE"
	case STENCIL_COMPONENTS:
		name = "GL_STENCIL_COMPONENTS"
	case R16_SNORM:
		name = "GL_R16_SNORM"
	case MAX_COMBINED_IMAGE_UNIFORMS:
		name = "GL_MAX_COMBINED_IMAGE_UNIFORMS"
	case TEXTURE_COMPARE_MODE:
		name = "GL_TEXTURE_COMPARE_MODE"
	case UNSIGNED_BYTE_2_3_3_REV:
		name = "GL_UNSIGNED_BYTE_2_3_3_REV"
	case FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL:
		name = "GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL"
	case COLOR_ATTACHMENT7:
		name = "GL_COLOR_ATTACHMENT7"
	case EQUAL:
		name = "GL_EQUAL"
	case IMAGE_1D:
		name = "GL_IMAGE_1D"
	case TEXTURE_MAX_ANISOTROPY_EXT:
		name = "GL_TEXTURE_MAX_ANISOTROPY_EXT"
	case INT_IMAGE_2D_MULTISAMPLE_ARRAY:
		name = "GL_INT_IMAGE_2D_MULTISAMPLE_ARRAY"
	case RGBA_INTEGER:
		name = "GL_RGBA_INTEGER"
	case CLAMP_TO_BORDER:
		name = "GL_CLAMP_TO_BORDER"
	case WRITE_ONLY:
		name = "GL_WRITE_ONLY"
	case MAX_LAYERS:
		name = "GL_MAX_LAYERS"
	case TEXTURE_GREEN_SIZE:
		name = "GL_TEXTURE_GREEN_SIZE"
	case STENCIL_BACK_PASS_DEPTH_PASS:
		name = "GL_STENCIL_BACK_PASS_DEPTH_PASS"
	case TEXTURE_BASE_LEVEL:
		name = "GL_TEXTURE_BASE_LEVEL"
	case UNSIGNED_INT_24_8:
		name = "GL_UNSIGNED_INT_24_8"
	case TEXTURE6:
		name = "GL_TEXTURE6"
	case GREEN:
		name = "GL_GREEN"
	case SLUMINANCE8_EXT:
		name = "GL_SLUMINANCE8_EXT"
	case TRIANGLE_FAN:
		name = "GL_TRIANGLE_FAN"
	case STENCIL_INDEX1:
		name = "GL_STENCIL_INDEX1"
	case TEXTURE_CUBE_MAP:
		name = "GL_TEXTURE_CUBE_MAP"
	case MAX_GEOMETRY_ATOMIC_COUNTERS:
		name = "GL_MAX_GEOMETRY_ATOMIC_COUNTERS"
	case MAX_FRAGMENT_INTERPOLATION_OFFSET:
		name = "GL_MAX_FRAGMENT_INTERPOLATION_OFFSET"
	case DEBUG_OUTPUT_SYNCHRONOUS:
		name = "GL_DEBUG_OUTPUT_SYNCHRONOUS"
	case STENCIL_PASS_DEPTH_FAIL:
		name = "GL_STENCIL_PASS_DEPTH_FAIL"
	case INT_SAMPLER_2D_MULTISAMPLE_ARRAY:
		name = "GL_INT_SAMPLER_2D_MULTISAMPLE_ARRAY"
	case FRAMEBUFFER_BLEND:
		name = "GL_FRAMEBUFFER_BLEND"
	case BLEND_EQUATION_RGB:
		name = "GL_BLEND_EQUATION_RGB"
	case SAMPLER_CUBE:
		name = "GL_SAMPLER_CUBE"
	case AND_INVERTED:
		name = "GL_AND_INVERTED"
	case TEXTURE3:
		name = "GL_TEXTURE3"
	case INTERNALFORMAT_DEPTH_SIZE:
		name = "GL_INTERNALFORMAT_DEPTH_SIZE"
	case SRGB_READ:
		name = "GL_SRGB_READ"
	case POINTS:
		name = "GL_POINTS"
	case COMPRESSED_RG_RGTC2:
		name = "GL_COMPRESSED_RG_RGTC2"
	case GEOMETRY_SHADER_BIT:
		name = "GL_GEOMETRY_SHADER_BIT"
	case VERTEX_PROGRAM_POINT_SIZE:
		name = "GL_VERTEX_PROGRAM_POINT_SIZE"
	case MAX_VERTEX_TEXTURE_IMAGE_UNITS:
		name = "GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS"
	case UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER:
		name = "GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER"
	case ALREADY_SIGNALED:
		name = "GL_ALREADY_SIGNALED"
	case MAX_TESS_CONTROL_INPUT_COMPONENTS:
		name = "GL_MAX_TESS_CONTROL_INPUT_COMPONENTS"
	case TEXTURE11:
		name = "GL_TEXTURE11"
	case SHADER_IMAGE_ACCESS_BARRIER_BIT:
		name = "GL_SHADER_IMAGE_ACCESS_BARRIER_BIT"
	case MAX_FRAMEBUFFER_SAMPLES:
		name = "GL_MAX_FRAMEBUFFER_SAMPLES"
	case TIMESTAMP:
		name = "GL_TIMESTAMP"
	case TEXTURE_VIEW_NUM_LAYERS:
		name = "GL_TEXTURE_VIEW_NUM_LAYERS"
	case SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE:
		name = "GL_SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE"
	case SLUMINANCE8_ALPHA8_EXT:
		name = "GL_SLUMINANCE8_ALPHA8_EXT"
	case MAX_VARYING_VECTORS:
		name = "GL_MAX_VARYING_VECTORS"
	case UNSIGNALED:
		name = "GL_UNSIGNALED"
	case FLOAT_VEC2:
		name = "GL_FLOAT_VEC2"
	case TRANSFORM_FEEDBACK_VARYING:
		name = "GL_TRANSFORM_FEEDBACK_VARYING"
	case BOOL_VEC3:
		name = "GL_BOOL_VEC3"
	case TEXTURE_FIXED_SAMPLE_LOCATIONS:
		name = "GL_TEXTURE_FIXED_SAMPLE_LOCATIONS"
	case SAMPLE_ALPHA_TO_COVERAGE:
		name = "GL_SAMPLE_ALPHA_TO_COVERAGE"
	case MAX_DUAL_SOURCE_DRAW_BUFFERS:
		name = "GL_MAX_DUAL_SOURCE_DRAW_BUFFERS"
	case MEDIUM_FLOAT:
		name = "GL_MEDIUM_FLOAT"
	case RENDERBUFFER_STENCIL_SIZE:
		name = "GL_RENDERBUFFER_STENCIL_SIZE"
	case FRAMEBUFFER_ATTACHMENT_BLUE_SIZE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_BLUE_SIZE"
	case VERTEX_SUBROUTINE_UNIFORM:
		name = "GL_VERTEX_SUBROUTINE_UNIFORM"
	case COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2:
		name = "GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2"
	case DECR_WRAP:
		name = "GL_DECR_WRAP"
	case VIEW_CLASS_S3TC_DXT1_RGBA:
		name = "GL_VIEW_CLASS_S3TC_DXT1_RGBA"
	case BOOL:
		name = "GL_BOOL"
	case UNSIGNED_INT_SAMPLER_BUFFER:
		name = "GL_UNSIGNED_INT_SAMPLER_BUFFER"
	case VIEW_CLASS_RGTC2_RG:
		name = "GL_VIEW_CLASS_RGTC2_RG"
	case LOW_INT:
		name = "GL_LOW_INT"
	case REFERENCED_BY_VERTEX_SHADER:
		name = "GL_REFERENCED_BY_VERTEX_SHADER"
	case ELEMENT_ARRAY_BUFFER_BINDING:
		name = "GL_ELEMENT_ARRAY_BUFFER_BINDING"
	case POLYGON_SMOOTH:
		name = "GL_POLYGON_SMOOTH"
	case VERTEX_ATTRIB_BINDING:
		name = "GL_VERTEX_ATTRIB_BINDING"
	case COMPRESSED_SIGNED_RED_RGTC1:
		name = "GL_COMPRESSED_SIGNED_RED_RGTC1"
	case UNIFORM_BLOCK:
		name = "GL_UNIFORM_BLOCK"
	case FRAGMENT_TEXTURE:
		name = "GL_FRAGMENT_TEXTURE"
	case RGBA16:
		name = "GL_RGBA16"
	case LEFT:
		name = "GL_LEFT"
	case IMAGE_FORMAT_COMPATIBILITY_BY_SIZE:
		name = "GL_IMAGE_FORMAT_COMPATIBILITY_BY_SIZE"
	case BLUE:
		name = "GL_BLUE"
	case FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE"
	case POLYGON_OFFSET_POINT:
		name = "GL_POLYGON_OFFSET_POINT"
	case DOUBLE_MAT2:
		name = "GL_DOUBLE_MAT2"
	case MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS:
		name = "GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS"
	case ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS:
		name = "GL_ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS"
	case CLIP_DISTANCE0:
		name = "GL_CLIP_DISTANCE0"
	case COPY_WRITE_BUFFER:
		name = "GL_COPY_WRITE_BUFFER"
	case RGB10:
		name = "GL_RGB10"
	case IMAGE_2D:
		name = "GL_IMAGE_2D"
	case SAMPLER_2D_RECT_SHADOW:
		name = "GL_SAMPLER_2D_RECT_SHADOW"
	case TEXTURE_BLUE_TYPE:
		name = "GL_TEXTURE_BLUE_TYPE"
	case INT_IMAGE_CUBE_MAP_ARRAY:
		name = "GL_INT_IMAGE_CUBE_MAP_ARRAY"
	case SAMPLE_BUFFERS:
		name = "GL_SAMPLE_BUFFERS"
	case MAX_COMPUTE_WORK_GROUP_COUNT:
		name = "GL_MAX_COMPUTE_WORK_GROUP_COUNT"
	case POLYGON_OFFSET_LINE:
		name = "GL_POLYGON_OFFSET_LINE"
	case FRAMEBUFFER_DEFAULT:
		name = "GL_FRAMEBUFFER_DEFAULT"
	case MAX_ARRAY_TEXTURE_LAYERS:
		name = "GL_MAX_ARRAY_TEXTURE_LAYERS"
	case PACK_COMPRESSED_BLOCK_WIDTH:
		name = "GL_PACK_COMPRESSED_BLOCK_WIDTH"
	case R8UI:
		name = "GL_R8UI"
	case INTERNALFORMAT_ALPHA_TYPE:
		name = "GL_INTERNALFORMAT_ALPHA_TYPE"
	case POLYGON_OFFSET_FILL:
		name = "GL_POLYGON_OFFSET_FILL"
	case TESS_GEN_SPACING:
		name = "GL_TESS_GEN_SPACING"
	case FRONT_AND_BACK:
		name = "GL_FRONT_AND_BACK"
	case COMPUTE_SUBROUTINE_UNIFORM:
		name = "GL_COMPUTE_SUBROUTINE_UNIFORM"
	case UNSIGNED_BYTE:
		name = "GL_UNSIGNED_BYTE"
	case TEXTURE_MIN_FILTER:
		name = "GL_TEXTURE_MIN_FILTER"
	case MAX_VERTEX_SHADER_STORAGE_BLOCKS:
		name = "GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS"
	case ACTIVE_UNIFORMS:
		name = "GL_ACTIVE_UNIFORMS"
	case FRAMEBUFFER:
		name = "GL_FRAMEBUFFER"
	case QUERY:
		name = "GL_QUERY"
	case SRGB_WRITE:
		name = "GL_SRGB_WRITE"
	case INTERNALFORMAT_BLUE_SIZE:
		name = "GL_INTERNALFORMAT_BLUE_SIZE"
	case UNSIGNED_INT_IMAGE_2D_ARRAY:
		name = "GL_UNSIGNED_INT_IMAGE_2D_ARRAY"
	case TEXTURE_STENCIL_SIZE:
		name = "GL_TEXTURE_STENCIL_SIZE"
	case MAX_NUM_COMPATIBLE_SUBROUTINES:
		name = "GL_MAX_NUM_COMPATIBLE_SUBROUTINES"
	case SRGB_ALPHA:
		name = "GL_SRGB_ALPHA"
	case MAP_INVALIDATE_BUFFER_BIT:
		name = "GL_MAP_INVALIDATE_BUFFER_BIT"
	case MAX_VERTEX_ATOMIC_COUNTERS:
		name = "GL_MAX_VERTEX_ATOMIC_COUNTERS"
	case DOUBLE_VEC2:
		name = "GL_DOUBLE_VEC2"
	case DRAW_BUFFER9:
		name = "GL_DRAW_BUFFER9"
	case UNIFORM_BLOCK_ACTIVE_UNIFORMS:
		name = "GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS"
	case MIN_MAP_BUFFER_ALIGNMENT:
		name = "GL_MIN_MAP_BUFFER_ALIGNMENT"
	case INTERNALFORMAT_SHARED_SIZE:
		name = "GL_INTERNALFORMAT_SHARED_SIZE"
	case DEBUG_CALLBACK_USER_PARAM:
		name = "GL_DEBUG_CALLBACK_USER_PARAM"
	case TEXTURE24:
		name = "GL_TEXTURE24"
	case COMPRESSED_RG11_EAC:
		name = "GL_COMPRESSED_RG11_EAC"
	case SAMPLE_MASK:
		name = "GL_SAMPLE_MASK"
	case DRAW_BUFFER11:
		name = "GL_DRAW_BUFFER11"
	case MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS:
		name = "GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS"
	case IMAGE_BUFFER:
		name = "GL_IMAGE_BUFFER"
	case MAX_DEBUG_GROUP_STACK_DEPTH:
		name = "GL_MAX_DEBUG_GROUP_STACK_DEPTH"
	case LOCATION_INDEX:
		name = "GL_LOCATION_INDEX"
	case ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH:
		name = "GL_ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH"
	case COLOR_ATTACHMENT6:
		name = "GL_COLOR_ATTACHMENT6"
	case MAX_TESS_CONTROL_IMAGE_UNIFORMS:
		name = "GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS"
	case PACK_SKIP_IMAGES:
		name = "GL_PACK_SKIP_IMAGES"
	case MAX_COMBINED_DIMENSIONS:
		name = "GL_MAX_COMBINED_DIMENSIONS"
	case UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES:
		name = "GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES"
	case EQUIV:
		name = "GL_EQUIV"
	case TESS_EVALUATION_SHADER_BIT:
		name = "GL_TESS_EVALUATION_SHADER_BIT"
	case VERTEX_ATTRIB_ARRAY_TYPE:
		name = "GL_VERTEX_ATTRIB_ARRAY_TYPE"
	case TEXTURE_CUBE_MAP_SEAMLESS:
		name = "GL_TEXTURE_CUBE_MAP_SEAMLESS"
	case STENCIL_FAIL:
		name = "GL_STENCIL_FAIL"
	case INTERNALFORMAT_PREFERRED:
		name = "GL_INTERNALFORMAT_PREFERRED"
	case DOUBLE_MAT3:
		name = "GL_DOUBLE_MAT3"
	case FRONT_RIGHT:
		name = "GL_FRONT_RIGHT"
	case INTERNALFORMAT_SUPPORTED:
		name = "GL_INTERNALFORMAT_SUPPORTED"
	case RGB10_A2UI:
		name = "GL_RGB10_A2UI"
	case ONE_MINUS_DST_ALPHA:
		name = "GL_ONE_MINUS_DST_ALPHA"
	case SAMPLER_3D:
		name = "GL_SAMPLER_3D"
	case TEXTURE_COMPRESSED_IMAGE_SIZE:
		name = "GL_TEXTURE_COMPRESSED_IMAGE_SIZE"
	case STENCIL_INDEX4:
		name = "GL_STENCIL_INDEX4"
	case PACK_COMPRESSED_BLOCK_DEPTH:
		name = "GL_PACK_COMPRESSED_BLOCK_DEPTH"
	case MAX_TESS_CONTROL_UNIFORM_BLOCKS:
		name = "GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS"
	case VERTEX_ATTRIB_ARRAY_SIZE:
		name = "GL_VERTEX_ATTRIB_ARRAY_SIZE"
	case POINT_SPRITE_COORD_ORIGIN:
		name = "GL_POINT_SPRITE_COORD_ORIGIN"
	case INT_IMAGE_BUFFER:
		name = "GL_INT_IMAGE_BUFFER"
	case COMPRESSED_RGBA_S3TC_DXT3_EXT:
		name = "GL_COMPRESSED_RGBA_S3TC_DXT3_EXT"
	case PROXY_TEXTURE_2D_MULTISAMPLE:
		name = "GL_PROXY_TEXTURE_2D_MULTISAMPLE"
	case OBJECT_TYPE:
		name = "GL_OBJECT_TYPE"
	case MAX_VERTEX_ATOMIC_COUNTER_BUFFERS:
		name = "GL_MAX_VERTEX_ATOMIC_COUNTER_BUFFERS"
	case ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES:
		name = "GL_ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES"
	case COLOR_ATTACHMENT3:
		name = "GL_COLOR_ATTACHMENT3"
	case RGB5:
		name = "GL_RGB5"
	case TESS_EVALUATION_SUBROUTINE_UNIFORM:
		name = "GL_TESS_EVALUATION_SUBROUTINE_UNIFORM"
	case TEXTURE_2D_MULTISAMPLE:
		name = "GL_TEXTURE_2D_MULTISAMPLE"
	case POINT_FADE_THRESHOLD_SIZE:
		name = "GL_POINT_FADE_THRESHOLD_SIZE"
	case MAX_HEIGHT:
		name = "GL_MAX_HEIGHT"
	case NUM_SHADER_BINARY_FORMATS:
		name = "GL_NUM_SHADER_BINARY_FORMATS"
	case MAX_CUBE_MAP_TEXTURE_SIZE:
		name = "GL_MAX_CUBE_MAP_TEXTURE_SIZE"
	case DEPTH_BUFFER_BIT:
		name = "GL_DEPTH_BUFFER_BIT"
	case FRAMEBUFFER_SRGB:
		name = "GL_FRAMEBUFFER_SRGB"
	case SAMPLE_COVERAGE:
		name = "GL_SAMPLE_COVERAGE"
	case MAX_COMBINED_SHADER_STORAGE_BLOCKS:
		name = "GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS"
	case TEXTURE_CUBE_MAP_POSITIVE_X:
		name = "GL_TEXTURE_CUBE_MAP_POSITIVE_X"
	case TEXTURE_WRAP_T:
		name = "GL_TEXTURE_WRAP_T"
	case OFFSET:
		name = "GL_OFFSET"
	case PROXY_TEXTURE_RECTANGLE:
		name = "GL_PROXY_TEXTURE_RECTANGLE"
	case RG_INTEGER:
		name = "GL_RG_INTEGER"
	case LINES_ADJACENCY:
		name = "GL_LINES_ADJACENCY"
	case ALL_BARRIER_BITS:
		name = "GL_ALL_BARRIER_BITS"
	case TEXTURE_BINDING_RECTANGLE:
		name = "GL_TEXTURE_BINDING_RECTANGLE"
	case NEAREST_MIPMAP_NEAREST:
		name = "GL_NEAREST_MIPMAP_NEAREST"
	case TEXTURE12:
		name = "GL_TEXTURE12"
	case FRAMEBUFFER_ATTACHMENT_RED_SIZE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_RED_SIZE"
	case UNSIGNED_INT_SAMPLER_1D_ARRAY:
		name = "GL_UNSIGNED_INT_SAMPLER_1D_ARRAY"
	case READ_FRAMEBUFFER_BINDING:
		name = "GL_READ_FRAMEBUFFER_BINDING"
	case PROGRAM:
		name = "GL_PROGRAM"
	case TEXTURE_INTERNAL_FORMAT:
		name = "GL_TEXTURE_INTERNAL_FORMAT"
	case STENCIL_BACK_PASS_DEPTH_FAIL:
		name = "GL_STENCIL_BACK_PASS_DEPTH_FAIL"
	case BUFFER_BINDING:
		name = "GL_BUFFER_BINDING"
	case MAX_ELEMENT_INDEX:
		name = "GL_MAX_ELEMENT_INDEX"
	case FRAGMENT_SHADER_DERIVATIVE_HINT:
		name = "GL_FRAGMENT_SHADER_DERIVATIVE_HINT"
	case DRAW_BUFFER14:
		name = "GL_DRAW_BUFFER14"
	case MAX_ELEMENTS_VERTICES:
		name = "GL_MAX_ELEMENTS_VERTICES"
	case VERTEX_ATTRIB_ARRAY_INTEGER:
		name = "GL_VERTEX_ATTRIB_ARRAY_INTEGER"
	case PROGRAM_INPUT:
		name = "GL_PROGRAM_INPUT"
	case R8_SNORM:
		name = "GL_R8_SNORM"
	case DRAW_INDIRECT_BUFFER_BINDING:
		name = "GL_DRAW_INDIRECT_BUFFER_BINDING"
	case MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS:
		name = "GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS"
	case INT_IMAGE_2D_RECT:
		name = "GL_INT_IMAGE_2D_RECT"
	case POINT:
		name = "GL_POINT"
	case LOW_FLOAT:
		name = "GL_LOW_FLOAT"
	case IMAGE_BINDING_LEVEL:
		name = "GL_IMAGE_BINDING_LEVEL"
	case RGB16_SNORM:
		name = "GL_RGB16_SNORM"
	case MAX_FRAGMENT_UNIFORM_BLOCKS:
		name = "GL_MAX_FRAGMENT_UNIFORM_BLOCKS"
	case MAX_DEBUG_MESSAGE_LENGTH:
		name = "GL_MAX_DEBUG_MESSAGE_LENGTH"
	case INT_IMAGE_1D_ARRAY:
		name = "GL_INT_IMAGE_1D_ARRAY"
	case FRAMEBUFFER_RENDERABLE:
		name = "GL_FRAMEBUFFER_RENDERABLE"
	case PROGRAM_PIPELINE_BINDING:
		name = "GL_PROGRAM_PIPELINE_BINDING"
	case TEXTURE17:
		name = "GL_TEXTURE17"
	case DRAW_BUFFER:
		name = "GL_DRAW_BUFFER"
	case FRAMEBUFFER_DEFAULT_HEIGHT:
		name = "GL_FRAMEBUFFER_DEFAULT_HEIGHT"
	case GEOMETRY_VERTICES_OUT:
		name = "GL_GEOMETRY_VERTICES_OUT"
	case MAX_VERTEX_UNIFORM_VECTORS:
		name = "GL_MAX_VERTEX_UNIFORM_VECTORS"
	case MAX_VERTEX_UNIFORM_COMPONENTS:
		name = "GL_MAX_VERTEX_UNIFORM_COMPONENTS"
	case TEXTURE_BINDING_1D:
		name = "GL_TEXTURE_BINDING_1D"
	case DST_COLOR:
		name = "GL_DST_COLOR"
	case DEBUG_GROUP_STACK_DEPTH:
		name = "GL_DEBUG_GROUP_STACK_DEPTH"
	case COMPRESSED_SRGB8_ALPHA8_ETC2_EAC:
		name = "GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC"
	case BLUE_INTEGER:
		name = "GL_BLUE_INTEGER"
	case UNPACK_ALIGNMENT:
		name = "GL_UNPACK_ALIGNMENT"
	case R8:
		name = "GL_R8"
	case MIN_PROGRAM_TEXEL_OFFSET:
		name = "GL_MIN_PROGRAM_TEXEL_OFFSET"
	case RGB5_A1:
		name = "GL_RGB5_A1"
	case TESS_CONTROL_TEXTURE:
		name = "GL_TESS_CONTROL_TEXTURE"
	case MAX_PATCH_VERTICES:
		name = "GL_MAX_PATCH_VERTICES"
	case BACK_LEFT:
		name = "GL_BACK_LEFT"
	case TEXTURE_CUBE_MAP_NEGATIVE_Z:
		name = "GL_TEXTURE_CUBE_MAP_NEGATIVE_Z"
	case COMPRESSED_RGBA8_ETC2_EAC:
		name = "GL_COMPRESSED_RGBA8_ETC2_EAC"
	case LOWER_LEFT:
		name = "GL_LOWER_LEFT"
	case PACK_IMAGE_HEIGHT:
		name = "GL_PACK_IMAGE_HEIGHT"
	case TEXTURE4:
		name = "GL_TEXTURE4"
	case READ_ONLY:
		name = "GL_READ_ONLY"
	case MAX_COLOR_TEXTURE_SAMPLES:
		name = "GL_MAX_COLOR_TEXTURE_SAMPLES"
	case TEXTURE31:
		name = "GL_TEXTURE31"
	case BLEND_SRC_RGB:
		name = "GL_BLEND_SRC_RGB"
	case SRC_COLOR:
		name = "GL_SRC_COLOR"
	case PACK_SKIP_ROWS:
		name = "GL_PACK_SKIP_ROWS"
	case UNPACK_COMPRESSED_BLOCK_HEIGHT:
		name = "GL_UNPACK_COMPRESSED_BLOCK_HEIGHT"
	case TEXTURE_SWIZZLE_G:
		name = "GL_TEXTURE_SWIZZLE_G"
	case LOCATION:
		name = "GL_LOCATION"
	case VIEW_CLASS_RGTC1_RED:
		name = "GL_VIEW_CLASS_RGTC1_RED"
	case INT_SAMPLER_CUBE_MAP_ARRAY:
		name = "GL_INT_SAMPLER_CUBE_MAP_ARRAY"
	case UNIFORM:
		name = "GL_UNIFORM"
	case CURRENT_VERTEX_ATTRIB:
		name = "GL_CURRENT_VERTEX_ATTRIB"
	case SAMPLE_SHADING:
		name = "GL_SAMPLE_SHADING"
	case STENCIL_TEST:
		name = "GL_STENCIL_TEST"
	case VIEW_CLASS_32_BITS:
		name = "GL_VIEW_CLASS_32_BITS"
	case ATOMIC_COUNTER_BUFFER_INDEX:
		name = "GL_ATOMIC_COUNTER_BUFFER_INDEX"
	case MATRIX_STRIDE:
		name = "GL_MATRIX_STRIDE"
	case COLOR_BUFFER_BIT:
		name = "GL_COLOR_BUFFER_BIT"
	case ACTIVE_VARIABLES:
		name = "GL_ACTIVE_VARIABLES"
	case TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		name = "GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT"
	case DEBUG_TYPE_PUSH_GROUP:
		name = "GL_DEBUG_TYPE_PUSH_GROUP"
	case R3_G3_B2:
		name = "GL_R3_G3_B2"
	case DRAW_FRAMEBUFFER_BINDING:
		name = "GL_DRAW_FRAMEBUFFER_BINDING"
	case PATCH_VERTICES:
		name = "GL_PATCH_VERTICES"
	case INTERNALFORMAT_RED_SIZE:
		name = "GL_INTERNALFORMAT_RED_SIZE"
	case PROXY_TEXTURE_1D:
		name = "GL_PROXY_TEXTURE_1D"
	case SAMPLER_2D_MULTISAMPLE_ARRAY:
		name = "GL_SAMPLER_2D_MULTISAMPLE_ARRAY"
	case SMOOTH_POINT_SIZE_RANGE:
		name = "GL_SMOOTH_POINT_SIZE_RANGE"
	case RGBA16I:
		name = "GL_RGBA16I"
	case MANUAL_GENERATE_MIPMAP:
		name = "GL_MANUAL_GENERATE_MIPMAP"
	case MAX_GEOMETRY_OUTPUT_COMPONENTS:
		name = "GL_MAX_GEOMETRY_OUTPUT_COMPONENTS"
	case BLEND_SRC:
		name = "GL_BLEND_SRC"
	case RGBA_SNORM:
		name = "GL_RGBA_SNORM"
	case MAX_GEOMETRY_SHADER_INVOCATIONS:
		name = "GL_MAX_GEOMETRY_SHADER_INVOCATIONS"
	case NAME_LENGTH:
		name = "GL_NAME_LENGTH"
	case INFO_LOG_LENGTH:
		name = "GL_INFO_LOG_LENGTH"
	case BLEND_DST:
		name = "GL_BLEND_DST"
	case SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST:
		name = "GL_SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST"
	case VIEW_CLASS_S3TC_DXT5_RGBA:
		name = "GL_VIEW_CLASS_S3TC_DXT5_RGBA"
	case SHADER_STORAGE_BUFFER:
		name = "GL_SHADER_STORAGE_BUFFER"
	case TEXTURE_CUBE_MAP_POSITIVE_Z:
		name = "GL_TEXTURE_CUBE_MAP_POSITIVE_Z"
	case MAX_DRAW_BUFFERS:
		name = "GL_MAX_DRAW_BUFFERS"
	case DRAW_BUFFER4:
		name = "GL_DRAW_BUFFER4"
	case FLOAT_MAT4:
		name = "GL_FLOAT_MAT4"
	case DEPTH_COMPONENTS:
		name = "GL_DEPTH_COMPONENTS"
	case REFERENCED_BY_COMPUTE_SHADER:
		name = "GL_REFERENCED_BY_COMPUTE_SHADER"
	case MAX_GEOMETRY_IMAGE_UNIFORMS:
		name = "GL_MAX_GEOMETRY_IMAGE_UNIFORMS"
	case ATOMIC_COUNTER_BUFFER_START:
		name = "GL_ATOMIC_COUNTER_BUFFER_START"
	case COLOR_ATTACHMENT1:
		name = "GL_COLOR_ATTACHMENT1"
	case UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX:
		name = "GL_UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX"
	case FILL:
		name = "GL_FILL"
	case MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS:
		name = "GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS"
	case DEBUG_TYPE_PERFORMANCE:
		name = "GL_DEBUG_TYPE_PERFORMANCE"
	case RGB16F:
		name = "GL_RGB16F"
	case TEXTURE_ALPHA_SIZE:
		name = "GL_TEXTURE_ALPHA_SIZE"
	case DEPTH_CLEAR_VALUE:
		name = "GL_DEPTH_CLEAR_VALUE"
	case TRANSFORM_FEEDBACK_VARYINGS:
		name = "GL_TRANSFORM_FEEDBACK_VARYINGS"
	case SHADER_IMAGE_LOAD:
		name = "GL_SHADER_IMAGE_LOAD"
	case MAX_WIDTH:
		name = "GL_MAX_WIDTH"
	case MAX_COLOR_ATTACHMENTS:
		name = "GL_MAX_COLOR_ATTACHMENTS"
	case ISOLINES:
		name = "GL_ISOLINES"
	case ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER:
		name = "GL_ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER"
	case RGB32F:
		name = "GL_RGB32F"
	case IMAGE_BINDING_NAME:
		name = "GL_IMAGE_BINDING_NAME"
	case MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS:
		name = "GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS"
	case DEBUG_TYPE_ERROR:
		name = "GL_DEBUG_TYPE_ERROR"
	case DOUBLE_MAT3x4:
		name = "GL_DOUBLE_MAT3x4"
	case AND:
		name = "GL_AND"
	case DEBUG_SOURCE_API:
		name = "GL_DEBUG_SOURCE_API"
	case MAX_VERTEX_UNIFORM_BLOCKS:
		name = "GL_MAX_VERTEX_UNIFORM_BLOCKS"
	case TEXTURE13:
		name = "GL_TEXTURE13"
	case TIMEOUT_EXPIRED:
		name = "GL_TIMEOUT_EXPIRED"
	case GEOMETRY_INPUT_TYPE:
		name = "GL_GEOMETRY_INPUT_TYPE"
	case DEBUG_TYPE_DEPRECATED_BEHAVIOR:
		name = "GL_DEBUG_TYPE_DEPRECATED_BEHAVIOR"
	case UNSIGNED_SHORT_5_6_5_REV:
		name = "GL_UNSIGNED_SHORT_5_6_5_REV"
	case NOTEQUAL:
		name = "GL_NOTEQUAL"
	case IMAGE_BINDING_LAYERED:
		name = "GL_IMAGE_BINDING_LAYERED"
	case STENCIL:
		name = "GL_STENCIL"
	case COLOR_LOGIC_OP:
		name = "GL_COLOR_LOGIC_OP"
	case INTERNALFORMAT_RED_TYPE:
		name = "GL_INTERNALFORMAT_RED_TYPE"
	case TRIANGLE_STRIP:
		name = "GL_TRIANGLE_STRIP"
	case TEXTURE_BORDER_COLOR:
		name = "GL_TEXTURE_BORDER_COLOR"
	case MAX_TRANSFORM_FEEDBACK_BUFFERS:
		name = "GL_MAX_TRANSFORM_FEEDBACK_BUFFERS"
	case IMAGE_BINDING_ACCESS:
		name = "GL_IMAGE_BINDING_ACCESS"
	case IMAGE_2D_RECT:
		name = "GL_IMAGE_2D_RECT"
	case NOR:
		name = "GL_NOR"
	case SAMPLE_COVERAGE_INVERT:
		name = "GL_SAMPLE_COVERAGE_INVERT"
	case CW:
		name = "GL_CW"
	case ARRAY_STRIDE:
		name = "GL_ARRAY_STRIDE"
	case VIEW_CLASS_S3TC_DXT1_RGB:
		name = "GL_VIEW_CLASS_S3TC_DXT1_RGB"
	case TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY:
		name = "GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY"
	case COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT:
		name = "GL_COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT"
	case IMAGE_CLASS_4_X_8:
		name = "GL_IMAGE_CLASS_4_X_8"
	case ONE_MINUS_SRC_ALPHA:
		name = "GL_ONE_MINUS_SRC_ALPHA"
	case DEBUG_SEVERITY_MEDIUM:
		name = "GL_DEBUG_SEVERITY_MEDIUM"
	case DYNAMIC_COPY:
		name = "GL_DYNAMIC_COPY"
	case MAX_SAMPLE_MASK_WORDS:
		name = "GL_MAX_SAMPLE_MASK_WORDS"
	case GEQUAL:
		name = "GL_GEQUAL"
	case IMAGE_CLASS_2_X_16:
		name = "GL_IMAGE_CLASS_2_X_16"
	case NEAREST_MIPMAP_LINEAR:
		name = "GL_NEAREST_MIPMAP_LINEAR"
	case COMPUTE_TEXTURE:
		name = "GL_COMPUTE_TEXTURE"
	case DEBUG_SEVERITY_NOTIFICATION:
		name = "GL_DEBUG_SEVERITY_NOTIFICATION"
	case DEPTH32F_STENCIL8:
		name = "GL_DEPTH32F_STENCIL8"
	case UNSIGNED_SHORT_5_5_5_1:
		name = "GL_UNSIGNED_SHORT_5_5_5_1"
	case UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY:
		name = "GL_UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY"
	case SRGB8_ALPHA8:
		name = "GL_SRGB8_ALPHA8"
	case R16F:
		name = "GL_R16F"
	case COMPRESSED_SIGNED_RG11_EAC:
		name = "GL_COMPRESSED_SIGNED_RG11_EAC"
	case FIRST_VERTEX_CONVENTION:
		name = "GL_FIRST_VERTEX_CONVENTION"
	case ONE_MINUS_DST_COLOR:
		name = "GL_ONE_MINUS_DST_COLOR"
	case ANY_SAMPLES_PASSED:
		name = "GL_ANY_SAMPLES_PASSED"
	case UNSIGNED_INT_SAMPLER_3D:
		name = "GL_UNSIGNED_INT_SAMPLER_3D"
	case BUFFER:
		name = "GL_BUFFER"
	case UNIFORM_TYPE:
		name = "GL_UNIFORM_TYPE"
	case TEXTURE_COMPRESSION_HINT:
		name = "GL_TEXTURE_COMPRESSION_HINT"
	case UNIFORM_BUFFER_START:
		name = "GL_UNIFORM_BUFFER_START"
	case TEXTURE14:
		name = "GL_TEXTURE14"
	case BUFFER_USAGE:
		name = "GL_BUFFER_USAGE"
	case AUTO_GENERATE_MIPMAP:
		name = "GL_AUTO_GENERATE_MIPMAP"
	case TEXTURE10:
		name = "GL_TEXTURE10"
	case DOUBLE_VEC3:
		name = "GL_DOUBLE_VEC3"
	case NUM_PROGRAM_BINARY_FORMATS:
		name = "GL_NUM_PROGRAM_BINARY_FORMATS"
	case UNPACK_SKIP_PIXELS:
		name = "GL_UNPACK_SKIP_PIXELS"
	case COPY:
		name = "GL_COPY"
	case TESS_CONTROL_SUBROUTINE:
		name = "GL_TESS_CONTROL_SUBROUTINE"
	case SLUMINANCE_EXT:
		name = "GL_SLUMINANCE_EXT"
	case FRONT_FACE:
		name = "GL_FRONT_FACE"
	case MAX_FRAGMENT_INPUT_COMPONENTS:
		name = "GL_MAX_FRAGMENT_INPUT_COMPONENTS"
	case FLOAT_MAT4x2:
		name = "GL_FLOAT_MAT4x2"
	case UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER:
		name = "GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER"
	case DEBUG_SOURCE_THIRD_PARTY:
		name = "GL_DEBUG_SOURCE_THIRD_PARTY"
	case GET_TEXTURE_IMAGE_FORMAT:
		name = "GL_GET_TEXTURE_IMAGE_FORMAT"
	case RGB:
		name = "GL_RGB"
	case LAST_VERTEX_CONVENTION:
		name = "GL_LAST_VERTEX_CONVENTION"
	case INTERNALFORMAT_GREEN_TYPE:
		name = "GL_INTERNALFORMAT_GREEN_TYPE"
	case SHADER_SOURCE_LENGTH:
		name = "GL_SHADER_SOURCE_LENGTH"
	case MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS:
		name = "GL_MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS"
	case UNSIGNED_INT_8_8_8_8_REV:
		name = "GL_UNSIGNED_INT_8_8_8_8_REV"
	case SAMPLER_BUFFER:
		name = "GL_SAMPLER_BUFFER"
	case RGBA16F:
		name = "GL_RGBA16F"
	case MAX_VIEWPORT_DIMS:
		name = "GL_MAX_VIEWPORT_DIMS"
	case R32UI:
		name = "GL_R32UI"
	case OR_REVERSE:
		name = "GL_OR_REVERSE"
	case UNSIGNED_INT_IMAGE_3D:
		name = "GL_UNSIGNED_INT_IMAGE_3D"
	case SAMPLER_2D_SHADOW:
		name = "GL_SAMPLER_2D_SHADOW"
	case DEBUG_SOURCE_OTHER:
		name = "GL_DEBUG_SOURCE_OTHER"
	case RGB565:
		name = "GL_RGB565"
	case DYNAMIC_DRAW:
		name = "GL_DYNAMIC_DRAW"
	case ARRAY_BUFFER:
		name = "GL_ARRAY_BUFFER"
	case RG32I:
		name = "GL_RG32I"
	case UNPACK_COMPRESSED_BLOCK_SIZE:
		name = "GL_UNPACK_COMPRESSED_BLOCK_SIZE"
	case CONDITION_SATISFIED:
		name = "GL_CONDITION_SATISFIED"
	case INT_SAMPLER_CUBE:
		name = "GL_INT_SAMPLER_CUBE"
	case INTERNALFORMAT_ALPHA_SIZE:
		name = "GL_INTERNALFORMAT_ALPHA_SIZE"
	case IS_ROW_MAJOR:
		name = "GL_IS_ROW_MAJOR"
	case RIGHT:
		name = "GL_RIGHT"
	case DRAW_FRAMEBUFFER:
		name = "GL_DRAW_FRAMEBUFFER"
	case RGB8UI:
		name = "GL_RGB8UI"
	case TEXTURE30:
		name = "GL_TEXTURE30"
	case OR:
		name = "GL_OR"
	case UNPACK_COMPRESSED_BLOCK_WIDTH:
		name = "GL_UNPACK_COMPRESSED_BLOCK_WIDTH"
	case TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH:
		name = "GL_TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH"
	case R16:
		name = "GL_R16"
	case UNIFORM_IS_ROW_MAJOR:
		name = "GL_UNIFORM_IS_ROW_MAJOR"
	case BLOCK_INDEX:
		name = "GL_BLOCK_INDEX"
	case SHADER_TYPE:
		name = "GL_SHADER_TYPE"
	case INT_VEC3:
		name = "GL_INT_VEC3"
	case TEXTURE_BINDING_2D_MULTISAMPLE:
		name = "GL_TEXTURE_BINDING_2D_MULTISAMPLE"
	case MIPMAP:
		name = "GL_MIPMAP"
	case STENCIL_ATTACHMENT:
		name = "GL_STENCIL_ATTACHMENT"
	case AND_REVERSE:
		name = "GL_AND_REVERSE"
	case MAX_ATOMIC_COUNTER_BUFFER_BINDINGS:
		name = "GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS"
	case RGBA16_SNORM:
		name = "GL_RGBA16_SNORM"
	case UNIFORM_ARRAY_STRIDE:
		name = "GL_UNIFORM_ARRAY_STRIDE"
	case FRAGMENT_SUBROUTINE:
		name = "GL_FRAGMENT_SUBROUTINE"
	case RG16I:
		name = "GL_RG16I"
	case RGBA:
		name = "GL_RGBA"
	case ACTIVE_ATTRIBUTES:
		name = "GL_ACTIVE_ATTRIBUTES"
	case FLOAT_MAT2:
		name = "GL_FLOAT_MAT2"
	case MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS:
		name = "GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS"
	case HIGH_FLOAT:
		name = "GL_HIGH_FLOAT"
	case PRIMITIVE_RESTART_FIXED_INDEX:
		name = "GL_PRIMITIVE_RESTART_FIXED_INDEX"
	case MIN_SAMPLE_SHADING_VALUE:
		name = "GL_MIN_SAMPLE_SHADING_VALUE"
	case PATCHES:
		name = "GL_PATCHES"
	case UNIFORM_BUFFER_BINDING:
		name = "GL_UNIFORM_BUFFER_BINDING"
	case TEXTURE16:
		name = "GL_TEXTURE16"
	case MAX_COMPUTE_IMAGE_UNIFORMS:
		name = "GL_MAX_COMPUTE_IMAGE_UNIFORMS"
	case RGB9_E5:
		name = "GL_RGB9_E5"
	case LESS:
		name = "GL_LESS"
	case COPY_INVERTED:
		name = "GL_COPY_INVERTED"
	case UNSIGNED_INT_SAMPLER_2D:
		name = "GL_UNSIGNED_INT_SAMPLER_2D"
	case TRIANGLES_ADJACENCY:
		name = "GL_TRIANGLES_ADJACENCY"
	case TEXTURE_COMPRESSED:
		name = "GL_TEXTURE_COMPRESSED"
	case INT_IMAGE_2D_ARRAY:
		name = "GL_INT_IMAGE_2D_ARRAY"
	case DEPTH_ATTACHMENT:
		name = "GL_DEPTH_ATTACHMENT"
	case DOUBLE_MAT2x4:
		name = "GL_DOUBLE_MAT2x4"
	case MAX_COMPUTE_TEXTURE_IMAGE_UNITS:
		name = "GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS"
	case RG16UI:
		name = "GL_RG16UI"
	case DEPTH_RANGE:
		name = "GL_DEPTH_RANGE"
	case POLYGON_SMOOTH_HINT:
		name = "GL_POLYGON_SMOOTH_HINT"
	case RENDERBUFFER_RED_SIZE:
		name = "GL_RENDERBUFFER_RED_SIZE"
	case MAX_TESS_GEN_LEVEL:
		name = "GL_MAX_TESS_GEN_LEVEL"
	case UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE:
		name = "GL_UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE"
	case VERTEX_SUBROUTINE:
		name = "GL_VERTEX_SUBROUTINE"
	case PROGRAM_PIPELINE:
		name = "GL_PROGRAM_PIPELINE"
	case MAX_COMPUTE_SHARED_MEMORY_SIZE:
		name = "GL_MAX_COMPUTE_SHARED_MEMORY_SIZE"
	case MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS:
		name = "GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS"
	case COLOR_ATTACHMENT5:
		name = "GL_COLOR_ATTACHMENT5"
	case TESS_CONTROL_SUBROUTINE_UNIFORM:
		name = "GL_TESS_CONTROL_SUBROUTINE_UNIFORM"
	case MAX_SAMPLES:
		name = "GL_MAX_SAMPLES"
	case BUFFER_SIZE:
		name = "GL_BUFFER_SIZE"
	case COMPRESSED_SIGNED_RG_RGTC2:
		name = "GL_COMPRESSED_SIGNED_RG_RGTC2"
	case MAX_VIEWPORTS:
		name = "GL_MAX_VIEWPORTS"
	case STENCIL_BACK_VALUE_MASK:
		name = "GL_STENCIL_BACK_VALUE_MASK"
	case MAX_DEPTH:
		name = "GL_MAX_DEPTH"
	case FLOAT_MAT4x3:
		name = "GL_FLOAT_MAT4x3"
	case DOUBLE_MAT4:
		name = "GL_DOUBLE_MAT4"
	case DYNAMIC_READ:
		name = "GL_DYNAMIC_READ"
	case INT_SAMPLER_2D_RECT:
		name = "GL_INT_SAMPLER_2D_RECT"
	case PACK_COMPRESSED_BLOCK_HEIGHT:
		name = "GL_PACK_COMPRESSED_BLOCK_HEIGHT"
	case TEXTURE29:
		name = "GL_TEXTURE29"
	case IMAGE_FORMAT_COMPATIBILITY_BY_CLASS:
		name = "GL_IMAGE_FORMAT_COMPATIBILITY_BY_CLASS"
	case VERTEX_ATTRIB_ARRAY_LONG:
		name = "GL_VERTEX_ATTRIB_ARRAY_LONG"
	case BGRA:
		name = "GL_BGRA"
	case TEXTURE_2D_MULTISAMPLE_ARRAY:
		name = "GL_TEXTURE_2D_MULTISAMPLE_ARRAY"
	case MAX_COMBINED_SHADER_OUTPUT_RESOURCES:
		name = "GL_MAX_COMBINED_SHADER_OUTPUT_RESOURCES"
	case PRIMITIVE_RESTART_INDEX:
		name = "GL_PRIMITIVE_RESTART_INDEX"
	case DEPTH:
		name = "GL_DEPTH"
	case STENCIL_RENDERABLE:
		name = "GL_STENCIL_RENDERABLE"
	case READ_BUFFER:
		name = "GL_READ_BUFFER"
	case COLOR_ATTACHMENT12:
		name = "GL_COLOR_ATTACHMENT12"
	case TESS_EVALUATION_SUBROUTINE:
		name = "GL_TESS_EVALUATION_SUBROUTINE"
	case VERTEX_ATTRIB_ARRAY_ENABLED:
		name = "GL_VERTEX_ATTRIB_ARRAY_ENABLED"
	case SAMPLER_CUBE_MAP_ARRAY_SHADOW:
		name = "GL_SAMPLER_CUBE_MAP_ARRAY_SHADOW"
	case RGB16UI:
		name = "GL_RGB16UI"
	case INTERNALFORMAT_GREEN_SIZE:
		name = "GL_INTERNALFORMAT_GREEN_SIZE"
	case TEXTURE_VIEW:
		name = "GL_TEXTURE_VIEW"
	case ACTIVE_TEXTURE:
		name = "GL_ACTIVE_TEXTURE"
	case MIN_PROGRAM_TEXTURE_GATHER_OFFSET:
		name = "GL_MIN_PROGRAM_TEXTURE_GATHER_OFFSET"
	case LINEAR_MIPMAP_NEAREST:
		name = "GL_LINEAR_MIPMAP_NEAREST"
	case BGR_INTEGER:
		name = "GL_BGR_INTEGER"
	case COMPUTE_LOCAL_WORK_SIZE:
		name = "GL_COMPUTE_LOCAL_WORK_SIZE"
	case RG32UI:
		name = "GL_RG32UI"
	case RG8_SNORM:
		name = "GL_RG8_SNORM"
	case GEOMETRY_OUTPUT_TYPE:
		name = "GL_GEOMETRY_OUTPUT_TYPE"
	case NUM_ACTIVE_VARIABLES:
		name = "GL_NUM_ACTIVE_VARIABLES"
	case FLOAT:
		name = "GL_FLOAT"
	case COLOR_ATTACHMENT0:
		name = "GL_COLOR_ATTACHMENT0"
	case TEXTURE_DEPTH_SIZE:
		name = "GL_TEXTURE_DEPTH_SIZE"
	case MULTISAMPLE:
		name = "GL_MULTISAMPLE"
	case COMPRESSED_SRGB_EXT:
		name = "GL_COMPRESSED_SRGB_EXT"
	case UNSIGNED_INT_VEC4:
		name = "GL_UNSIGNED_INT_VEC4"
	case VIEW_CLASS_64_BITS:
		name = "GL_VIEW_CLASS_64_BITS"
	case FLOAT_MAT2x4:
		name = "GL_FLOAT_MAT2x4"
	case INTERNALFORMAT_STENCIL_TYPE:
		name = "GL_INTERNALFORMAT_STENCIL_TYPE"
	case TRANSFORM_FEEDBACK_BUFFER_SIZE:
		name = "GL_TRANSFORM_FEEDBACK_BUFFER_SIZE"
	case RENDERBUFFER_WIDTH:
		name = "GL_RENDERBUFFER_WIDTH"
	case FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS:
		name = "GL_FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS"
	case KEEP:
		name = "GL_KEEP"
	case DOUBLE_MAT3x2:
		name = "GL_DOUBLE_MAT3x2"
	case TEXTURE_SWIZZLE_R:
		name = "GL_TEXTURE_SWIZZLE_R"
	case PROGRAM_BINARY_FORMATS:
		name = "GL_PROGRAM_BINARY_FORMATS"
	case STENCIL_BACK_FUNC:
		name = "GL_STENCIL_BACK_FUNC"
	case R32I:
		name = "GL_R32I"
	case UNSIGNED_INT_SAMPLER_CUBE:
		name = "GL_UNSIGNED_INT_SAMPLER_CUBE"
	case RG16F:
		name = "GL_RG16F"
	case GEOMETRY_TEXTURE:
		name = "GL_GEOMETRY_TEXTURE"
	case UNIFORM_MATRIX_STRIDE:
		name = "GL_UNIFORM_MATRIX_STRIDE"
	case RENDERBUFFER:
		name = "GL_RENDERBUFFER"
	case UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY:
		name = "GL_UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY"
	case SHADER_STORAGE_BUFFER_START:
		name = "GL_SHADER_STORAGE_BUFFER_START"
	case SAMPLES:
		name = "GL_SAMPLES"
	case VIEW_CLASS_24_BITS:
		name = "GL_VIEW_CLASS_24_BITS"
	case TEXTURE_RED_SIZE:
		name = "GL_TEXTURE_RED_SIZE"
	case RG32F:
		name = "GL_RG32F"
	case SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST:
		name = "GL_SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST"
	case DEBUG_SOURCE_SHADER_COMPILER:
		name = "GL_DEBUG_SOURCE_SHADER_COMPILER"
	case SHADER_IMAGE_STORE:
		name = "GL_SHADER_IMAGE_STORE"
	case R11F_G11F_B10F:
		name = "GL_R11F_G11F_B10F"
	case DOUBLEBUFFER:
		name = "GL_DOUBLEBUFFER"
	case GEOMETRY_SHADER_INVOCATIONS:
		name = "GL_GEOMETRY_SHADER_INVOCATIONS"
	case INTERNALFORMAT_BLUE_TYPE:
		name = "GL_INTERNALFORMAT_BLUE_TYPE"
	case PACK_ROW_LENGTH:
		name = "GL_PACK_ROW_LENGTH"
	case MAX_COMBINED_UNIFORM_BLOCKS:
		name = "GL_MAX_COMBINED_UNIFORM_BLOCKS"
	case INVERT:
		name = "GL_INVERT"
	case SYNC_FENCE:
		name = "GL_SYNC_FENCE"
	case SAMPLER_1D_ARRAY_SHADOW:
		name = "GL_SAMPLER_1D_ARRAY_SHADOW"
	case SHORT:
		name = "GL_SHORT"
	case TEXTURE_BINDING_CUBE_MAP_ARRAY:
		name = "GL_TEXTURE_BINDING_CUBE_MAP_ARRAY"
	case MAX_PROGRAM_TEXEL_OFFSET:
		name = "GL_MAX_PROGRAM_TEXEL_OFFSET"
	case POLYGON_MODE:
		name = "GL_POLYGON_MODE"
	case MAX_TESS_EVALUATION_UNIFORM_COMPONENTS:
		name = "GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS"
	case FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS:
		name = "GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS"
	case SCISSOR_BOX:
		name = "GL_SCISSOR_BOX"
	case INT_SAMPLER_1D_ARRAY:
		name = "GL_INT_SAMPLER_1D_ARRAY"
	case VERTEX_ATTRIB_ARRAY_DIVISOR:
		name = "GL_VERTEX_ATTRIB_ARRAY_DIVISOR"
	case UNPACK_COMPRESSED_BLOCK_DEPTH:
		name = "GL_UNPACK_COMPRESSED_BLOCK_DEPTH"
	case DEBUG_TYPE_OTHER:
		name = "GL_DEBUG_TYPE_OTHER"
	case LINE_STRIP:
		name = "GL_LINE_STRIP"
	case VENDOR:
		name = "GL_VENDOR"
	case UNSIGNED_INT_IMAGE_CUBE:
		name = "GL_UNSIGNED_INT_IMAGE_CUBE"
	case RGB8_SNORM:
		name = "GL_RGB8_SNORM"
	case DEPTH_STENCIL_ATTACHMENT:
		name = "GL_DEPTH_STENCIL_ATTACHMENT"
	case MAX_TEXTURE_BUFFER_SIZE:
		name = "GL_MAX_TEXTURE_BUFFER_SIZE"
	case FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE"
	case FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE"
	case TEXTURE_3D:
		name = "GL_TEXTURE_3D"
	case BUFFER_UPDATE_BARRIER_BIT:
		name = "GL_BUFFER_UPDATE_BARRIER_BIT"
	case PROXY_TEXTURE_2D_ARRAY:
		name = "GL_PROXY_TEXTURE_2D_ARRAY"
	case FIXED_ONLY:
		name = "GL_FIXED_ONLY"
	case UNSIGNED_INT_IMAGE_1D_ARRAY:
		name = "GL_UNSIGNED_INT_IMAGE_1D_ARRAY"
	case STENCIL_WRITEMASK:
		name = "GL_STENCIL_WRITEMASK"
	case VERTEX_ARRAY_BINDING:
		name = "GL_VERTEX_ARRAY_BINDING"
	case MAX_FRAMEBUFFER_HEIGHT:
		name = "GL_MAX_FRAMEBUFFER_HEIGHT"
	case BUFFER_VARIABLE:
		name = "GL_BUFFER_VARIABLE"
	case IMAGE_COMPATIBILITY_CLASS:
		name = "GL_IMAGE_COMPATIBILITY_CLASS"
	case ATOMIC_COUNTER_BUFFER_DATA_SIZE:
		name = "GL_ATOMIC_COUNTER_BUFFER_DATA_SIZE"
	case RGBA12:
		name = "GL_RGBA12"
	case QUERY_WAIT:
		name = "GL_QUERY_WAIT"
	case TEXTURE_BUFFER_DATA_STORE_BINDING:
		name = "GL_TEXTURE_BUFFER_DATA_STORE_BINDING"
	case LINE:
		name = "GL_LINE"
	case RGB16:
		name = "GL_RGB16"
	case DELETE_STATUS:
		name = "GL_DELETE_STATUS"
	case DOUBLE:
		name = "GL_DOUBLE"
	case DRAW_BUFFER0:
		name = "GL_DRAW_BUFFER0"
	case DEBUG_SEVERITY_HIGH:
		name = "GL_DEBUG_SEVERITY_HIGH"
	case MAX_FRAGMENT_SHADER_STORAGE_BLOCKS:
		name = "GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS"
	case MAX_CLIP_DISTANCES:
		name = "GL_MAX_CLIP_DISTANCES"
	case COLOR_WRITEMASK:
		name = "GL_COLOR_WRITEMASK"
	case CLIP_DISTANCE3:
		name = "GL_CLIP_DISTANCE3"
	case FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE"
	case LAYER_PROVOKING_VERTEX:
		name = "GL_LAYER_PROVOKING_VERTEX"
	case UNIFORM_SIZE:
		name = "GL_UNIFORM_SIZE"
	case ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH:
		name = "GL_ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH"
	case UNSIGNED_INT_5_9_9_9_REV:
		name = "GL_UNSIGNED_INT_5_9_9_9_REV"
	case DRAW_BUFFER15:
		name = "GL_DRAW_BUFFER15"
	case UNIFORM_BUFFER_OFFSET_ALIGNMENT:
		name = "GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT"
	case DEBUG_TYPE_UNDEFINED_BEHAVIOR:
		name = "GL_DEBUG_TYPE_UNDEFINED_BEHAVIOR"
	case PATCH_DEFAULT_INNER_LEVEL:
		name = "GL_PATCH_DEFAULT_INNER_LEVEL"
	case IMAGE_BINDING_LAYER:
		name = "GL_IMAGE_BINDING_LAYER"
	case TEXTURE_WRAP_S:
		name = "GL_TEXTURE_WRAP_S"
	case UNSIGNED_INT_IMAGE_BUFFER:
		name = "GL_UNSIGNED_INT_IMAGE_BUFFER"
	case IS_PER_PATCH:
		name = "GL_IS_PER_PATCH"
	case RGB8:
		name = "GL_RGB8"
	case MAX_FRAGMENT_UNIFORM_VECTORS:
		name = "GL_MAX_FRAGMENT_UNIFORM_VECTORS"
	case COMPRESSED_SIGNED_R11_EAC:
		name = "GL_COMPRESSED_SIGNED_R11_EAC"
	case SHADER_STORAGE_BARRIER_BIT:
		name = "GL_SHADER_STORAGE_BARRIER_BIT"
	case RED_INTEGER:
		name = "GL_RED_INTEGER"
	case DRAW_BUFFER1:
		name = "GL_DRAW_BUFFER1"
	case ELEMENT_ARRAY_BUFFER:
		name = "GL_ELEMENT_ARRAY_BUFFER"
	case UNSIGNED_INT_SAMPLER_2D_ARRAY:
		name = "GL_UNSIGNED_INT_SAMPLER_2D_ARRAY"
	case ACTIVE_SUBROUTINE_MAX_LENGTH:
		name = "GL_ACTIVE_SUBROUTINE_MAX_LENGTH"
	case READ_PIXELS_TYPE:
		name = "GL_READ_PIXELS_TYPE"
	case TEXTURE:
		name = "GL_TEXTURE"
	case INT_SAMPLER_2D:
		name = "GL_INT_SAMPLER_2D"
	case PROGRAM_OUTPUT:
		name = "GL_PROGRAM_OUTPUT"
	case GEOMETRY_SUBROUTINE:
		name = "GL_GEOMETRY_SUBROUTINE"
	case UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER:
		name = "GL_UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER"
	case UNPACK_SKIP_ROWS:
		name = "GL_UNPACK_SKIP_ROWS"
	case IMAGE_CLASS_11_11_10:
		name = "GL_IMAGE_CLASS_11_11_10"
	case DRAW_BUFFER13:
		name = "GL_DRAW_BUFFER13"
	case BOOL_VEC2:
		name = "GL_BOOL_VEC2"
	case COMPILE_STATUS:
		name = "GL_COMPILE_STATUS"
	case MAX_FRAGMENT_UNIFORM_COMPONENTS:
		name = "GL_MAX_FRAGMENT_UNIFORM_COMPONENTS"
	case PROVOKING_VERTEX:
		name = "GL_PROVOKING_VERTEX"
	case TRANSFORM_FEEDBACK_BUFFER:
		name = "GL_TRANSFORM_FEEDBACK_BUFFER"
	case CLAMP_TO_EDGE:
		name = "GL_CLAMP_TO_EDGE"
	case COLOR_COMPONENTS:
		name = "GL_COLOR_COMPONENTS"
	case INT_SAMPLER_3D:
		name = "GL_INT_SAMPLER_3D"
	case STATIC_COPY:
		name = "GL_STATIC_COPY"
	case BLEND:
		name = "GL_BLEND"
	case MAX_RECTANGLE_TEXTURE_SIZE:
		name = "GL_MAX_RECTANGLE_TEXTURE_SIZE"
	case SMOOTH_LINE_WIDTH_GRANULARITY:
		name = "GL_SMOOTH_LINE_WIDTH_GRANULARITY"
	case INT_SAMPLER_2D_ARRAY:
		name = "GL_INT_SAMPLER_2D_ARRAY"
	case DEPTH_STENCIL_TEXTURE_MODE:
		name = "GL_DEPTH_STENCIL_TEXTURE_MODE"
	case MAX_GEOMETRY_INPUT_COMPONENTS:
		name = "GL_MAX_GEOMETRY_INPUT_COMPONENTS"
	case BUFFER_MAP_LENGTH:
		name = "GL_BUFFER_MAP_LENGTH"
	case VERTEX_ATTRIB_ARRAY_NORMALIZED:
		name = "GL_VERTEX_ATTRIB_ARRAY_NORMALIZED"
	case TEXTURE28:
		name = "GL_TEXTURE28"
	case GREATER:
		name = "GL_GREATER"
	case RGBA32UI:
		name = "GL_RGBA32UI"
	case IMPLEMENTATION_COLOR_READ_FORMAT:
		name = "GL_IMPLEMENTATION_COLOR_READ_FORMAT"
	case UNPACK_ROW_LENGTH:
		name = "GL_UNPACK_ROW_LENGTH"
	case TEXTURE_MAX_LOD:
		name = "GL_TEXTURE_MAX_LOD"
	case DOUBLE_VEC4:
		name = "GL_DOUBLE_VEC4"
	case TEXTURE_CUBE_MAP_ARRAY:
		name = "GL_TEXTURE_CUBE_MAP_ARRAY"
	case COLOR_CLEAR_VALUE:
		name = "GL_COLOR_CLEAR_VALUE"
	case TEXTURE_BINDING_2D:
		name = "GL_TEXTURE_BINDING_2D"
	case MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS:
		name = "GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS"
	case NUM_COMPRESSED_TEXTURE_FORMATS:
		name = "GL_NUM_COMPRESSED_TEXTURE_FORMATS"
	case TEXTURE_WRAP_R:
		name = "GL_TEXTURE_WRAP_R"
	case MAX_COMPUTE_LOCAL_INVOCATIONS:
		name = "GL_MAX_COMPUTE_LOCAL_INVOCATIONS"
	case FLOAT_MAT3x2:
		name = "GL_FLOAT_MAT3x2"
	case LINEAR:
		name = "GL_LINEAR"
	case IMAGE_CLASS_1_X_16:
		name = "GL_IMAGE_CLASS_1_X_16"
	case FRAMEBUFFER_UNSUPPORTED:
		name = "GL_FRAMEBUFFER_UNSUPPORTED"
	case PROXY_TEXTURE_3D:
		name = "GL_PROXY_TEXTURE_3D"
	case ACTIVE_ATOMIC_COUNTER_BUFFERS:
		name = "GL_ACTIVE_ATOMIC_COUNTER_BUFFERS"
	case MAX_TESS_EVALUATION_IMAGE_UNIFORMS:
		name = "GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS"
	case DEPTH_CLAMP:
		name = "GL_DEPTH_CLAMP"
	case RENDERBUFFER_HEIGHT:
		name = "GL_RENDERBUFFER_HEIGHT"
	case TEXTURE_BINDING_CUBE_MAP:
		name = "GL_TEXTURE_BINDING_CUBE_MAP"
	case DOUBLE_MAT2x3:
		name = "GL_DOUBLE_MAT2x3"
	case TEXTURE_IMAGE_TYPE:
		name = "GL_TEXTURE_IMAGE_TYPE"
	case DST_ALPHA:
		name = "GL_DST_ALPHA"
	case TEXTURE_2D_ARRAY:
		name = "GL_TEXTURE_2D_ARRAY"
	case DEBUG_TYPE_POP_GROUP:
		name = "GL_DEBUG_TYPE_POP_GROUP"
	case MAX_TESS_CONTROL_UNIFORM_COMPONENTS:
		name = "GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS"
	case RG16:
		name = "GL_RG16"
	case MAX_VERTEX_OUTPUT_COMPONENTS:
		name = "GL_MAX_VERTEX_OUTPUT_COMPONENTS"
	case IMAGE_CLASS_4_X_16:
		name = "GL_IMAGE_CLASS_4_X_16"
	case IMAGE_FORMAT_COMPATIBILITY_TYPE:
		name = "GL_IMAGE_FORMAT_COMPATIBILITY_TYPE"
	case UNIFORM_BLOCK_INDEX:
		name = "GL_UNIFORM_BLOCK_INDEX"
	case MAX_DEBUG_LOGGED_MESSAGES:
		name = "GL_MAX_DEBUG_LOGGED_MESSAGES"
	case VERTEX_BINDING_STRIDE:
		name = "GL_VERTEX_BINDING_STRIDE"
	case COLOR_ATTACHMENT11:
		name = "GL_COLOR_ATTACHMENT11"
	case UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER:
		name = "GL_UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER"
	case DOUBLE_MAT4x2:
		name = "GL_DOUBLE_MAT4x2"
	case ACTIVE_UNIFORM_MAX_LENGTH:
		name = "GL_ACTIVE_UNIFORM_MAX_LENGTH"
	case PACK_SKIP_PIXELS:
		name = "GL_PACK_SKIP_PIXELS"
	case DEPTH_FUNC:
		name = "GL_DEPTH_FUNC"
	case COLOR_ATTACHMENT9:
		name = "GL_COLOR_ATTACHMENT9"
	case MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS:
		name = "GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS"
	case MAX_SUBROUTINES:
		name = "GL_MAX_SUBROUTINES"
	case DRAW_BUFFER5:
		name = "GL_DRAW_BUFFER5"
	case MAX_TEXTURE_LOD_BIAS:
		name = "GL_MAX_TEXTURE_LOD_BIAS"
	case RENDERBUFFER_INTERNAL_FORMAT:
		name = "GL_RENDERBUFFER_INTERNAL_FORMAT"
	case STEREO:
		name = "GL_STEREO"
	case PROXY_TEXTURE_1D_ARRAY:
		name = "GL_PROXY_TEXTURE_1D_ARRAY"
	case SAMPLER_2D:
		name = "GL_SAMPLER_2D"
	case MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		name = "GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT"
	case NUM_SHADING_LANGUAGE_VERSIONS:
		name = "GL_NUM_SHADING_LANGUAGE_VERSIONS"
	case RENDERBUFFER_DEPTH_SIZE:
		name = "GL_RENDERBUFFER_DEPTH_SIZE"
	case MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS:
		name = "GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS"
	case VERTEX_BINDING_DIVISOR:
		name = "GL_VERTEX_BINDING_DIVISOR"
	case MAX_VERTEX_IMAGE_UNIFORMS:
		name = "GL_MAX_VERTEX_IMAGE_UNIFORMS"
	case STENCIL_INDEX:
		name = "GL_STENCIL_INDEX"
	case UNIFORM_BLOCK_BINDING:
		name = "GL_UNIFORM_BLOCK_BINDING"
	case TESS_EVALUATION_TEXTURE:
		name = "GL_TESS_EVALUATION_TEXTURE"
	case MAX_SHADER_STORAGE_BUFFER_BINDINGS:
		name = "GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS"
	case CONTEXT_PROFILE_MASK:
		name = "GL_CONTEXT_PROFILE_MASK"
	case BYTE:
		name = "GL_BYTE"
	case PROGRAM_SEPARABLE:
		name = "GL_PROGRAM_SEPARABLE"
	case R32F:
		name = "GL_R32F"
	case TRANSFORM_FEEDBACK_BUFFER_START:
		name = "GL_TRANSFORM_FEEDBACK_BUFFER_START"
	case VIEW_CLASS_48_BITS:
		name = "GL_VIEW_CLASS_48_BITS"
	case UNSIGNED_NORMALIZED:
		name = "GL_UNSIGNED_NORMALIZED"
	case RGBA16UI:
		name = "GL_RGBA16UI"
	case FRAGMENT_INTERPOLATION_OFFSET_BITS:
		name = "GL_FRAGMENT_INTERPOLATION_OFFSET_BITS"
	case FRAMEBUFFER_INCOMPLETE_READ_BUFFER:
		name = "GL_FRAMEBUFFER_INCOMPLETE_READ_BUFFER"
	case LINK_STATUS:
		name = "GL_LINK_STATUS"
	case TEXTURE22:
		name = "GL_TEXTURE22"
	case R8I:
		name = "GL_R8I"
	case MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS:
		name = "GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS"
	case MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS:
		name = "GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS"
	case MIRRORED_REPEAT:
		name = "GL_MIRRORED_REPEAT"
	case SAMPLER_2D_ARRAY_SHADOW:
		name = "GL_SAMPLER_2D_ARRAY_SHADOW"
	case SAMPLER_BINDING:
		name = "GL_SAMPLER_BINDING"
	case DEPTH_COMPONENT32:
		name = "GL_DEPTH_COMPONENT32"
	case VERTEX_TEXTURE:
		name = "GL_VERTEX_TEXTURE"
	case SAMPLE_MASK_VALUE:
		name = "GL_SAMPLE_MASK_VALUE"
	case UNSIGNED_INT_8_8_8_8:
		name = "GL_UNSIGNED_INT_8_8_8_8"
	case PACK_SWAP_BYTES:
		name = "GL_PACK_SWAP_BYTES"
	case QUERY_BY_REGION_WAIT:
		name = "GL_QUERY_BY_REGION_WAIT"
	case FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER:
		name = "GL_FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER"
	case SYNC_FLAGS:
		name = "GL_SYNC_FLAGS"
	case MAX_GEOMETRY_OUTPUT_VERTICES:
		name = "GL_MAX_GEOMETRY_OUTPUT_VERTICES"
	case BACK:
		name = "GL_BACK"
	case CURRENT_PROGRAM:
		name = "GL_CURRENT_PROGRAM"
	case DISPATCH_INDIRECT_BUFFER:
		name = "GL_DISPATCH_INDIRECT_BUFFER"
	case SHADER_COMPILER:
		name = "GL_SHADER_COMPILER"
	case UNIFORM_BLOCK_DATA_SIZE:
		name = "GL_UNIFORM_BLOCK_DATA_SIZE"
	case UNSIGNED_SHORT_5_6_5:
		name = "GL_UNSIGNED_SHORT_5_6_5"
	case VERTEX_BINDING_OFFSET:
		name = "GL_VERTEX_BINDING_OFFSET"
	case QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION:
		name = "GL_QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION"
	case MAX_COMPUTE_SHADER_STORAGE_BLOCKS:
		name = "GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS"
	case QUERY_NO_WAIT:
		name = "GL_QUERY_NO_WAIT"
	case CLIP_DISTANCE6:
		name = "GL_CLIP_DISTANCE6"
	case COMPRESSED_RGBA_S3TC_DXT1_EXT:
		name = "GL_COMPRESSED_RGBA_S3TC_DXT1_EXT"
	case COMPRESSED_RED:
		name = "GL_COMPRESSED_RED"
	case MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS:
		name = "GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS"
	case RGB_INTEGER:
		name = "GL_RGB_INTEGER"
	case COMPRESSED_RED_RGTC1:
		name = "GL_COMPRESSED_RED_RGTC1"
	case DEPTH_COMPONENT16:
		name = "GL_DEPTH_COMPONENT16"
	case BACK_RIGHT:
		name = "GL_BACK_RIGHT"
	case NUM_EXTENSIONS:
		name = "GL_NUM_EXTENSIONS"
	case SHADER_IMAGE_ATOMIC:
		name = "GL_SHADER_IMAGE_ATOMIC"
	case SLUMINANCE_ALPHA_EXT:
		name = "GL_SLUMINANCE_ALPHA_EXT"
	case COMPRESSED_RGBA_S3TC_DXT5_EXT:
		name = "GL_COMPRESSED_RGBA_S3TC_DXT5_EXT"
	case TEXTURE23:
		name = "GL_TEXTURE23"
	case SAMPLER_1D_SHADOW:
		name = "GL_SAMPLER_1D_SHADOW"
	case SUBPIXEL_BITS:
		name = "GL_SUBPIXEL_BITS"
	case MAX_LABEL_LENGTH:
		name = "GL_MAX_LABEL_LENGTH"
	case TEXTURE_BINDING_3D:
		name = "GL_TEXTURE_BINDING_3D"
	case UNSIGNED_INT_10F_11F_11F_REV:
		name = "GL_UNSIGNED_INT_10F_11F_11F_REV"
	case STENCIL_VALUE_MASK:
		name = "GL_STENCIL_VALUE_MASK"
	case TOP_LEVEL_ARRAY_SIZE:
		name = "GL_TOP_LEVEL_ARRAY_SIZE"
	case SYNC_GPU_COMMANDS_COMPLETE:
		name = "GL_SYNC_GPU_COMMANDS_COMPLETE"
	case DONT_CARE:
		name = "GL_DONT_CARE"
	case UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER:
		name = "GL_UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER"
	case TEXTURE_MAG_FILTER:
		name = "GL_TEXTURE_MAG_FILTER"
	case UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY:
		name = "GL_UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY"
	case IMAGE_CLASS_2_X_8:
		name = "GL_IMAGE_CLASS_2_X_8"
	case SHADER_STORAGE_BUFFER_BINDING:
		name = "GL_SHADER_STORAGE_BUFFER_BINDING"
	case DITHER:
		name = "GL_DITHER"
	case INT_2_10_10_10_REV:
		name = "GL_INT_2_10_10_10_REV"
	case CLIP_DISTANCE1:
		name = "GL_CLIP_DISTANCE1"
	case TEXTURE25:
		name = "GL_TEXTURE25"
	case ACTIVE_RESOURCES:
		name = "GL_ACTIVE_RESOURCES"
	case MAX_UNIFORM_BLOCK_SIZE:
		name = "GL_MAX_UNIFORM_BLOCK_SIZE"
	case FRONT:
		name = "GL_FRONT"
	case NUM_SAMPLE_COUNTS:
		name = "GL_NUM_SAMPLE_COUNTS"
	case TOP_LEVEL_ARRAY_STRIDE:
		name = "GL_TOP_LEVEL_ARRAY_STRIDE"
	case PACK_COMPRESSED_BLOCK_SIZE:
		name = "GL_PACK_COMPRESSED_BLOCK_SIZE"
	case PIXEL_UNPACK_BUFFER_BINDING:
		name = "GL_PIXEL_UNPACK_BUFFER_BINDING"
	case MAX_COMPUTE_ATOMIC_COUNTERS:
		name = "GL_MAX_COMPUTE_ATOMIC_COUNTERS"
	case UNSIGNED_INT_VEC2:
		name = "GL_UNSIGNED_INT_VEC2"
	case TEXTURE19:
		name = "GL_TEXTURE19"
	case COMPRESSED_RGB:
		name = "GL_COMPRESSED_RGB"
	case COMPATIBLE_SUBROUTINES:
		name = "GL_COMPATIBLE_SUBROUTINES"
	case GEOMETRY_SUBROUTINE_UNIFORM:
		name = "GL_GEOMETRY_SUBROUTINE_UNIFORM"
	case TEXTURE1:
		name = "GL_TEXTURE1"
	case SAMPLER_CUBE_MAP_ARRAY:
		name = "GL_SAMPLER_CUBE_MAP_ARRAY"
	case TEXTURE_BINDING_1D_ARRAY:
		name = "GL_TEXTURE_BINDING_1D_ARRAY"
	case VIEWPORT_BOUNDS_RANGE:
		name = "GL_VIEWPORT_BOUNDS_RANGE"
	case ATTACHED_SHADERS:
		name = "GL_ATTACHED_SHADERS"
	case HALF_FLOAT:
		name = "GL_HALF_FLOAT"
	case COMPRESSED_R11_EAC:
		name = "GL_COMPRESSED_R11_EAC"
	case SAMPLER_1D_ARRAY:
		name = "GL_SAMPLER_1D_ARRAY"
	case TEXTURE_SHARED_SIZE:
		name = "GL_TEXTURE_SHARED_SIZE"
	case IMAGE_PIXEL_FORMAT:
		name = "GL_IMAGE_PIXEL_FORMAT"
	case MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS:
		name = "GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS"
	case MAX_GEOMETRY_TEXTURE_IMAGE_UNITS:
		name = "GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS"
	case IMAGE_2D_ARRAY:
		name = "GL_IMAGE_2D_ARRAY"
	case FRAMEBUFFER_ATTACHMENT_GREEN_SIZE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_GREEN_SIZE"
	case ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER:
		name = "GL_ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER"
	case TEXTURE8:
		name = "GL_TEXTURE8"
	case READ_PIXELS:
		name = "GL_READ_PIXELS"
	case MAX_RENDERBUFFER_SIZE:
		name = "GL_MAX_RENDERBUFFER_SIZE"
	case TEXTURE_MIN_LOD:
		name = "GL_TEXTURE_MIN_LOD"
	case COMPRESSED_SLUMINANCE_ALPHA_EXT:
		name = "GL_COMPRESSED_SLUMINANCE_ALPHA_EXT"
	case MAX_TEXTURE_IMAGE_UNITS:
		name = "GL_MAX_TEXTURE_IMAGE_UNITS"
	case BGRA_INTEGER:
		name = "GL_BGRA_INTEGER"
	case SRC1_COLOR:
		name = "GL_SRC1_COLOR"
	case TEXTURE_DEPTH_TYPE:
		name = "GL_TEXTURE_DEPTH_TYPE"
	case INT_IMAGE_2D_MULTISAMPLE:
		name = "GL_INT_IMAGE_2D_MULTISAMPLE"
	case UNSIGNED_INT_IMAGE_2D:
		name = "GL_UNSIGNED_INT_IMAGE_2D"
	case MAX_FRAGMENT_IMAGE_UNIFORMS:
		name = "GL_MAX_FRAGMENT_IMAGE_UNIFORMS"
	case MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS:
		name = "GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS"
	case INCR_WRAP:
		name = "GL_INCR_WRAP"
	case DEBUG_NEXT_LOGGED_MESSAGE_LENGTH:
		name = "GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH"
	case UNSIGNED_INT_VEC3:
		name = "GL_UNSIGNED_INT_VEC3"
	case SAMPLER_2D_RECT:
		name = "GL_SAMPLER_2D_RECT"
	case FLOAT_VEC4:
		name = "GL_FLOAT_VEC4"
	case MAX_ATOMIC_COUNTER_BUFFER_SIZE:
		name = "GL_MAX_ATOMIC_COUNTER_BUFFER_SIZE"
	case PATCH_DEFAULT_OUTER_LEVEL:
		name = "GL_PATCH_DEFAULT_OUTER_LEVEL"
	case MAX_VERTEX_ATTRIB_BINDINGS:
		name = "GL_MAX_VERTEX_ATTRIB_BINDINGS"
	case RENDERBUFFER_BLUE_SIZE:
		name = "GL_RENDERBUFFER_BLUE_SIZE"
	case MAX_SUBROUTINE_UNIFORM_LOCATIONS:
		name = "GL_MAX_SUBROUTINE_UNIFORM_LOCATIONS"
	case TEXTURE_SWIZZLE_B:
		name = "GL_TEXTURE_SWIZZLE_B"
	case CLAMP_READ_COLOR:
		name = "GL_CLAMP_READ_COLOR"
	case INTERLEAVED_ATTRIBS:
		name = "GL_INTERLEAVED_ATTRIBS"
	case RG:
		name = "GL_RG"
	case NOOP:
		name = "GL_NOOP"
	case UNSIGNED_SHORT_4_4_4_4_REV:
		name = "GL_UNSIGNED_SHORT_4_4_4_4_REV"
	case ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER:
		name = "GL_ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER"
	case TEXTURE9:
		name = "GL_TEXTURE9"
	case STENCIL_BACK_WRITEMASK:
		name = "GL_STENCIL_BACK_WRITEMASK"
	case SET:
		name = "GL_SET"
	case R16I:
		name = "GL_R16I"
	case TEXTURE0:
		name = "GL_TEXTURE0"
	case RGBA8_SNORM:
		name = "GL_RGBA8_SNORM"
	case MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS:
		name = "GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS"
	case COLOR_ATTACHMENT15:
		name = "GL_COLOR_ATTACHMENT15"
	case MAX_VERTEX_ATTRIBS:
		name = "GL_MAX_VERTEX_ATTRIBS"
	case DEPTH_COMPONENT32F:
		name = "GL_DEPTH_COMPONENT32F"
	case FRACTIONAL_EVEN:
		name = "GL_FRACTIONAL_EVEN"
	case SHADING_LANGUAGE_VERSION:
		name = "GL_SHADING_LANGUAGE_VERSION"
	case SHADER_BINARY_FORMATS:
		name = "GL_SHADER_BINARY_FORMATS"
	case UNSIGNED_INT_IMAGE_2D_RECT:
		name = "GL_UNSIGNED_INT_IMAGE_2D_RECT"
	case DRAW_BUFFER10:
		name = "GL_DRAW_BUFFER10"
	case MEDIUM_INT:
		name = "GL_MEDIUM_INT"
	case DEPTH_STENCIL:
		name = "GL_DEPTH_STENCIL"
	case READ_FRAMEBUFFER:
		name = "GL_READ_FRAMEBUFFER"
	case TEXTURE_BUFFER:
		name = "GL_TEXTURE_BUFFER"
	case CLEAR:
		name = "GL_CLEAR"
	case ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER:
		name = "GL_ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER"
	case MAX_COMBINED_TEXTURE_IMAGE_UNITS:
		name = "GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS"
	case DEBUG_TYPE_PORTABILITY:
		name = "GL_DEBUG_TYPE_PORTABILITY"
	case UNDEFINED_VERTEX:
		name = "GL_UNDEFINED_VERTEX"
	case DRAW_BUFFER6:
		name = "GL_DRAW_BUFFER6"
	case IMAGE_CLASS_10_10_10_2:
		name = "GL_IMAGE_CLASS_10_10_10_2"
	case TESS_GEN_POINT_MODE:
		name = "GL_TESS_GEN_POINT_MODE"
	case INT_IMAGE_2D:
		name = "GL_INT_IMAGE_2D"
	case FRAMEBUFFER_DEFAULT_SAMPLES:
		name = "GL_FRAMEBUFFER_DEFAULT_SAMPLES"
	case BUFFER_ACCESS:
		name = "GL_BUFFER_ACCESS"
	case ACTIVE_ATTRIBUTE_MAX_LENGTH:
		name = "GL_ACTIVE_ATTRIBUTE_MAX_LENGTH"
	case MIN_FRAGMENT_INTERPOLATION_OFFSET:
		name = "GL_MIN_FRAGMENT_INTERPOLATION_OFFSET"
	case ATOMIC_COUNTER_BUFFER_BINDING:
		name = "GL_ATOMIC_COUNTER_BUFFER_BINDING"
	case FRAMEBUFFER_INCOMPLETE_MULTISAMPLE:
		name = "GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE"
	case PIXEL_PACK_BUFFER:
		name = "GL_PIXEL_PACK_BUFFER"
	case STENCIL_INDEX8:
		name = "GL_STENCIL_INDEX8"
	case DEPTH_RENDERABLE:
		name = "GL_DEPTH_RENDERABLE"
	case DRAW_BUFFER3:
		name = "GL_DRAW_BUFFER3"
	case DEBUG_SOURCE_WINDOW_SYSTEM:
		name = "GL_DEBUG_SOURCE_WINDOW_SYSTEM"
	case IMAGE_3D:
		name = "GL_IMAGE_3D"
	case MAX_FRAGMENT_ATOMIC_COUNTERS:
		name = "GL_MAX_FRAGMENT_ATOMIC_COUNTERS"
	case RED_SNORM:
		name = "GL_RED_SNORM"
	case LINE_WIDTH:
		name = "GL_LINE_WIDTH"
	case VERTEX_ATTRIB_ARRAY_POINTER:
		name = "GL_VERTEX_ATTRIB_ARRAY_POINTER"
	case QUERY_RESULT:
		name = "GL_QUERY_RESULT"
	case MAX_VERTEX_ATTRIB_RELATIVE_OFFSET:
		name = "GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET"
	case VIEW_COMPATIBILITY_CLASS:
		name = "GL_VIEW_COMPATIBILITY_CLASS"
	case UNSIGNED_INT:
		name = "GL_UNSIGNED_INT"
	case DRAW_BUFFER2:
		name = "GL_DRAW_BUFFER2"
	case FLOAT_MAT2x3:
		name = "GL_FLOAT_MAT2x3"
	case UNSIGNED_INT_IMAGE_1D:
		name = "GL_UNSIGNED_INT_IMAGE_1D"
	case DEPTH_COMPONENT24:
		name = "GL_DEPTH_COMPONENT24"
	case PIXEL_PACK_BUFFER_BINDING:
		name = "GL_PIXEL_PACK_BUFFER_BINDING"
	case IMAGE_CUBE_MAP_ARRAY:
		name = "GL_IMAGE_CUBE_MAP_ARRAY"
	case STENCIL_INDEX16:
		name = "GL_STENCIL_INDEX16"
	case DEPTH_COMPONENT:
		name = "GL_DEPTH_COMPONENT"
	case TEXTURE_LOD_BIAS:
		name = "GL_TEXTURE_LOD_BIAS"
	case TEXTURE_SAMPLES:
		name = "GL_TEXTURE_SAMPLES"
	case COMPRESSED_RGBA:
		name = "GL_COMPRESSED_RGBA"
	case ALWAYS:
		name = "GL_ALWAYS"
	case MAX_TESS_CONTROL_OUTPUT_COMPONENTS:
		name = "GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS"
	case ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER:
		name = "GL_ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER"
	case RED:
		name = "GL_RED"
	case UNPACK_IMAGE_HEIGHT:
		name = "GL_UNPACK_IMAGE_HEIGHT"
	case UNPACK_LSB_FIRST:
		name = "GL_UNPACK_LSB_FIRST"
	case POLYGON_OFFSET_FACTOR:
		name = "GL_POLYGON_OFFSET_FACTOR"
	case BGR:
		name = "GL_BGR"
	case CLIP_DISTANCE4:
		name = "GL_CLIP_DISTANCE4"
	case POLYGON_OFFSET_UNITS:
		name = "GL_POLYGON_OFFSET_UNITS"
	case PROGRAM_BINARY_LENGTH:
		name = "GL_PROGRAM_BINARY_LENGTH"
	case RENDERBUFFER_BINDING:
		name = "GL_RENDERBUFFER_BINDING"
	case TEXTURE_RECTANGLE:
		name = "GL_TEXTURE_RECTANGLE"
	case FLOAT_MAT3x4:
		name = "GL_FLOAT_MAT3x4"
	case MAX_TESS_EVALUATION_INPUT_COMPONENTS:
		name = "GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS"
	case CLIP_DISTANCE5:
		name = "GL_CLIP_DISTANCE5"
	case ACTIVE_SUBROUTINES:
		name = "GL_ACTIVE_SUBROUTINES"
	case STATIC_DRAW:
		name = "GL_STATIC_DRAW"
	case IMAGE_CLASS_2_X_32:
		name = "GL_IMAGE_CLASS_2_X_32"
	case COMPRESSED_SLUMINANCE_EXT:
		name = "GL_COMPRESSED_SLUMINANCE_EXT"
	case BLEND_EQUATION_ALPHA:
		name = "GL_BLEND_EQUATION_ALPHA"
	case IMAGE_CUBE:
		name = "GL_IMAGE_CUBE"
	case SCISSOR_TEST:
		name = "GL_SCISSOR_TEST"
	case TEXTURE_DEPTH:
		name = "GL_TEXTURE_DEPTH"
	case SIGNED_NORMALIZED:
		name = "GL_SIGNED_NORMALIZED"
	case DRAW_BUFFER8:
		name = "GL_DRAW_BUFFER8"
	case REFERENCED_BY_TESS_EVALUATION_SHADER:
		name = "GL_REFERENCED_BY_TESS_EVALUATION_SHADER"
	case DEPTH24_STENCIL8:
		name = "GL_DEPTH24_STENCIL8"
	case IMAGE_2D_MULTISAMPLE:
		name = "GL_IMAGE_2D_MULTISAMPLE"
	case HIGH_INT:
		name = "GL_HIGH_INT"
	case UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY:
		name = "GL_UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY"
	case TEXTURE_IMAGE_FORMAT:
		name = "GL_TEXTURE_IMAGE_FORMAT"
	case READ_PIXELS_FORMAT:
		name = "GL_READ_PIXELS_FORMAT"
	case MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS:
		name = "GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS"
	case INT_VEC2:
		name = "GL_INT_VEC2"
	case COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
		name = "GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2"
	case COLOR_ATTACHMENT13:
		name = "GL_COLOR_ATTACHMENT13"
	case MAX_ELEMENTS_INDICES:
		name = "GL_MAX_ELEMENTS_INDICES"
	case COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT:
		name = "GL_COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT"
	case TIME_ELAPSED:
		name = "GL_TIME_ELAPSED"
	case INT_IMAGE_3D:
		name = "GL_INT_IMAGE_3D"
	case ONE_MINUS_SRC1_ALPHA:
		name = "GL_ONE_MINUS_SRC1_ALPHA"
	case PIXEL_UNPACK_BUFFER:
		name = "GL_PIXEL_UNPACK_BUFFER"
	case MAX_SERVER_WAIT_TIMEOUT:
		name = "GL_MAX_SERVER_WAIT_TIMEOUT"
	case COMPRESSED_RG:
		name = "GL_COMPRESSED_RG"
	case TRIANGLE_STRIP_ADJACENCY:
		name = "GL_TRIANGLE_STRIP_ADJACENCY"
	case BLEND_DST_ALPHA:
		name = "GL_BLEND_DST_ALPHA"
	case MAX_GEOMETRY_SHADER_STORAGE_BLOCKS:
		name = "GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS"
	case UNIFORM_BLOCK_NAME_LENGTH:
		name = "GL_UNIFORM_BLOCK_NAME_LENGTH"
	case TEXTURE_COMPRESSED_BLOCK_HEIGHT:
		name = "GL_TEXTURE_COMPRESSED_BLOCK_HEIGHT"
	case INTERNALFORMAT_STENCIL_SIZE:
		name = "GL_INTERNALFORMAT_STENCIL_SIZE"
	case IMAGE_1D_ARRAY:
		name = "GL_IMAGE_1D_ARRAY"
	case LEQUAL:
		name = "GL_LEQUAL"
	case PACK_LSB_FIRST:
		name = "GL_PACK_LSB_FIRST"
	case UNSIGNED_SHORT_4_4_4_4:
		name = "GL_UNSIGNED_SHORT_4_4_4_4"
	case MAX_FRAMEBUFFER_LAYERS:
		name = "GL_MAX_FRAMEBUFFER_LAYERS"
	case TEXTURE_IMMUTABLE_FORMAT:
		name = "GL_TEXTURE_IMMUTABLE_FORMAT"
	case INT_SAMPLER_2D_MULTISAMPLE:
		name = "GL_INT_SAMPLER_2D_MULTISAMPLE"
	case MAX_IMAGE_SAMPLES:
		name = "GL_MAX_IMAGE_SAMPLES"
	case TEXTURE_CUBE_MAP_POSITIVE_Y:
		name = "GL_TEXTURE_CUBE_MAP_POSITIVE_Y"
	case MAX_3D_TEXTURE_SIZE:
		name = "GL_MAX_3D_TEXTURE_SIZE"
	case ATOMIC_COUNTER_BUFFER_SIZE:
		name = "GL_ATOMIC_COUNTER_BUFFER_SIZE"
	case TRANSFORM_FEEDBACK_BARRIER_BIT:
		name = "GL_TRANSFORM_FEEDBACK_BARRIER_BIT"
	case STENCIL_CLEAR_VALUE:
		name = "GL_STENCIL_CLEAR_VALUE"
	case TEXTURE_COMPRESSED_BLOCK_SIZE:
		name = "GL_TEXTURE_COMPRESSED_BLOCK_SIZE"
	case ALIASED_LINE_WIDTH_RANGE:
		name = "GL_ALIASED_LINE_WIDTH_RANGE"
	case STATIC_READ:
		name = "GL_STATIC_READ"
	case CAVEAT_SUPPORT:
		name = "GL_CAVEAT_SUPPORT"
	case UNIFORM_NAME_LENGTH:
		name = "GL_UNIFORM_NAME_LENGTH"
	case RGBA32F:
		name = "GL_RGBA32F"
	case MAX_INTEGER_SAMPLES:
		name = "GL_MAX_INTEGER_SAMPLES"
	case STREAM_READ:
		name = "GL_STREAM_READ"
	case LINE_SMOOTH:
		name = "GL_LINE_SMOOTH"
	case RGBA4:
		name = "GL_RGBA4"
	case CLIP_DISTANCE7:
		name = "GL_CLIP_DISTANCE7"
	case RG16_SNORM:
		name = "GL_RG16_SNORM"
	case VERSION:
		name = "GL_VERSION"
	case FRAMEBUFFER_ATTACHMENT_OBJECT_NAME:
		name = "GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME"
	case ARRAY_BUFFER_BINDING:
		name = "GL_ARRAY_BUFFER_BINDING"
	case FRAGMENT_SUBROUTINE_UNIFORM:
		name = "GL_FRAGMENT_SUBROUTINE_UNIFORM"
	case NEAREST:
		name = "GL_NEAREST"
	case PRIMITIVE_RESTART:
		name = "GL_PRIMITIVE_RESTART"
	case INT_SAMPLER_1D:
		name = "GL_INT_SAMPLER_1D"
	case TEXTURE_SHADOW:
		name = "GL_TEXTURE_SHADOW"
	case STENCIL_BACK_REF:
		name = "GL_STENCIL_BACK_REF"
	case COLOR_ATTACHMENT10:
		name = "GL_COLOR_ATTACHMENT10"
	case RG8:
		name = "GL_RG8"
	case LINEAR_MIPMAP_LINEAR:
		name = "GL_LINEAR_MIPMAP_LINEAR"
	case RENDERBUFFER_GREEN_SIZE:
		name = "GL_RENDERBUFFER_GREEN_SIZE"
	case PROXY_TEXTURE_2D:
		name = "GL_PROXY_TEXTURE_2D"
	case LINE_SMOOTH_HINT:
		name = "GL_LINE_SMOOTH_HINT"
	case MAX_NAME_LENGTH:
		name = "GL_MAX_NAME_LENGTH"
	default:
		name = fmt.Sprintf("GL_ENUM_%v", enum)
	}
	return
}
