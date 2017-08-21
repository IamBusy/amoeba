package amoeba

import (
	"github.com/drgomesp/cargo/container"
)

var app *container.Container

func init() {
	app = container.New()
}

func Set(id string, tranformer interface{}) {
	app.Set(id, tranformer)
}

func Get(id string) (service interface{}, err error) {
	return app.Get(id)
}

func MustGet(id string) interface{} {
	return app.MustGet(id)
}
