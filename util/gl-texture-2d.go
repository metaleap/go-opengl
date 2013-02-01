package glutil

import (
	"fmt"
	"image"
	"math"
	"reflect"

	gl "github.com/go3d/go-opengl/gogl"
)

//	Represents a 2-dimensional texture image.
type Texture2D struct {
	//	Common texture params
	TextureBase

	//	Width and height of mip-level 0 for this 2-dimensional texture object.
	Width, Height gl.Sizei
}

//	Initializes --but does not Recreate()-- a new Texture2D, calls its Init() method and returns it.
func NewTexture2D() (me *Texture2D) {
	me = &Texture2D{}
	me.Init()
	return
}

//	Returns the maximum number of possible MIP map levels for this 2-dimensional texture image according to its current Width and Height.
func (me *Texture2D) MaxNumMipLevels() gl.Sizei {
	return gl.Sizei(maxNumOfTexture2DMipLevels(float64(me.Width), float64(me.Height)))
}

//	Sets GlTarget to gl.TEXTURE_2D and initializes TextureBase with defaults.
func (me *Texture2D) Init() {
	me.GlTarget = gl.TEXTURE_2D
	me.TextureBase.init()
}

//	Deletes and (re)creates the texture object based on its current params.
//	Uploads the image data at me.PixelData.Ptr, if any.
//	Generates the MIP map if me.MipMap.AutoGen is true and me.MipMap.NumLevels isn't 1.
//	If me.MipMap.NumLevels is 0 (or just smaller than 1), then me.MaxNumMipLevels() is used.
func (me *Texture2D) Recreate() {
	me.onBeforeRecreate()
	hasPixData := (me.PixelData.Ptr != gl.Pointer(nil))
	numLevels := me.MipMap.NumLevels
	if numLevels < 1 {
		numLevels = me.MaxNumMipLevels()
	}
	if me.immutable() {
		if gl.TexStorage2D(me.GlTarget, numLevels, me.SizedInternalFormat, me.Width, me.Height); hasPixData {
			gl.TexSubImage2D(me.GlTarget, 0, 0, 0, me.Width, me.Height, me.PixelData.Format, me.PixelData.Type, me.PixelData.Ptr)
		}
	} else {
		gl.TexImage2D(me.GlTarget, 0, gl.Int(me.SizedInternalFormat), me.Width, me.Height, 0, me.PixelData.Format, me.PixelData.Type, me.PixelData.Ptr)
	}
	if hasPixData && me.MipMap.AutoGen && (me.GlTarget != gl.TEXTURE_RECTANGLE) {
		gl.GenerateMipmap(me.GlTarget)
	}
	me.onAfterRecreate()
}

//	Prepares this Texture2D for uploading the specified Image via Recreate().
//	This sets all of the following fields to applicable values:
//	me.PixelData.Type, me.PixelData.Format, me.PixelData.Ptr, me.Width, me.Height, me.MipMap.NumLevels, me.SizedInternalFormat
func (me *Texture2D) SetFromImage(img image.Image) (err error) {
	me.PixelData.Type = gl.UNSIGNED_BYTE
	me.Width, me.Height = gl.Sizei(img.Bounds().Dx()), gl.Sizei(img.Bounds().Dy())
	me.MipMap.NumLevels = me.MaxNumMipLevels()
	switch pic := img.(type) {
	case *image.Alpha:
		me.SizedInternalFormat = gl.R8
		me.PixelData.Format = gl.RED
		me.PixelData.Ptr = gl.Pointer(&pic.Pix[0])
	case *image.Alpha16:
		me.SizedInternalFormat = gl.R16
		me.PixelData.Format = gl.RED
		me.PixelData.Ptr = gl.Pointer(&pic.Pix[0])
	case *image.Gray:
		me.SizedInternalFormat = gl.R8
		me.PixelData.Format = gl.RED
		me.PixelData.Ptr = gl.Pointer(&pic.Pix[0])
	case *image.Gray16:
		me.SizedInternalFormat = gl.R16
		me.PixelData.Format = gl.RED
		me.PixelData.Ptr = gl.Pointer(&pic.Pix[0])
	case *image.NRGBA:
		me.SizedInternalFormat = gl.RGBA8
		me.PixelData.Format = gl.RGBA
		me.PixelData.Ptr = gl.Pointer(&pic.Pix[0])
	case *image.NRGBA64:
		me.SizedInternalFormat = gl.RGBA16
		me.PixelData.Format = gl.RGBA
		me.PixelData.Ptr = gl.Pointer(&pic.Pix[0])
	case *image.RGBA:
		me.SizedInternalFormat = gl.RGBA8
		me.PixelData.Format = gl.RGBA
		me.PixelData.Ptr = gl.Pointer(&pic.Pix[0])
	case *image.RGBA64:
		me.SizedInternalFormat = gl.RGBA16
		me.PixelData.Format = gl.RGBA
		me.PixelData.Ptr = gl.Pointer(&pic.Pix[0])
	default:
		err = fmt.Errorf("Unsupported image.Image type (%v) for use as OpenGL texture", reflect.TypeOf(pic))
	}
	return
}

//	Returns the maximum number of possible MIP map levels for a 2-dimensional texture image with the specified width and height.
func maxNumOfTexture2DMipLevels(width, height float64) float64 {
	return math.Log2(math.Max(width, height)) + 1
}
