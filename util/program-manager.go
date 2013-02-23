package ugl

import (
	"time"

	ustr "github.com/metaleap/go-util/str"
)

//	A simple container to collect multiple program objects by name and compile them all in one go.
//	Prior to using a ProgramManager, you need to call its Reset() method to allocate/initialize its map fields.
//	After using a ProgramManager, Reset() cleans up and deletes resources allocated by it.
type ProgramManager struct {
	//	These name-value pairs are prepended to shader sources
	//	as #define directives in MakeProgramsFromRawSources().
	Defines map[string]interface{}

	//	The names of the Programs in this ProgramManager.
	//	The MakeProgramsFromRawSources() method uses this to populate the Programs hash-table
	//	and to locate associated shaders in RawSources for each Program.
	Names []string

	//	To be populated by the MakeProgramsFromRawSources() method.
	Programs map[string]*Program

	Sources struct {
		//	Shader sources to be linked to named programs by MakeProgramsFromRawSources(),
		//	each to be associated with a name that is also in Names.
		In ProgramSources

		//	Repopulated by each MakeProgramsFromRawSources() call, contains the actual
		//	full GLSL source strings as sent to GL, including #define and #version directives.
		Out ProgramSources
	}

	UniformBlocks map[string]*UniformBlock
}

//	Creates copies of all RawSources with srcName as dstName, and adds dstName to me.Names.
//	Typical use-case: compile an existing program anew but with different #defines (likely resulting in a different binary) under a different name.
func (me *ProgramManager) CloneRawSources(srcName, dstName string) (cloned bool) {
	if ustr.IsInSlice(me.Names, srcName) && !ustr.IsInSlice(me.Names, dstName) {
		rs := &me.Sources.In
		for _, m := range []map[string]string{rs.Compute, rs.Fragment, rs.Geometry, rs.TessCtl, rs.TessEval, rs.Vertex} {
			if len(m[srcName]) > 0 {
				m[dstName] = m[srcName]
				cloned = true
			}
		}
		if cloned {
			me.Names = append(me.Names, dstName)
		}
	}
	return
}

//	For each program name in me.Names: creates a new Program, compiles all Shader stages for that name in me.RawSources, and
//	links the Program with these compiled Shader stages. If successful, the Program is added to me.Programs under that name.
//	If forceAll is true, a Program already existing under a name is deleted and recreated; otherwise, it is kept and not recreated.
//	If forceSome has values, only those names (instead of all me.Names) are processed.
//	The name-value pairs in me.Defines (if any) are prepended to the shader sources as #define directives after the #version directive.
//	If an error occurs, this method returns it immediately; otherwise, dur contains the total time taken to make all Programs.
func (me *ProgramManager) MakeProgramsFromRawSources(forceAll bool, forceSome ...string) (dur time.Duration, err error) {
	var prog *Program
	var srcCx, srcFx, srcGx, srcHx, srcDx, srcVx, srcTmp string
	var tmpMap *map[string]string
	rs, timeStart, names := &me.Sources.In, time.Now(), me.Names
	me.Sources.Out.reset()
	if len(forceSome) > 0 {
		names, forceAll = forceSome, true
	}
	for _, name := range names {
		if prog = me.Programs[name]; prog != nil {
			if !forceAll {
				continue
			}
		}
		if prog == nil {
			prog = NewProgram(name)
			if err = prog.Create(); err != nil {
				return
			}
		}
		srcCx, srcFx, srcGx, srcHx, srcDx, srcVx = rs.Compute[name], rs.Fragment[name], rs.Geometry[name], rs.TessCtl[name], rs.TessEval[name], rs.Vertex[name]
		for tmpMap, srcTmp = range map[*map[string]string]string{
			&me.Sources.Out.Compute:  srcCx,
			&me.Sources.Out.Fragment: srcFx,
			&me.Sources.Out.Geometry: srcGx,
			&me.Sources.Out.TessCtl:  srcHx,
			&me.Sources.Out.TessEval: srcDx,
			&me.Sources.Out.Vertex:   srcVx,
		} {
			(*tmpMap)[name] = srcTmp
		}
		if err = prog.CompileAndLinkShaders(&srcCx, &srcFx, &srcGx, &srcHx, &srcDx, &srcVx, me.Defines); err != nil {
			prog.Dispose()
			return
		} else {
			me.Programs[name] = prog
		}
	}
	dur = time.Now().Sub(timeStart)
	return
}

//	Allocates/re-initializes all maps in me (Programs, Defines, RawSources) as new empty maps.
//	Also deletes all existing me.Programs (if any) from OpenGL.
func (me *ProgramManager) Reset() {
	if me.Programs != nil {
		for _, prog := range me.Programs {
			prog.Dispose()
		}
	}
	if me.UniformBlocks != nil {
		for _, ubo := range me.UniformBlocks {
			ubo.Dispose()
		}
	}
	me.UniformBlocks = map[string]*UniformBlock{}
	me.Programs = map[string]*Program{}
	me.Defines = map[string]interface{}{}
	me.Sources.In.reset()
	me.Sources.Out.reset()
}

//	Used for ProgramManager.RawSources and ProgramManager.FinalRealSources.
type ProgramSources struct {
	//	Named shader sources for the Compute Shader stage.
	Compute map[string]string

	//	Named shader sources for the Fragment Shader stage.
	Fragment map[string]string

	//	Named shader sources for the Geometry Shader stage.
	Geometry map[string]string

	//	Named shader sources for the Tessellation Control Shader stage.
	TessCtl map[string]string

	//	Named shader sources for the Tessellation Evaluation Shader stage.
	TessEval map[string]string

	//	Named shader sources for the Vertex Shader stage.
	Vertex map[string]string
}

func (me *ProgramSources) reset() {
	for _, field := range []*map[string]string{&me.Compute, &me.Fragment, &me.Geometry, &me.TessCtl, &me.TessEval, &me.Vertex} {
		*field = map[string]string{}
	}
}
