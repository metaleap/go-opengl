package ugl

import (
	"fmt"

	"github.com/metaleap/go-util-misc"
)

var (
	Cache GlCache

	//	Provides access to miscellaneous GL functionality.
	Util GlUtil

	//	Provides access to miscellaneous type-specific GL functionality.
	Typed TypeUtils

	//	Provides methods that each wrap a GL function call and immediately check for GL errors to be returned as Go errors.
	Try GlTry
)

//	DO call this immediately upon successful OpenGL context initialization to prepare this package for use.
//	Fully populates all fields and sub-fields in the Support package global.
func Init() {
	setVersion()
	setSupportInfos()
}

//	log.Println()s the error returned by Util.Error(), if any.
func LogLastError(stepFmt string, stepFmtArgs ...interface{}) {
	if err := Util.LastError(stepFmt, stepFmtArgs...); err != nil {
		ugo.LogError(err)
	}
}

func errf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

func strf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
