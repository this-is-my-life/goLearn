package main

import (
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func main() {
	app, err := astilectron.New(astilectron.Options{
		AppName: "Go Learn - Astilectron",
	})
	erring(err)

	defer app.Close()

	err = app.Start()
	erring(err)

	w, err := app.NewWindow("src/index.html", &astilectron.WindowOptions{Frame: astikit.BoolPtr(false)})
	erring(err)

	w.Create()

	app.Wait()
}

func erring(err error) {
	if err != nil {
		panic(err)
	}
}
