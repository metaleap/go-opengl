go-opengl
=========

Go packages for working with OpenGL:


core:
=====


An OpenGL binding inspired by and 90% similar to [GoGL](https://github.com/chsc/gogl) but generated by [my own binding-generator](https://github.com/go3d/go-opengl/tree/master/cmd/gen-opengl-bindings) and fine-tuned to the needs of [go:ngine](http://github.com/go3d/go-ngine).

(Better to call it a blatant rip-off as a little hacking exercise to grok what's going on behind the curtains in a CGO binding and to be able to customize it slightly.)

If you need an OpenGL binding, you may first want to try [the pre-generated GoGL bindings](https://github.com/chsc/gogl) or their own binding-generator.

This OpenGL binding differs from GoGL's bindings as follows:

- deprecated/removed API functions/enums are not included (though seems like the newest GoGL can also skip these)
- does not include "compatibility profile"-only functions/enums (GoGL can also gen core-only bindings but their gl42 pre-made pkg wasn't, last I checked)
- support for GL features up to 4.3 (~~GoGL is currently at 4.2~~)
- has select GL extensions (those deemed worthy of support in go:ngine) included in the binding package without needing to generate/build/link additional *ext*, *arb*, etc. binding packages.
- Init() and stolen-from-GoGL utility functions moved to a *Util* struct: so every exported (non-method) function is a direct CGO binding function. (Just minor cosmetics really.)
- added to *Util*: EnumName(), ErrorFlags(), Error() methods
- runtime checking for individual function availability via a *Supports* struct
- function wrappers that check for GL errors and return Go errors via a *Try* struct
- the above items add massively to the generated binding pkg (core.a is 7.8MB whereas GoGL's is only 2MB), but hey, the Go linker strips unused code anyway.


util:
=====


Based on the above **core** GL binding, provides higher-level utility constructs and a few sane slim wrappers (without going full-on "OO"-overboard) for Go.



cmd:
====


- gen-opengl-bindings: can be used to generate  OpenGL bindings like the above **core** package
- opengl-minimal-app: a "minimal" program using OpenGL core profile to draw a triangle and a quad, to test the above binding with GLFW
- gogl-minimal-app: same but with the GoGL (gl42) binding
