package resources

import (
	_ "embed"
	"path"
)

const GAppID = "com.github.ppvan.Cquan"

const gResourceAppPath = "/com/github/ppvan/Cquan/"

var GResourceStyleCSSPath = path.Join(gResourceAppPath, "style.css")
var GResourceMetainfoPath = path.Join(gResourceAppPath, "com.github.ppvan.Cquan.metainfo.xml")

//go:generate blueprint-compiler compile --output window.ui window.blp
var GResourceWindowPath = path.Join(gResourceAppPath, "window.ui")

//go:generate glib-compile-resources com.github.ppvan.Cquan.gresource.xml
//go:embed com.github.ppvan.Cquan.gresource
var GResource []byte
