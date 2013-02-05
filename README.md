go-opengl
=========

Go packages for working with OpenGL. While more or less adequate for general use, they're being developed in conjunction with [**go:ngine**](http://github.com/go3d/go-ngine) and thus are often slightly tweaked or optimized for **go:ngine**'s usage needs and patterns.


core:
=====


A (very) [**GoGL**](https://github.com/chsc/gogl)-like OpenGL binding for [**go:ngine**](http://github.com/go3d/go-ngine).

- includes strictly only GL core profile functionality from version 3.3 onwards (up until 4.3) -- note, this does not mean that a core profile context is *required*, although it is probably advisable on any current-gen GPU with recent drivers
- no compatibility-profile-only functionality or any that was deprecated or removed with 3.3 or earlier
- Init() and "stolen-from-GoGL" utility functions moved to a *Util* struct: so every exported (non-method) function is a direct CGO binding function. (Just minor cosmetics really.)
- added to *Util*: ErrorFlags() method
- embeds the following GL extensions right inside the same binding package:
	- [EXT_texture_filter_anisotropic](http://www.opengl.org/registry/specs/EXT/texture_filter_anisotropic.txt)
	- [EXT_texture_compression_s3tc](http://www.opengl.org/registry/specs/EXT/texture_compression_s3tc.txt)
	- [EXT_texture_sRGB](http://www.opengl.org/registry/specs/EXT/texture_sRGB.txt)


util:
=====


Based on the above **core** GL binding package, provides higher-level utility constructs and a few sane lean minimalist wrappers (without going full-on "OO"-overboard).



cmd:
====


- **opengl-minimal-app**: a "minimal" program using OpenGL core profile to draw a triangle and a quad, to test the above **core** GL binding package with [**GLFW**](https://github.com/go-gl/glfw)
- **gogl-minimal-app**: same but with the GoGL (gl43) binding
- **gen-opengl-bindings**: used to generate the above **core** GL binding package
