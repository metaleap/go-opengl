package gl

import "fmt"

const (
	FRAGMENT_SHADER                                            = 0x8B30
	TEXTURE_LOD_BIAS                                           = 0x8501
	RG32UI                                                     = 0x823C
	IMAGE_CLASS_2_X_32                                         = 0x82BA
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              = 0x8DD7
	IMAGE_BINDING_NAME                                         = 0x8F3A
	PROXY_TEXTURE_CUBE_MAP                                     = 0x851B
	MAX_COMPUTE_UNIFORM_BLOCKS                                 = 0x91BB
	MAX_COMBINED_IMAGE_UNIFORMS                                = 0x90CF
	ACTIVE_PROGRAM                                             = 0x8259
	COLOR_RENDERABLE                                           = 0x8286
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             = 0x886A
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            = 0x8A34
	BUFFER                                                     = 0x82E0
	UNIFORM_MATRIX_STRIDE                                      = 0x8A3D
	UNSIGNED_SHORT_5_5_5_1                                     = 0x8034
	INTERNALFORMAT_GREEN_SIZE                                  = 0x8272
	RGBA4                                                      = 0x8056
	SAMPLE_COVERAGE_VALUE                                      = 0x80AA
	LOGIC_OP_MODE                                              = 0x0BF0
	BUFFER_DATA_SIZE                                           = 0x9303
	LESS                                                       = 0x0201
	ONE_MINUS_DST_ALPHA                                        = 0x0305
	STENCIL_RENDERABLE                                         = 0x8288
	SAMPLE_ALPHA_TO_COVERAGE                                   = 0x809E
	STENCIL_REF                                                = 0x0B97
	COPY_INVERTED                                              = 0x150C
	UNSIGNED_INT_SAMPLER_2D                                    = 0x8DD2
	POINT_SIZE                                                 = 0x0B11
	MAX_DEPTH_TEXTURE_SAMPLES                                  = 0x910F
	INVALID_INDEX                                              = 0xFFFFFFFF
	RGBA32I                                                    = 0x8D82
	SAMPLER_BINDING                                            = 0x8919
	AND_REVERSE                                                = 0x1502
	DONT_CARE                                                  = 0x1100
	COLOR_LOGIC_OP                                             = 0x0BF2
	INTERNALFORMAT_STENCIL_TYPE                                = 0x827D
	MAX_VIEWPORT_DIMS                                          = 0x0D3A
	DEBUG_SEVERITY_LOW                                         = 0x9148
	CAVEAT_SUPPORT                                             = 0x82B8
	DOUBLE_MAT3                                                = 0x8F47
	VERTEX_ATTRIB_ARRAY_ENABLED                                = 0x8622
	UNSIGNED_INT_IMAGE_CUBE                                    = 0x9066
	INT_IMAGE_1D_ARRAY                                         = 0x905D
	FRAMEBUFFER_UNDEFINED                                      = 0x8219
	TESS_EVALUATION_TEXTURE                                    = 0x829D
	WRITE_ONLY                                                 = 0x88B9
	MAX_CLIP_DISTANCES                                         = 0x0D32
	MAX_HEIGHT                                                 = 0x827F
	COMPRESSED_R11_EAC                                         = 0x9270
	TEXTURE_COMPRESSION_HINT                                   = 0x84EF
	TEXTURE_BUFFER_DATA_STORE_BINDING                          = 0x8C2D
	STENCIL_BACK_PASS_DEPTH_PASS                               = 0x8803
	ALL_BARRIER_BITS                                           = 0xFFFFFFFF
	SAMPLER_2D_ARRAY                                           = 0x8DC1
	TRANSFORM_FEEDBACK_BUFFER                                  = 0x8C8E
	TEXTURE7                                                   = 0x84C7
	GREEN                                                      = 0x1904
	DRAW_BUFFER9                                               = 0x882E
	TIMESTAMP                                                  = 0x8E28
	CCW                                                        = 0x0901
	POLYGON_OFFSET_UNITS                                       = 0x2A00
	UNPACK_COMPRESSED_BLOCK_DEPTH                              = 0x9129
	QUERY_COUNTER_BITS                                         = 0x8864
	STENCIL_BACK_FAIL                                          = 0x8801
	STENCIL_BACK_REF                                           = 0x8CA3
	SYNC_STATUS                                                = 0x9114
	LINES_ADJACENCY                                            = 0x000A
	RGB16_SNORM                                                = 0x8F9A
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           = 0x9279
	RED_INTEGER                                                = 0x8D94
	SRC1_ALPHA                                                 = 0x8589
	MAX_PROGRAM_TEXEL_OFFSET                                   = 0x8905
	SRGB_ALPHA                                                 = 0x8C42
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    = 0x906C
	TEXTURE_CUBE_MAP                                           = 0x8513
	IMAGE_CLASS_11_11_10                                       = 0x82C2
	ACTIVE_UNIFORM_BLOCKS                                      = 0x8A36
	DEPTH_COMPONENT32F                                         = 0x8CAC
	INT_IMAGE_1D                                               = 0x9057
	UNIFORM_BUFFER_BINDING                                     = 0x8A28
	PROXY_TEXTURE_2D_ARRAY                                     = 0x8C1B
	TEXTURE4                                                   = 0x84C4
	COLOR_WRITEMASK                                            = 0x0C23
	INT_IMAGE_2D_ARRAY                                         = 0x905E
	DEPTH_ATTACHMENT                                           = 0x8D00
	VERTEX_SHADER_BIT                                          = 0x00000001
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                = 0x8A46
	BUFFER_SIZE                                                = 0x8764
	UNSIGNED_INT_ATOMIC_COUNTER                                = 0x92DB
	TEXTURE_BUFFER                                             = 0x8C2A
	IMAGE_PIXEL_FORMAT                                         = 0x82A9
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT                        = 0x8C4D
	RENDERBUFFER_WIDTH                                         = 0x8D42
	IS_PER_PATCH                                               = 0x92E7
	DST_COLOR                                                  = 0x0306
	BUFFER_MAP_OFFSET                                          = 0x9121
	DRAW_BUFFER1                                               = 0x8826
	CLAMP_TO_EDGE                                              = 0x812F
	SAMPLES_PASSED                                             = 0x8914
	QUERY_BY_REGION_WAIT                                       = 0x8E15
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             = 0x9124
	TEXTURE_BLUE_SIZE                                          = 0x805E
	TEXTURE_BINDING_RECTANGLE                                  = 0x84F6
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           = 0x910C
	DECR_WRAP                                                  = 0x8508
	NUM_SAMPLE_COUNTS                                          = 0x9380
	FLOAT_VEC2                                                 = 0x8B50
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 = 0x8C80
	GEOMETRY_SHADER_BIT                                        = 0x00000004
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            = 0x84F0
	COMPRESSED_RGBA_S3TC_DXT1_EXT                              = 0x83F1
	INTERLEAVED_ATTRIBS                                        = 0x8C8C
	TEXTURE6                                                   = 0x84C6
	RGBA2                                                      = 0x8055
	MAX_TEXTURE_LOD_BIAS                                       = 0x84FD
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 = 0x90EC
	OR                                                         = 0x1507
	TEXTURE_COMPRESSED_IMAGE_SIZE                              = 0x86A0
	STENCIL_BACK_VALUE_MASK                                    = 0x8CA4
	POINT_FADE_THRESHOLD_SIZE                                  = 0x8128
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         = 0x92F0
	BOOL                                                       = 0x8B56
	STACK_UNDERFLOW                                            = 0x0504
	RGB32F                                                     = 0x8815
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     = 0x8E80
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          = 0x8213
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          = 0x8C8F
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            = 0x8DDF
	GEOMETRY_SUBROUTINE                                        = 0x92EB
	VALIDATE_STATUS                                            = 0x8B83
	FULL_SUPPORT                                               = 0x82B7
	RENDERER                                                   = 0x1F01
	UNSIGNED_INT_SAMPLER_CUBE                                  = 0x8DD4
	UNSIGNED_INT_10F_11F_11F_REV                               = 0x8C3B
	INT_SAMPLER_BUFFER                                         = 0x8DD0
	BUFFER_UPDATE_BARRIER_BIT                                  = 0x00000200
	MAX_UNIFORM_LOCATIONS                                      = 0x826E
	OUT_OF_MEMORY                                              = 0x0505
	TESS_EVALUATION_SHADER                                     = 0x8E87
	FRAMEBUFFER                                                = 0x8D40
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         = 0x8CD0
	TEXTURE_RED_SIZE                                           = 0x805C
	RG32I                                                      = 0x823B
	UNSIGNED_INT_2_10_10_10_REV                                = 0x8368
	VERTEX_ARRAY_BINDING                                       = 0x85B5
	IMAGE_CLASS_1_X_16                                         = 0x82BE
	MAX_FRAGMENT_INPUT_COMPONENTS                              = 0x9125
	NEAREST                                                    = 0x2600
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       = 0x82AE
	MIN_MAP_BUFFER_ALIGNMENT                                   = 0x90BC
	TEXTURE2                                                   = 0x84C2
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    = 0x92CD
	INTERNALFORMAT_RED_TYPE                                    = 0x8278
	RGB8I                                                      = 0x8D8F
	PROXY_TEXTURE_RECTANGLE                                    = 0x84F7
	STENCIL_INDEX8                                             = 0x8D48
	RGB5                                                       = 0x8050
	GEOMETRY_SUBROUTINE_UNIFORM                                = 0x92F1
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         = 0x8264
	DISPATCH_INDIRECT_BUFFER_BINDING                           = 0x90EF
	COPY_READ_BUFFER_BINDING                                   = 0x8F36
	DEPTH_RANGE                                                = 0x0B70
	UNIFORM                                                    = 0x92E1
	TEXTURE_GATHER                                             = 0x82A2
	CONDITION_SATISFIED                                        = 0x911C
	UNSIGNED_INT_10_10_10_2                                    = 0x8036
	REFERENCED_BY_GEOMETRY_SHADER                              = 0x9309
	NUM_SHADING_LANGUAGE_VERSIONS                              = 0x82E9
	SRGB8_ALPHA8_EXT                                           = 0x8C43
	RG16F                                                      = 0x822F
	MAX_COMBINED_ATOMIC_COUNTERS                               = 0x92D7
	PROGRAM_BINARY_RETRIEVABLE_HINT                            = 0x8257
	MAX_COMPUTE_LOCAL_INVOCATIONS                              = 0x90EB
	MAX_IMAGE_SAMPLES                                          = 0x906D
	UPPER_LEFT                                                 = 0x8CA2
	PRIMITIVE_RESTART_FIXED_INDEX                              = 0x8D69
	DRAW_BUFFER                                                = 0x0C01
	VIEW_CLASS_RGTC2_RG                                        = 0x82D1
	RENDERBUFFER_INTERNAL_FORMAT                               = 0x8D44
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    = 0x92C8
	TEXTURE_CUBE_MAP_NEGATIVE_X                                = 0x8516
	PROXY_TEXTURE_2D                                           = 0x8064
	INTERNALFORMAT_SUPPORTED                                   = 0x826F
	RASTERIZER_DISCARD                                         = 0x8C89
	CLEAR                                                      = 0x1500
	DYNAMIC_COPY                                               = 0x88EA
	COMPRESSED_SIGNED_RG11_EAC                                 = 0x9273
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         = 0x9103
	TEXTURE_1D_ARRAY                                           = 0x8C18
	SLUMINANCE8_ALPHA8_EXT                                     = 0x8C45
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              = 0x8C8A
	MAX_UNIFORM_BLOCK_SIZE                                     = 0x8A30
	STENCIL_INDEX                                              = 0x1901
	TEXTURE_MAG_FILTER                                         = 0x2800
	INT_IMAGE_2D                                               = 0x9058
	ONE_MINUS_SRC_COLOR                                        = 0x0301
	TEXTURE14                                                  = 0x84CE
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              = 0x8F39
	STENCIL_PASS_DEPTH_PASS                                    = 0x0B96
	VIEW_CLASS_48_BITS                                         = 0x82C7
	DEPTH                                                      = 0x1801
	DRAW_BUFFER6                                               = 0x882B
	INVALID_ENUM                                               = 0x0500
	RG_INTEGER                                                 = 0x8228
	TIMEOUT_EXPIRED                                            = 0x911B
	FRONT_FACE                                                 = 0x0B46
	INT_SAMPLER_1D                                             = 0x8DC9
	TEXTURE_SWIZZLE_G                                          = 0x8E43
	INT_SAMPLER_2D                                             = 0x8DCA
	DEBUG_TYPE_OTHER                                           = 0x8251
	SLUMINANCE_ALPHA_EXT                                       = 0x8C44
	LINEAR_MIPMAP_NEAREST                                      = 0x2701
	MAX_VERTEX_ATOMIC_COUNTERS                                 = 0x92D2
	COLOR_ATTACHMENT14                                         = 0x8CEE
	SRGB8_ALPHA8                                               = 0x8C43
	BOOL_VEC2                                                  = 0x8B57
	TEXTURE_GREEN_TYPE                                         = 0x8C11
	DECR                                                       = 0x1E03
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          = 0x8CD6
	TRANSFORM_FEEDBACK_BUFFER_START                            = 0x8C84
	RG16_SNORM                                                 = 0x8F99
	ELEMENT_ARRAY_BUFFER                                       = 0x8893
	TRIANGLE_STRIP_ADJACENCY                                   = 0x000D
	INTERNALFORMAT_SHARED_SIZE                                 = 0x8277
	TEXTURE_COMPARE_MODE                                       = 0x884C
	NO_ERROR                                                   = 0
	COMPUTE_LOCAL_WORK_SIZE                                    = 0x8267
	CLIP_DISTANCE7                                             = 0x3007
	DEBUG_OUTPUT_SYNCHRONOUS                                   = 0x8242
	TEXTURE_MIN_LOD                                            = 0x813A
	MAP_UNSYNCHRONIZED_BIT                                     = 0x0020
	ACTIVE_SUBROUTINE_UNIFORMS                                 = 0x8DE6
	ANY_SAMPLES_PASSED                                         = 0x8C2F
	SHADER_STORAGE_BUFFER_BINDING                              = 0x90D3
	VERTEX_TEXTURE                                             = 0x829B
	QUERY_BY_REGION_NO_WAIT                                    = 0x8E16
	TESS_EVALUATION_SHADER_BIT                                 = 0x00000010
	ATTACHED_SHADERS                                           = 0x8B85
	UNIFORM_BLOCK_INDEX                                        = 0x8A3A
	ZERO                                                       = 0
	COMPRESSED_RGBA_S3TC_DXT3_EXT                              = 0x83F2
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        = 0x92D0
	ATOMIC_COUNTER_BUFFER_BINDING                              = 0x92C1
	SAMPLER_2D_ARRAY_SHADOW                                    = 0x8DC4
	INVERT                                                     = 0x150A
	BGRA                                                       = 0x80E1
	NOR                                                        = 0x1508
	GET_TEXTURE_IMAGE_TYPE                                     = 0x8292
	TEXTURE_VIEW_MIN_LEVEL                                     = 0x82DB
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       = 0x8A35
	ATOMIC_COUNTER_BUFFER_SIZE                                 = 0x92C3
	DRAW_BUFFER14                                              = 0x8833
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  = 0x8A44
	R8_SNORM                                                   = 0x8F94
	SAMPLER_2D_MULTISAMPLE                                     = 0x9108
	POLYGON_MODE                                               = 0x0B40
	UNSIGNED_NORMALIZED                                        = 0x8C17
	IMAGE_CLASS_10_10_10_2                                     = 0x82C3
	DRAW_INDIRECT_BUFFER_BINDING                               = 0x8F43
	COMPRESSED_RG_RGTC2                                        = 0x8DBD
	PROGRAM_BINARY_LENGTH                                      = 0x8741
	TESS_GEN_POINT_MODE                                        = 0x8E79
	UNIFORM_BUFFER_SIZE                                        = 0x8A2A
	COMMAND_BARRIER_BIT                                        = 0x00000040
	READ_FRAMEBUFFER                                           = 0x8CA8
	MAX_RECTANGLE_TEXTURE_SIZE                                 = 0x84F8
	UNIFORM_SIZE                                               = 0x8A38
	IS_ROW_MAJOR                                               = 0x9300
	TEXTURE_GATHER_SHADOW                                      = 0x82A3
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        = 0x92CA
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             = 0x8262
	TESS_CONTROL_SUBROUTINE                                    = 0x92E9
	RGBA_INTEGER                                               = 0x8D99
	TEXTURE_UPDATE_BARRIER_BIT                                 = 0x00000100
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          = 0x906B
	STEREO                                                     = 0x0C33
	FRAGMENT_SHADER_DERIVATIVE_HINT                            = 0x8B8B
	INT_IMAGE_2D_MULTISAMPLE                                   = 0x9060
	RGB32I                                                     = 0x8D83
	MAX_NUM_COMPATIBLE_SUBROUTINES                             = 0x92F8
	TEXTURE_BORDER_COLOR                                       = 0x1004
	UNIFORM_BUFFER                                             = 0x8A11
	RG8UI                                                      = 0x8238
	SIGNED_NORMALIZED                                          = 0x8F9C
	INT_2_10_10_10_REV                                         = 0x8D9F
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           = 0x90D6
	MAX_VARYING_COMPONENTS                                     = 0x8B4B
	FRONT                                                      = 0x0404
	NUM_ACTIVE_VARIABLES                                       = 0x9304
	VERSION                                                    = 0x1F02
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          = 0x886C
	DEPTH_COMPONENT32                                          = 0x81A7
	IMAGE_TEXEL_SIZE                                           = 0x82A7
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       = 0x9105
	TEXTURE_SWIZZLE_R                                          = 0x8E42
	SRGB_READ                                                  = 0x8297
	COLOR_BUFFER_BIT                                           = 0x00004000
	FRAMEBUFFER_BARRIER_BIT                                    = 0x00000400
	TEXTURE_ALPHA_TYPE                                         = 0x8C13
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                = 0x8B8A
	RGBA12                                                     = 0x805A
	IMAGE_2D_RECT                                              = 0x904F
	IMAGE_CLASS_4_X_32                                         = 0x82B9
	FLOAT_MAT3x2                                               = 0x8B67
	MAX_UNIFORM_BUFFER_BINDINGS                                = 0x8A2F
	MAX_VERTEX_IMAGE_UNIFORMS                                  = 0x90CA
	CLIP_DISTANCE0                                             = 0x3000
	RENDERBUFFER_DEPTH_SIZE                                    = 0x8D54
	COMPRESSED_TEXTURE_FORMATS                                 = 0x86A3
	TRIANGLES                                                  = 0x0004
	TEXTURE23                                                  = 0x84D7
	TEXTURE22                                                  = 0x84D6
	COMPRESSED_SRGB_ALPHA_EXT                                  = 0x8C49
	TESS_CONTROL_TEXTURE                                       = 0x829C
	COMPRESSED_RED_RGTC1                                       = 0x8DBB
	DRAW_BUFFER7                                               = 0x882C
	LINE_WIDTH_GRANULARITY                                     = 0x0B23
	TEXTURE_CUBE_MAP_POSITIVE_X                                = 0x8515
	POLYGON_SMOOTH                                             = 0x0B41
	GEQUAL                                                     = 0x0206
	RENDERBUFFER_BLUE_SIZE                                     = 0x8D52
	COMPRESSED_RGBA_S3TC_DXT5_EXT                              = 0x83F3
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   = 0x8A33
	PRIMITIVE_RESTART                                          = 0x8F9D
	SYNC_GPU_COMMANDS_COMPLETE                                 = 0x9117
	MAX_GEOMETRY_SHADER_INVOCATIONS                            = 0x8E5A
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              = 0x900D
	TEXTURE26                                                  = 0x84DA
	BLEND_SRC_ALPHA                                            = 0x80CB
	R32F                                                       = 0x822E
	IMAGE_1D                                                   = 0x904C
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        = 0x92CB
	TEXTURE_COMPRESSED_BLOCK_SIZE                              = 0x82B3
	AUTO_GENERATE_MIPMAP                                       = 0x8295
	TEXTURE_BINDING_CUBE_MAP                                   = 0x8514
	VIEW_CLASS_S3TC_DXT5_RGBA                                  = 0x82CF
	MAX_CUBE_MAP_TEXTURE_SIZE                                  = 0x851C
	STENCIL_FAIL                                               = 0x0B94
	INT_SAMPLER_CUBE                                           = 0x8DCC
	ISOLINES                                                   = 0x8E7A
	UNIFORM_BLOCK_DATA_SIZE                                    = 0x8A40
	COMPRESSED_RGBA8_ETC2_EAC                                  = 0x9278
	TEXTURE_WRAP_S                                             = 0x2802
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     = 0x8A31
	DEBUG_TYPE_PORTABILITY                                     = 0x824F
	EQUIV                                                      = 0x1509
	TEXTURE_SWIZZLE_RGBA                                       = 0x8E46
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       = 0x8E49
	BUFFER_MAPPED                                              = 0x88BC
	STENCIL                                                    = 0x1802
	PACK_SKIP_IMAGES                                           = 0x806B
	TRIANGLE_STRIP                                             = 0x0005
	VIEW_CLASS_S3TC_DXT3_RGBA                                  = 0x82CE
	TEXTURE_BINDING_2D_MULTISAMPLE                             = 0x9104
	INTERNALFORMAT_DEPTH_TYPE                                  = 0x827C
	TEXTURE30                                                  = 0x84DE
	VERTEX_ATTRIB_ARRAY_STRIDE                                 = 0x8624
	VIEW_CLASS_16_BITS                                         = 0x82CA
	MAX_SHADER_STORAGE_BLOCK_SIZE                              = 0x90DE
	HIGH_FLOAT                                                 = 0x8DF2
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             = 0x9107
	MAX_VERTEX_ATTRIBS                                         = 0x8869
	DEPTH_CLEAR_VALUE                                          = 0x0B73
	MAX_FRAMEBUFFER_WIDTH                                      = 0x9315
	MAX_3D_TEXTURE_SIZE                                        = 0x8073
	FALSE                                                      = 0
	VERTEX_SUBROUTINE                                          = 0x92E8
	UNSIGNED_BYTE_2_3_3_REV                                    = 0x8362
	PACK_SKIP_ROWS                                             = 0x0D03
	LINEAR                                                     = 0x2601
	RGBA                                                       = 0x1908
	IMAGE_COMPATIBILITY_CLASS                                  = 0x82A8
	KEEP                                                       = 0x1E00
	COMPUTE_SUBROUTINE_UNIFORM                                 = 0x92F3
	HIGH_INT                                                   = 0x8DF5
	TEXTURE8                                                   = 0x84C8
	TEXTURE_DEPTH_TYPE                                         = 0x8C16
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  = 0x9277
	COMPRESSED_RGB                                             = 0x84ED
	TEXTURE_BINDING_3D                                         = 0x806A
	FRAGMENT_SUBROUTINE_UNIFORM                                = 0x92F2
	ATOMIC_COUNTER_BUFFER                                      = 0x92C0
	TEXTURE_IMAGE_TYPE                                         = 0x8290
	INT_SAMPLER_2D_RECT                                        = 0x8DCD
	MAX_VARYING_VECTORS                                        = 0x8DFC
	IMAGE_3D                                                   = 0x904E
	DOUBLE_MAT2                                                = 0x8F46
	ONE_MINUS_DST_COLOR                                        = 0x0307
	VERTEX_ATTRIB_ARRAY_LONG                                   = 0x874E
	AND                                                        = 0x1501
	GEOMETRY_SHADER_INVOCATIONS                                = 0x887F
	TYPE                                                       = 0x92FA
	SRC_COLOR                                                  = 0x0300
	MAX_SUBROUTINES                                            = 0x8DE7
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         = 0x90DD
	ATOMIC_COUNTER_BUFFER_INDEX                                = 0x9301
	RGB10                                                      = 0x8052
	POLYGON_SMOOTH_HINT                                        = 0x0C53
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         = 0x90DC
	FIXED                                                      = 0x140C
	RG                                                         = 0x8227
	SRGB_EXT                                                   = 0x8C40
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            = 0x82B2
	INTERNALFORMAT_BLUE_TYPE                                   = 0x827A
	UNSIGNED_INT_VEC3                                          = 0x8DC7
	VIEW_CLASS_8_BITS                                          = 0x82CB
	MAX_DEBUG_GROUP_STACK_DEPTH                                = 0x826C
	SAMPLE_COVERAGE                                            = 0x80A0
	NUM_COMPATIBLE_SUBROUTINES                                 = 0x8E4A
	MAX_WIDTH                                                  = 0x827E
	ALL_SHADER_BITS                                            = 0xFFFFFFFF
	READ_PIXELS_FORMAT                                         = 0x828D
	SAMPLER                                                    = 0x82E6
	DISPATCH_INDIRECT_BUFFER                                   = 0x90EE
	MAX_GEOMETRY_INPUT_COMPONENTS                              = 0x9123
	FLOAT_32_UNSIGNED_INT_24_8_REV                             = 0x8DAD
	NAME_LENGTH                                                = 0x92F9
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      = 0x8C88
	TEXTURE_COMPRESSED                                         = 0x86A1
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         = 0x84F1
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        = 0x8E7F
	RG8I                                                       = 0x8237
	SRC1_COLOR                                                 = 0x88F9
	DOUBLE_MAT4x2                                              = 0x8F4D
	LINE                                                       = 0x1B01
	SYNC_FLUSH_COMMANDS_BIT                                    = 0x00000001
	DEBUG_SOURCE_APPLICATION                                   = 0x824A
	SAMPLE_SHADING                                             = 0x8C36
	TEXTURE_VIEW_NUM_LEVELS                                    = 0x82DC
	ACTIVE_RESOURCES                                           = 0x92F5
	DEBUG_TYPE_PUSH_GROUP                                      = 0x8269
	UNIFORM_NAME_LENGTH                                        = 0x8A39
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           = 0x8C29
	FLOAT_MAT2                                                 = 0x8B5A
	BLEND_SRC_RGB                                              = 0x80C9
	COMPUTE_SHADER                                             = 0x91B9
	TEXTURE27                                                  = 0x84DB
	FIXED_ONLY                                                 = 0x891D
	TRIANGLES_ADJACENCY                                        = 0x000C
	R16_SNORM                                                  = 0x8F98
	DOUBLE_MAT4                                                = 0x8F48
	STREAM_READ                                                = 0x88E1
	FRONT_RIGHT                                                = 0x0401
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             = 0x8E70
	EQUAL                                                      = 0x0202
	UNSIGNED_SHORT_4_4_4_4_REV                                 = 0x8365
	MAX_DEPTH                                                  = 0x8280
	RGBA32UI                                                   = 0x8D70
	MAX_TESS_GEN_LEVEL                                         = 0x8E7E
	TEXTURE_SHADOW                                             = 0x82A1
	RENDERBUFFER                                               = 0x8D41
	DEBUG_TYPE_POP_GROUP                                       = 0x826A
	MAJOR_VERSION                                              = 0x821B
	COMPRESSED_SIGNED_R11_EAC                                  = 0x9271
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 = 0x9314
	TEXTURE_MAX_LEVEL                                          = 0x813D
	IMAGE_1D_ARRAY                                             = 0x9052
	STENCIL_INDEX16                                            = 0x8D49
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         = 0x8CDB
	TEXTURE17                                                  = 0x84D1
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            = 0x92C4
	RG8                                                        = 0x822B
	RG_SNORM                                                   = 0x8F91
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              = 0x92D9
	UNPACK_LSB_FIRST                                           = 0x0CF1
	SEPARATE_ATTRIBS                                           = 0x8C8D
	STENCIL_BUFFER_BIT                                         = 0x00000400
	SRGB_ALPHA_EXT                                             = 0x8C42
	SMOOTH_LINE_WIDTH_RANGE                                    = 0x0B22
	VIEW_CLASS_24_BITS                                         = 0x82C9
	MAX_FRAMEBUFFER_HEIGHT                                     = 0x9316
	RGBA8                                                      = 0x8058
	RGB_SNORM                                                  = 0x8F92
	SCISSOR_BOX                                                = 0x0C10
	DRAW_BUFFER10                                              = 0x882F
	TEXTURE19                                                  = 0x84D3
	TEXTURE_RED_TYPE                                           = 0x8C10
	DEBUG_CALLBACK_USER_PARAM                                  = 0x8245
	MAX_VERTEX_UNIFORM_BLOCKS                                  = 0x8A2B
	IMAGE_BINDING_LAYER                                        = 0x8F3D
	UNIFORM_OFFSET                                             = 0x8A3B
	DRAW_BUFFER13                                              = 0x8832
	RGB5_A1                                                    = 0x8057
	POINT_SIZE_RANGE                                           = 0x0B12
	PROVOKING_VERTEX                                           = 0x8E4F
	SHADER_STORAGE_BARRIER_BIT                                 = 0x2000
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                = 0x851A
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         = 0x8CDC
	CONTEXT_PROFILE_MASK                                       = 0x9126
	SAMPLER_1D_ARRAY_SHADOW                                    = 0x8DC3
	DEBUG_TYPE_MARKER                                          = 0x8268
	TEXTURE_WRAP_T                                             = 0x2803
	TEXTURE_BUFFER_OFFSET                                      = 0x919D
	COMPUTE_SUBROUTINE                                         = 0x92ED
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       = MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS
	MAX_COMPUTE_IMAGE_UNIFORMS                                 = 0x91BD
	STENCIL_ATTACHMENT                                         = 0x8D20
	DEPTH_STENCIL_TEXTURE_MODE                                 = 0x90EA
	UNDEFINED_VERTEX                                           = 0x8260
	COLOR_ATTACHMENT13                                         = 0x8CED
	R8I                                                        = 0x8231
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        = 0x910A
	TEXTURE_SHARED_SIZE                                        = 0x8C3F
	TEXTURE_ALPHA_SIZE                                         = 0x805F
	IMPLEMENTATION_COLOR_READ_TYPE                             = 0x8B9A
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  = 0x910D
	MAX_FRAGMENT_ATOMIC_COUNTERS                               = 0x92D6
	DEPTH_STENCIL_ATTACHMENT                                   = 0x821A
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        = 0x92D1
	NOTEQUAL                                                   = 0x0205
	SHORT                                                      = 0x1402
	UNSIGNED_INT_IMAGE_2D_RECT                                 = 0x9065
	DEPTH_BUFFER_BIT                                           = 0x00000100
	ELEMENT_ARRAY_BARRIER_BIT                                  = 0x00000002
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               = 0x88FC
	PROGRAM                                                    = 0x82E2
	DEBUG_SOURCE_OTHER                                         = 0x824B
	VERTEX_ATTRIB_ARRAY_SIZE                                   = 0x8623
	RGBA16F                                                    = 0x881A
	PROXY_TEXTURE_3D                                           = 0x8070
	SYNC_FLAGS                                                 = 0x9115
	ALREADY_SIGNALED                                           = 0x911A
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        = 0x900F
	REFERENCED_BY_TESS_CONTROL_SHADER                          = 0x9307
	GET_TEXTURE_IMAGE_FORMAT                                   = 0x8291
	SRC_ALPHA_SATURATE                                         = 0x0308
	COLOR_ATTACHMENT10                                         = 0x8CEA
	FRACTIONAL_EVEN                                            = 0x8E7C
	TRANSFORM_FEEDBACK                                         = 0x8E22
	TEXTURE_BINDING_BUFFER                                     = 0x8C2C
	VIEW_CLASS_128_BITS                                        = 0x82C4
	IMAGE_2D_MULTISAMPLE                                       = 0x9055
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             = 0x82B1
	MAX_COLOR_ATTACHMENTS                                      = 0x8CDF
	SAMPLER_CUBE                                               = 0x8B60
	VERTEX_ATTRIB_RELATIVE_OFFSET                              = 0x82D5
	DOUBLE_VEC3                                                = 0x8FFD
	INCR_WRAP                                                  = 0x8507
	INT_SAMPLER_2D_ARRAY                                       = 0x8DCF
	MAX_SAMPLES                                                = 0x8D57
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          = 0x906A
	TEXTURE_BINDING_2D_ARRAY                                   = 0x8C1D
	READ_PIXELS_TYPE                                           = 0x828E
	UNSIGNED_INT_IMAGE_1D_ARRAY                                = 0x9068
	INTERNALFORMAT_DEPTH_SIZE                                  = 0x8275
	PIXEL_BUFFER_BARRIER_BIT                                   = 0x00000080
	UNSIGNED_INT                                               = 0x1405
	DEBUG_TYPE_PERFORMANCE                                     = 0x8250
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                = 0x8518
	INCR                                                       = 0x1E02
	DITHER                                                     = 0x0BD0
	PACK_SKIP_PIXELS                                           = 0x0D04
	BLEND_EQUATION_ALPHA                                       = 0x883D
	QUADS                                                      = 0x0007
	NONE                                                       = 0
	VERTEX_SHADER                                              = 0x8B31
	ACTIVE_ATTRIBUTES                                          = 0x8B89
	PACK_COMPRESSED_BLOCK_SIZE                                 = 0x912E
	LINE_LOOP                                                  = 0x0002
	TEXTURE_2D_MULTISAMPLE_ARRAY                               = 0x9102
	IMAGE_2D                                                   = 0x904D
	COMPRESSED_RED                                             = 0x8225
	PIXEL_PACK_BUFFER_BINDING                                  = 0x88ED
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            = 0x90CB
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        = 0x8217
	FRAMEBUFFER_ATTACHMENT_LAYERED                             = 0x8DA7
	VERTEX_BINDING_STRIDE                                      = 0x82D8
	UNSIGNED_INT_IMAGE_2D_ARRAY                                = 0x9069
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            = 0x00000001
	COPY                                                       = 0x1503
	SAMPLE_COVERAGE_INVERT                                     = 0x80AB
	FRAMEBUFFER_RENDERABLE                                     = 0x8289
	MAX_ELEMENTS_VERTICES                                      = 0x80E8
	ARRAY_SIZE                                                 = 0x92FB
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         = 0x8E5D
	SAMPLER_2D_RECT                                            = 0x8B63
	RENDERBUFFER_HEIGHT                                        = 0x8D43
	UNIFORM_ARRAY_STRIDE                                       = 0x8A3C
	MIN_PROGRAM_TEXEL_OFFSET                                   = 0x8904
	MAX_INTEGER_SAMPLES                                        = 0x9110
	VIEWPORT_INDEX_PROVOKING_VERTEX                            = 0x825F
	SHADING_LANGUAGE_VERSION                                   = 0x8B8C
	VIEW_CLASS_BPTC_FLOAT                                      = 0x82D3
	UNPACK_SKIP_ROWS                                           = 0x0CF3
	INTERNALFORMAT_RED_SIZE                                    = 0x8271
	DOUBLE_VEC4                                                = 0x8FFE
	STATIC_READ                                                = 0x88E5
	RG16UI                                                     = 0x823A
	IMAGE_2D_ARRAY                                             = 0x9053
	NEAREST_MIPMAP_LINEAR                                      = 0x2702
	BOOL_VEC3                                                  = 0x8B58
	FRAMEBUFFER_COMPLETE                                       = 0x8CD5
	MAX_VERTEX_STREAMS                                         = 0x8E71
	DOUBLE                                                     = 0x140A
	QUERY_RESULT_AVAILABLE                                     = 0x8867
	RG32F                                                      = 0x8230
	DOUBLE_MAT2x4                                              = 0x8F4A
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       = 0x886D
	POLYGON_OFFSET_FILL                                        = 0x8037
	MAX_COLOR_TEXTURE_SAMPLES                                  = 0x910E
	MIRRORED_REPEAT                                            = 0x8370
	FRAMEBUFFER_SRGB                                           = 0x8DB9
	TEXTURE_MAX_LOD                                            = 0x813B
	NICEST                                                     = 0x1102
	FRONT_AND_BACK                                             = 0x0408
	GEOMETRY_INPUT_TYPE                                        = 0x8917
	DOUBLE_MAT3x4                                              = 0x8F4C
	DEBUG_SEVERITY_MEDIUM                                      = 0x9147
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     = 0x90DF
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            = 0x8E1F
	LINK_STATUS                                                = 0x8B82
	XOR                                                        = 0x1506
	BYTE                                                       = 0x1400
	FRAMEBUFFER_BINDING                                        = 0x8CA6
	SRC_ALPHA                                                  = 0x0302
	TRANSFORM_FEEDBACK_BINDING                                 = 0x8E25
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  = 0x8CD7
	TEXTURE13                                                  = 0x84CD
	TEXTURE10                                                  = 0x84CA
	MAX_TEXTURE_SIZE                                           = 0x0D33
	COLOR_CLEAR_VALUE                                          = 0x0C22
	FLOAT_MAT4x2                                               = 0x8B69
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          = 0x92C7
	MAX_COMBINED_UNIFORM_BLOCKS                                = 0x8A2E
	SYNC_CONDITION                                             = 0x9113
	ONE_MINUS_SRC_ALPHA                                        = 0x0303
	INT_IMAGE_BUFFER                                           = 0x905C
	TEXTURE                                                    = 0x1702
	CONTEXT_FLAG_DEBUG_BIT                                     = 0x00000002
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             = 0x9128
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        = 0x82AC
	REFERENCED_BY_COMPUTE_SHADER                               = 0x930B
	FLOAT_VEC4                                                 = 0x8B52
	READ_BUFFER                                                = 0x0C02
	TEXTURE_2D_ARRAY                                           = 0x8C1A
	CONTEXT_FLAGS                                              = 0x821E
	RGBA16                                                     = 0x805B
	IMAGE_CUBE_MAP_ARRAY                                       = 0x9054
	INT_SAMPLER_2D_MULTISAMPLE                                 = 0x9109
	UNPACK_COMPRESSED_BLOCK_SIZE                               = 0x912A
	CULL_FACE                                                  = 0x0B44
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          = 0x00000002
	UNSIGNED_INT_IMAGE_3D                                      = 0x9064
	BACK                                                       = 0x0405
	TEXTURE_BINDING_2D                                         = 0x8069
	SIGNALED                                                   = 0x9119
	FLOAT                                                      = 0x1406
	TRANSFORM_FEEDBACK_ACTIVE                                  = 0x8E24
	TEXTURE_SAMPLES                                            = 0x9106
	TRANSFORM_FEEDBACK_VARYINGS                                = 0x8C83
	R3_G3_B2                                                   = 0x2A10
	UNPACK_ROW_LENGTH                                          = 0x0CF2
	RGB_INTEGER                                                = 0x8D98
	DEBUG_SOURCE_API                                           = 0x8246
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    = 0x8C8B
	TEXTURE20                                                  = 0x84D4
	MAX_SAMPLE_MASK_WORDS                                      = 0x8E59
	TEXTURE_GREEN_SIZE                                         = 0x805D
	BLOCK_INDEX                                                = 0x92FD
	RIGHT                                                      = 0x0407
	DEPTH_COMPONENT24                                          = 0x81A6
	MAX_VERTEX_UNIFORM_COMPONENTS                              = 0x8B4A
	TEXTURE_INTERNAL_FORMAT                                    = 0x1003
	SAMPLER_2D                                                 = 0x8B5E
	SMOOTH_LINE_WIDTH_GRANULARITY                              = 0x0B23
	CURRENT_QUERY                                              = 0x8865
	BUFFER_ACCESS                                              = 0x88BB
	TEXTURE_COMPARE_FUNC                                       = 0x884D
	STENCIL_BACK_WRITEMASK                                     = 0x8CA5
	COMPUTE_TEXTURE                                            = 0x82A0
	TEXTURE_WIDTH                                              = 0x1000
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               = 0x8CD3
	COLOR_ATTACHMENT12                                         = 0x8CEC
	TESS_GEN_MODE                                              = 0x8E76
	DRAW_BUFFER2                                               = 0x8827
	RG16                                                       = 0x822C
	COPY_WRITE_BUFFER_BINDING                                  = 0x8F37
	UNIFORM_BLOCK                                              = 0x92E2
	SRGB                                                       = 0x8C40
	MAX_DRAW_BUFFERS                                           = 0x8824
	CLAMP_TO_BORDER                                            = 0x812D
	CLIP_DISTANCE6                                             = 0x3006
	STENCIL_INDEX1                                             = 0x8D46
	RG8_SNORM                                                  = 0x8F95
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           = 0x8214
	SMOOTH_POINT_SIZE_GRANULARITY                              = 0x0B13
	CLIP_DISTANCE4                                             = 0x3004
	TEXTURE_3D                                                 = 0x806F
	SHADER_IMAGE_LOAD                                          = 0x82A4
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   = 0x9276
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               = 0x8E1E
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5C
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             = 0x9061
	GREEN_INTEGER                                              = 0x8D95
	COMPRESSED_RGB8_ETC2                                       = 0x9274
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           = TRANSFORM_FEEDBACK_PAUSED
	SHADER_BINARY_FORMATS                                      = 0x8DF8
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        = 0x92CF
	R16UI                                                      = 0x8234
	UNPACK_ALIGNMENT                                           = 0x0CF5
	ANY_SAMPLES_PASSED_CONSERVATIVE                            = 0x8D6A
	MAP_READ_BIT                                               = 0x0001
	DEPTH_STENCIL                                              = 0x84F9
	MEDIUM_INT                                                 = 0x8DF4
	VIEW_CLASS_64_BITS                                         = 0x82C6
	DEPTH_FUNC                                                 = 0x0B74
	ALWAYS                                                     = 0x0207
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           = 0x8DE8
	ARRAY_BUFFER_BINDING                                       = 0x8894
	DEPTH_RENDERABLE                                           = 0x8287
	RG16I                                                      = 0x8239
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            = 0x8212
	INT_VEC4                                                   = 0x8B55
	ONE_MINUS_SRC1_ALPHA                                       = 0x88FB
	QUERY                                                      = 0x82E3
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         = 0x8E8A
	PATCHES                                                    = 0x000E
	UNPACK_COMPRESSED_BLOCK_WIDTH                              = 0x9127
	DRAW_BUFFER3                                               = 0x8828
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              = 0x8A42
	DEPTH_TEST                                                 = 0x0B71
	DEPTH24_STENCIL8                                           = 0x88F0
	IMAGE_CLASS_2_X_16                                         = 0x82BD
	RGBA16_SNORM                                               = 0x8F9B
	UNSIGNED_INT_IMAGE_1D                                      = 0x9062
	BGR                                                        = 0x80E0
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT                        = 0x8C4F
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        = 0x0001
	MAX_COMPUTE_UNIFORM_COMPONENTS                             = 0x8263
	MAX_DEBUG_MESSAGE_LENGTH                                   = 0x9143
	UNSIGNED_INT_24_8                                          = 0x84FA
	RENDERBUFFER_SAMPLES                                       = 0x8CAB
	CURRENT_VERTEX_ATTRIB                                      = 0x8626
	STENCIL_CLEAR_VALUE                                        = 0x0B91
	QUERY_RESULT                                               = 0x8866
	RENDERBUFFER_ALPHA_SIZE                                    = 0x8D53
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             = 0x8C85
	OBJECT_TYPE                                                = 0x9112
	VIEW_CLASS_S3TC_DXT1_RGB                                   = 0x82CC
	COLOR_ATTACHMENT0                                          = 0x8CE0
	TEXTURE12                                                  = 0x84CC
	NUM_EXTENSIONS                                             = 0x821D
	STENCIL_BACK_FUNC                                          = 0x8800
	MAX_NUM_ACTIVE_VARIABLES                                   = 0x92F7
	TEXTURE1                                                   = 0x84C1
	TEXTURE_MIN_FILTER                                         = 0x2801
	MAX_NAME_LENGTH                                            = 0x92F6
	RENDERBUFFER_GREEN_SIZE                                    = 0x8D51
	MAX_GEOMETRY_IMAGE_UNIFORMS                                = 0x90CD
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5E
	PROGRAM_PIPELINE                                           = 0x82E4
	PROXY_TEXTURE_1D                                           = 0x8063
	UNSIGNED_SHORT_5_6_5                                       = 0x8363
	PACK_COMPRESSED_BLOCK_WIDTH                                = 0x912B
	MAX_COMPUTE_ATOMIC_COUNTERS                                = 0x8265
	MAX_RENDERBUFFER_SIZE                                      = 0x84E8
	UNSIGNED_INT_IMAGE_2D                                      = 0x9063
	COMPRESSED_RGB_S3TC_DXT1_EXT                               = 0x83F0
	IMAGE_BINDING_ACCESS                                       = 0x8F3E
	ACTIVE_SUBROUTINES                                         = 0x8DE5
	ATOMIC_COUNTER_BUFFER_START                                = 0x92C2
	SLUMINANCE8_EXT                                            = 0x8C47
	IMAGE_BINDING_LEVEL                                        = 0x8F3B
	COLOR_ATTACHMENT6                                          = 0x8CE6
	LINES                                                      = 0x0001
	TEXTURE_CUBE_MAP_SEAMLESS                                  = 0x884F
	IMPLEMENTATION_COLOR_READ_FORMAT                           = 0x8B9B
	FLOAT_MAT3x4                                               = 0x8B68
	DEBUG_SOURCE_THIRD_PARTY                                   = 0x8249
	MAX_FRAGMENT_UNIFORM_VECTORS                               = 0x8DFD
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            = 0x90C7
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x92C9
	LINE_SMOOTH                                                = 0x0B20
	DYNAMIC_DRAW                                               = 0x88E8
	MAX_GEOMETRY_ATOMIC_COUNTERS                               = 0x92D5
	VIEW_CLASS_RGTC1_RED                                       = 0x82D0
	DOUBLE_VEC2                                                = 0x8FFC
	FLOAT_MAT2x4                                               = 0x8B66
	TOP_LEVEL_ARRAY_SIZE                                       = 0x930C
	INTERNALFORMAT_PREFERRED                                   = 0x8270
	PATCH_DEFAULT_INNER_LEVEL                                  = 0x8E73
	VIEWPORT_SUBPIXEL_BITS                                     = 0x825C
	RED                                                        = 0x1903
	SMOOTH_POINT_SIZE_RANGE                                    = 0x0B12
	MAX_ELEMENT_INDEX                                          = 0x8D6B
	BACK_RIGHT                                                 = 0x0403
	TEXTURE_2D                                                 = 0x0DE1
	ONE_MINUS_SRC1_COLOR                                       = 0x88FA
	READ_WRITE                                                 = 0x88BA
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          = 0x82D9
	RGB9_E5                                                    = 0x8C3D
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      = 0x8210
	TEXTURE_1D                                                 = 0x0DE0
	VERTEX_ATTRIB_ARRAY_TYPE                                   = 0x8625
	DRAW_BUFFER12                                              = 0x8831
	INTERNALFORMAT_STENCIL_SIZE                                = 0x8276
	UNPACK_SKIP_IMAGES                                         = 0x806D
	RENDERBUFFER_BINDING                                       = 0x8CA7
	LAST_VERTEX_CONVENTION                                     = 0x8E4E
	COMPRESSED_RG                                              = 0x8226
	RGBA8_SNORM                                                = 0x8F97
	MAX_LAYERS                                                 = 0x8281
	VERTEX_SUBROUTINE_UNIFORM                                  = 0x92EE
	INT_SAMPLER_CUBE_MAP_ARRAY                                 = 0x900E
	BUFFER_VARIABLE                                            = 0x92E5
	COMPRESSED_SRGB                                            = 0x8C48
	SUBPIXEL_BITS                                              = 0x0D50
	MAX_COMPUTE_WORK_GROUP_COUNT                               = 0x91BE
	DEBUG_LOGGED_MESSAGES                                      = 0x9145
	COMPRESSED_SRGB_S3TC_DXT1_EXT                              = 0x8C4C
	INTERNALFORMAT_GREEN_TYPE                                  = 0x8279
	DRAW_BUFFER5                                               = 0x882A
	MAX_TEXTURE_BUFFER_SIZE                                    = 0x8C2B
	MAP_INVALIDATE_BUFFER_BIT                                  = 0x0008
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         = 0x90C8
	TEXTURE_WRAP_R                                             = 0x8072
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         = 0x92DC
	DEPTH_WRITEMASK                                            = 0x0B72
	FILTER                                                     = 0x829A
	REFERENCED_BY_VERTEX_SHADER                                = 0x9306
	VIEWPORT_BOUNDS_RANGE                                      = 0x825D
	RGBA16UI                                                   = 0x8D76
	MAX_TEXTURE_MAX_ANISOTROPY_EXT                             = 0x84FF
	TEXTURE_CUBE_MAP_POSITIVE_Y                                = 0x8517
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       = 0x8E81
	UNPACK_SWAP_BYTES                                          = 0x0CF0
	TEXTURE_BUFFER_SIZE                                        = 0x919E
	SRGB8                                                      = 0x8C41
	IMAGE_CLASS_2_X_8                                          = 0x82C0
	INTERNALFORMAT_ALPHA_SIZE                                  = 0x8274
	UNSIGNED_INT_VEC2                                          = 0x8DC6
	READ_ONLY                                                  = 0x88B8
	PROGRAM_OUTPUT                                             = 0x92E4
	COMPRESSED_SIGNED_RED_RGTC1                                = 0x8DBC
	MAX_DEBUG_LOGGED_MESSAGES                                  = 0x9144
	MAX_ARRAY_TEXTURE_LAYERS                                   = 0x88FF
	COLOR_ATTACHMENT11                                         = 0x8CEB
	REFERENCED_BY_TESS_EVALUATION_SHADER                       = 0x9308
	TRANSFORM_FEEDBACK_BARRIER_BIT                             = 0x00000800
	INT                                                        = 0x1404
	FRAGMENT_SUBROUTINE                                        = 0x92EC
	COMPILE_STATUS                                             = 0x8B81
	TESS_CONTROL_SUBROUTINE_UNIFORM                            = 0x92EF
	TRUE                                                       = 1
	DEBUG_OUTPUT                                               = 0x92E0
	TESS_CONTROL_SHADER                                        = 0x8E88
	CURRENT_PROGRAM                                            = 0x8B8D
	INVALID_FRAMEBUFFER_OPERATION                              = 0x0506
	STENCIL_BACK_PASS_DEPTH_FAIL                               = 0x8802
	INT_SAMPLER_1D_ARRAY                                       = 0x8DCE
	CLIP_DISTANCE5                                             = 0x3005
	RENDERBUFFER_RED_SIZE                                      = 0x8D50
	NEAREST_MIPMAP_NEAREST                                     = 0x2700
	TEXTURE29                                                  = 0x84DD
	VIEW_CLASS_96_BITS                                         = 0x82C5
	R16F                                                       = 0x822D
	TEXTURE31                                                  = 0x84DF
	BGRA_INTEGER                                               = 0x8D9B
	SAMPLE_MASK                                                = 0x8E51
	UNIFORM_BLOCK_BINDING                                      = 0x8A3F
	SAMPLER_BUFFER                                             = 0x8DC2
	FLOAT_MAT4                                                 = 0x8B5C
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         = 0x90DA
	RGB16UI                                                    = 0x8D77
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             = 0x900A
	TEXTURE_BLUE_TYPE                                          = 0x8C12
	TESS_EVALUATION_SUBROUTINE                                 = 0x92EA
	MULTISAMPLE                                                = 0x809D
	GEOMETRY_SHADER                                            = 0x8DD9
	BUFFER_MAP_POINTER                                         = 0x88BD
	RGB10_A2UI                                                 = 0x906F
	TEXTURE_IMAGE_FORMAT                                       = 0x828F
	VERTEX_PROGRAM_POINT_SIZE                                  = 0x8642
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       = 0x8DA8
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      = 0x8C76
	LEFT                                                       = 0x0406
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          = 0x8216
	IMAGE_CLASS_1_X_8                                          = 0x82C1
	AND_INVERTED                                               = 0x1504
	ALIASED_LINE_WIDTH_RANGE                                   = 0x846E
	UNIFORM_BLOCK_NAME_LENGTH                                  = 0x8A41
	LAYER_PROVOKING_VERTEX                                     = 0x825E
	TEXTURE_BASE_LEVEL                                         = 0x813C
	TEXTURE_BINDING_1D_ARRAY                                   = 0x8C1C
	SHADER_STORAGE_BLOCK                                       = 0x92E6
	TEXTURE15                                                  = 0x84CF
	UNSIGNED_SHORT_5_6_5_REV                                   = 0x8364
	NAND                                                       = 0x150E
	COLOR_ATTACHMENT7                                          = 0x8CE7
	PIXEL_UNPACK_BUFFER                                        = 0x88EC
	INTERNALFORMAT_ALPHA_TYPE                                  = 0x827B
	RENDERBUFFER_STENCIL_SIZE                                  = 0x8D55
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            = 0x00000020
	COMPRESSED_SLUMINANCE_EXT                                  = 0x8C4A
	PACK_LSB_FIRST                                             = 0x0D01
	TEXTURE9                                                   = 0x84C9
	DEBUG_CALLBACK_FUNCTION                                    = 0x8244
	CW                                                         = 0x0900
	MAX_FRAMEBUFFER_SAMPLES                                    = 0x9318
	PATCH_DEFAULT_OUTER_LEVEL                                  = 0x8E74
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    = 0x8266
	SHADER_STORAGE_BUFFER                                      = 0x90D2
	VIEW_CLASS_BPTC_UNORM                                      = 0x82D2
	COMPRESSED_SIGNED_RG_RGTC2                                 = 0x8DBE
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   = 0x8E85
	TEXTURE_DEPTH                                              = 0x8071
	COLOR_ATTACHMENT15                                         = 0x8CEF
	LOCATION_INDEX                                             = 0x930F
	MAX_VERTEX_UNIFORM_VECTORS                                 = 0x8DFB
	VIEW_CLASS_32_BITS                                         = 0x82C8
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         = 0x90ED
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           = 0x92D3
	STENCIL_COMPONENTS                                         = 0x8285
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      = 0x8211
	DELETE_STATUS                                              = 0x8B80
	CLIP_DISTANCE2                                             = 0x3002
	FRACTIONAL_ODD                                             = 0x8E7B
	MAX_FRAMEBUFFER_LAYERS                                     = 0x9317
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            = 0x8E89
	BLEND_DST_RGB                                              = 0x80C8
	CLIP_DISTANCE1                                             = 0x3001
	TEXTURE_IMMUTABLE_LEVELS                                   = 0x82DF
	DRAW_FRAMEBUFFER                                           = 0x8CA9
	FRAMEBUFFER_DEFAULT_SAMPLES                                = 0x9313
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   = 0x8E4C
	DEPTH_CLAMP                                                = 0x864F
	DEBUG_GROUP_STACK_DEPTH                                    = 0x826D
	STENCIL_TEST                                               = 0x0B90
	LINE_WIDTH_RANGE                                           = 0x0B22
	COMPRESSED_SRGB_ALPHA                                      = 0x8C49
	R16I                                                       = 0x8233
	TEXTURE_VIEW_NUM_LAYERS                                    = 0x82DE
	RGBA8UI                                                    = 0x8D7C
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS                      = 0x8F9F
	DOUBLE_MAT2x3                                              = 0x8F49
	UNSIGNED_INT_5_9_9_9_REV                                   = 0x8C3E
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    = 0x8E82
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          = 0x90DB
	RGB                                                        = 0x1907
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       = 0x8DE1
	MAX_IMAGE_UNITS                                            = 0x8F38
	DOUBLE_MAT3x2                                              = 0x8F4B
	COLOR_COMPONENTS                                           = 0x8283
	UNIFORM_TYPE                                               = 0x8A37
	MAX_FRAGMENT_IMAGE_UNIFORMS                                = 0x90CE
	COMPARE_REF_TO_TEXTURE                                     = 0x884E
	TEXTURE25                                                  = 0x84D9
	VIEW_COMPATIBILITY_CLASS                                   = 0x82B6
	UNSIGNED_INT_SAMPLER_1D                                    = 0x8DD1
	SHADER_IMAGE_ATOMIC                                        = 0x82A6
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             = 0x824D
	COLOR_ENCODING                                             = 0x8296
	MAX_TESS_PATCH_COMPONENTS                                  = 0x8E84
	NUM_COMPRESSED_TEXTURE_FORMATS                             = 0x86A2
	MAX_LABEL_LENGTH                                           = 0x82E8
	SAMPLER_1D_SHADOW                                          = 0x8B61
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         = 0x8CD1
	MEDIUM_FLOAT                                               = 0x8DF1
	PACK_SWAP_BYTES                                            = 0x0D00
	MIPMAP                                                     = 0x8293
	POINTS                                                     = 0x0000
	FLOAT_VEC3                                                 = 0x8B51
	ONE                                                        = 1
	POINT                                                      = 0x1B00
	IMAGE_PIXEL_TYPE                                           = 0x82AA
	RGB4                                                       = 0x804F
	UNSIGNED_BYTE                                              = 0x1401
	DEPTH_COMPONENT16                                          = 0x81A5
	LINE_SMOOTH_HINT                                           = 0x0C52
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           = TRANSFORM_FEEDBACK_ACTIVE
	ACTIVE_UNIFORM_MAX_LENGTH                                  = 0x8B87
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             = 0x8B4C
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     = 0x90D8
	TEXTURE5                                                   = 0x84C5
	UNSIGNED_INT_IMAGE_BUFFER                                  = 0x9067
	NOOP                                                       = 0x1505
	IMAGE_CUBE                                                 = 0x9050
	PACK_ALIGNMENT                                             = 0x0D05
	PIXEL_UNPACK_BUFFER_BINDING                                = 0x88EF
	SAMPLER_CUBE_MAP_ARRAY                                     = 0x900C
	TEXTURE21                                                  = 0x84D5
	PROGRAM_BINARY_FORMATS                                     = 0x87FF
	IMAGE_2D_MULTISAMPLE_ARRAY                                 = 0x9056
	COPY_WRITE_BUFFER                                          = COPY_WRITE_BUFFER_BINDING
	REFERENCED_BY_FRAGMENT_SHADER                              = 0x930A
	TEXTURE_FETCH_BARRIER_BIT                                  = 0x00000008
	NEVER                                                      = 0x0200
	RGBA32F                                                    = 0x8814
	RGB8UI                                                     = 0x8D7D
	PROGRAM_INPUT                                              = 0x92E3
	DRAW_BUFFER8                                               = 0x882D
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         = 0x90D7
	TEXTURE_VIEW                                               = 0x82B5
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       = 0x8CD2
	DRAW_FRAMEBUFFER_BINDING                                   = FRAMEBUFFER_BINDING
	MIN_SAMPLE_SHADING_VALUE                                   = 0x8C37
	GEOMETRY_TEXTURE                                           = 0x829E
	COLOR_ATTACHMENT3                                          = 0x8CE3
	RGB16I                                                     = 0x8D89
	INT_IMAGE_CUBE                                             = 0x905B
	UNPACK_IMAGE_HEIGHT                                        = 0x806E
	VERTEX_ATTRIB_ARRAY_DIVISOR                                = 0x88FE
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       = 0x8A43
	MANUAL_GENERATE_MIPMAP                                     = 0x8294
	BLEND                                                      = 0x0BE2
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            = 0x919F
	DEPTH_COMPONENT                                            = 0x1902
	VENDOR                                                     = 0x1F00
	REPLACE                                                    = 0x1E01
	ELEMENT_ARRAY_BUFFER_BINDING                               = 0x8895
	FILL                                                       = 0x1B02
	INVALID_VALUE                                              = 0x0501
	QUERY_NO_WAIT                                              = 0x8E14
	SAMPLE_POSITION                                            = 0x8E50
	RGB8_SNORM                                                 = 0x8F96
	FRAMEBUFFER_RENDERABLE_LAYERED                             = 0x828A
	INT_IMAGE_2D_RECT                                          = 0x905A
	VERTEX_ATTRIB_ARRAY_INTEGER                                = 0x88FD
	REPEAT                                                     = 0x2901
	COLOR_ATTACHMENT4                                          = 0x8CE4
	UNSIGNED_SHORT_4_4_4_4                                     = 0x8033
	DOUBLE_MAT4x3                                              = 0x8F4E
	STATIC_COPY                                                = 0x88E6
	ARRAY_STRIDE                                               = 0x92FE
	TRIANGLE_FAN                                               = 0x0006
	BLEND_DST_ALPHA                                            = 0x80CA
	VIEW_CLASS_S3TC_DXT1_RGBA                                  = 0x82CD
	LOWER_LEFT                                                 = 0x8CA1
	FRAMEBUFFER_DEFAULT_WIDTH                                  = 0x9310
	IMAGE_CLASS_4_X_8                                          = 0x82BF
	FRONT_LEFT                                                 = 0x0400
	TEXTURE_2D_MULTISAMPLE                                     = 0x9100
	UNSIGNED_INT_SAMPLER_3D                                    = 0x8DD3
	TEXTURE_STENCIL_SIZE                                       = 0x88F1
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               = 0x900B
	MAP_INVALIDATE_RANGE_BIT                                   = 0x0004
	FLOAT_MAT3                                                 = 0x8B5B
	ACTIVE_SUBROUTINE_MAX_LENGTH                               = 0x8E48
	TESS_CONTROL_SHADER_BIT                                    = 0x00000008
	BUFFER_BINDING                                             = 0x9302
	TEXTURE_HEIGHT                                             = 0x1001
	MAX_VIEWPORTS                                              = 0x825B
	TEXTURE_SWIZZLE_A                                          = 0x8E45
	COMPRESSED_SLUMINANCE_ALPHA_EXT                            = 0x8C4B
	PROGRAM_POINT_SIZE                                         = 0x8642
	SAMPLER_3D                                                 = 0x8B5F
	STENCIL_VALUE_MASK                                         = 0x0B93
	OFFSET                                                     = 0x92FC
	WAIT_FAILED                                                = 0x911D
	LOW_INT                                                    = 0x8DF3
	DEPTH_COMPONENTS                                           = 0x8284
	CULL_FACE_MODE                                             = 0x0B45
	DRAW_BUFFER0                                               = 0x8825
	PATCH_VERTICES                                             = 0x8E72
	PROGRAM_SEPARABLE                                          = 0x8258
	DEPTH32F_STENCIL8                                          = 0x8CAD
	PROXY_TEXTURE_2D_MULTISAMPLE                               = 0x9101
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         = 0x889F
	INTERNALFORMAT_BLUE_SIZE                                   = 0x8273
	CONTEXT_CORE_PROFILE_BIT                                   = 0x00000001
	DEBUG_SOURCE_SHADER_COMPILER                               = 0x8248
	DRAW_BUFFER11                                              = 0x8830
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5F
	STREAM_DRAW                                                = 0x88E0
	PRIMITIVES_GENERATED                                       = 0x8C87
	UNSIGNED_SHORT_1_5_5_5_REV                                 = 0x8366
	MAX_VARYING_FLOATS                                         = 0x8B4B
	R8UI                                                       = 0x8232
	DEBUG_SEVERITY_HIGH                                        = 0x9146
	LINE_STRIP                                                 = 0x0003
	RED_SNORM                                                  = 0x8F90
	TEXTURE18                                                  = 0x84D2
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         = 0x8D56
	UNPACK_SKIP_PIXELS                                         = 0x0CF4
	FLOAT_MAT2x3                                               = 0x8B65
	R32UI                                                      = 0x8236
	UNSIGNED_INT_SAMPLER_2D_RECT                               = 0x8DD5
	SYNC_FENCE                                                 = 0x9116
	INT_SAMPLER_3D                                             = 0x8DCB
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            = 0x91BC
	UNSIGNED_BYTE_3_3_2                                        = 0x8032
	TEXTURE_CUBE_MAP_POSITIVE_Z                                = 0x8519
	MATRIX_STRIDE                                              = 0x92FF
	SLUMINANCE_EXT                                             = 0x8C46
	CLIP_DISTANCE3                                             = 0x3003
	BLEND_EQUATION_RGB                                         = 0x8009
	MAX_COMPUTE_WORK_GROUP_SIZE                                = 0x91BF
	UNSIGNED_INT_8_8_8_8_REV                                   = 0x8367
	IMAGE_BINDING_FORMAT                                       = 0x906E
	ACTIVE_UNIFORMS                                            = 0x8B86
	DEBUG_SEVERITY_NOTIFICATION                                = 0x826B
	OR_INVERTED                                                = 0x150D
	FRAGMENT_TEXTURE                                           = 0x829F
	DEBUG_SOURCE_WINDOW_SYSTEM                                 = 0x8247
	GEOMETRY_VERTICES_OUT                                      = 0x8916
	BOOL_VEC4                                                  = 0x8B59
	SHADER_STORAGE_BUFFER_SIZE                                 = 0x90D5
	SAMPLER_1D_ARRAY                                           = 0x8DC0
	TRANSFORM_FEEDBACK_PAUSED                                  = 0x8E23
	IMAGE_BINDING_LAYERED                                      = 0x8F3C
	COLOR_ATTACHMENT8                                          = 0x8CE8
	SAMPLES                                                    = 0x80A9
	BUFFER_ACCESS_FLAGS                                        = 0x911F
	TEXTURE0                                                   = 0x84C0
	FASTEST                                                    = 0x1101
	UNSIGNED_INT_SAMPLER_BUFFER                                = 0x8DD8
	PIXEL_PACK_BUFFER                                          = 0x88EB
	STENCIL_FUNC                                               = 0x0B92
	BGR_INTEGER                                                = 0x8D9A
	FRAMEBUFFER_DEFAULT_LAYERS                                 = 0x9312
	ALPHA                                                      = 0x1906
	INT_IMAGE_3D                                               = 0x9059
	VERTEX_BINDING_OFFSET                                      = 0x82D7
	SHADER_SOURCE_LENGTH                                       = 0x8B88
	HALF_FLOAT                                                 = 0x140B
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        = 0x92D4
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      = 0x8E86
	SAMPLER_CUBE_SHADOW                                        = 0x8DC5
	R16                                                        = 0x822A
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           = 0x8243
	TEXTURE_SWIZZLE_B                                          = 0x8E44
	TEXTURE_CUBE_MAP_ARRAY                                     = 0x9009
	ATOMIC_COUNTER_BARRIER_BIT                                 = 0x00001000
	TRANSFORM_FEEDBACK_VARYING                                 = 0x92F4
	COMPRESSED_SRGB8_ETC2                                      = 0x9275
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        = 0x92DA
	IMAGE_CLASS_4_X_16                                         = 0x82BC
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     = 0x82AF
	COMPATIBLE_SUBROUTINES                                     = 0x8E4B
	RGBA_SNORM                                                 = 0x8F93
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 = 0x92CE
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         = 0x8E83
	PRIMITIVE_RESTART_INDEX                                    = 0x8F9E
	ACTIVE_TEXTURE                                             = 0x84E0
	FLOAT_MAT4x3                                               = 0x8B6A
	SAMPLER_1D                                                 = 0x8B5D
	ARRAY_BUFFER                                               = 0x8892
	DRAW_BUFFER15                                              = 0x8834
	TEXTURE_DEPTH_SIZE                                         = 0x884A
	FRAGMENT_SHADER_BIT                                        = 0x00000002
	GEOMETRY_OUTPUT_TYPE                                       = 0x8918
	LEQUAL                                                     = 0x0203
	TESS_GEN_VERTEX_ORDER                                      = 0x8E78
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         = 0x90CC
	SHADER_COMPILER                                            = 0x8DFA
	RGB32UI                                                    = 0x8D71
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              = 0x8DD6
	RGB565                                                     = 0x8D62
	INVALID_OPERATION                                          = 0x0502
	UNIFORM_IS_ROW_MAJOR                                       = 0x8A3E
	DISPLAY_LIST                                               = 0x82E7
	TEXTURE_IMMUTABLE_FORMAT                                   = 0x912F
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT                        = 0x8C4E
	INT_VEC2                                                   = 0x8B53
	TEXTURE_RECTANGLE                                          = 0x84F5
	IMAGE_CLASS_1_X_32                                         = 0x82BB
	BLEND_DST                                                  = 0x0BE0
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            = 0x8B49
	BLUE                                                       = 0x1905
	UNSIGNALED                                                 = 0x9118
	SHADER                                                     = 0x82E1
	SCISSOR_TEST                                               = 0x0C11
	TESS_CONTROL_OUTPUT_VERTICES                               = 0x8E75
	DRAW_BUFFER4                                               = 0x8829
	MINOR_VERSION                                              = 0x821C
	READ_FRAMEBUFFER_BINDING                                   = 0x8CAA
	TIME_ELAPSED                                               = 0x88BF
	MAX_COMBINED_DIMENSIONS                                    = 0x8282
	FRAMEBUFFER_DEFAULT_HEIGHT                                 = 0x9311
	VERTEX_ATTRIB_BINDING                                      = 0x82D4
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       = 0x8CD4
	MAX_TEXTURE_IMAGE_UNITS                                    = 0x8872
	PROXY_TEXTURE_1D_ARRAY                                     = 0x8C19
	COPY_READ_BUFFER                                           = COPY_READ_BUFFER_BINDING
	DRAW_INDIRECT_BUFFER                                       = 0x8F3F
	RGB16                                                      = 0x8054
	RGB16F                                                     = 0x881B
	MAP_WRITE_BIT                                              = 0x0002
	COLOR                                                      = 0x1800
	R11F_G11F_B10F                                             = 0x8C3A
	R32I                                                       = 0x8235
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           = 0x8B4D
	DEBUG_TYPE_ERROR                                           = 0x824C
	R8                                                         = 0x8229
	UNSIGNED_SHORT                                             = 0x1403
	SRGB_WRITE                                                 = 0x8298
	NUM_SHADER_BINARY_FORMATS                                  = 0x8DF9
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        = 0x92C6
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              = 0x824E
	TEXTURE_BINDING_1D                                         = 0x8068
	TESS_GEN_SPACING                                           = 0x8E77
	TEXTURE3                                                   = 0x84C3
	SAMPLER_2D_RECT_SHADOW                                     = 0x8B64
	NUM_PROGRAM_BINARY_FORMATS                                 = 0x87FE
	UNSIGNED_INT_8_8_8_8                                       = 0x8035
	READ_PIXELS                                                = 0x828C
	UNSIGNED_INT_VEC4                                          = 0x8DC8
	RGB10_A2                                                   = 0x8059
	STENCIL_PASS_DEPTH_FAIL                                    = 0x0B95
	LOCATION                                                   = 0x930E
	SHADER_IMAGE_STORE                                         = 0x82A5
	UNIFORM_BUFFER_START                                       = 0x8A29
	FRAMEBUFFER_DEFAULT                                        = 0x8218
	RGBA16I                                                    = 0x8D88
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        = 0x8E47
	INT_IMAGE_CUBE_MAP_ARRAY                                   = 0x905F
	RGBA8I                                                     = 0x8D8E
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          = 0x8215
	IMAGE_BUFFER                                               = 0x9051
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        = 0x90C9
	COLOR_ATTACHMENT5                                          = 0x8CE5
	STATIC_DRAW                                                = 0x88E4
	QUERY_WAIT                                                 = 0x8E13
	BACK_LEFT                                                  = 0x0402
	MAX_GEOMETRY_OUTPUT_VERTICES                               = 0x8DE0
	PROGRAM_PIPELINE_BINDING                                   = 0x825A
	INT_VEC3                                                   = 0x8B54
	BLUE_INTEGER                                               = 0x8D96
	TEXTURE16                                                  = 0x84D0
	ACTIVE_VARIABLES                                           = 0x9305
	BUFFER_USAGE                                               = 0x8765
	TOP_LEVEL_ARRAY_STRIDE                                     = 0x930D
	COMPRESSED_RGBA                                            = 0x84EE
	MAX_VERTEX_OUTPUT_COMPONENTS                               = 0x9122
	RGB12                                                      = 0x8053
	STENCIL_INDEX4                                             = 0x8D47
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  = 0x90D9
	DST_ALPHA                                                  = 0x0304
	DYNAMIC_READ                                               = 0x88E9
	FRAMEBUFFER_UNSUPPORTED                                    = 0x8CDD
	POINT_SIZE_GRANULARITY                                     = 0x0B13
	PACK_ROW_LENGTH                                            = 0x0D02
	TRANSFORM_FEEDBACK_BUFFER_MODE                             = 0x8C7F
	PACK_COMPRESSED_BLOCK_DEPTH                                = 0x912D
	SAMPLE_MASK_VALUE                                          = 0x8E52
	SHADER_TYPE                                                = 0x8B4F
	MAX_SERVER_WAIT_TIMEOUT                                    = 0x9111
	STENCIL_WRITEMASK                                          = 0x0B98
	BLEND_SRC                                                  = 0x0BE1
	FRAMEBUFFER_BLEND                                          = 0x828B
	FIRST_VERTEX_CONVENTION                                    = 0x8E4D
	LINE_STRIP_ADJACENCY                                       = 0x000B
	STREAM_COPY                                                = 0x88E2
	POLYGON_OFFSET_LINE                                        = 0x2A02
	SAMPLE_BUFFERS                                             = 0x80A8
	STACK_OVERFLOW                                             = 0x0503
	COLOR_ATTACHMENT1                                          = 0x8CE1
	MAX_VERTEX_ATTRIB_BINDINGS                                 = 0x82DA
	MAX_FRAGMENT_UNIFORM_BLOCKS                                = 0x8A2D
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      = 0x82AD
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          = 0x92CC
	COMPRESSED_SRGB_EXT                                        = 0x8C48
	COLOR_ATTACHMENT2                                          = 0x8CE2
	OR_REVERSE                                                 = 0x150B
	VERTEX_ATTRIB_ARRAY_POINTER                                = 0x8645
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5B
	LINE_WIDTH                                                 = 0x0B21
	MAP_FLUSH_EXPLICIT_BIT                                     = 0x0010
	SHADER_STORAGE_BUFFER_START                                = 0x90D4
	SAMPLER_2D_MULTISAMPLE_ARRAY                               = 0x910B
	POLYGON_OFFSET_POINT                                       = 0x2A01
	TEXTURE_MAX_ANISOTROPY_EXT                                 = 0x84FE
	SAMPLER_2D_SHADOW                                          = 0x8B62
	INFO_LOG_LENGTH                                            = 0x8B84
	SRGB8_EXT                                                  = 0x8C41
	POLYGON_OFFSET_FACTOR                                      = 0x8038
	BUFFER_MAP_LENGTH                                          = 0x9120
	SAMPLE_ALPHA_TO_ONE                                        = 0x809F
	VIEWPORT                                                   = 0x0BA2
	EXTENSIONS                                                 = 0x1F03
	CLEAR_BUFFER                                               = 0x82B4
	MAX_ELEMENTS_INDICES                                       = 0x80E9
	POINT_SPRITE_COORD_ORIGIN                                  = 0x8CA0
	GREATER                                                    = 0x0204
	TEXTURE24                                                  = 0x84D8
	UNIFORM_BARRIER_BIT                                        = 0x00000004
	PACK_IMAGE_HEIGHT                                          = 0x806C
	TEXTURE11                                                  = 0x84CB
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             = 0x92D8
	MAX_PATCH_VERTICES                                         = 0x8E7D
	CLAMP_READ_COLOR                                           = 0x891C
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               = 0x92C5
	TIMEOUT_IGNORED                                            = 0xFFFFFFFFFFFFFFFF
	VERTEX_BINDING_DIVISOR                                     = 0x82D6
	TEXTURE28                                                  = 0x84DC
	DOUBLEBUFFER                                               = 0x0C32
	RGB8                                                       = 0x8051
	LINEAR_MIPMAP_LINEAR                                       = 0x2703
	PACK_COMPRESSED_BLOCK_HEIGHT                               = 0x912C
	LOW_FLOAT                                                  = 0x8DF0
	SET                                                        = 0x150F
	COMPRESSED_RG11_EAC                                        = 0x9272
	COLOR_ATTACHMENT9                                          = 0x8CE9
	TEXTURE_VIEW_MIN_LAYER                                     = 0x82DD
)

//	Returns the name of the specified Enum. Not recommended in a real-time loop.
func (_ GlUtil) EnumName(enum Enum) (name string) {
	switch enum {
	case COMPUTE_SHADER:
		name = "GL_COMPUTE_SHADER"
	case VERTEX_SHADER:
		name = "GL_VERTEX_SHADER"
	case TESS_CONTROL_SHADER:
		name = "GL_TESS_CONTROL_SHADER"
	case INVALID_FRAMEBUFFER_OPERATION:
		name = "GL_INVALID_FRAMEBUFFER_OPERATION"
	case GEOMETRY_SHADER:
		name = "GL_GEOMETRY_SHADER"
	case INVALID_VALUE:
		name = "GL_INVALID_VALUE"
	case INVALID_OPERATION:
		name = "GL_INVALID_OPERATION"
	case STACK_OVERFLOW:
		name = "GL_STACK_OVERFLOW"
	case FRAGMENT_SHADER:
		name = "GL_FRAGMENT_SHADER"
	case STACK_UNDERFLOW:
		name = "GL_STACK_UNDERFLOW"
	case OUT_OF_MEMORY:
		name = "GL_OUT_OF_MEMORY"
	case TESS_EVALUATION_SHADER:
		name = "GL_TESS_EVALUATION_SHADER"
	case INVALID_ENUM:
		name = "GL_INVALID_ENUM"
	case BUFFER:
		name = "GL_BUFFER"
	case UNIFORM_MATRIX_STRIDE:
		name = "GL_UNIFORM_MATRIX_STRIDE"
	case UNSIGNED_SHORT_5_5_5_1:
		name = "GL_UNSIGNED_SHORT_5_5_5_1"
	case INTERNALFORMAT_GREEN_SIZE:
		name = "GL_INTERNALFORMAT_GREEN_SIZE"
	case RGBA4:
		name = "GL_RGBA4"
	case SAMPLE_COVERAGE_VALUE:
		name = "GL_SAMPLE_COVERAGE_VALUE"
	case LOGIC_OP_MODE:
		name = "GL_LOGIC_OP_MODE"
	case BUFFER_DATA_SIZE:
		name = "GL_BUFFER_DATA_SIZE"
	case LESS:
		name = "GL_LESS"
	case ONE_MINUS_DST_ALPHA:
		name = "GL_ONE_MINUS_DST_ALPHA"
	case STENCIL_RENDERABLE:
		name = "GL_STENCIL_RENDERABLE"
	case SAMPLE_ALPHA_TO_COVERAGE:
		name = "GL_SAMPLE_ALPHA_TO_COVERAGE"
	case STENCIL_REF:
		name = "GL_STENCIL_REF"
	case COPY_INVERTED:
		name = "GL_COPY_INVERTED"
	case UNSIGNED_INT_SAMPLER_2D:
		name = "GL_UNSIGNED_INT_SAMPLER_2D"
	case POINT_SIZE:
		name = "GL_POINT_SIZE"
	case MAX_DEPTH_TEXTURE_SAMPLES:
		name = "GL_MAX_DEPTH_TEXTURE_SAMPLES"
	case INVALID_INDEX:
		name = "GL_INVALID_INDEX"
	case RGBA32I:
		name = "GL_RGBA32I"
	case SAMPLER_BINDING:
		name = "GL_SAMPLER_BINDING"
	case AND_REVERSE:
		name = "GL_AND_REVERSE"
	case DONT_CARE:
		name = "GL_DONT_CARE"
	case COLOR_LOGIC_OP:
		name = "GL_COLOR_LOGIC_OP"
	case INTERNALFORMAT_STENCIL_TYPE:
		name = "GL_INTERNALFORMAT_STENCIL_TYPE"
	case MAX_VIEWPORT_DIMS:
		name = "GL_MAX_VIEWPORT_DIMS"
	case DEBUG_SEVERITY_LOW:
		name = "GL_DEBUG_SEVERITY_LOW"
	case CAVEAT_SUPPORT:
		name = "GL_CAVEAT_SUPPORT"
	case DOUBLE_MAT3:
		name = "GL_DOUBLE_MAT3"
	case VERTEX_ATTRIB_ARRAY_ENABLED:
		name = "GL_VERTEX_ATTRIB_ARRAY_ENABLED"
	case UNSIGNED_INT_IMAGE_CUBE:
		name = "GL_UNSIGNED_INT_IMAGE_CUBE"
	case INT_IMAGE_1D_ARRAY:
		name = "GL_INT_IMAGE_1D_ARRAY"
	case FRAMEBUFFER_UNDEFINED:
		name = "GL_FRAMEBUFFER_UNDEFINED"
	case TESS_EVALUATION_TEXTURE:
		name = "GL_TESS_EVALUATION_TEXTURE"
	case WRITE_ONLY:
		name = "GL_WRITE_ONLY"
	case MAX_CLIP_DISTANCES:
		name = "GL_MAX_CLIP_DISTANCES"
	case MAX_HEIGHT:
		name = "GL_MAX_HEIGHT"
	case COMPRESSED_R11_EAC:
		name = "GL_COMPRESSED_R11_EAC"
	case TEXTURE_COMPRESSION_HINT:
		name = "GL_TEXTURE_COMPRESSION_HINT"
	case TEXTURE_BUFFER_DATA_STORE_BINDING:
		name = "GL_TEXTURE_BUFFER_DATA_STORE_BINDING"
	case STENCIL_BACK_PASS_DEPTH_PASS:
		name = "GL_STENCIL_BACK_PASS_DEPTH_PASS"
	case SAMPLER_2D_ARRAY:
		name = "GL_SAMPLER_2D_ARRAY"
	case TRANSFORM_FEEDBACK_BUFFER:
		name = "GL_TRANSFORM_FEEDBACK_BUFFER"
	case TEXTURE7:
		name = "GL_TEXTURE7"
	case GREEN:
		name = "GL_GREEN"
	case DRAW_BUFFER9:
		name = "GL_DRAW_BUFFER9"
	case TIMESTAMP:
		name = "GL_TIMESTAMP"
	case CCW:
		name = "GL_CCW"
	case POLYGON_OFFSET_UNITS:
		name = "GL_POLYGON_OFFSET_UNITS"
	case UNPACK_COMPRESSED_BLOCK_DEPTH:
		name = "GL_UNPACK_COMPRESSED_BLOCK_DEPTH"
	case QUERY_COUNTER_BITS:
		name = "GL_QUERY_COUNTER_BITS"
	case STENCIL_BACK_FAIL:
		name = "GL_STENCIL_BACK_FAIL"
	case STENCIL_BACK_REF:
		name = "GL_STENCIL_BACK_REF"
	case SYNC_STATUS:
		name = "GL_SYNC_STATUS"
	case LINES_ADJACENCY:
		name = "GL_LINES_ADJACENCY"
	case RGB16_SNORM:
		name = "GL_RGB16_SNORM"
	case COMPRESSED_SRGB8_ALPHA8_ETC2_EAC:
		name = "GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC"
	case RED_INTEGER:
		name = "GL_RED_INTEGER"
	case SRC1_ALPHA:
		name = "GL_SRC1_ALPHA"
	case MAX_PROGRAM_TEXEL_OFFSET:
		name = "GL_MAX_PROGRAM_TEXEL_OFFSET"
	case SRGB_ALPHA:
		name = "GL_SRGB_ALPHA"
	case UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY:
		name = "GL_UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY"
	case TEXTURE_CUBE_MAP:
		name = "GL_TEXTURE_CUBE_MAP"
	case IMAGE_CLASS_11_11_10:
		name = "GL_IMAGE_CLASS_11_11_10"
	case ACTIVE_UNIFORM_BLOCKS:
		name = "GL_ACTIVE_UNIFORM_BLOCKS"
	case DEPTH_COMPONENT32F:
		name = "GL_DEPTH_COMPONENT32F"
	case INT_IMAGE_1D:
		name = "GL_INT_IMAGE_1D"
	case UNIFORM_BUFFER_BINDING:
		name = "GL_UNIFORM_BUFFER_BINDING"
	case PROXY_TEXTURE_2D_ARRAY:
		name = "GL_PROXY_TEXTURE_2D_ARRAY"
	case TEXTURE4:
		name = "GL_TEXTURE4"
	case COLOR_WRITEMASK:
		name = "GL_COLOR_WRITEMASK"
	case INT_IMAGE_2D_ARRAY:
		name = "GL_INT_IMAGE_2D_ARRAY"
	case DEPTH_ATTACHMENT:
		name = "GL_DEPTH_ATTACHMENT"
	case VERTEX_SHADER_BIT:
		name = "GL_VERTEX_SHADER_BIT"
	case UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER:
		name = "GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER"
	case BUFFER_SIZE:
		name = "GL_BUFFER_SIZE"
	case UNSIGNED_INT_ATOMIC_COUNTER:
		name = "GL_UNSIGNED_INT_ATOMIC_COUNTER"
	case TEXTURE_BUFFER:
		name = "GL_TEXTURE_BUFFER"
	case IMAGE_PIXEL_FORMAT:
		name = "GL_IMAGE_PIXEL_FORMAT"
	case COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT:
		name = "GL_COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT"
	case RENDERBUFFER_WIDTH:
		name = "GL_RENDERBUFFER_WIDTH"
	case IS_PER_PATCH:
		name = "GL_IS_PER_PATCH"
	case DST_COLOR:
		name = "GL_DST_COLOR"
	case BUFFER_MAP_OFFSET:
		name = "GL_BUFFER_MAP_OFFSET"
	case DRAW_BUFFER1:
		name = "GL_DRAW_BUFFER1"
	case CLAMP_TO_EDGE:
		name = "GL_CLAMP_TO_EDGE"
	case SAMPLES_PASSED:
		name = "GL_SAMPLES_PASSED"
	case QUERY_BY_REGION_WAIT:
		name = "GL_QUERY_BY_REGION_WAIT"
	case MAX_GEOMETRY_OUTPUT_COMPONENTS:
		name = "GL_MAX_GEOMETRY_OUTPUT_COMPONENTS"
	case TEXTURE_BLUE_SIZE:
		name = "GL_TEXTURE_BLUE_SIZE"
	case TEXTURE_BINDING_RECTANGLE:
		name = "GL_TEXTURE_BINDING_RECTANGLE"
	case INT_SAMPLER_2D_MULTISAMPLE_ARRAY:
		name = "GL_INT_SAMPLER_2D_MULTISAMPLE_ARRAY"
	case DECR_WRAP:
		name = "GL_DECR_WRAP"
	case NUM_SAMPLE_COUNTS:
		name = "GL_NUM_SAMPLE_COUNTS"
	case FLOAT_VEC2:
		name = "GL_FLOAT_VEC2"
	case MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS:
		name = "GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS"
	case GEOMETRY_SHADER_BIT:
		name = "GL_GEOMETRY_SHADER_BIT"
	case UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER:
		name = "GL_UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER"
	case COMPRESSED_RGBA_S3TC_DXT1_EXT:
		name = "GL_COMPRESSED_RGBA_S3TC_DXT1_EXT"
	case INTERLEAVED_ATTRIBS:
		name = "GL_INTERLEAVED_ATTRIBS"
	case TEXTURE6:
		name = "GL_TEXTURE6"
	case RGBA2:
		name = "GL_RGBA2"
	case MAX_TEXTURE_LOD_BIAS:
		name = "GL_MAX_TEXTURE_LOD_BIAS"
	case UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER:
		name = "GL_UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER"
	case OR:
		name = "GL_OR"
	case TEXTURE_COMPRESSED_IMAGE_SIZE:
		name = "GL_TEXTURE_COMPRESSED_IMAGE_SIZE"
	case STENCIL_BACK_VALUE_MASK:
		name = "GL_STENCIL_BACK_VALUE_MASK"
	case POINT_FADE_THRESHOLD_SIZE:
		name = "GL_POINT_FADE_THRESHOLD_SIZE"
	case TESS_EVALUATION_SUBROUTINE_UNIFORM:
		name = "GL_TESS_EVALUATION_SUBROUTINE_UNIFORM"
	case BOOL:
		name = "GL_BOOL"
	case RGB32F:
		name = "GL_RGB32F"
	case MAX_TESS_EVALUATION_UNIFORM_COMPONENTS:
		name = "GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS"
	case FRAMEBUFFER_ATTACHMENT_GREEN_SIZE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_GREEN_SIZE"
	case TRANSFORM_FEEDBACK_BUFFER_BINDING:
		name = "GL_TRANSFORM_FEEDBACK_BUFFER_BINDING"
	case MAX_GEOMETRY_UNIFORM_COMPONENTS:
		name = "GL_MAX_GEOMETRY_UNIFORM_COMPONENTS"
	case GEOMETRY_SUBROUTINE:
		name = "GL_GEOMETRY_SUBROUTINE"
	case VALIDATE_STATUS:
		name = "GL_VALIDATE_STATUS"
	case FULL_SUPPORT:
		name = "GL_FULL_SUPPORT"
	case RENDERER:
		name = "GL_RENDERER"
	case UNSIGNED_INT_SAMPLER_CUBE:
		name = "GL_UNSIGNED_INT_SAMPLER_CUBE"
	case UNSIGNED_INT_10F_11F_11F_REV:
		name = "GL_UNSIGNED_INT_10F_11F_11F_REV"
	case INT_SAMPLER_BUFFER:
		name = "GL_INT_SAMPLER_BUFFER"
	case BUFFER_UPDATE_BARRIER_BIT:
		name = "GL_BUFFER_UPDATE_BARRIER_BIT"
	case MAX_UNIFORM_LOCATIONS:
		name = "GL_MAX_UNIFORM_LOCATIONS"
	case FRAMEBUFFER:
		name = "GL_FRAMEBUFFER"
	case FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE"
	case TEXTURE_RED_SIZE:
		name = "GL_TEXTURE_RED_SIZE"
	case RG32I:
		name = "GL_RG32I"
	case UNSIGNED_INT_2_10_10_10_REV:
		name = "GL_UNSIGNED_INT_2_10_10_10_REV"
	case VERTEX_ARRAY_BINDING:
		name = "GL_VERTEX_ARRAY_BINDING"
	case IMAGE_CLASS_1_X_16:
		name = "GL_IMAGE_CLASS_1_X_16"
	case MAX_FRAGMENT_INPUT_COMPONENTS:
		name = "GL_MAX_FRAGMENT_INPUT_COMPONENTS"
	case NEAREST:
		name = "GL_NEAREST"
	case SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE:
		name = "GL_SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE"
	case MIN_MAP_BUFFER_ALIGNMENT:
		name = "GL_MIN_MAP_BUFFER_ALIGNMENT"
	case TEXTURE2:
		name = "GL_TEXTURE2"
	case MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS:
		name = "GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS"
	case INTERNALFORMAT_RED_TYPE:
		name = "GL_INTERNALFORMAT_RED_TYPE"
	case RGB8I:
		name = "GL_RGB8I"
	case PROXY_TEXTURE_RECTANGLE:
		name = "GL_PROXY_TEXTURE_RECTANGLE"
	case STENCIL_INDEX8:
		name = "GL_STENCIL_INDEX8"
	case RGB5:
		name = "GL_RGB5"
	case GEOMETRY_SUBROUTINE_UNIFORM:
		name = "GL_GEOMETRY_SUBROUTINE_UNIFORM"
	case MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS:
		name = "GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS"
	case DISPATCH_INDIRECT_BUFFER_BINDING:
		name = "GL_DISPATCH_INDIRECT_BUFFER_BINDING"
	case COPY_READ_BUFFER_BINDING:
		name = "GL_COPY_READ_BUFFER_BINDING"
	case DEPTH_RANGE:
		name = "GL_DEPTH_RANGE"
	case UNIFORM:
		name = "GL_UNIFORM"
	case TEXTURE_GATHER:
		name = "GL_TEXTURE_GATHER"
	case CONDITION_SATISFIED:
		name = "GL_CONDITION_SATISFIED"
	case UNSIGNED_INT_10_10_10_2:
		name = "GL_UNSIGNED_INT_10_10_10_2"
	case REFERENCED_BY_GEOMETRY_SHADER:
		name = "GL_REFERENCED_BY_GEOMETRY_SHADER"
	case NUM_SHADING_LANGUAGE_VERSIONS:
		name = "GL_NUM_SHADING_LANGUAGE_VERSIONS"
	case SRGB8_ALPHA8_EXT:
		name = "GL_SRGB8_ALPHA8_EXT"
	case RG16F:
		name = "GL_RG16F"
	case MAX_COMBINED_ATOMIC_COUNTERS:
		name = "GL_MAX_COMBINED_ATOMIC_COUNTERS"
	case PROGRAM_BINARY_RETRIEVABLE_HINT:
		name = "GL_PROGRAM_BINARY_RETRIEVABLE_HINT"
	case MAX_COMPUTE_LOCAL_INVOCATIONS:
		name = "GL_MAX_COMPUTE_LOCAL_INVOCATIONS"
	case MAX_IMAGE_SAMPLES:
		name = "GL_MAX_IMAGE_SAMPLES"
	case UPPER_LEFT:
		name = "GL_UPPER_LEFT"
	case PRIMITIVE_RESTART_FIXED_INDEX:
		name = "GL_PRIMITIVE_RESTART_FIXED_INDEX"
	case DRAW_BUFFER:
		name = "GL_DRAW_BUFFER"
	case VIEW_CLASS_RGTC2_RG:
		name = "GL_VIEW_CLASS_RGTC2_RG"
	case RENDERBUFFER_INTERNAL_FORMAT:
		name = "GL_RENDERBUFFER_INTERNAL_FORMAT"
	case ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER:
		name = "GL_ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER"
	case TEXTURE_CUBE_MAP_NEGATIVE_X:
		name = "GL_TEXTURE_CUBE_MAP_NEGATIVE_X"
	case PROXY_TEXTURE_2D:
		name = "GL_PROXY_TEXTURE_2D"
	case INTERNALFORMAT_SUPPORTED:
		name = "GL_INTERNALFORMAT_SUPPORTED"
	case RASTERIZER_DISCARD:
		name = "GL_RASTERIZER_DISCARD"
	case CLEAR:
		name = "GL_CLEAR"
	case DYNAMIC_COPY:
		name = "GL_DYNAMIC_COPY"
	case COMPRESSED_SIGNED_RG11_EAC:
		name = "GL_COMPRESSED_SIGNED_RG11_EAC"
	case PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY:
		name = "GL_PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY"
	case TEXTURE_1D_ARRAY:
		name = "GL_TEXTURE_1D_ARRAY"
	case SLUMINANCE8_ALPHA8_EXT:
		name = "GL_SLUMINANCE8_ALPHA8_EXT"
	case MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS:
		name = "GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS"
	case MAX_UNIFORM_BLOCK_SIZE:
		name = "GL_MAX_UNIFORM_BLOCK_SIZE"
	case STENCIL_INDEX:
		name = "GL_STENCIL_INDEX"
	case TEXTURE_MAG_FILTER:
		name = "GL_TEXTURE_MAG_FILTER"
	case INT_IMAGE_2D:
		name = "GL_INT_IMAGE_2D"
	case ONE_MINUS_SRC_COLOR:
		name = "GL_ONE_MINUS_SRC_COLOR"
	case TEXTURE14:
		name = "GL_TEXTURE14"
	case MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS:
		name = "GL_MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS"
	case STENCIL_PASS_DEPTH_PASS:
		name = "GL_STENCIL_PASS_DEPTH_PASS"
	case VIEW_CLASS_48_BITS:
		name = "GL_VIEW_CLASS_48_BITS"
	case DEPTH:
		name = "GL_DEPTH"
	case DRAW_BUFFER6:
		name = "GL_DRAW_BUFFER6"
	case RG_INTEGER:
		name = "GL_RG_INTEGER"
	case TIMEOUT_EXPIRED:
		name = "GL_TIMEOUT_EXPIRED"
	case FRONT_FACE:
		name = "GL_FRONT_FACE"
	case INT_SAMPLER_1D:
		name = "GL_INT_SAMPLER_1D"
	case TEXTURE_SWIZZLE_G:
		name = "GL_TEXTURE_SWIZZLE_G"
	case INT_SAMPLER_2D:
		name = "GL_INT_SAMPLER_2D"
	case DEBUG_TYPE_OTHER:
		name = "GL_DEBUG_TYPE_OTHER"
	case SLUMINANCE_ALPHA_EXT:
		name = "GL_SLUMINANCE_ALPHA_EXT"
	case LINEAR_MIPMAP_NEAREST:
		name = "GL_LINEAR_MIPMAP_NEAREST"
	case MAX_VERTEX_ATOMIC_COUNTERS:
		name = "GL_MAX_VERTEX_ATOMIC_COUNTERS"
	case COLOR_ATTACHMENT14:
		name = "GL_COLOR_ATTACHMENT14"
	case BOOL_VEC2:
		name = "GL_BOOL_VEC2"
	case TEXTURE_GREEN_TYPE:
		name = "GL_TEXTURE_GREEN_TYPE"
	case DECR:
		name = "GL_DECR"
	case FRAMEBUFFER_INCOMPLETE_ATTACHMENT:
		name = "GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT"
	case TRANSFORM_FEEDBACK_BUFFER_START:
		name = "GL_TRANSFORM_FEEDBACK_BUFFER_START"
	case RG16_SNORM:
		name = "GL_RG16_SNORM"
	case ELEMENT_ARRAY_BUFFER:
		name = "GL_ELEMENT_ARRAY_BUFFER"
	case TRIANGLE_STRIP_ADJACENCY:
		name = "GL_TRIANGLE_STRIP_ADJACENCY"
	case INTERNALFORMAT_SHARED_SIZE:
		name = "GL_INTERNALFORMAT_SHARED_SIZE"
	case TEXTURE_COMPARE_MODE:
		name = "GL_TEXTURE_COMPARE_MODE"
	case NO_ERROR:
		name = "GL_NO_ERROR"
	case COMPUTE_LOCAL_WORK_SIZE:
		name = "GL_COMPUTE_LOCAL_WORK_SIZE"
	case CLIP_DISTANCE7:
		name = "GL_CLIP_DISTANCE7"
	case DEBUG_OUTPUT_SYNCHRONOUS:
		name = "GL_DEBUG_OUTPUT_SYNCHRONOUS"
	case TEXTURE_MIN_LOD:
		name = "GL_TEXTURE_MIN_LOD"
	case MAP_UNSYNCHRONIZED_BIT:
		name = "GL_MAP_UNSYNCHRONIZED_BIT"
	case ACTIVE_SUBROUTINE_UNIFORMS:
		name = "GL_ACTIVE_SUBROUTINE_UNIFORMS"
	case ANY_SAMPLES_PASSED:
		name = "GL_ANY_SAMPLES_PASSED"
	case SHADER_STORAGE_BUFFER_BINDING:
		name = "GL_SHADER_STORAGE_BUFFER_BINDING"
	case VERTEX_TEXTURE:
		name = "GL_VERTEX_TEXTURE"
	case QUERY_BY_REGION_NO_WAIT:
		name = "GL_QUERY_BY_REGION_NO_WAIT"
	case TESS_EVALUATION_SHADER_BIT:
		name = "GL_TESS_EVALUATION_SHADER_BIT"
	case ATTACHED_SHADERS:
		name = "GL_ATTACHED_SHADERS"
	case UNIFORM_BLOCK_INDEX:
		name = "GL_UNIFORM_BLOCK_INDEX"
	case COMPRESSED_RGBA_S3TC_DXT3_EXT:
		name = "GL_COMPRESSED_RGBA_S3TC_DXT3_EXT"
	case MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS:
		name = "GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS"
	case ATOMIC_COUNTER_BUFFER_BINDING:
		name = "GL_ATOMIC_COUNTER_BUFFER_BINDING"
	case SAMPLER_2D_ARRAY_SHADOW:
		name = "GL_SAMPLER_2D_ARRAY_SHADOW"
	case INVERT:
		name = "GL_INVERT"
	case BGRA:
		name = "GL_BGRA"
	case NOR:
		name = "GL_NOR"
	case GET_TEXTURE_IMAGE_TYPE:
		name = "GL_GET_TEXTURE_IMAGE_TYPE"
	case TEXTURE_VIEW_MIN_LEVEL:
		name = "GL_TEXTURE_VIEW_MIN_LEVEL"
	case ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH:
		name = "GL_ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH"
	case ATOMIC_COUNTER_BUFFER_SIZE:
		name = "GL_ATOMIC_COUNTER_BUFFER_SIZE"
	case DRAW_BUFFER14:
		name = "GL_DRAW_BUFFER14"
	case UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER:
		name = "GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER"
	case R8_SNORM:
		name = "GL_R8_SNORM"
	case SAMPLER_2D_MULTISAMPLE:
		name = "GL_SAMPLER_2D_MULTISAMPLE"
	case POLYGON_MODE:
		name = "GL_POLYGON_MODE"
	case UNSIGNED_NORMALIZED:
		name = "GL_UNSIGNED_NORMALIZED"
	case IMAGE_CLASS_10_10_10_2:
		name = "GL_IMAGE_CLASS_10_10_10_2"
	case DRAW_INDIRECT_BUFFER_BINDING:
		name = "GL_DRAW_INDIRECT_BUFFER_BINDING"
	case COMPRESSED_RG_RGTC2:
		name = "GL_COMPRESSED_RG_RGTC2"
	case PROGRAM_BINARY_LENGTH:
		name = "GL_PROGRAM_BINARY_LENGTH"
	case TESS_GEN_POINT_MODE:
		name = "GL_TESS_GEN_POINT_MODE"
	case UNIFORM_BUFFER_SIZE:
		name = "GL_UNIFORM_BUFFER_SIZE"
	case COMMAND_BARRIER_BIT:
		name = "GL_COMMAND_BARRIER_BIT"
	case READ_FRAMEBUFFER:
		name = "GL_READ_FRAMEBUFFER"
	case MAX_RECTANGLE_TEXTURE_SIZE:
		name = "GL_MAX_RECTANGLE_TEXTURE_SIZE"
	case UNIFORM_SIZE:
		name = "GL_UNIFORM_SIZE"
	case IS_ROW_MAJOR:
		name = "GL_IS_ROW_MAJOR"
	case TEXTURE_GATHER_SHADOW:
		name = "GL_TEXTURE_GATHER_SHADOW"
	case ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER:
		name = "GL_ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER"
	case MAX_COMPUTE_SHARED_MEMORY_SIZE:
		name = "GL_MAX_COMPUTE_SHARED_MEMORY_SIZE"
	case TESS_CONTROL_SUBROUTINE:
		name = "GL_TESS_CONTROL_SUBROUTINE"
	case RGBA_INTEGER:
		name = "GL_RGBA_INTEGER"
	case TEXTURE_UPDATE_BARRIER_BIT:
		name = "GL_TEXTURE_UPDATE_BARRIER_BIT"
	case UNSIGNED_INT_IMAGE_2D_MULTISAMPLE:
		name = "GL_UNSIGNED_INT_IMAGE_2D_MULTISAMPLE"
	case STEREO:
		name = "GL_STEREO"
	case FRAGMENT_SHADER_DERIVATIVE_HINT:
		name = "GL_FRAGMENT_SHADER_DERIVATIVE_HINT"
	case INT_IMAGE_2D_MULTISAMPLE:
		name = "GL_INT_IMAGE_2D_MULTISAMPLE"
	case RGB32I:
		name = "GL_RGB32I"
	case MAX_NUM_COMPATIBLE_SUBROUTINES:
		name = "GL_MAX_NUM_COMPATIBLE_SUBROUTINES"
	case TEXTURE_BORDER_COLOR:
		name = "GL_TEXTURE_BORDER_COLOR"
	case UNIFORM_BUFFER:
		name = "GL_UNIFORM_BUFFER"
	case RG8UI:
		name = "GL_RG8UI"
	case SIGNED_NORMALIZED:
		name = "GL_SIGNED_NORMALIZED"
	case INT_2_10_10_10_REV:
		name = "GL_INT_2_10_10_10_REV"
	case MAX_VERTEX_SHADER_STORAGE_BLOCKS:
		name = "GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS"
	case MAX_VARYING_COMPONENTS:
		name = "GL_MAX_VARYING_COMPONENTS"
	case FRONT:
		name = "GL_FRONT"
	case NUM_ACTIVE_VARIABLES:
		name = "GL_NUM_ACTIVE_VARIABLES"
	case VERSION:
		name = "GL_VERSION"
	case MAX_TESS_CONTROL_INPUT_COMPONENTS:
		name = "GL_MAX_TESS_CONTROL_INPUT_COMPONENTS"
	case DEPTH_COMPONENT32:
		name = "GL_DEPTH_COMPONENT32"
	case IMAGE_TEXEL_SIZE:
		name = "GL_IMAGE_TEXEL_SIZE"
	case TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY:
		name = "GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY"
	case TEXTURE_SWIZZLE_R:
		name = "GL_TEXTURE_SWIZZLE_R"
	case SRGB_READ:
		name = "GL_SRGB_READ"
	case COLOR_BUFFER_BIT:
		name = "GL_COLOR_BUFFER_BIT"
	case FRAMEBUFFER_BARRIER_BIT:
		name = "GL_FRAMEBUFFER_BARRIER_BIT"
	case TEXTURE_ALPHA_TYPE:
		name = "GL_TEXTURE_ALPHA_TYPE"
	case ACTIVE_ATTRIBUTE_MAX_LENGTH:
		name = "GL_ACTIVE_ATTRIBUTE_MAX_LENGTH"
	case RGBA12:
		name = "GL_RGBA12"
	case IMAGE_2D_RECT:
		name = "GL_IMAGE_2D_RECT"
	case IMAGE_CLASS_4_X_32:
		name = "GL_IMAGE_CLASS_4_X_32"
	case FLOAT_MAT3x2:
		name = "GL_FLOAT_MAT3x2"
	case MAX_UNIFORM_BUFFER_BINDINGS:
		name = "GL_MAX_UNIFORM_BUFFER_BINDINGS"
	case MAX_VERTEX_IMAGE_UNIFORMS:
		name = "GL_MAX_VERTEX_IMAGE_UNIFORMS"
	case CLIP_DISTANCE0:
		name = "GL_CLIP_DISTANCE0"
	case RENDERBUFFER_DEPTH_SIZE:
		name = "GL_RENDERBUFFER_DEPTH_SIZE"
	case COMPRESSED_TEXTURE_FORMATS:
		name = "GL_COMPRESSED_TEXTURE_FORMATS"
	case TEXTURE23:
		name = "GL_TEXTURE23"
	case TEXTURE22:
		name = "GL_TEXTURE22"
	case COMPRESSED_SRGB_ALPHA_EXT:
		name = "GL_COMPRESSED_SRGB_ALPHA_EXT"
	case TESS_CONTROL_TEXTURE:
		name = "GL_TESS_CONTROL_TEXTURE"
	case COMPRESSED_RED_RGTC1:
		name = "GL_COMPRESSED_RED_RGTC1"
	case DRAW_BUFFER7:
		name = "GL_DRAW_BUFFER7"
	case LINE_WIDTH_GRANULARITY:
		name = "GL_LINE_WIDTH_GRANULARITY"
	case TEXTURE_CUBE_MAP_POSITIVE_X:
		name = "GL_TEXTURE_CUBE_MAP_POSITIVE_X"
	case POLYGON_SMOOTH:
		name = "GL_POLYGON_SMOOTH"
	case GEQUAL:
		name = "GL_GEQUAL"
	case RENDERBUFFER_BLUE_SIZE:
		name = "GL_RENDERBUFFER_BLUE_SIZE"
	case COMPRESSED_RGBA_S3TC_DXT5_EXT:
		name = "GL_COMPRESSED_RGBA_S3TC_DXT5_EXT"
	case MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS:
		name = "GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS"
	case PRIMITIVE_RESTART:
		name = "GL_PRIMITIVE_RESTART"
	case SYNC_GPU_COMMANDS_COMPLETE:
		name = "GL_SYNC_GPU_COMMANDS_COMPLETE"
	case MAX_GEOMETRY_SHADER_INVOCATIONS:
		name = "GL_MAX_GEOMETRY_SHADER_INVOCATIONS"
	case SAMPLER_CUBE_MAP_ARRAY_SHADOW:
		name = "GL_SAMPLER_CUBE_MAP_ARRAY_SHADOW"
	case TEXTURE26:
		name = "GL_TEXTURE26"
	case BLEND_SRC_ALPHA:
		name = "GL_BLEND_SRC_ALPHA"
	case R32F:
		name = "GL_R32F"
	case IMAGE_1D:
		name = "GL_IMAGE_1D"
	case ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER:
		name = "GL_ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER"
	case TEXTURE_COMPRESSED_BLOCK_SIZE:
		name = "GL_TEXTURE_COMPRESSED_BLOCK_SIZE"
	case AUTO_GENERATE_MIPMAP:
		name = "GL_AUTO_GENERATE_MIPMAP"
	case TEXTURE_BINDING_CUBE_MAP:
		name = "GL_TEXTURE_BINDING_CUBE_MAP"
	case VIEW_CLASS_S3TC_DXT5_RGBA:
		name = "GL_VIEW_CLASS_S3TC_DXT5_RGBA"
	case MAX_CUBE_MAP_TEXTURE_SIZE:
		name = "GL_MAX_CUBE_MAP_TEXTURE_SIZE"
	case STENCIL_FAIL:
		name = "GL_STENCIL_FAIL"
	case INT_SAMPLER_CUBE:
		name = "GL_INT_SAMPLER_CUBE"
	case ISOLINES:
		name = "GL_ISOLINES"
	case UNIFORM_BLOCK_DATA_SIZE:
		name = "GL_UNIFORM_BLOCK_DATA_SIZE"
	case COMPRESSED_RGBA8_ETC2_EAC:
		name = "GL_COMPRESSED_RGBA8_ETC2_EAC"
	case TEXTURE_WRAP_S:
		name = "GL_TEXTURE_WRAP_S"
	case MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS:
		name = "GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS"
	case DEBUG_TYPE_PORTABILITY:
		name = "GL_DEBUG_TYPE_PORTABILITY"
	case EQUIV:
		name = "GL_EQUIV"
	case TEXTURE_SWIZZLE_RGBA:
		name = "GL_TEXTURE_SWIZZLE_RGBA"
	case ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH:
		name = "GL_ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH"
	case BUFFER_MAPPED:
		name = "GL_BUFFER_MAPPED"
	case STENCIL:
		name = "GL_STENCIL"
	case PACK_SKIP_IMAGES:
		name = "GL_PACK_SKIP_IMAGES"
	case TRIANGLE_STRIP:
		name = "GL_TRIANGLE_STRIP"
	case VIEW_CLASS_S3TC_DXT3_RGBA:
		name = "GL_VIEW_CLASS_S3TC_DXT3_RGBA"
	case TEXTURE_BINDING_2D_MULTISAMPLE:
		name = "GL_TEXTURE_BINDING_2D_MULTISAMPLE"
	case INTERNALFORMAT_DEPTH_TYPE:
		name = "GL_INTERNALFORMAT_DEPTH_TYPE"
	case TEXTURE30:
		name = "GL_TEXTURE30"
	case VERTEX_ATTRIB_ARRAY_STRIDE:
		name = "GL_VERTEX_ATTRIB_ARRAY_STRIDE"
	case VIEW_CLASS_16_BITS:
		name = "GL_VIEW_CLASS_16_BITS"
	case MAX_SHADER_STORAGE_BLOCK_SIZE:
		name = "GL_MAX_SHADER_STORAGE_BLOCK_SIZE"
	case HIGH_FLOAT:
		name = "GL_HIGH_FLOAT"
	case TEXTURE_FIXED_SAMPLE_LOCATIONS:
		name = "GL_TEXTURE_FIXED_SAMPLE_LOCATIONS"
	case MAX_VERTEX_ATTRIBS:
		name = "GL_MAX_VERTEX_ATTRIBS"
	case DEPTH_CLEAR_VALUE:
		name = "GL_DEPTH_CLEAR_VALUE"
	case MAX_FRAMEBUFFER_WIDTH:
		name = "GL_MAX_FRAMEBUFFER_WIDTH"
	case MAX_3D_TEXTURE_SIZE:
		name = "GL_MAX_3D_TEXTURE_SIZE"
	case VERTEX_SUBROUTINE:
		name = "GL_VERTEX_SUBROUTINE"
	case UNSIGNED_BYTE_2_3_3_REV:
		name = "GL_UNSIGNED_BYTE_2_3_3_REV"
	case PACK_SKIP_ROWS:
		name = "GL_PACK_SKIP_ROWS"
	case LINEAR:
		name = "GL_LINEAR"
	case RGBA:
		name = "GL_RGBA"
	case IMAGE_COMPATIBILITY_CLASS:
		name = "GL_IMAGE_COMPATIBILITY_CLASS"
	case KEEP:
		name = "GL_KEEP"
	case COMPUTE_SUBROUTINE_UNIFORM:
		name = "GL_COMPUTE_SUBROUTINE_UNIFORM"
	case HIGH_INT:
		name = "GL_HIGH_INT"
	case TEXTURE8:
		name = "GL_TEXTURE8"
	case TEXTURE_DEPTH_TYPE:
		name = "GL_TEXTURE_DEPTH_TYPE"
	case COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
		name = "GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2"
	case COMPRESSED_RGB:
		name = "GL_COMPRESSED_RGB"
	case TEXTURE_BINDING_3D:
		name = "GL_TEXTURE_BINDING_3D"
	case FRAGMENT_SUBROUTINE_UNIFORM:
		name = "GL_FRAGMENT_SUBROUTINE_UNIFORM"
	case ATOMIC_COUNTER_BUFFER:
		name = "GL_ATOMIC_COUNTER_BUFFER"
	case TEXTURE_IMAGE_TYPE:
		name = "GL_TEXTURE_IMAGE_TYPE"
	case INT_SAMPLER_2D_RECT:
		name = "GL_INT_SAMPLER_2D_RECT"
	case MAX_VARYING_VECTORS:
		name = "GL_MAX_VARYING_VECTORS"
	case IMAGE_3D:
		name = "GL_IMAGE_3D"
	case DOUBLE_MAT2:
		name = "GL_DOUBLE_MAT2"
	case ONE_MINUS_DST_COLOR:
		name = "GL_ONE_MINUS_DST_COLOR"
	case VERTEX_ATTRIB_ARRAY_LONG:
		name = "GL_VERTEX_ATTRIB_ARRAY_LONG"
	case AND:
		name = "GL_AND"
	case GEOMETRY_SHADER_INVOCATIONS:
		name = "GL_GEOMETRY_SHADER_INVOCATIONS"
	case TYPE:
		name = "GL_TYPE"
	case SRC_COLOR:
		name = "GL_SRC_COLOR"
	case MAX_SUBROUTINES:
		name = "GL_MAX_SUBROUTINES"
	case MAX_SHADER_STORAGE_BUFFER_BINDINGS:
		name = "GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS"
	case ATOMIC_COUNTER_BUFFER_INDEX:
		name = "GL_ATOMIC_COUNTER_BUFFER_INDEX"
	case RGB10:
		name = "GL_RGB10"
	case POLYGON_SMOOTH_HINT:
		name = "GL_POLYGON_SMOOTH_HINT"
	case MAX_COMBINED_SHADER_STORAGE_BLOCKS:
		name = "GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS"
	case FIXED:
		name = "GL_FIXED"
	case RG:
		name = "GL_RG"
	case SRGB_EXT:
		name = "GL_SRGB_EXT"
	case TEXTURE_COMPRESSED_BLOCK_HEIGHT:
		name = "GL_TEXTURE_COMPRESSED_BLOCK_HEIGHT"
	case INTERNALFORMAT_BLUE_TYPE:
		name = "GL_INTERNALFORMAT_BLUE_TYPE"
	case UNSIGNED_INT_VEC3:
		name = "GL_UNSIGNED_INT_VEC3"
	case VIEW_CLASS_8_BITS:
		name = "GL_VIEW_CLASS_8_BITS"
	case MAX_DEBUG_GROUP_STACK_DEPTH:
		name = "GL_MAX_DEBUG_GROUP_STACK_DEPTH"
	case SAMPLE_COVERAGE:
		name = "GL_SAMPLE_COVERAGE"
	case NUM_COMPATIBLE_SUBROUTINES:
		name = "GL_NUM_COMPATIBLE_SUBROUTINES"
	case MAX_WIDTH:
		name = "GL_MAX_WIDTH"
	case READ_PIXELS_FORMAT:
		name = "GL_READ_PIXELS_FORMAT"
	case SAMPLER:
		name = "GL_SAMPLER"
	case DISPATCH_INDIRECT_BUFFER:
		name = "GL_DISPATCH_INDIRECT_BUFFER"
	case MAX_GEOMETRY_INPUT_COMPONENTS:
		name = "GL_MAX_GEOMETRY_INPUT_COMPONENTS"
	case FLOAT_32_UNSIGNED_INT_24_8_REV:
		name = "GL_FLOAT_32_UNSIGNED_INT_24_8_REV"
	case NAME_LENGTH:
		name = "GL_NAME_LENGTH"
	case TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:
		name = "GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN"
	case TEXTURE_COMPRESSED:
		name = "GL_TEXTURE_COMPRESSED"
	case UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER:
		name = "GL_UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER"
	case MAX_TESS_CONTROL_UNIFORM_COMPONENTS:
		name = "GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS"
	case RG8I:
		name = "GL_RG8I"
	case SRC1_COLOR:
		name = "GL_SRC1_COLOR"
	case DOUBLE_MAT4x2:
		name = "GL_DOUBLE_MAT4x2"
	case LINE:
		name = "GL_LINE"
	case DEBUG_SOURCE_APPLICATION:
		name = "GL_DEBUG_SOURCE_APPLICATION"
	case SAMPLE_SHADING:
		name = "GL_SAMPLE_SHADING"
	case TEXTURE_VIEW_NUM_LEVELS:
		name = "GL_TEXTURE_VIEW_NUM_LEVELS"
	case ACTIVE_RESOURCES:
		name = "GL_ACTIVE_RESOURCES"
	case DEBUG_TYPE_PUSH_GROUP:
		name = "GL_DEBUG_TYPE_PUSH_GROUP"
	case UNIFORM_NAME_LENGTH:
		name = "GL_UNIFORM_NAME_LENGTH"
	case MAX_GEOMETRY_TEXTURE_IMAGE_UNITS:
		name = "GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS"
	case FLOAT_MAT2:
		name = "GL_FLOAT_MAT2"
	case BLEND_SRC_RGB:
		name = "GL_BLEND_SRC_RGB"
	case TEXTURE27:
		name = "GL_TEXTURE27"
	case FIXED_ONLY:
		name = "GL_FIXED_ONLY"
	case TRIANGLES_ADJACENCY:
		name = "GL_TRIANGLES_ADJACENCY"
	case R16_SNORM:
		name = "GL_R16_SNORM"
	case DOUBLE_MAT4:
		name = "GL_DOUBLE_MAT4"
	case STREAM_READ:
		name = "GL_STREAM_READ"
	case FRONT_RIGHT:
		name = "GL_FRONT_RIGHT"
	case MAX_TRANSFORM_FEEDBACK_BUFFERS:
		name = "GL_MAX_TRANSFORM_FEEDBACK_BUFFERS"
	case EQUAL:
		name = "GL_EQUAL"
	case UNSIGNED_SHORT_4_4_4_4_REV:
		name = "GL_UNSIGNED_SHORT_4_4_4_4_REV"
	case MAX_DEPTH:
		name = "GL_MAX_DEPTH"
	case RGBA32UI:
		name = "GL_RGBA32UI"
	case MAX_TESS_GEN_LEVEL:
		name = "GL_MAX_TESS_GEN_LEVEL"
	case TEXTURE_SHADOW:
		name = "GL_TEXTURE_SHADOW"
	case RENDERBUFFER:
		name = "GL_RENDERBUFFER"
	case DEBUG_TYPE_POP_GROUP:
		name = "GL_DEBUG_TYPE_POP_GROUP"
	case MAJOR_VERSION:
		name = "GL_MAJOR_VERSION"
	case COMPRESSED_SIGNED_R11_EAC:
		name = "GL_COMPRESSED_SIGNED_R11_EAC"
	case FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS:
		name = "GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS"
	case TEXTURE_MAX_LEVEL:
		name = "GL_TEXTURE_MAX_LEVEL"
	case IMAGE_1D_ARRAY:
		name = "GL_IMAGE_1D_ARRAY"
	case STENCIL_INDEX16:
		name = "GL_STENCIL_INDEX16"
	case FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER:
		name = "GL_FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER"
	case TEXTURE17:
		name = "GL_TEXTURE17"
	case ATOMIC_COUNTER_BUFFER_DATA_SIZE:
		name = "GL_ATOMIC_COUNTER_BUFFER_DATA_SIZE"
	case RG8:
		name = "GL_RG8"
	case RG_SNORM:
		name = "GL_RG_SNORM"
	case ACTIVE_ATOMIC_COUNTER_BUFFERS:
		name = "GL_ACTIVE_ATOMIC_COUNTER_BUFFERS"
	case UNPACK_LSB_FIRST:
		name = "GL_UNPACK_LSB_FIRST"
	case SEPARATE_ATTRIBS:
		name = "GL_SEPARATE_ATTRIBS"
	case SMOOTH_LINE_WIDTH_RANGE:
		name = "GL_SMOOTH_LINE_WIDTH_RANGE"
	case VIEW_CLASS_24_BITS:
		name = "GL_VIEW_CLASS_24_BITS"
	case MAX_FRAMEBUFFER_HEIGHT:
		name = "GL_MAX_FRAMEBUFFER_HEIGHT"
	case RGBA8:
		name = "GL_RGBA8"
	case RGB_SNORM:
		name = "GL_RGB_SNORM"
	case SCISSOR_BOX:
		name = "GL_SCISSOR_BOX"
	case DRAW_BUFFER10:
		name = "GL_DRAW_BUFFER10"
	case TEXTURE19:
		name = "GL_TEXTURE19"
	case TEXTURE_RED_TYPE:
		name = "GL_TEXTURE_RED_TYPE"
	case DEBUG_CALLBACK_USER_PARAM:
		name = "GL_DEBUG_CALLBACK_USER_PARAM"
	case MAX_VERTEX_UNIFORM_BLOCKS:
		name = "GL_MAX_VERTEX_UNIFORM_BLOCKS"
	case IMAGE_BINDING_LAYER:
		name = "GL_IMAGE_BINDING_LAYER"
	case UNIFORM_OFFSET:
		name = "GL_UNIFORM_OFFSET"
	case DRAW_BUFFER13:
		name = "GL_DRAW_BUFFER13"
	case RGB5_A1:
		name = "GL_RGB5_A1"
	case POINT_SIZE_RANGE:
		name = "GL_POINT_SIZE_RANGE"
	case PROVOKING_VERTEX:
		name = "GL_PROVOKING_VERTEX"
	case SHADER_STORAGE_BARRIER_BIT:
		name = "GL_SHADER_STORAGE_BARRIER_BIT"
	case TEXTURE_CUBE_MAP_NEGATIVE_Z:
		name = "GL_TEXTURE_CUBE_MAP_NEGATIVE_Z"
	case FRAMEBUFFER_INCOMPLETE_READ_BUFFER:
		name = "GL_FRAMEBUFFER_INCOMPLETE_READ_BUFFER"
	case CONTEXT_PROFILE_MASK:
		name = "GL_CONTEXT_PROFILE_MASK"
	case SAMPLER_1D_ARRAY_SHADOW:
		name = "GL_SAMPLER_1D_ARRAY_SHADOW"
	case DEBUG_TYPE_MARKER:
		name = "GL_DEBUG_TYPE_MARKER"
	case TEXTURE_WRAP_T:
		name = "GL_TEXTURE_WRAP_T"
	case TEXTURE_BUFFER_OFFSET:
		name = "GL_TEXTURE_BUFFER_OFFSET"
	case COMPUTE_SUBROUTINE:
		name = "GL_COMPUTE_SUBROUTINE"
	case MAX_COMPUTE_IMAGE_UNIFORMS:
		name = "GL_MAX_COMPUTE_IMAGE_UNIFORMS"
	case STENCIL_ATTACHMENT:
		name = "GL_STENCIL_ATTACHMENT"
	case DEPTH_STENCIL_TEXTURE_MODE:
		name = "GL_DEPTH_STENCIL_TEXTURE_MODE"
	case UNDEFINED_VERTEX:
		name = "GL_UNDEFINED_VERTEX"
	case COLOR_ATTACHMENT13:
		name = "GL_COLOR_ATTACHMENT13"
	case R8I:
		name = "GL_R8I"
	case UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE:
		name = "GL_UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE"
	case TEXTURE_SHARED_SIZE:
		name = "GL_TEXTURE_SHARED_SIZE"
	case TEXTURE_ALPHA_SIZE:
		name = "GL_TEXTURE_ALPHA_SIZE"
	case IMPLEMENTATION_COLOR_READ_TYPE:
		name = "GL_IMPLEMENTATION_COLOR_READ_TYPE"
	case UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY:
		name = "GL_UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY"
	case MAX_FRAGMENT_ATOMIC_COUNTERS:
		name = "GL_MAX_FRAGMENT_ATOMIC_COUNTERS"
	case DEPTH_STENCIL_ATTACHMENT:
		name = "GL_DEPTH_STENCIL_ATTACHMENT"
	case MAX_COMBINED_ATOMIC_COUNTER_BUFFERS:
		name = "GL_MAX_COMBINED_ATOMIC_COUNTER_BUFFERS"
	case NOTEQUAL:
		name = "GL_NOTEQUAL"
	case SHORT:
		name = "GL_SHORT"
	case UNSIGNED_INT_IMAGE_2D_RECT:
		name = "GL_UNSIGNED_INT_IMAGE_2D_RECT"
	case ELEMENT_ARRAY_BARRIER_BIT:
		name = "GL_ELEMENT_ARRAY_BARRIER_BIT"
	case MAX_DUAL_SOURCE_DRAW_BUFFERS:
		name = "GL_MAX_DUAL_SOURCE_DRAW_BUFFERS"
	case PROGRAM:
		name = "GL_PROGRAM"
	case DEBUG_SOURCE_OTHER:
		name = "GL_DEBUG_SOURCE_OTHER"
	case VERTEX_ATTRIB_ARRAY_SIZE:
		name = "GL_VERTEX_ATTRIB_ARRAY_SIZE"
	case RGBA16F:
		name = "GL_RGBA16F"
	case PROXY_TEXTURE_3D:
		name = "GL_PROXY_TEXTURE_3D"
	case SYNC_FLAGS:
		name = "GL_SYNC_FLAGS"
	case ALREADY_SIGNALED:
		name = "GL_ALREADY_SIGNALED"
	case UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY:
		name = "GL_UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY"
	case REFERENCED_BY_TESS_CONTROL_SHADER:
		name = "GL_REFERENCED_BY_TESS_CONTROL_SHADER"
	case GET_TEXTURE_IMAGE_FORMAT:
		name = "GL_GET_TEXTURE_IMAGE_FORMAT"
	case SRC_ALPHA_SATURATE:
		name = "GL_SRC_ALPHA_SATURATE"
	case COLOR_ATTACHMENT10:
		name = "GL_COLOR_ATTACHMENT10"
	case FRACTIONAL_EVEN:
		name = "GL_FRACTIONAL_EVEN"
	case TRANSFORM_FEEDBACK:
		name = "GL_TRANSFORM_FEEDBACK"
	case TEXTURE_BINDING_BUFFER:
		name = "GL_TEXTURE_BINDING_BUFFER"
	case VIEW_CLASS_128_BITS:
		name = "GL_VIEW_CLASS_128_BITS"
	case IMAGE_2D_MULTISAMPLE:
		name = "GL_IMAGE_2D_MULTISAMPLE"
	case TEXTURE_COMPRESSED_BLOCK_WIDTH:
		name = "GL_TEXTURE_COMPRESSED_BLOCK_WIDTH"
	case MAX_COLOR_ATTACHMENTS:
		name = "GL_MAX_COLOR_ATTACHMENTS"
	case SAMPLER_CUBE:
		name = "GL_SAMPLER_CUBE"
	case VERTEX_ATTRIB_RELATIVE_OFFSET:
		name = "GL_VERTEX_ATTRIB_RELATIVE_OFFSET"
	case DOUBLE_VEC3:
		name = "GL_DOUBLE_VEC3"
	case INCR_WRAP:
		name = "GL_INCR_WRAP"
	case INT_SAMPLER_2D_ARRAY:
		name = "GL_INT_SAMPLER_2D_ARRAY"
	case MAX_SAMPLES:
		name = "GL_MAX_SAMPLES"
	case UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY:
		name = "GL_UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY"
	case TEXTURE_BINDING_2D_ARRAY:
		name = "GL_TEXTURE_BINDING_2D_ARRAY"
	case READ_PIXELS_TYPE:
		name = "GL_READ_PIXELS_TYPE"
	case UNSIGNED_INT_IMAGE_1D_ARRAY:
		name = "GL_UNSIGNED_INT_IMAGE_1D_ARRAY"
	case INTERNALFORMAT_DEPTH_SIZE:
		name = "GL_INTERNALFORMAT_DEPTH_SIZE"
	case PIXEL_BUFFER_BARRIER_BIT:
		name = "GL_PIXEL_BUFFER_BARRIER_BIT"
	case UNSIGNED_INT:
		name = "GL_UNSIGNED_INT"
	case DEBUG_TYPE_PERFORMANCE:
		name = "GL_DEBUG_TYPE_PERFORMANCE"
	case TEXTURE_CUBE_MAP_NEGATIVE_Y:
		name = "GL_TEXTURE_CUBE_MAP_NEGATIVE_Y"
	case INCR:
		name = "GL_INCR"
	case DITHER:
		name = "GL_DITHER"
	case PACK_SKIP_PIXELS:
		name = "GL_PACK_SKIP_PIXELS"
	case BLEND_EQUATION_ALPHA:
		name = "GL_BLEND_EQUATION_ALPHA"
	case QUADS:
		name = "GL_QUADS"
	case ACTIVE_ATTRIBUTES:
		name = "GL_ACTIVE_ATTRIBUTES"
	case PACK_COMPRESSED_BLOCK_SIZE:
		name = "GL_PACK_COMPRESSED_BLOCK_SIZE"
	case TEXTURE_2D_MULTISAMPLE_ARRAY:
		name = "GL_TEXTURE_2D_MULTISAMPLE_ARRAY"
	case IMAGE_2D:
		name = "GL_IMAGE_2D"
	case COMPRESSED_RED:
		name = "GL_COMPRESSED_RED"
	case PIXEL_PACK_BUFFER_BINDING:
		name = "GL_PIXEL_PACK_BUFFER_BINDING"
	case MAX_TESS_CONTROL_IMAGE_UNIFORMS:
		name = "GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS"
	case FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE"
	case FRAMEBUFFER_ATTACHMENT_LAYERED:
		name = "GL_FRAMEBUFFER_ATTACHMENT_LAYERED"
	case VERTEX_BINDING_STRIDE:
		name = "GL_VERTEX_BINDING_STRIDE"
	case UNSIGNED_INT_IMAGE_2D_ARRAY:
		name = "GL_UNSIGNED_INT_IMAGE_2D_ARRAY"
	case COPY:
		name = "GL_COPY"
	case SAMPLE_COVERAGE_INVERT:
		name = "GL_SAMPLE_COVERAGE_INVERT"
	case FRAMEBUFFER_RENDERABLE:
		name = "GL_FRAMEBUFFER_RENDERABLE"
	case MAX_ELEMENTS_VERTICES:
		name = "GL_MAX_ELEMENTS_VERTICES"
	case ARRAY_SIZE:
		name = "GL_ARRAY_SIZE"
	case FRAGMENT_INTERPOLATION_OFFSET_BITS:
		name = "GL_FRAGMENT_INTERPOLATION_OFFSET_BITS"
	case SAMPLER_2D_RECT:
		name = "GL_SAMPLER_2D_RECT"
	case RENDERBUFFER_HEIGHT:
		name = "GL_RENDERBUFFER_HEIGHT"
	case UNIFORM_ARRAY_STRIDE:
		name = "GL_UNIFORM_ARRAY_STRIDE"
	case MIN_PROGRAM_TEXEL_OFFSET:
		name = "GL_MIN_PROGRAM_TEXEL_OFFSET"
	case MAX_INTEGER_SAMPLES:
		name = "GL_MAX_INTEGER_SAMPLES"
	case VIEWPORT_INDEX_PROVOKING_VERTEX:
		name = "GL_VIEWPORT_INDEX_PROVOKING_VERTEX"
	case SHADING_LANGUAGE_VERSION:
		name = "GL_SHADING_LANGUAGE_VERSION"
	case VIEW_CLASS_BPTC_FLOAT:
		name = "GL_VIEW_CLASS_BPTC_FLOAT"
	case UNPACK_SKIP_ROWS:
		name = "GL_UNPACK_SKIP_ROWS"
	case INTERNALFORMAT_RED_SIZE:
		name = "GL_INTERNALFORMAT_RED_SIZE"
	case DOUBLE_VEC4:
		name = "GL_DOUBLE_VEC4"
	case STATIC_READ:
		name = "GL_STATIC_READ"
	case RG16UI:
		name = "GL_RG16UI"
	case IMAGE_2D_ARRAY:
		name = "GL_IMAGE_2D_ARRAY"
	case NEAREST_MIPMAP_LINEAR:
		name = "GL_NEAREST_MIPMAP_LINEAR"
	case BOOL_VEC3:
		name = "GL_BOOL_VEC3"
	case FRAMEBUFFER_COMPLETE:
		name = "GL_FRAMEBUFFER_COMPLETE"
	case MAX_VERTEX_STREAMS:
		name = "GL_MAX_VERTEX_STREAMS"
	case DOUBLE:
		name = "GL_DOUBLE"
	case QUERY_RESULT_AVAILABLE:
		name = "GL_QUERY_RESULT_AVAILABLE"
	case RG32F:
		name = "GL_RG32F"
	case DOUBLE_MAT2x4:
		name = "GL_DOUBLE_MAT2x4"
	case MAX_TESS_EVALUATION_INPUT_COMPONENTS:
		name = "GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS"
	case POLYGON_OFFSET_FILL:
		name = "GL_POLYGON_OFFSET_FILL"
	case MAX_COLOR_TEXTURE_SAMPLES:
		name = "GL_MAX_COLOR_TEXTURE_SAMPLES"
	case MIRRORED_REPEAT:
		name = "GL_MIRRORED_REPEAT"
	case FRAMEBUFFER_SRGB:
		name = "GL_FRAMEBUFFER_SRGB"
	case TEXTURE_MAX_LOD:
		name = "GL_TEXTURE_MAX_LOD"
	case NICEST:
		name = "GL_NICEST"
	case FRONT_AND_BACK:
		name = "GL_FRONT_AND_BACK"
	case GEOMETRY_INPUT_TYPE:
		name = "GL_GEOMETRY_INPUT_TYPE"
	case DOUBLE_MAT3x4:
		name = "GL_DOUBLE_MAT3x4"
	case DEBUG_SEVERITY_MEDIUM:
		name = "GL_DEBUG_SEVERITY_MEDIUM"
	case SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT:
		name = "GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT"
	case MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS:
		name = "GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS"
	case LINK_STATUS:
		name = "GL_LINK_STATUS"
	case XOR:
		name = "GL_XOR"
	case BYTE:
		name = "GL_BYTE"
	case FRAMEBUFFER_BINDING:
		name = "GL_FRAMEBUFFER_BINDING"
	case SRC_ALPHA:
		name = "GL_SRC_ALPHA"
	case TRANSFORM_FEEDBACK_BINDING:
		name = "GL_TRANSFORM_FEEDBACK_BINDING"
	case FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT:
		name = "GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT"
	case TEXTURE13:
		name = "GL_TEXTURE13"
	case TEXTURE10:
		name = "GL_TEXTURE10"
	case MAX_TEXTURE_SIZE:
		name = "GL_MAX_TEXTURE_SIZE"
	case COLOR_CLEAR_VALUE:
		name = "GL_COLOR_CLEAR_VALUE"
	case FLOAT_MAT4x2:
		name = "GL_FLOAT_MAT4x2"
	case ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER:
		name = "GL_ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER"
	case MAX_COMBINED_UNIFORM_BLOCKS:
		name = "GL_MAX_COMBINED_UNIFORM_BLOCKS"
	case SYNC_CONDITION:
		name = "GL_SYNC_CONDITION"
	case ONE_MINUS_SRC_ALPHA:
		name = "GL_ONE_MINUS_SRC_ALPHA"
	case INT_IMAGE_BUFFER:
		name = "GL_INT_IMAGE_BUFFER"
	case TEXTURE:
		name = "GL_TEXTURE"
	case UNPACK_COMPRESSED_BLOCK_HEIGHT:
		name = "GL_UNPACK_COMPRESSED_BLOCK_HEIGHT"
	case SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST:
		name = "GL_SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST"
	case REFERENCED_BY_COMPUTE_SHADER:
		name = "GL_REFERENCED_BY_COMPUTE_SHADER"
	case FLOAT_VEC4:
		name = "GL_FLOAT_VEC4"
	case READ_BUFFER:
		name = "GL_READ_BUFFER"
	case TEXTURE_2D_ARRAY:
		name = "GL_TEXTURE_2D_ARRAY"
	case CONTEXT_FLAGS:
		name = "GL_CONTEXT_FLAGS"
	case RGBA16:
		name = "GL_RGBA16"
	case IMAGE_CUBE_MAP_ARRAY:
		name = "GL_IMAGE_CUBE_MAP_ARRAY"
	case INT_SAMPLER_2D_MULTISAMPLE:
		name = "GL_INT_SAMPLER_2D_MULTISAMPLE"
	case UNPACK_COMPRESSED_BLOCK_SIZE:
		name = "GL_UNPACK_COMPRESSED_BLOCK_SIZE"
	case CULL_FACE:
		name = "GL_CULL_FACE"
	case UNSIGNED_INT_IMAGE_3D:
		name = "GL_UNSIGNED_INT_IMAGE_3D"
	case BACK:
		name = "GL_BACK"
	case TEXTURE_BINDING_2D:
		name = "GL_TEXTURE_BINDING_2D"
	case SIGNALED:
		name = "GL_SIGNALED"
	case FLOAT:
		name = "GL_FLOAT"
	case TRANSFORM_FEEDBACK_ACTIVE:
		name = "GL_TRANSFORM_FEEDBACK_ACTIVE"
	case TEXTURE_SAMPLES:
		name = "GL_TEXTURE_SAMPLES"
	case TRANSFORM_FEEDBACK_VARYINGS:
		name = "GL_TRANSFORM_FEEDBACK_VARYINGS"
	case R3_G3_B2:
		name = "GL_R3_G3_B2"
	case UNPACK_ROW_LENGTH:
		name = "GL_UNPACK_ROW_LENGTH"
	case RGB_INTEGER:
		name = "GL_RGB_INTEGER"
	case DEBUG_SOURCE_API:
		name = "GL_DEBUG_SOURCE_API"
	case MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS:
		name = "GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS"
	case TEXTURE20:
		name = "GL_TEXTURE20"
	case MAX_SAMPLE_MASK_WORDS:
		name = "GL_MAX_SAMPLE_MASK_WORDS"
	case TEXTURE_GREEN_SIZE:
		name = "GL_TEXTURE_GREEN_SIZE"
	case BLOCK_INDEX:
		name = "GL_BLOCK_INDEX"
	case RIGHT:
		name = "GL_RIGHT"
	case DEPTH_COMPONENT24:
		name = "GL_DEPTH_COMPONENT24"
	case MAX_VERTEX_UNIFORM_COMPONENTS:
		name = "GL_MAX_VERTEX_UNIFORM_COMPONENTS"
	case TEXTURE_INTERNAL_FORMAT:
		name = "GL_TEXTURE_INTERNAL_FORMAT"
	case SAMPLER_2D:
		name = "GL_SAMPLER_2D"
	case CURRENT_QUERY:
		name = "GL_CURRENT_QUERY"
	case BUFFER_ACCESS:
		name = "GL_BUFFER_ACCESS"
	case TEXTURE_COMPARE_FUNC:
		name = "GL_TEXTURE_COMPARE_FUNC"
	case STENCIL_BACK_WRITEMASK:
		name = "GL_STENCIL_BACK_WRITEMASK"
	case COMPUTE_TEXTURE:
		name = "GL_COMPUTE_TEXTURE"
	case TEXTURE_WIDTH:
		name = "GL_TEXTURE_WIDTH"
	case FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE"
	case COLOR_ATTACHMENT12:
		name = "GL_COLOR_ATTACHMENT12"
	case TESS_GEN_MODE:
		name = "GL_TESS_GEN_MODE"
	case DRAW_BUFFER2:
		name = "GL_DRAW_BUFFER2"
	case RG16:
		name = "GL_RG16"
	case COPY_WRITE_BUFFER_BINDING:
		name = "GL_COPY_WRITE_BUFFER_BINDING"
	case UNIFORM_BLOCK:
		name = "GL_UNIFORM_BLOCK"
	case MAX_DRAW_BUFFERS:
		name = "GL_MAX_DRAW_BUFFERS"
	case CLAMP_TO_BORDER:
		name = "GL_CLAMP_TO_BORDER"
	case CLIP_DISTANCE6:
		name = "GL_CLIP_DISTANCE6"
	case STENCIL_INDEX1:
		name = "GL_STENCIL_INDEX1"
	case RG8_SNORM:
		name = "GL_RG8_SNORM"
	case FRAMEBUFFER_ATTACHMENT_BLUE_SIZE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_BLUE_SIZE"
	case SMOOTH_POINT_SIZE_GRANULARITY:
		name = "GL_SMOOTH_POINT_SIZE_GRANULARITY"
	case CLIP_DISTANCE4:
		name = "GL_CLIP_DISTANCE4"
	case TEXTURE_3D:
		name = "GL_TEXTURE_3D"
	case SHADER_IMAGE_LOAD:
		name = "GL_SHADER_IMAGE_LOAD"
	case COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2:
		name = "GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2"
	case MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS:
		name = "GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS"
	case MAX_FRAGMENT_INTERPOLATION_OFFSET:
		name = "GL_MAX_FRAGMENT_INTERPOLATION_OFFSET"
	case INT_IMAGE_2D_MULTISAMPLE_ARRAY:
		name = "GL_INT_IMAGE_2D_MULTISAMPLE_ARRAY"
	case GREEN_INTEGER:
		name = "GL_GREEN_INTEGER"
	case COMPRESSED_RGB8_ETC2:
		name = "GL_COMPRESSED_RGB8_ETC2"
	case SHADER_BINARY_FORMATS:
		name = "GL_SHADER_BINARY_FORMATS"
	case MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS:
		name = "GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS"
	case R16UI:
		name = "GL_R16UI"
	case UNPACK_ALIGNMENT:
		name = "GL_UNPACK_ALIGNMENT"
	case ANY_SAMPLES_PASSED_CONSERVATIVE:
		name = "GL_ANY_SAMPLES_PASSED_CONSERVATIVE"
	case DEPTH_STENCIL:
		name = "GL_DEPTH_STENCIL"
	case MEDIUM_INT:
		name = "GL_MEDIUM_INT"
	case VIEW_CLASS_64_BITS:
		name = "GL_VIEW_CLASS_64_BITS"
	case DEPTH_FUNC:
		name = "GL_DEPTH_FUNC"
	case ALWAYS:
		name = "GL_ALWAYS"
	case MAX_SUBROUTINE_UNIFORM_LOCATIONS:
		name = "GL_MAX_SUBROUTINE_UNIFORM_LOCATIONS"
	case ARRAY_BUFFER_BINDING:
		name = "GL_ARRAY_BUFFER_BINDING"
	case DEPTH_RENDERABLE:
		name = "GL_DEPTH_RENDERABLE"
	case RG16I:
		name = "GL_RG16I"
	case FRAMEBUFFER_ATTACHMENT_RED_SIZE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_RED_SIZE"
	case INT_VEC4:
		name = "GL_INT_VEC4"
	case ONE_MINUS_SRC1_ALPHA:
		name = "GL_ONE_MINUS_SRC1_ALPHA"
	case QUERY:
		name = "GL_QUERY"
	case MAX_TESS_EVALUATION_UNIFORM_BLOCKS:
		name = "GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS"
	case PATCHES:
		name = "GL_PATCHES"
	case UNPACK_COMPRESSED_BLOCK_WIDTH:
		name = "GL_UNPACK_COMPRESSED_BLOCK_WIDTH"
	case DRAW_BUFFER3:
		name = "GL_DRAW_BUFFER3"
	case UNIFORM_BLOCK_ACTIVE_UNIFORMS:
		name = "GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS"
	case DEPTH_TEST:
		name = "GL_DEPTH_TEST"
	case DEPTH24_STENCIL8:
		name = "GL_DEPTH24_STENCIL8"
	case IMAGE_CLASS_2_X_16:
		name = "GL_IMAGE_CLASS_2_X_16"
	case RGBA16_SNORM:
		name = "GL_RGBA16_SNORM"
	case UNSIGNED_INT_IMAGE_1D:
		name = "GL_UNSIGNED_INT_IMAGE_1D"
	case BGR:
		name = "GL_BGR"
	case COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT:
		name = "GL_COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT"
	case MAX_COMPUTE_UNIFORM_COMPONENTS:
		name = "GL_MAX_COMPUTE_UNIFORM_COMPONENTS"
	case MAX_DEBUG_MESSAGE_LENGTH:
		name = "GL_MAX_DEBUG_MESSAGE_LENGTH"
	case UNSIGNED_INT_24_8:
		name = "GL_UNSIGNED_INT_24_8"
	case RENDERBUFFER_SAMPLES:
		name = "GL_RENDERBUFFER_SAMPLES"
	case CURRENT_VERTEX_ATTRIB:
		name = "GL_CURRENT_VERTEX_ATTRIB"
	case STENCIL_CLEAR_VALUE:
		name = "GL_STENCIL_CLEAR_VALUE"
	case QUERY_RESULT:
		name = "GL_QUERY_RESULT"
	case RENDERBUFFER_ALPHA_SIZE:
		name = "GL_RENDERBUFFER_ALPHA_SIZE"
	case TRANSFORM_FEEDBACK_BUFFER_SIZE:
		name = "GL_TRANSFORM_FEEDBACK_BUFFER_SIZE"
	case OBJECT_TYPE:
		name = "GL_OBJECT_TYPE"
	case VIEW_CLASS_S3TC_DXT1_RGB:
		name = "GL_VIEW_CLASS_S3TC_DXT1_RGB"
	case COLOR_ATTACHMENT0:
		name = "GL_COLOR_ATTACHMENT0"
	case TEXTURE12:
		name = "GL_TEXTURE12"
	case NUM_EXTENSIONS:
		name = "GL_NUM_EXTENSIONS"
	case STENCIL_BACK_FUNC:
		name = "GL_STENCIL_BACK_FUNC"
	case MAX_NUM_ACTIVE_VARIABLES:
		name = "GL_MAX_NUM_ACTIVE_VARIABLES"
	case TEXTURE1:
		name = "GL_TEXTURE1"
	case TEXTURE_MIN_FILTER:
		name = "GL_TEXTURE_MIN_FILTER"
	case MAX_NAME_LENGTH:
		name = "GL_MAX_NAME_LENGTH"
	case RENDERBUFFER_GREEN_SIZE:
		name = "GL_RENDERBUFFER_GREEN_SIZE"
	case MAX_GEOMETRY_IMAGE_UNIFORMS:
		name = "GL_MAX_GEOMETRY_IMAGE_UNIFORMS"
	case MIN_PROGRAM_TEXTURE_GATHER_OFFSET:
		name = "GL_MIN_PROGRAM_TEXTURE_GATHER_OFFSET"
	case PROGRAM_PIPELINE:
		name = "GL_PROGRAM_PIPELINE"
	case PROXY_TEXTURE_1D:
		name = "GL_PROXY_TEXTURE_1D"
	case UNSIGNED_SHORT_5_6_5:
		name = "GL_UNSIGNED_SHORT_5_6_5"
	case PACK_COMPRESSED_BLOCK_WIDTH:
		name = "GL_PACK_COMPRESSED_BLOCK_WIDTH"
	case MAX_COMPUTE_ATOMIC_COUNTERS:
		name = "GL_MAX_COMPUTE_ATOMIC_COUNTERS"
	case MAX_RENDERBUFFER_SIZE:
		name = "GL_MAX_RENDERBUFFER_SIZE"
	case UNSIGNED_INT_IMAGE_2D:
		name = "GL_UNSIGNED_INT_IMAGE_2D"
	case COMPRESSED_RGB_S3TC_DXT1_EXT:
		name = "GL_COMPRESSED_RGB_S3TC_DXT1_EXT"
	case IMAGE_BINDING_ACCESS:
		name = "GL_IMAGE_BINDING_ACCESS"
	case ACTIVE_SUBROUTINES:
		name = "GL_ACTIVE_SUBROUTINES"
	case ATOMIC_COUNTER_BUFFER_START:
		name = "GL_ATOMIC_COUNTER_BUFFER_START"
	case SLUMINANCE8_EXT:
		name = "GL_SLUMINANCE8_EXT"
	case IMAGE_BINDING_LEVEL:
		name = "GL_IMAGE_BINDING_LEVEL"
	case COLOR_ATTACHMENT6:
		name = "GL_COLOR_ATTACHMENT6"
	case TEXTURE_CUBE_MAP_SEAMLESS:
		name = "GL_TEXTURE_CUBE_MAP_SEAMLESS"
	case IMPLEMENTATION_COLOR_READ_FORMAT:
		name = "GL_IMPLEMENTATION_COLOR_READ_FORMAT"
	case FLOAT_MAT3x4:
		name = "GL_FLOAT_MAT3x4"
	case DEBUG_SOURCE_THIRD_PARTY:
		name = "GL_DEBUG_SOURCE_THIRD_PARTY"
	case MAX_FRAGMENT_UNIFORM_VECTORS:
		name = "GL_MAX_FRAGMENT_UNIFORM_VECTORS"
	case IMAGE_FORMAT_COMPATIBILITY_TYPE:
		name = "GL_IMAGE_FORMAT_COMPATIBILITY_TYPE"
	case ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER:
		name = "GL_ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER"
	case LINE_SMOOTH:
		name = "GL_LINE_SMOOTH"
	case DYNAMIC_DRAW:
		name = "GL_DYNAMIC_DRAW"
	case MAX_GEOMETRY_ATOMIC_COUNTERS:
		name = "GL_MAX_GEOMETRY_ATOMIC_COUNTERS"
	case VIEW_CLASS_RGTC1_RED:
		name = "GL_VIEW_CLASS_RGTC1_RED"
	case DOUBLE_VEC2:
		name = "GL_DOUBLE_VEC2"
	case FLOAT_MAT2x4:
		name = "GL_FLOAT_MAT2x4"
	case TOP_LEVEL_ARRAY_SIZE:
		name = "GL_TOP_LEVEL_ARRAY_SIZE"
	case INTERNALFORMAT_PREFERRED:
		name = "GL_INTERNALFORMAT_PREFERRED"
	case PATCH_DEFAULT_INNER_LEVEL:
		name = "GL_PATCH_DEFAULT_INNER_LEVEL"
	case VIEWPORT_SUBPIXEL_BITS:
		name = "GL_VIEWPORT_SUBPIXEL_BITS"
	case RED:
		name = "GL_RED"
	case MAX_ELEMENT_INDEX:
		name = "GL_MAX_ELEMENT_INDEX"
	case BACK_RIGHT:
		name = "GL_BACK_RIGHT"
	case TEXTURE_2D:
		name = "GL_TEXTURE_2D"
	case ONE_MINUS_SRC1_COLOR:
		name = "GL_ONE_MINUS_SRC1_COLOR"
	case READ_WRITE:
		name = "GL_READ_WRITE"
	case MAX_VERTEX_ATTRIB_RELATIVE_OFFSET:
		name = "GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET"
	case RGB9_E5:
		name = "GL_RGB9_E5"
	case FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING:
		name = "GL_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING"
	case TEXTURE_1D:
		name = "GL_TEXTURE_1D"
	case VERTEX_ATTRIB_ARRAY_TYPE:
		name = "GL_VERTEX_ATTRIB_ARRAY_TYPE"
	case DRAW_BUFFER12:
		name = "GL_DRAW_BUFFER12"
	case INTERNALFORMAT_STENCIL_SIZE:
		name = "GL_INTERNALFORMAT_STENCIL_SIZE"
	case UNPACK_SKIP_IMAGES:
		name = "GL_UNPACK_SKIP_IMAGES"
	case RENDERBUFFER_BINDING:
		name = "GL_RENDERBUFFER_BINDING"
	case LAST_VERTEX_CONVENTION:
		name = "GL_LAST_VERTEX_CONVENTION"
	case COMPRESSED_RG:
		name = "GL_COMPRESSED_RG"
	case RGBA8_SNORM:
		name = "GL_RGBA8_SNORM"
	case MAX_LAYERS:
		name = "GL_MAX_LAYERS"
	case VERTEX_SUBROUTINE_UNIFORM:
		name = "GL_VERTEX_SUBROUTINE_UNIFORM"
	case INT_SAMPLER_CUBE_MAP_ARRAY:
		name = "GL_INT_SAMPLER_CUBE_MAP_ARRAY"
	case BUFFER_VARIABLE:
		name = "GL_BUFFER_VARIABLE"
	case COMPRESSED_SRGB:
		name = "GL_COMPRESSED_SRGB"
	case SUBPIXEL_BITS:
		name = "GL_SUBPIXEL_BITS"
	case MAX_COMPUTE_WORK_GROUP_COUNT:
		name = "GL_MAX_COMPUTE_WORK_GROUP_COUNT"
	case DEBUG_LOGGED_MESSAGES:
		name = "GL_DEBUG_LOGGED_MESSAGES"
	case COMPRESSED_SRGB_S3TC_DXT1_EXT:
		name = "GL_COMPRESSED_SRGB_S3TC_DXT1_EXT"
	case INTERNALFORMAT_GREEN_TYPE:
		name = "GL_INTERNALFORMAT_GREEN_TYPE"
	case DRAW_BUFFER5:
		name = "GL_DRAW_BUFFER5"
	case MAX_TEXTURE_BUFFER_SIZE:
		name = "GL_MAX_TEXTURE_BUFFER_SIZE"
	case MAP_INVALIDATE_BUFFER_BIT:
		name = "GL_MAP_INVALIDATE_BUFFER_BIT"
	case IMAGE_FORMAT_COMPATIBILITY_BY_SIZE:
		name = "GL_IMAGE_FORMAT_COMPATIBILITY_BY_SIZE"
	case TEXTURE_WRAP_R:
		name = "GL_TEXTURE_WRAP_R"
	case MAX_ATOMIC_COUNTER_BUFFER_BINDINGS:
		name = "GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS"
	case DEPTH_WRITEMASK:
		name = "GL_DEPTH_WRITEMASK"
	case FILTER:
		name = "GL_FILTER"
	case REFERENCED_BY_VERTEX_SHADER:
		name = "GL_REFERENCED_BY_VERTEX_SHADER"
	case VIEWPORT_BOUNDS_RANGE:
		name = "GL_VIEWPORT_BOUNDS_RANGE"
	case RGBA16UI:
		name = "GL_RGBA16UI"
	case MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		name = "GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT"
	case TEXTURE_CUBE_MAP_POSITIVE_Y:
		name = "GL_TEXTURE_CUBE_MAP_POSITIVE_Y"
	case MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS:
		name = "GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS"
	case UNPACK_SWAP_BYTES:
		name = "GL_UNPACK_SWAP_BYTES"
	case TEXTURE_BUFFER_SIZE:
		name = "GL_TEXTURE_BUFFER_SIZE"
	case SRGB8:
		name = "GL_SRGB8"
	case IMAGE_CLASS_2_X_8:
		name = "GL_IMAGE_CLASS_2_X_8"
	case INTERNALFORMAT_ALPHA_SIZE:
		name = "GL_INTERNALFORMAT_ALPHA_SIZE"
	case UNSIGNED_INT_VEC2:
		name = "GL_UNSIGNED_INT_VEC2"
	case READ_ONLY:
		name = "GL_READ_ONLY"
	case PROGRAM_OUTPUT:
		name = "GL_PROGRAM_OUTPUT"
	case COMPRESSED_SIGNED_RED_RGTC1:
		name = "GL_COMPRESSED_SIGNED_RED_RGTC1"
	case MAX_DEBUG_LOGGED_MESSAGES:
		name = "GL_MAX_DEBUG_LOGGED_MESSAGES"
	case MAX_ARRAY_TEXTURE_LAYERS:
		name = "GL_MAX_ARRAY_TEXTURE_LAYERS"
	case COLOR_ATTACHMENT11:
		name = "GL_COLOR_ATTACHMENT11"
	case REFERENCED_BY_TESS_EVALUATION_SHADER:
		name = "GL_REFERENCED_BY_TESS_EVALUATION_SHADER"
	case TRANSFORM_FEEDBACK_BARRIER_BIT:
		name = "GL_TRANSFORM_FEEDBACK_BARRIER_BIT"
	case INT:
		name = "GL_INT"
	case FRAGMENT_SUBROUTINE:
		name = "GL_FRAGMENT_SUBROUTINE"
	case COMPILE_STATUS:
		name = "GL_COMPILE_STATUS"
	case TESS_CONTROL_SUBROUTINE_UNIFORM:
		name = "GL_TESS_CONTROL_SUBROUTINE_UNIFORM"
	case DEBUG_OUTPUT:
		name = "GL_DEBUG_OUTPUT"
	case CURRENT_PROGRAM:
		name = "GL_CURRENT_PROGRAM"
	case STENCIL_BACK_PASS_DEPTH_FAIL:
		name = "GL_STENCIL_BACK_PASS_DEPTH_FAIL"
	case INT_SAMPLER_1D_ARRAY:
		name = "GL_INT_SAMPLER_1D_ARRAY"
	case CLIP_DISTANCE5:
		name = "GL_CLIP_DISTANCE5"
	case RENDERBUFFER_RED_SIZE:
		name = "GL_RENDERBUFFER_RED_SIZE"
	case NEAREST_MIPMAP_NEAREST:
		name = "GL_NEAREST_MIPMAP_NEAREST"
	case TEXTURE29:
		name = "GL_TEXTURE29"
	case VIEW_CLASS_96_BITS:
		name = "GL_VIEW_CLASS_96_BITS"
	case R16F:
		name = "GL_R16F"
	case TEXTURE31:
		name = "GL_TEXTURE31"
	case BGRA_INTEGER:
		name = "GL_BGRA_INTEGER"
	case SAMPLE_MASK:
		name = "GL_SAMPLE_MASK"
	case UNIFORM_BLOCK_BINDING:
		name = "GL_UNIFORM_BLOCK_BINDING"
	case SAMPLER_BUFFER:
		name = "GL_SAMPLER_BUFFER"
	case FLOAT_MAT4:
		name = "GL_FLOAT_MAT4"
	case MAX_FRAGMENT_SHADER_STORAGE_BLOCKS:
		name = "GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS"
	case RGB16UI:
		name = "GL_RGB16UI"
	case TEXTURE_BINDING_CUBE_MAP_ARRAY:
		name = "GL_TEXTURE_BINDING_CUBE_MAP_ARRAY"
	case TEXTURE_BLUE_TYPE:
		name = "GL_TEXTURE_BLUE_TYPE"
	case TESS_EVALUATION_SUBROUTINE:
		name = "GL_TESS_EVALUATION_SUBROUTINE"
	case MULTISAMPLE:
		name = "GL_MULTISAMPLE"
	case BUFFER_MAP_POINTER:
		name = "GL_BUFFER_MAP_POINTER"
	case RGB10_A2UI:
		name = "GL_RGB10_A2UI"
	case TEXTURE_IMAGE_FORMAT:
		name = "GL_TEXTURE_IMAGE_FORMAT"
	case VERTEX_PROGRAM_POINT_SIZE:
		name = "GL_VERTEX_PROGRAM_POINT_SIZE"
	case FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS:
		name = "GL_FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS"
	case TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH:
		name = "GL_TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH"
	case LEFT:
		name = "GL_LEFT"
	case FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE"
	case IMAGE_CLASS_1_X_8:
		name = "GL_IMAGE_CLASS_1_X_8"
	case AND_INVERTED:
		name = "GL_AND_INVERTED"
	case ALIASED_LINE_WIDTH_RANGE:
		name = "GL_ALIASED_LINE_WIDTH_RANGE"
	case UNIFORM_BLOCK_NAME_LENGTH:
		name = "GL_UNIFORM_BLOCK_NAME_LENGTH"
	case LAYER_PROVOKING_VERTEX:
		name = "GL_LAYER_PROVOKING_VERTEX"
	case TEXTURE_BASE_LEVEL:
		name = "GL_TEXTURE_BASE_LEVEL"
	case TEXTURE_BINDING_1D_ARRAY:
		name = "GL_TEXTURE_BINDING_1D_ARRAY"
	case SHADER_STORAGE_BLOCK:
		name = "GL_SHADER_STORAGE_BLOCK"
	case TEXTURE15:
		name = "GL_TEXTURE15"
	case UNSIGNED_SHORT_5_6_5_REV:
		name = "GL_UNSIGNED_SHORT_5_6_5_REV"
	case NAND:
		name = "GL_NAND"
	case COLOR_ATTACHMENT7:
		name = "GL_COLOR_ATTACHMENT7"
	case PIXEL_UNPACK_BUFFER:
		name = "GL_PIXEL_UNPACK_BUFFER"
	case INTERNALFORMAT_ALPHA_TYPE:
		name = "GL_INTERNALFORMAT_ALPHA_TYPE"
	case RENDERBUFFER_STENCIL_SIZE:
		name = "GL_RENDERBUFFER_STENCIL_SIZE"
	case COMPRESSED_SLUMINANCE_EXT:
		name = "GL_COMPRESSED_SLUMINANCE_EXT"
	case PACK_LSB_FIRST:
		name = "GL_PACK_LSB_FIRST"
	case TEXTURE9:
		name = "GL_TEXTURE9"
	case DEBUG_CALLBACK_FUNCTION:
		name = "GL_DEBUG_CALLBACK_FUNCTION"
	case CW:
		name = "GL_CW"
	case MAX_FRAMEBUFFER_SAMPLES:
		name = "GL_MAX_FRAMEBUFFER_SAMPLES"
	case PATCH_DEFAULT_OUTER_LEVEL:
		name = "GL_PATCH_DEFAULT_OUTER_LEVEL"
	case MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS:
		name = "GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS"
	case SHADER_STORAGE_BUFFER:
		name = "GL_SHADER_STORAGE_BUFFER"
	case VIEW_CLASS_BPTC_UNORM:
		name = "GL_VIEW_CLASS_BPTC_UNORM"
	case COMPRESSED_SIGNED_RG_RGTC2:
		name = "GL_COMPRESSED_SIGNED_RG_RGTC2"
	case MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS:
		name = "GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS"
	case TEXTURE_DEPTH:
		name = "GL_TEXTURE_DEPTH"
	case COLOR_ATTACHMENT15:
		name = "GL_COLOR_ATTACHMENT15"
	case LOCATION_INDEX:
		name = "GL_LOCATION_INDEX"
	case MAX_VERTEX_UNIFORM_VECTORS:
		name = "GL_MAX_VERTEX_UNIFORM_VECTORS"
	case VIEW_CLASS_32_BITS:
		name = "GL_VIEW_CLASS_32_BITS"
	case ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER:
		name = "GL_ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER"
	case MAX_TESS_CONTROL_ATOMIC_COUNTERS:
		name = "GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS"
	case STENCIL_COMPONENTS:
		name = "GL_STENCIL_COMPONENTS"
	case FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE"
	case DELETE_STATUS:
		name = "GL_DELETE_STATUS"
	case CLIP_DISTANCE2:
		name = "GL_CLIP_DISTANCE2"
	case FRACTIONAL_ODD:
		name = "GL_FRACTIONAL_ODD"
	case MAX_FRAMEBUFFER_LAYERS:
		name = "GL_MAX_FRAMEBUFFER_LAYERS"
	case MAX_TESS_CONTROL_UNIFORM_BLOCKS:
		name = "GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS"
	case BLEND_DST_RGB:
		name = "GL_BLEND_DST_RGB"
	case CLIP_DISTANCE1:
		name = "GL_CLIP_DISTANCE1"
	case TEXTURE_IMMUTABLE_LEVELS:
		name = "GL_TEXTURE_IMMUTABLE_LEVELS"
	case DRAW_FRAMEBUFFER:
		name = "GL_DRAW_FRAMEBUFFER"
	case FRAMEBUFFER_DEFAULT_SAMPLES:
		name = "GL_FRAMEBUFFER_DEFAULT_SAMPLES"
	case QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION:
		name = "GL_QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION"
	case DEPTH_CLAMP:
		name = "GL_DEPTH_CLAMP"
	case DEBUG_GROUP_STACK_DEPTH:
		name = "GL_DEBUG_GROUP_STACK_DEPTH"
	case STENCIL_TEST:
		name = "GL_STENCIL_TEST"
	case R16I:
		name = "GL_R16I"
	case TEXTURE_VIEW_NUM_LAYERS:
		name = "GL_TEXTURE_VIEW_NUM_LAYERS"
	case RGBA8UI:
		name = "GL_RGBA8UI"
	case MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS:
		name = "GL_MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS"
	case DOUBLE_MAT2x3:
		name = "GL_DOUBLE_MAT2x3"
	case UNSIGNED_INT_5_9_9_9_REV:
		name = "GL_UNSIGNED_INT_5_9_9_9_REV"
	case MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS:
		name = "GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS"
	case MAX_COMPUTE_SHADER_STORAGE_BLOCKS:
		name = "GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS"
	case RGB:
		name = "GL_RGB"
	case MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS:
		name = "GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS"
	case MAX_IMAGE_UNITS:
		name = "GL_MAX_IMAGE_UNITS"
	case DOUBLE_MAT3x2:
		name = "GL_DOUBLE_MAT3x2"
	case COLOR_COMPONENTS:
		name = "GL_COLOR_COMPONENTS"
	case UNIFORM_TYPE:
		name = "GL_UNIFORM_TYPE"
	case MAX_FRAGMENT_IMAGE_UNIFORMS:
		name = "GL_MAX_FRAGMENT_IMAGE_UNIFORMS"
	case COMPARE_REF_TO_TEXTURE:
		name = "GL_COMPARE_REF_TO_TEXTURE"
	case TEXTURE25:
		name = "GL_TEXTURE25"
	case VIEW_COMPATIBILITY_CLASS:
		name = "GL_VIEW_COMPATIBILITY_CLASS"
	case UNSIGNED_INT_SAMPLER_1D:
		name = "GL_UNSIGNED_INT_SAMPLER_1D"
	case SHADER_IMAGE_ATOMIC:
		name = "GL_SHADER_IMAGE_ATOMIC"
	case DEBUG_TYPE_DEPRECATED_BEHAVIOR:
		name = "GL_DEBUG_TYPE_DEPRECATED_BEHAVIOR"
	case COLOR_ENCODING:
		name = "GL_COLOR_ENCODING"
	case MAX_TESS_PATCH_COMPONENTS:
		name = "GL_MAX_TESS_PATCH_COMPONENTS"
	case NUM_COMPRESSED_TEXTURE_FORMATS:
		name = "GL_NUM_COMPRESSED_TEXTURE_FORMATS"
	case MAX_LABEL_LENGTH:
		name = "GL_MAX_LABEL_LENGTH"
	case SAMPLER_1D_SHADOW:
		name = "GL_SAMPLER_1D_SHADOW"
	case FRAMEBUFFER_ATTACHMENT_OBJECT_NAME:
		name = "GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME"
	case MEDIUM_FLOAT:
		name = "GL_MEDIUM_FLOAT"
	case PACK_SWAP_BYTES:
		name = "GL_PACK_SWAP_BYTES"
	case MIPMAP:
		name = "GL_MIPMAP"
	case FLOAT_VEC3:
		name = "GL_FLOAT_VEC3"
	case POINT:
		name = "GL_POINT"
	case IMAGE_PIXEL_TYPE:
		name = "GL_IMAGE_PIXEL_TYPE"
	case RGB4:
		name = "GL_RGB4"
	case UNSIGNED_BYTE:
		name = "GL_UNSIGNED_BYTE"
	case DEPTH_COMPONENT16:
		name = "GL_DEPTH_COMPONENT16"
	case LINE_SMOOTH_HINT:
		name = "GL_LINE_SMOOTH_HINT"
	case ACTIVE_UNIFORM_MAX_LENGTH:
		name = "GL_ACTIVE_UNIFORM_MAX_LENGTH"
	case MAX_VERTEX_TEXTURE_IMAGE_UNITS:
		name = "GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS"
	case MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS:
		name = "GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS"
	case TEXTURE5:
		name = "GL_TEXTURE5"
	case UNSIGNED_INT_IMAGE_BUFFER:
		name = "GL_UNSIGNED_INT_IMAGE_BUFFER"
	case NOOP:
		name = "GL_NOOP"
	case IMAGE_CUBE:
		name = "GL_IMAGE_CUBE"
	case PACK_ALIGNMENT:
		name = "GL_PACK_ALIGNMENT"
	case PIXEL_UNPACK_BUFFER_BINDING:
		name = "GL_PIXEL_UNPACK_BUFFER_BINDING"
	case SAMPLER_CUBE_MAP_ARRAY:
		name = "GL_SAMPLER_CUBE_MAP_ARRAY"
	case TEXTURE21:
		name = "GL_TEXTURE21"
	case PROGRAM_BINARY_FORMATS:
		name = "GL_PROGRAM_BINARY_FORMATS"
	case IMAGE_2D_MULTISAMPLE_ARRAY:
		name = "GL_IMAGE_2D_MULTISAMPLE_ARRAY"
	case REFERENCED_BY_FRAGMENT_SHADER:
		name = "GL_REFERENCED_BY_FRAGMENT_SHADER"
	case RGBA32F:
		name = "GL_RGBA32F"
	case RGB8UI:
		name = "GL_RGB8UI"
	case PROGRAM_INPUT:
		name = "GL_PROGRAM_INPUT"
	case DRAW_BUFFER8:
		name = "GL_DRAW_BUFFER8"
	case MAX_GEOMETRY_SHADER_STORAGE_BLOCKS:
		name = "GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS"
	case TEXTURE_VIEW:
		name = "GL_TEXTURE_VIEW"
	case FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL:
		name = "GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL"
	case MIN_SAMPLE_SHADING_VALUE:
		name = "GL_MIN_SAMPLE_SHADING_VALUE"
	case GEOMETRY_TEXTURE:
		name = "GL_GEOMETRY_TEXTURE"
	case COLOR_ATTACHMENT3:
		name = "GL_COLOR_ATTACHMENT3"
	case RGB16I:
		name = "GL_RGB16I"
	case INT_IMAGE_CUBE:
		name = "GL_INT_IMAGE_CUBE"
	case UNPACK_IMAGE_HEIGHT:
		name = "GL_UNPACK_IMAGE_HEIGHT"
	case VERTEX_ATTRIB_ARRAY_DIVISOR:
		name = "GL_VERTEX_ATTRIB_ARRAY_DIVISOR"
	case UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES:
		name = "GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES"
	case MANUAL_GENERATE_MIPMAP:
		name = "GL_MANUAL_GENERATE_MIPMAP"
	case BLEND:
		name = "GL_BLEND"
	case TEXTURE_BUFFER_OFFSET_ALIGNMENT:
		name = "GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT"
	case DEPTH_COMPONENT:
		name = "GL_DEPTH_COMPONENT"
	case VENDOR:
		name = "GL_VENDOR"
	case REPLACE:
		name = "GL_REPLACE"
	case ELEMENT_ARRAY_BUFFER_BINDING:
		name = "GL_ELEMENT_ARRAY_BUFFER_BINDING"
	case FILL:
		name = "GL_FILL"
	case QUERY_NO_WAIT:
		name = "GL_QUERY_NO_WAIT"
	case SAMPLE_POSITION:
		name = "GL_SAMPLE_POSITION"
	case RGB8_SNORM:
		name = "GL_RGB8_SNORM"
	case FRAMEBUFFER_RENDERABLE_LAYERED:
		name = "GL_FRAMEBUFFER_RENDERABLE_LAYERED"
	case INT_IMAGE_2D_RECT:
		name = "GL_INT_IMAGE_2D_RECT"
	case VERTEX_ATTRIB_ARRAY_INTEGER:
		name = "GL_VERTEX_ATTRIB_ARRAY_INTEGER"
	case REPEAT:
		name = "GL_REPEAT"
	case COLOR_ATTACHMENT4:
		name = "GL_COLOR_ATTACHMENT4"
	case UNSIGNED_SHORT_4_4_4_4:
		name = "GL_UNSIGNED_SHORT_4_4_4_4"
	case DOUBLE_MAT4x3:
		name = "GL_DOUBLE_MAT4x3"
	case STATIC_COPY:
		name = "GL_STATIC_COPY"
	case ARRAY_STRIDE:
		name = "GL_ARRAY_STRIDE"
	case TRIANGLE_FAN:
		name = "GL_TRIANGLE_FAN"
	case BLEND_DST_ALPHA:
		name = "GL_BLEND_DST_ALPHA"
	case VIEW_CLASS_S3TC_DXT1_RGBA:
		name = "GL_VIEW_CLASS_S3TC_DXT1_RGBA"
	case LOWER_LEFT:
		name = "GL_LOWER_LEFT"
	case FRAMEBUFFER_DEFAULT_WIDTH:
		name = "GL_FRAMEBUFFER_DEFAULT_WIDTH"
	case IMAGE_CLASS_4_X_8:
		name = "GL_IMAGE_CLASS_4_X_8"
	case TEXTURE_2D_MULTISAMPLE:
		name = "GL_TEXTURE_2D_MULTISAMPLE"
	case UNSIGNED_INT_SAMPLER_3D:
		name = "GL_UNSIGNED_INT_SAMPLER_3D"
	case TEXTURE_STENCIL_SIZE:
		name = "GL_TEXTURE_STENCIL_SIZE"
	case PROXY_TEXTURE_CUBE_MAP_ARRAY:
		name = "GL_PROXY_TEXTURE_CUBE_MAP_ARRAY"
	case FLOAT_MAT3:
		name = "GL_FLOAT_MAT3"
	case ACTIVE_SUBROUTINE_MAX_LENGTH:
		name = "GL_ACTIVE_SUBROUTINE_MAX_LENGTH"
	case BUFFER_BINDING:
		name = "GL_BUFFER_BINDING"
	case TEXTURE_HEIGHT:
		name = "GL_TEXTURE_HEIGHT"
	case MAX_VIEWPORTS:
		name = "GL_MAX_VIEWPORTS"
	case TEXTURE_SWIZZLE_A:
		name = "GL_TEXTURE_SWIZZLE_A"
	case COMPRESSED_SLUMINANCE_ALPHA_EXT:
		name = "GL_COMPRESSED_SLUMINANCE_ALPHA_EXT"
	case SAMPLER_3D:
		name = "GL_SAMPLER_3D"
	case STENCIL_VALUE_MASK:
		name = "GL_STENCIL_VALUE_MASK"
	case OFFSET:
		name = "GL_OFFSET"
	case WAIT_FAILED:
		name = "GL_WAIT_FAILED"
	case LOW_INT:
		name = "GL_LOW_INT"
	case DEPTH_COMPONENTS:
		name = "GL_DEPTH_COMPONENTS"
	case CULL_FACE_MODE:
		name = "GL_CULL_FACE_MODE"
	case DRAW_BUFFER0:
		name = "GL_DRAW_BUFFER0"
	case PATCH_VERTICES:
		name = "GL_PATCH_VERTICES"
	case PROGRAM_SEPARABLE:
		name = "GL_PROGRAM_SEPARABLE"
	case DEPTH32F_STENCIL8:
		name = "GL_DEPTH32F_STENCIL8"
	case PROXY_TEXTURE_2D_MULTISAMPLE:
		name = "GL_PROXY_TEXTURE_2D_MULTISAMPLE"
	case VERTEX_ATTRIB_ARRAY_BUFFER_BINDING:
		name = "GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING"
	case INTERNALFORMAT_BLUE_SIZE:
		name = "GL_INTERNALFORMAT_BLUE_SIZE"
	case DEBUG_SOURCE_SHADER_COMPILER:
		name = "GL_DEBUG_SOURCE_SHADER_COMPILER"
	case DRAW_BUFFER11:
		name = "GL_DRAW_BUFFER11"
	case MAX_PROGRAM_TEXTURE_GATHER_OFFSET:
		name = "GL_MAX_PROGRAM_TEXTURE_GATHER_OFFSET"
	case STREAM_DRAW:
		name = "GL_STREAM_DRAW"
	case PRIMITIVES_GENERATED:
		name = "GL_PRIMITIVES_GENERATED"
	case UNSIGNED_SHORT_1_5_5_5_REV:
		name = "GL_UNSIGNED_SHORT_1_5_5_5_REV"
	case R8UI:
		name = "GL_R8UI"
	case DEBUG_SEVERITY_HIGH:
		name = "GL_DEBUG_SEVERITY_HIGH"
	case LINE_STRIP:
		name = "GL_LINE_STRIP"
	case RED_SNORM:
		name = "GL_RED_SNORM"
	case TEXTURE18:
		name = "GL_TEXTURE18"
	case FRAMEBUFFER_INCOMPLETE_MULTISAMPLE:
		name = "GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE"
	case UNPACK_SKIP_PIXELS:
		name = "GL_UNPACK_SKIP_PIXELS"
	case FLOAT_MAT2x3:
		name = "GL_FLOAT_MAT2x3"
	case R32UI:
		name = "GL_R32UI"
	case UNSIGNED_INT_SAMPLER_2D_RECT:
		name = "GL_UNSIGNED_INT_SAMPLER_2D_RECT"
	case SYNC_FENCE:
		name = "GL_SYNC_FENCE"
	case INT_SAMPLER_3D:
		name = "GL_INT_SAMPLER_3D"
	case MAX_COMPUTE_TEXTURE_IMAGE_UNITS:
		name = "GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS"
	case UNSIGNED_BYTE_3_3_2:
		name = "GL_UNSIGNED_BYTE_3_3_2"
	case TEXTURE_CUBE_MAP_POSITIVE_Z:
		name = "GL_TEXTURE_CUBE_MAP_POSITIVE_Z"
	case MATRIX_STRIDE:
		name = "GL_MATRIX_STRIDE"
	case SLUMINANCE_EXT:
		name = "GL_SLUMINANCE_EXT"
	case CLIP_DISTANCE3:
		name = "GL_CLIP_DISTANCE3"
	case BLEND_EQUATION_RGB:
		name = "GL_BLEND_EQUATION_RGB"
	case MAX_COMPUTE_WORK_GROUP_SIZE:
		name = "GL_MAX_COMPUTE_WORK_GROUP_SIZE"
	case UNSIGNED_INT_8_8_8_8_REV:
		name = "GL_UNSIGNED_INT_8_8_8_8_REV"
	case IMAGE_BINDING_FORMAT:
		name = "GL_IMAGE_BINDING_FORMAT"
	case ACTIVE_UNIFORMS:
		name = "GL_ACTIVE_UNIFORMS"
	case DEBUG_SEVERITY_NOTIFICATION:
		name = "GL_DEBUG_SEVERITY_NOTIFICATION"
	case OR_INVERTED:
		name = "GL_OR_INVERTED"
	case FRAGMENT_TEXTURE:
		name = "GL_FRAGMENT_TEXTURE"
	case DEBUG_SOURCE_WINDOW_SYSTEM:
		name = "GL_DEBUG_SOURCE_WINDOW_SYSTEM"
	case GEOMETRY_VERTICES_OUT:
		name = "GL_GEOMETRY_VERTICES_OUT"
	case BOOL_VEC4:
		name = "GL_BOOL_VEC4"
	case SHADER_STORAGE_BUFFER_SIZE:
		name = "GL_SHADER_STORAGE_BUFFER_SIZE"
	case SAMPLER_1D_ARRAY:
		name = "GL_SAMPLER_1D_ARRAY"
	case TRANSFORM_FEEDBACK_PAUSED:
		name = "GL_TRANSFORM_FEEDBACK_PAUSED"
	case IMAGE_BINDING_LAYERED:
		name = "GL_IMAGE_BINDING_LAYERED"
	case COLOR_ATTACHMENT8:
		name = "GL_COLOR_ATTACHMENT8"
	case SAMPLES:
		name = "GL_SAMPLES"
	case BUFFER_ACCESS_FLAGS:
		name = "GL_BUFFER_ACCESS_FLAGS"
	case TEXTURE0:
		name = "GL_TEXTURE0"
	case FASTEST:
		name = "GL_FASTEST"
	case UNSIGNED_INT_SAMPLER_BUFFER:
		name = "GL_UNSIGNED_INT_SAMPLER_BUFFER"
	case PIXEL_PACK_BUFFER:
		name = "GL_PIXEL_PACK_BUFFER"
	case STENCIL_FUNC:
		name = "GL_STENCIL_FUNC"
	case BGR_INTEGER:
		name = "GL_BGR_INTEGER"
	case FRAMEBUFFER_DEFAULT_LAYERS:
		name = "GL_FRAMEBUFFER_DEFAULT_LAYERS"
	case ALPHA:
		name = "GL_ALPHA"
	case INT_IMAGE_3D:
		name = "GL_INT_IMAGE_3D"
	case VERTEX_BINDING_OFFSET:
		name = "GL_VERTEX_BINDING_OFFSET"
	case SHADER_SOURCE_LENGTH:
		name = "GL_SHADER_SOURCE_LENGTH"
	case HALF_FLOAT:
		name = "GL_HALF_FLOAT"
	case MAX_TESS_EVALUATION_ATOMIC_COUNTERS:
		name = "GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS"
	case MAX_TESS_EVALUATION_OUTPUT_COMPONENTS:
		name = "GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS"
	case SAMPLER_CUBE_SHADOW:
		name = "GL_SAMPLER_CUBE_SHADOW"
	case R16:
		name = "GL_R16"
	case DEBUG_NEXT_LOGGED_MESSAGE_LENGTH:
		name = "GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH"
	case TEXTURE_SWIZZLE_B:
		name = "GL_TEXTURE_SWIZZLE_B"
	case TEXTURE_CUBE_MAP_ARRAY:
		name = "GL_TEXTURE_CUBE_MAP_ARRAY"
	case TRANSFORM_FEEDBACK_VARYING:
		name = "GL_TRANSFORM_FEEDBACK_VARYING"
	case COMPRESSED_SRGB8_ETC2:
		name = "GL_COMPRESSED_SRGB8_ETC2"
	case UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX:
		name = "GL_UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX"
	case IMAGE_CLASS_4_X_16:
		name = "GL_IMAGE_CLASS_4_X_16"
	case SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE:
		name = "GL_SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE"
	case COMPATIBLE_SUBROUTINES:
		name = "GL_COMPATIBLE_SUBROUTINES"
	case RGBA_SNORM:
		name = "GL_RGBA_SNORM"
	case MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS:
		name = "GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS"
	case MAX_TESS_CONTROL_OUTPUT_COMPONENTS:
		name = "GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS"
	case PRIMITIVE_RESTART_INDEX:
		name = "GL_PRIMITIVE_RESTART_INDEX"
	case ACTIVE_TEXTURE:
		name = "GL_ACTIVE_TEXTURE"
	case FLOAT_MAT4x3:
		name = "GL_FLOAT_MAT4x3"
	case SAMPLER_1D:
		name = "GL_SAMPLER_1D"
	case ARRAY_BUFFER:
		name = "GL_ARRAY_BUFFER"
	case DRAW_BUFFER15:
		name = "GL_DRAW_BUFFER15"
	case TEXTURE_DEPTH_SIZE:
		name = "GL_TEXTURE_DEPTH_SIZE"
	case GEOMETRY_OUTPUT_TYPE:
		name = "GL_GEOMETRY_OUTPUT_TYPE"
	case LEQUAL:
		name = "GL_LEQUAL"
	case TESS_GEN_VERTEX_ORDER:
		name = "GL_TESS_GEN_VERTEX_ORDER"
	case MAX_TESS_EVALUATION_IMAGE_UNIFORMS:
		name = "GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS"
	case SHADER_COMPILER:
		name = "GL_SHADER_COMPILER"
	case RGB32UI:
		name = "GL_RGB32UI"
	case UNSIGNED_INT_SAMPLER_1D_ARRAY:
		name = "GL_UNSIGNED_INT_SAMPLER_1D_ARRAY"
	case RGB565:
		name = "GL_RGB565"
	case UNIFORM_IS_ROW_MAJOR:
		name = "GL_UNIFORM_IS_ROW_MAJOR"
	case DISPLAY_LIST:
		name = "GL_DISPLAY_LIST"
	case TEXTURE_IMMUTABLE_FORMAT:
		name = "GL_TEXTURE_IMMUTABLE_FORMAT"
	case COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT:
		name = "GL_COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT"
	case INT_VEC2:
		name = "GL_INT_VEC2"
	case TEXTURE_RECTANGLE:
		name = "GL_TEXTURE_RECTANGLE"
	case IMAGE_CLASS_1_X_32:
		name = "GL_IMAGE_CLASS_1_X_32"
	case BLEND_DST:
		name = "GL_BLEND_DST"
	case MAX_FRAGMENT_UNIFORM_COMPONENTS:
		name = "GL_MAX_FRAGMENT_UNIFORM_COMPONENTS"
	case BLUE:
		name = "GL_BLUE"
	case UNSIGNALED:
		name = "GL_UNSIGNALED"
	case SHADER:
		name = "GL_SHADER"
	case SCISSOR_TEST:
		name = "GL_SCISSOR_TEST"
	case TESS_CONTROL_OUTPUT_VERTICES:
		name = "GL_TESS_CONTROL_OUTPUT_VERTICES"
	case DRAW_BUFFER4:
		name = "GL_DRAW_BUFFER4"
	case MINOR_VERSION:
		name = "GL_MINOR_VERSION"
	case READ_FRAMEBUFFER_BINDING:
		name = "GL_READ_FRAMEBUFFER_BINDING"
	case TIME_ELAPSED:
		name = "GL_TIME_ELAPSED"
	case MAX_COMBINED_DIMENSIONS:
		name = "GL_MAX_COMBINED_DIMENSIONS"
	case FRAMEBUFFER_DEFAULT_HEIGHT:
		name = "GL_FRAMEBUFFER_DEFAULT_HEIGHT"
	case VERTEX_ATTRIB_BINDING:
		name = "GL_VERTEX_ATTRIB_BINDING"
	case FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER:
		name = "GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER"
	case MAX_TEXTURE_IMAGE_UNITS:
		name = "GL_MAX_TEXTURE_IMAGE_UNITS"
	case PROXY_TEXTURE_1D_ARRAY:
		name = "GL_PROXY_TEXTURE_1D_ARRAY"
	case DRAW_INDIRECT_BUFFER:
		name = "GL_DRAW_INDIRECT_BUFFER"
	case RGB16:
		name = "GL_RGB16"
	case RGB16F:
		name = "GL_RGB16F"
	case COLOR:
		name = "GL_COLOR"
	case R11F_G11F_B10F:
		name = "GL_R11F_G11F_B10F"
	case R32I:
		name = "GL_R32I"
	case MAX_COMBINED_TEXTURE_IMAGE_UNITS:
		name = "GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS"
	case DEBUG_TYPE_ERROR:
		name = "GL_DEBUG_TYPE_ERROR"
	case R8:
		name = "GL_R8"
	case UNSIGNED_SHORT:
		name = "GL_UNSIGNED_SHORT"
	case SRGB_WRITE:
		name = "GL_SRGB_WRITE"
	case NUM_SHADER_BINARY_FORMATS:
		name = "GL_NUM_SHADER_BINARY_FORMATS"
	case ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES:
		name = "GL_ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES"
	case DEBUG_TYPE_UNDEFINED_BEHAVIOR:
		name = "GL_DEBUG_TYPE_UNDEFINED_BEHAVIOR"
	case TEXTURE_BINDING_1D:
		name = "GL_TEXTURE_BINDING_1D"
	case TESS_GEN_SPACING:
		name = "GL_TESS_GEN_SPACING"
	case TEXTURE3:
		name = "GL_TEXTURE3"
	case SAMPLER_2D_RECT_SHADOW:
		name = "GL_SAMPLER_2D_RECT_SHADOW"
	case NUM_PROGRAM_BINARY_FORMATS:
		name = "GL_NUM_PROGRAM_BINARY_FORMATS"
	case UNSIGNED_INT_8_8_8_8:
		name = "GL_UNSIGNED_INT_8_8_8_8"
	case READ_PIXELS:
		name = "GL_READ_PIXELS"
	case UNSIGNED_INT_VEC4:
		name = "GL_UNSIGNED_INT_VEC4"
	case RGB10_A2:
		name = "GL_RGB10_A2"
	case STENCIL_PASS_DEPTH_FAIL:
		name = "GL_STENCIL_PASS_DEPTH_FAIL"
	case LOCATION:
		name = "GL_LOCATION"
	case SHADER_IMAGE_STORE:
		name = "GL_SHADER_IMAGE_STORE"
	case UNIFORM_BUFFER_START:
		name = "GL_UNIFORM_BUFFER_START"
	case FRAMEBUFFER_DEFAULT:
		name = "GL_FRAMEBUFFER_DEFAULT"
	case RGBA16I:
		name = "GL_RGBA16I"
	case ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS:
		name = "GL_ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS"
	case INT_IMAGE_CUBE_MAP_ARRAY:
		name = "GL_INT_IMAGE_CUBE_MAP_ARRAY"
	case RGBA8I:
		name = "GL_RGBA8I"
	case FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE:
		name = "GL_FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE"
	case IMAGE_BUFFER:
		name = "GL_IMAGE_BUFFER"
	case IMAGE_FORMAT_COMPATIBILITY_BY_CLASS:
		name = "GL_IMAGE_FORMAT_COMPATIBILITY_BY_CLASS"
	case COLOR_ATTACHMENT5:
		name = "GL_COLOR_ATTACHMENT5"
	case STATIC_DRAW:
		name = "GL_STATIC_DRAW"
	case QUERY_WAIT:
		name = "GL_QUERY_WAIT"
	case BACK_LEFT:
		name = "GL_BACK_LEFT"
	case MAX_GEOMETRY_OUTPUT_VERTICES:
		name = "GL_MAX_GEOMETRY_OUTPUT_VERTICES"
	case PROGRAM_PIPELINE_BINDING:
		name = "GL_PROGRAM_PIPELINE_BINDING"
	case INT_VEC3:
		name = "GL_INT_VEC3"
	case BLUE_INTEGER:
		name = "GL_BLUE_INTEGER"
	case TEXTURE16:
		name = "GL_TEXTURE16"
	case ACTIVE_VARIABLES:
		name = "GL_ACTIVE_VARIABLES"
	case BUFFER_USAGE:
		name = "GL_BUFFER_USAGE"
	case TOP_LEVEL_ARRAY_STRIDE:
		name = "GL_TOP_LEVEL_ARRAY_STRIDE"
	case COMPRESSED_RGBA:
		name = "GL_COMPRESSED_RGBA"
	case MAX_VERTEX_OUTPUT_COMPONENTS:
		name = "GL_MAX_VERTEX_OUTPUT_COMPONENTS"
	case RGB12:
		name = "GL_RGB12"
	case STENCIL_INDEX4:
		name = "GL_STENCIL_INDEX4"
	case MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS:
		name = "GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS"
	case DST_ALPHA:
		name = "GL_DST_ALPHA"
	case DYNAMIC_READ:
		name = "GL_DYNAMIC_READ"
	case FRAMEBUFFER_UNSUPPORTED:
		name = "GL_FRAMEBUFFER_UNSUPPORTED"
	case PACK_ROW_LENGTH:
		name = "GL_PACK_ROW_LENGTH"
	case TRANSFORM_FEEDBACK_BUFFER_MODE:
		name = "GL_TRANSFORM_FEEDBACK_BUFFER_MODE"
	case PACK_COMPRESSED_BLOCK_DEPTH:
		name = "GL_PACK_COMPRESSED_BLOCK_DEPTH"
	case SAMPLE_MASK_VALUE:
		name = "GL_SAMPLE_MASK_VALUE"
	case SHADER_TYPE:
		name = "GL_SHADER_TYPE"
	case MAX_SERVER_WAIT_TIMEOUT:
		name = "GL_MAX_SERVER_WAIT_TIMEOUT"
	case STENCIL_WRITEMASK:
		name = "GL_STENCIL_WRITEMASK"
	case BLEND_SRC:
		name = "GL_BLEND_SRC"
	case FRAMEBUFFER_BLEND:
		name = "GL_FRAMEBUFFER_BLEND"
	case FIRST_VERTEX_CONVENTION:
		name = "GL_FIRST_VERTEX_CONVENTION"
	case LINE_STRIP_ADJACENCY:
		name = "GL_LINE_STRIP_ADJACENCY"
	case STREAM_COPY:
		name = "GL_STREAM_COPY"
	case POLYGON_OFFSET_LINE:
		name = "GL_POLYGON_OFFSET_LINE"
	case SAMPLE_BUFFERS:
		name = "GL_SAMPLE_BUFFERS"
	case COLOR_ATTACHMENT1:
		name = "GL_COLOR_ATTACHMENT1"
	case MAX_VERTEX_ATTRIB_BINDINGS:
		name = "GL_MAX_VERTEX_ATTRIB_BINDINGS"
	case MAX_FRAGMENT_UNIFORM_BLOCKS:
		name = "GL_MAX_FRAGMENT_UNIFORM_BLOCKS"
	case SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST:
		name = "GL_SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST"
	case MAX_VERTEX_ATOMIC_COUNTER_BUFFERS:
		name = "GL_MAX_VERTEX_ATOMIC_COUNTER_BUFFERS"
	case COLOR_ATTACHMENT2:
		name = "GL_COLOR_ATTACHMENT2"
	case OR_REVERSE:
		name = "GL_OR_REVERSE"
	case VERTEX_ATTRIB_ARRAY_POINTER:
		name = "GL_VERTEX_ATTRIB_ARRAY_POINTER"
	case MIN_FRAGMENT_INTERPOLATION_OFFSET:
		name = "GL_MIN_FRAGMENT_INTERPOLATION_OFFSET"
	case LINE_WIDTH:
		name = "GL_LINE_WIDTH"
	case SHADER_STORAGE_BUFFER_START:
		name = "GL_SHADER_STORAGE_BUFFER_START"
	case SAMPLER_2D_MULTISAMPLE_ARRAY:
		name = "GL_SAMPLER_2D_MULTISAMPLE_ARRAY"
	case POLYGON_OFFSET_POINT:
		name = "GL_POLYGON_OFFSET_POINT"
	case TEXTURE_MAX_ANISOTROPY_EXT:
		name = "GL_TEXTURE_MAX_ANISOTROPY_EXT"
	case SAMPLER_2D_SHADOW:
		name = "GL_SAMPLER_2D_SHADOW"
	case INFO_LOG_LENGTH:
		name = "GL_INFO_LOG_LENGTH"
	case POLYGON_OFFSET_FACTOR:
		name = "GL_POLYGON_OFFSET_FACTOR"
	case BUFFER_MAP_LENGTH:
		name = "GL_BUFFER_MAP_LENGTH"
	case SAMPLE_ALPHA_TO_ONE:
		name = "GL_SAMPLE_ALPHA_TO_ONE"
	case VIEWPORT:
		name = "GL_VIEWPORT"
	case EXTENSIONS:
		name = "GL_EXTENSIONS"
	case CLEAR_BUFFER:
		name = "GL_CLEAR_BUFFER"
	case MAX_ELEMENTS_INDICES:
		name = "GL_MAX_ELEMENTS_INDICES"
	case POINT_SPRITE_COORD_ORIGIN:
		name = "GL_POINT_SPRITE_COORD_ORIGIN"
	case GREATER:
		name = "GL_GREATER"
	case TEXTURE24:
		name = "GL_TEXTURE24"
	case PACK_IMAGE_HEIGHT:
		name = "GL_PACK_IMAGE_HEIGHT"
	case TEXTURE11:
		name = "GL_TEXTURE11"
	case MAX_ATOMIC_COUNTER_BUFFER_SIZE:
		name = "GL_MAX_ATOMIC_COUNTER_BUFFER_SIZE"
	case MAX_PATCH_VERTICES:
		name = "GL_MAX_PATCH_VERTICES"
	case CLAMP_READ_COLOR:
		name = "GL_CLAMP_READ_COLOR"
	case ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS:
		name = "GL_ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS"
	case VERTEX_BINDING_DIVISOR:
		name = "GL_VERTEX_BINDING_DIVISOR"
	case TEXTURE28:
		name = "GL_TEXTURE28"
	case DOUBLEBUFFER:
		name = "GL_DOUBLEBUFFER"
	case RGB8:
		name = "GL_RGB8"
	case LINEAR_MIPMAP_LINEAR:
		name = "GL_LINEAR_MIPMAP_LINEAR"
	case PACK_COMPRESSED_BLOCK_HEIGHT:
		name = "GL_PACK_COMPRESSED_BLOCK_HEIGHT"
	case LOW_FLOAT:
		name = "GL_LOW_FLOAT"
	case SET:
		name = "GL_SET"
	case COMPRESSED_RG11_EAC:
		name = "GL_COMPRESSED_RG11_EAC"
	case COLOR_ATTACHMENT9:
		name = "GL_COLOR_ATTACHMENT9"
	case TEXTURE_VIEW_MIN_LAYER:
		name = "GL_TEXTURE_VIEW_MIN_LAYER"
	case TEXTURE_LOD_BIAS:
		name = "GL_TEXTURE_LOD_BIAS"
	case RG32UI:
		name = "GL_RG32UI"
	case IMAGE_CLASS_2_X_32:
		name = "GL_IMAGE_CLASS_2_X_32"
	case UNSIGNED_INT_SAMPLER_2D_ARRAY:
		name = "GL_UNSIGNED_INT_SAMPLER_2D_ARRAY"
	case IMAGE_BINDING_NAME:
		name = "GL_IMAGE_BINDING_NAME"
	case PROXY_TEXTURE_CUBE_MAP:
		name = "GL_PROXY_TEXTURE_CUBE_MAP"
	case MAX_COMPUTE_UNIFORM_BLOCKS:
		name = "GL_MAX_COMPUTE_UNIFORM_BLOCKS"
	case MAX_COMBINED_IMAGE_UNIFORMS:
		name = "GL_MAX_COMBINED_IMAGE_UNIFORMS"
	case ACTIVE_PROGRAM:
		name = "GL_ACTIVE_PROGRAM"
	case COLOR_RENDERABLE:
		name = "GL_COLOR_RENDERABLE"
	case VERTEX_ATTRIB_ARRAY_NORMALIZED:
		name = "GL_VERTEX_ATTRIB_ARRAY_NORMALIZED"
	case UNIFORM_BUFFER_OFFSET_ALIGNMENT:
		name = "GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT"
	default:
		name = fmt.Sprintf("GL_ENUM_%v", enum)
	}
	return
}
