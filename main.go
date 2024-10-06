package main

import (
	"fmt"
	"os"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/core/gioutil"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/ppvan/cquan/internal/resources"
)

func main() {

	gresources, err := gio.NewResourceFromData(glib.NewBytesWithGo(resources.GResource))
	if err != nil {
		panic(err)
	}
	gio.ResourcesRegister(gresources)

	// settings := gio.NewSettings(resources.GAppID)

	app := adw.NewApplication(resources.GAppID, gio.ApplicationNonUnique)

	app.ConnectActivate(func() { activate(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

var stringListModel = gioutil.NewListModelType[string]()

func activate(app *adw.Application) {

	list := stringListModel.New()
	list.Append("Hello")
	list.Append("Listview")
	list.Append("from")
	list.Append("Go")

	builder := gtk.NewBuilderFromResource(resources.GResourceWindowPath)
	window := builder.GetObject("main-window").Cast().(*adw.ApplicationWindow)
	todo_model := builder.GetObject("todos_model").Cast().(*gtk.SingleSelection)
	todo_model.SetObjectProperty("model", list)
	listview := builder.GetObject("todos").Cast().(*gtk.ListView)

	// render list-item
	listview.SetFactory(newStringItemFactory(&todo_model.ListModel))

	app.AddWindow(&window.Window)
	window.Present()

}

func newStringItemFactory(model *gio.ListModel) *gtk.ListItemFactory {
	factory := gtk.NewSignalListItemFactory()

	factory.ConnectBind(func(object *glib.Object) {
		item := object.Cast().(*gtk.ListItem)
		_data := item.Item()
		// how to get original string from _data?
		fmt.Print(_data)
		// so I can bind it to here
		item.SetChild(gtk.NewLabel("hello"))
	})

	return &factory.ListItemFactory
}
