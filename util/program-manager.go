package ugl

import (
	"time"

	ustr "github.com/metaleap/go-util/str"
)

const progManCap = 16

type ProgramManager struct {
	Programs []Program

	names map[string]int
}

func (me *ProgramManager) Init() {
	me.Programs = make([]Program, 0, progManCap)
	me.names = make(map[string]int, progManCap)
}

func (me *ProgramManager) Dispose() {
	for i := 0; i < len(me.Programs); i++ {
		me.Programs[i].Dispose()
	}
}

func (me *ProgramManager) AddNew(name string) (prog *Program) {
	if len(me.Programs) == cap(me.Programs) {
		nuProgs := make([]Program, len(me.Programs), cap(me.Programs)+progManCap)
		copy(nuProgs, me.Programs)
		me.Programs = nuProgs
	}
	index := len(me.Programs)
	me.Programs = append(me.Programs, Program{})
	prog = &me.Programs[index]
	prog.Init(name)
	me.names[name] = index
	return
}

func (me *ProgramManager) Get(name string) (prog *Program) {
	if i, ok := me.names[name]; ok {
		prog = &me.Programs[i]
	}
	return
}

func (me *ProgramManager) MakeProgramsFromRawSources(defines map[string]interface{}, forceAll bool, forceSome ...string) (dur time.Duration, progsMade []bool, err error) {
	progsMade = make([]bool, len(me.Programs))
	timeStart := time.Now()
	for i := 0; i < len(me.Programs); i++ {
		if me.Programs[i].GlHandle == 0 || forceAll || ustr.IsInSlice(forceSome, me.Programs[i].Name) {
			progsMade[i] = true
			me.Programs[i].Create()
			if err = me.Programs[i].CompileAndLinkShaders(defines); err != nil {
				me.Programs[i].Dispose()
				return
			}
		}
	}
	dur = time.Now().Sub(timeStart)
	return
}
