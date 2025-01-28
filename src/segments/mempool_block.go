package segments

import (
	"github.com/jandedobbeleer/oh-my-posh/src/properties"
	"github.com/jandedobbeleer/oh-my-posh/src/runtime"
)

type MempoolBlock struct {
	base
}

const (
	//NewProp enables something
	NewProp properties.Property = "newprop"
)

func (n *MempoolBlock) Enabled() bool {
	return true
}

func (n *MempoolBlock) Template() string {
	return " {{.Text}} "
}

func (n *MempoolBlock) Init(props properties.Properties, env runtime.Environment) {
	n.props = props
	n.env = env
}
