package style

// Layout define the layout for printing the pretty format.
type Layout struct {
	// GoTemplate defines layout in with Go template syntax.
	GoTemplate string `json:"goTemplate,omitempty" yaml:"goTemplate,omitempty"`
}
