package assertly

import (
	"github.com/viant/toolbox"
)

//Context represent validation context
type Context struct {
	toolbox.Context
	Directives *Directives
	Evaluator  *toolbox.MacroEvaluator
}

//NewContext returns a context
func NewContext(ctx toolbox.Context, directives *Directives, evaluator *toolbox.MacroEvaluator) *Context {
	if ctx == nil {
		ctx = toolbox.NewContext()
	}
	if directives == nil {
		directives = NewDirectives()
	}
	if evaluator == nil {
		evaluator = toolbox.NewMacroEvaluator("<ds:", ">", ValueProviderRegistry)
	}
	return &Context{
		Context:    ctx,
		Directives: directives,
		Evaluator:  evaluator,
	}
}

//NewDefaultContext returns default context
func NewDefaultContext() *Context {
	return NewContext(nil, nil, nil)
}
