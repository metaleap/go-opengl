package ugl

import (
	"time"

	"github.com/metaleap/go-util-slice"
)

const progManCap = 16

type ProgramManager struct {
	All []Program

	names map[string]int
}

func (me *ProgramManager) Init() {
	me.All = make([]Program, 0, progManCap)
	me.names = make(map[string]int, progManCap)
}

func (me *ProgramManager) Dispose() {
	for i := 0; i < len(me.All); i++ {
		me.All[i].Dispose()
	}
}

func (me *ProgramManager) AddNew(name string) (prog *Program) {
	if len(me.All) == cap(me.All) {
		nuProgs := make([]Program, len(me.All), cap(me.All)+progManCap)
		copy(nuProgs, me.All)
		me.All = nuProgs
	}
	index := len(me.All)
	me.All = append(me.All, Program{})
	prog = &me.All[index]
	prog.Init(index, name)
	me.names[name] = index
	return
}

func (me *ProgramManager) Get(name string) (prog *Program) {
	if i, ok := me.names[name]; ok {
		prog = &me.All[i]
	}
	return
}

func (me *ProgramManager) Index(name string) (index int) {
	var ok bool
	if index, ok = me.names[name]; !ok {
		index = -1
	}
	return
}

func (me *ProgramManager) MakeProgramsFromRawSources(defines map[string]interface{}, forceAll bool, forceSome ...string) (dur time.Duration, progsMade []bool, err error) {
	progsMade = make([]bool, len(me.All))
	timeStart := time.Now()
	for i := 0; i < len(me.All); i++ {
		if me.All[i].GlHandle == 0 || forceAll || uslice.StrHas(forceSome, me.All[i].Name) {
			progsMade[i] = true
			me.All[i].Create()
			if err = me.All[i].CompileAndLinkShaders(defines); err != nil {
				me.All[i].Dispose()
				return
			}
		}
	}
	dur = time.Now().Sub(timeStart)
	return
}
