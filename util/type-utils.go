package ugl

import (
	gl "github.com/go3d/go-opengl/core"

	unum "github.com/metaleap/go-util/num"
)

const (
	//	The size of a gl.Float in bytes
	SizeOfGlFloat gl.Sizeiptr = 4

	//	The size of a gl.Uint in bytes
	SizeOfGlUint gl.Sizeiptr = 4
)

var (
	PtrNil = gl.Ptr(nil)
)

//	Represents a quaternion or 4-dimensional vector (32-bit gl.Float components)
type GlVec4 [4]gl.Float

//	Sets me to the first 4-or-less vals.
func (me *GlVec4) Set(vals ...gl.Float) {
	for i := 0; (i < len(me)) && (i < len(vals)); i++ {
		me[i] = vals[i]
	}
}

//	Represents a 3x3 matrix of 9 32-bit gl.Float values
type GlMat3 [9]gl.Float

//	Instantiates and returns a new 3x3 matrix (32-bit) from the specified 3x3 matrix (64-bit).
func NewGlMat3(mat *unum.Mat3) (glMat *GlMat3) {
	if glMat = (&GlMat3{}); mat != nil {
		glMat.Load(mat)
	}
	return
}

//	Sets all values in this 3x3 matrix (32-bit) from the specified 3x3 matrix (64-bit).
func (me *GlMat3) Load(mat *unum.Mat3) {
	me[0], me[3], me[6] = gl.Float(mat[0]), gl.Float(mat[3]), gl.Float(mat[6])
	me[1], me[4], me[7] = gl.Float(mat[1]), gl.Float(mat[4]), gl.Float(mat[7])
	me[2], me[5], me[8] = gl.Float(mat[2]), gl.Float(mat[5]), gl.Float(mat[8])
}

//	Represents a 3x3 matrix of 16 32-bit gl.Float values
type GlMat4 [16]gl.Float

//	Instantiates and returns a new 4x4 matrix (32-bit) from the specified 4x4 matrix (64-bit).
func NewGlMat4(mat *unum.Mat4) (glMat *GlMat4) {
	if glMat = (&GlMat4{}); mat != nil {
		glMat.Load(mat)
	}
	return
}

//	Sets all values in this 4x4 matrix (32-bit) from the specified 4x4 matrix (64-bit).
func (me *GlMat4) Load(mat *unum.Mat4) {
	me[0], me[4], me[8], me[12] = gl.Float(mat[0]), gl.Float(mat[4]), gl.Float(mat[8]), gl.Float(mat[12])
	me[1], me[5], me[9], me[13] = gl.Float(mat[1]), gl.Float(mat[5]), gl.Float(mat[9]), gl.Float(mat[13])
	me[2], me[6], me[10], me[14] = gl.Float(mat[2]), gl.Float(mat[6]), gl.Float(mat[10]), gl.Float(mat[14])
	me[3], me[7], me[11], me[15] = gl.Float(mat[3]), gl.Float(mat[7]), gl.Float(mat[11]), gl.Float(mat[15])
}

//	A singleton type, only used for the package-global ugl.Typed variable.
type TypeUtils struct {
}

//	Clamps the specified val between min and max.
func (_ TypeUtils) Clamp(val, min, max gl.Float) gl.Float {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
	// return MinF(MaxF(val, min), max)
}

//	If one but not two, return ifOne. If two but not one, return ifTwo. Else, return ifBoth.
func (_ TypeUtils) Ifb(one, two bool, ifOne, ifTwo, ifBoth gl.Enum) (val gl.Enum) {
	if one && !two {
		val = ifOne
	} else if two && !one {
		val = ifTwo
	} else {
		val = ifBoth
	}
	return
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func (_ TypeUtils) Ife(cond bool, ifTrue, ifFalse gl.Enum) gl.Enum {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func (_ TypeUtils) Ifi(cond bool, ifTrue, ifFalse gl.Int) gl.Int {
	if cond {
		return ifTrue
	}
	return ifFalse
}

/*
//	Returns the maximum of two values.
func Max(one, two gl.Float) gl.Float {
	if two > one {
		return two
	}
	return one
}

//	Returns the minimum of two values.
func Min(one, two gl.Float) gl.Float {
	if two < one {
		return two
	}
	return one
}
*/
