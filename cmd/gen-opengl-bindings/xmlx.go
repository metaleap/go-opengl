package main

import (
	xmlx "github.com/jteeuwen/go-pkg-xmlx"
	ustr "github.com/metaleap/go-util/str"
)

type nodeFunc func(*xmlx.Node)

func checkForUnknownAtts(xn *xmlx.Node, knownAttNames ...string) {
	for _, att := range xn.Attributes {
		if !ustr.IsInSlice(knownAttNames, att.Name.Local) {
			println("UNKNOWN <" + xn.Name.Local + "> ATT: " + att.Name.Local)
		}
	}
}

func isLegacy(xn *xmlx.Node) bool {
	asp, asd, asr := xas(xn, "profile"), xas(xn, "deprecated"), xas(xn, "removed")
	return asp == "compatibility" || (len(asd) > 0 /*&& asd <= cfg.minVer.dotted*/) || (len(asr) > 0 /*&& asr <= cfg.minVer.dotted*/)
}

func xmlWalkDoc(nodeName string, onNode nodeFunc) {
	for _, xn := range specDoc.SelectNodesRecursive(xmlns, nodeName) {
		onNode(xn)
	}
}

func xmlWalkDocN(onNode nodeFunc, names ...string) {
	for _, name := range names {
		xmlWalkDoc(name, onNode)
	}
}

func xmlWalkExts() {
	if cfg.genExtsAll {
		i, exts := 0, map[string]bool{}
		xmlWalkDoc("ext", func(xn *xmlx.Node) {
			exts[xas(xn, "name")] = true
		})
		cfg.genExts = make([]string, len(exts))
		for ext, _ := range exts {
			cfg.genExts[i] = ext
			i++
		}
	}
}

func xas(xn *xmlx.Node, name string) string {
	return xn.As("", name)
}
