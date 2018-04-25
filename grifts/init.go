package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/jcepedavillamayor/apiexample/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
