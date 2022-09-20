package printer

// SetOutputFormat is used only for testing purposes. It's in _test.go so it won't be available as public API.
func (r *Container) SetOutputFormat(in OutputFormat) {
	r.output = in
}
