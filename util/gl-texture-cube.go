package ugl

import (
	"image"

	gl "github.com/go3d/go-opengl/core"
)

//	Represents six 2-dimensional texture images of equal dimensions.
type TextureCube struct {
	TextureBase
	Width, Height gl.Sizei
}

//	Initializes --but does not Recreate()-- a new TextureCube, calls its Init() method and returns it.
func NewTextureCube() (me *TextureCube) {
	me = &TextureCube{}
	me.Init()
	return
}

//	Sets GlTarget to gl.TEXTURE_CUBE_MAP and initializes TextureBase with defaults.
func (me *TextureCube) Init() {
	me.GlTarget = gl.TEXTURE_CUBE_MAP
	me.TextureBase.init()
	me.PixelData.Ptrs = make([]gl.Ptr, 6)
}

//	Returns the maximum number of possible MIP map levels for any of this cube-map's 6 faces according to its current Width and Height.
func (me *TextureCube) MaxNumMipLevels() gl.Sizei {
	return gl.Sizei(texture2DMaxNumMipLevels(float64(me.Width), float64(me.Height)))
}

//	Prepares this TextureCube for uploading the specified Images via Recreate().
//	This sets all of the following fields to applicable values:
//	me.PixelData.Type, me.PixelData.Format, me.PixelData.Ptrs, me.Width, me.Height, me.MipMap.NumLevels, me.SizedInternalFormat
func (me *TextureCube) PrepFromImages(bgra, uintRev bool, images ...image.Image) (err error) {
	if len(images) < 6 {
		err = errf("TextureCube.PrepFromImages: expected 6 cube-faces, but only %v were supplied", len(images))
	} else {
		me.Width, me.Height = gl.Sizei(images[0].Bounds().Dx()), gl.Sizei(images[0].Bounds().Dy())
		me.MipMap.NumLevels = me.MaxNumMipLevels()
		err = me.prepFromImages(bgra, uintRev, images...)
	}
	return
}

//	Deletes and (re)creates the texture object based on its current params.
//	Uploads the image data pointed to in me.PixelData.Ptrs, if any.
//	Generates the MIP map if me.MipMap.AutoGen is true and me.MipMap.NumLevels isn't 1.
//	If me.MipMap.NumLevels is 0 (or just smaller than 1), then me.MaxNumMipLevels() is used.
func (me *TextureCube) Recreate() error {
	return me.recreate(6, gl.TEXTURE_CUBE_MAP_POSITIVE_X, me.MaxNumMipLevels(), me.Width, me.Height)
}

func (me *TextureCube) SubImage(glTargetFace gl.Enum, x, y gl.Int, width, height gl.Sizei, ptr gl.Ptr) error {
	return me.subImage(glTargetFace, x, y, width, height, ptr)
}
