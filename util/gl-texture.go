package ugl

import (
	"image"
	"reflect"

	gl "github.com/go3d/go-opengl/core"
)

//	Implemented by specialized texture types such as Texture2D.
type Texture interface {
	Bind()

	//	Deletes and (re)creates the texture object based on its current params.
	Recreate() error

	Unbind()
}

//	Embedded by specialized texture types such as Texture2D.
type TextureBase struct {
	//	The OpenGL handle for this texture object.
	//	This is 0 before calling Recreate() and after calling Dispose().
	GlHandle gl.Uint

	//	The type of texture, such as for example gl.TEXTURE_2D.
	GlTarget gl.Enum

	//	Specifies the sized internal format to be used to store texture image data,
	//	as per gl.TexStorageNN(). Defaults to gl.RGBA8.
	SizedInternalFormat gl.Enum

	//	If true, then (if Support.Textures.Immutable is also true) whenever this Texture
	//	is Recreate()d it is declared immutable in OpenGL, meaning its dimensions are
	//	locked at creation time and cannot be changed subsequently.
	//	Defaults to Support.Textures.Immutable.
	Immutable bool

	//	Settings for this texture's MIP map, if any.
	MipMap struct {
		//	If true (the default) and NumLevels isn't 1, all MIP map
		//	levels are automatically generated by Recreate()
		AutoGen bool

		//	The maximum number of MIP map levels for this texture object.
		//	Set to 0 (the default) to have Recreate() determine this
		//	automatically, set to 1 for a texture object with no MIP map.
		NumLevels gl.Sizei
	}

	//	Information regarding the pixel data stored by this texture object.
	PixelData struct {
		//	Specifies the format of the pixel data, as per
		//	gl.TexSubImageNN(). Defaults to gl.RGBA.
		Format gl.Enum

		//	Specifies the data type of the pixel data, as per
		//	gl.TexSubImageNN(). Defaults to gl.UNSIGNED_BYTE.
		Type gl.Enum

		//	Pointers (one per sub-image) to the first pixel
		//	of the data stream to be uploaded by Recreate().
		//	Initially defaults to []gl.Ptr { gl.Ptr(nil) }
		Ptrs []gl.Ptr
	}
}

//	Binds this texture object.
func (me *TextureBase) Bind() {
	gl.BindTexture(me.GlTarget, me.GlHandle)
}

//	Deletes this texture object from OpenGL.
func (me *TextureBase) Dispose() {
	if me.GlHandle != 0 {
		gl.DeleteTextures(1, &me.GlHandle)
		me.GlHandle = 0
	}
}

func (me *TextureBase) immutable() bool {
	return Support.Textures.Immutable && me.Immutable
}

func (me *TextureBase) init() {
	me.MipMap.AutoGen, me.Immutable = true, Support.Textures.Immutable
	me.SizedInternalFormat = gl.RGBA8
	me.PixelData.Format = gl.RGBA
	me.PixelData.Type = gl.UNSIGNED_BYTE
	me.PixelData.Ptrs = []gl.Ptr{gl.Ptr(nil)}
}

func (me *TextureBase) onAfterRecreate() {
	me.Unbind()
}

func (me *TextureBase) onBeforeRecreate() (err error) {
	if me.immutable() {
		me.Unbind()
		me.Dispose()
		err = Try.GenTextures(1, &me.GlHandle)
	} else if me.GlHandle == 0 {
		err = Try.GenTextures(1, &me.GlHandle)
	}
	if err == nil {
		me.Bind()
	}
	return
}

func (me *TextureBase) prepFromImages(images ...image.Image) (err error) {
	pixData := &me.PixelData
	pixData.Type = gl.UNSIGNED_BYTE
	if len(pixData.Ptrs) < len(images) {
		nuPtrs := make([]gl.Ptr, len(images))
		copy(nuPtrs, pixData.Ptrs)
		pixData.Ptrs = nuPtrs
	}
	for i, img := range images {
		switch pic := img.(type) {
		case *image.Alpha:
			me.SizedInternalFormat = gl.R8
			pixData.Format = gl.RED
			pixData.Ptrs[i] = gl.Ptr(&pic.Pix[0])
		case *image.Alpha16:
			me.SizedInternalFormat = gl.R16
			pixData.Format = gl.RED
			pixData.Ptrs[i] = gl.Ptr(&pic.Pix[0])
		case *image.Gray:
			me.SizedInternalFormat = gl.R8
			pixData.Format = gl.RED
			pixData.Ptrs[i] = gl.Ptr(&pic.Pix[0])
		case *image.Gray16:
			me.SizedInternalFormat = gl.R16
			pixData.Format = gl.RED
			pixData.Ptrs[i] = gl.Ptr(&pic.Pix[0])
		case *image.NRGBA:
			me.SizedInternalFormat = gl.RGBA8
			pixData.Format = gl.RGBA
			pixData.Ptrs[i] = gl.Ptr(&pic.Pix[0])
		case *image.NRGBA64:
			me.SizedInternalFormat = gl.RGBA16
			pixData.Format = gl.RGBA
			pixData.Ptrs[i] = gl.Ptr(&pic.Pix[0])
		case *image.RGBA:
			me.SizedInternalFormat = gl.RGBA8
			pixData.Format = gl.RGBA
			pixData.Ptrs[i] = gl.Ptr(&pic.Pix[0])
		case *image.RGBA64:
			me.SizedInternalFormat = gl.RGBA16
			pixData.Format = gl.RGBA
			pixData.Ptrs[i] = gl.Ptr(&pic.Pix[0])
		default:
			err = errf("Unsupported image.Image type (%v) for use as OpenGL texture", reflect.TypeOf(pic))
		}
	}
	return
}

//	Unbinds whatever texture is currently bound.
func (me *TextureBase) Unbind() {
	gl.BindTexture(me.GlTarget, 0)
}
