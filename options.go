package version

import "github.com/mszostok/version/style"

// Inspired by https://github.com/kubernetes-sigs/controller-runtime/blob/v0.12.3/pkg/client/options.go
// It allows to have the same functional opt func across all constructors. For example:
//
// - version.NewCobraCmd(version.WithPrettyRenderer(..))
// - version.NewPrinter(version.WithPrettyRenderer(..))
// - version.NewPrettyPrinter(version.WithPrettyRenderer(..))
//

// "Functional" Option Interfaces

type PrinterContainerOption interface {
	ApplyToPrinterContainer(*PrinterContainer)
}

type PrettyPrinterOption interface {
	ApplyPrettyPrinterOption(*Pretty)
}

type CobraExtensionOption interface {
	ApplyCobraExtensionOption(*CobraExtensionOptions)
}

// CustomPrettyRenderer

type CustomPrettyRenderer struct {
	rendererFn PrettyRenderFunc
}

// WithPrettyRenderer sets a custom renderer function.
func WithPrettyRenderer(renderer PrettyRenderFunc) *CustomPrettyRenderer {
	return &CustomPrettyRenderer{
		rendererFn: renderer,
	}
}

func (c *CustomPrettyRenderer) ApplyToPrinterContainer(cfg *PrinterContainer) {
	cfg.prettyOptions = append(cfg.prettyOptions, c)
}

func (c *CustomPrettyRenderer) ApplyCobraExtensionOption(cfg *CobraExtensionOptions) {
	cfg.PrinterOptions = append(cfg.PrinterOptions, c)
}

func (c *CustomPrettyRenderer) ApplyPrettyPrinterOption(cfg *Pretty) {
	cfg.customRenderFn = c.rendererFn
}

// CustomPrettyFormatting

type CustomPrettyFormatting struct {
	formatting *style.Formatting
}

// WithPrettyFormatting sets a custom pretty formatting.
func WithPrettyFormatting(formatting *style.Formatting) *CustomPrettyFormatting {
	return &CustomPrettyFormatting{
		formatting: formatting,
	}
}

func (c *CustomPrettyFormatting) ApplyToPrinterContainer(cfg *PrinterContainer) {
	cfg.prettyOptions = append(cfg.prettyOptions, c)
}

func (c *CustomPrettyFormatting) ApplyCobraExtensionOption(cfg *CobraExtensionOptions) {
	cfg.PrinterOptions = append(cfg.PrinterOptions, c)
}

func (c *CustomPrettyFormatting) ApplyPrettyPrinterOption(cfg *Pretty) {
	if c == nil || c.formatting == nil {
		return
	}
	cfg.defaultRenderConfig.Formatting = *c.formatting
}

// CustomPrettyLayout

type CustomPrettyLayout struct {
	layout style.Layout
}

// WithPrettyLayout sets a custom pretty layout.
func WithPrettyLayout(layout style.Layout) *CustomPrettyLayout {
	return &CustomPrettyLayout{
		layout: layout,
	}
}

func (c *CustomPrettyLayout) ApplyToPrinterContainer(cfg *PrinterContainer) {
	cfg.prettyOptions = append(cfg.prettyOptions, c)
}

func (c *CustomPrettyLayout) ApplyCobraExtensionOption(cfg *CobraExtensionOptions) {
	cfg.PrinterOptions = append(cfg.PrinterOptions, c)
}

func (c *CustomPrettyLayout) ApplyPrettyPrinterOption(cfg *Pretty) {
	cfg.defaultRenderConfig.Layout = c.layout
}

// CustomPrettyStyle

type CustomPrettyStyle struct {
	cfg *style.Config
}

// WithPrettyStyle sets a custom pretty style.
func WithPrettyStyle(cfg *style.Config) *CustomPrettyStyle {
	return &CustomPrettyStyle{
		cfg: cfg,
	}
}

func (c *CustomPrettyStyle) ApplyToPrinterContainer(cfg *PrinterContainer) {
	cfg.prettyOptions = append(cfg.prettyOptions, c)
}

func (c *CustomPrettyStyle) ApplyCobraExtensionOption(cfg *CobraExtensionOptions) {
	cfg.PrinterOptions = append(cfg.PrinterOptions, c)
}

func (c *CustomPrettyStyle) ApplyPrettyPrinterOption(cfg *Pretty) {
	cfg.defaultRenderConfig = c.cfg
}
