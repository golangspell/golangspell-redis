package template

import (
	"github.com/golangspell/golangspell-redis/appcontext"
	"github.com/golangspell/golangspell-redis/config"
	"github.com/golangspell/golangspell/gateway/template"
)

//getRenderer lazy loads a Renderer
func getRenderer() appcontext.Component {
	return &template.Renderer{}
}

func init() {
	if config.Values.TestRun {
		return
	}

	appcontext.Current.Add(appcontext.Renderer, getRenderer)
}
