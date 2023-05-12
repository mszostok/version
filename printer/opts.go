package printer

import (
	"encoding/json"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"go.szostok.io/version"
	"go.szostok.io/version/style"
	"go.szostok.io/version/upgrade"
)

// Inspired by https://github.com/kubernetes-sigs/controller-runtime/blob/v0.12.3/pkg/client/options.go
// It allows having the same functional opt func across all constructors. For example:
//
// - version.NewPrinter(version.WithPrettyRenderer(..))
// - version.NewPrettyPrinter(version.WithPrettyRenderer(..))

type (
	// PostHookFunc represents post execution function signature.
	PostHookFunc func() error

	// ContainerOption provides an option to set a Container options.
	ContainerOption interface {
		// ApplyToContainerOption sets a given option for Container.
		ApplyToContainerOption(*ContainerOptions)
	}

	// ContainerOptions holds Container possible customization settings.
	ContainerOptions struct {
		PrettyOptions []PrettyOption
		PostHookFunc  PostHookFunc
		UpgradeNotice *upgrade.GitHubDetector
		ExcludedField version.Field
	}

	// PrettyOption provides an option to set a Pretty printer options.
	PrettyOption interface {
		// ApplyToPrettyOption sets a given option to Pretty printer.
		ApplyToPrettyOption(*PrettyOptions)
	}

	// PrettyOptions holds Pretty possible customization settings.
	PrettyOptions struct {
		CustomRenderFn PrettyRenderFunc
		RenderConfig   *style.Config
		PostRenderFunc PrettyPostRenderFunc
	}
)

// CustomPrettyRenderer provides an option to set a custom renderer function across multiple constructors.
type CustomPrettyRenderer struct {
	rendererFn PrettyRenderFunc
}

// WithPrettyRenderer sets a custom renderer function.
func WithPrettyRenderer(renderer PrettyRenderFunc) *CustomPrettyRenderer {
	return &CustomPrettyRenderer{
		rendererFn: renderer,
	}
}

// ApplyToContainerOption sets a given option for Container.
func (c *CustomPrettyRenderer) ApplyToContainerOption(cfg *ContainerOptions) {
	cfg.PrettyOptions = append(cfg.PrettyOptions, c)
}

// ApplyToPrettyOption sets a given option to Pretty printer.
func (c *CustomPrettyRenderer) ApplyToPrettyOption(cfg *PrettyOptions) {
	cfg.CustomRenderFn = c.rendererFn
}

// CustomPrettyFormatting provides an option to set a custom pretty formatting across multiple constructors.
type CustomPrettyFormatting struct {
	formatting *style.Formatting
}

// WithPrettyFormatting sets a custom pretty formatting.
func WithPrettyFormatting(formatting *style.Formatting) *CustomPrettyFormatting {
	return &CustomPrettyFormatting{
		formatting: formatting,
	}
}

// ApplyToContainerOption sets a given option for Container.
func (c *CustomPrettyFormatting) ApplyToContainerOption(cfg *ContainerOptions) {
	cfg.PrettyOptions = append(cfg.PrettyOptions, c)
}

// ApplyToPrettyOption sets a given option to Pretty printer.
func (c *CustomPrettyFormatting) ApplyToPrettyOption(cfg *PrettyOptions) {
	if c == nil || c.formatting == nil {
		return
	}
	cfg.RenderConfig.Formatting = *c.formatting
}

// CustomPrettyLayout provides an option to set a custom pretty layout across multiple constructors.
type CustomPrettyLayout struct {
	layout *style.Layout
}

// WithPrettyLayout sets a custom pretty layout.
func WithPrettyLayout(layout *style.Layout) *CustomPrettyLayout {
	return &CustomPrettyLayout{
		layout: layout,
	}
}

// ApplyToContainerOption sets a given option for Container.
func (c *CustomPrettyLayout) ApplyToContainerOption(cfg *ContainerOptions) {
	cfg.PrettyOptions = append(cfg.PrettyOptions, c)
}

// ApplyToPrettyOption sets a given option to Pretty printer.
func (c *CustomPrettyLayout) ApplyToPrettyOption(cfg *PrettyOptions) {
	if c == nil || c.layout == nil {
		return
	}
	cfg.RenderConfig.Layout = *c.layout
}

// CustomPrettyStyle provides an option to set a custom pretty style across multiple constructors.
type CustomPrettyStyle struct {
	cfg *style.Config
}

// WithPrettyStyle sets a custom pretty style.
func WithPrettyStyle(cfg *style.Config) *CustomPrettyStyle {
	return &CustomPrettyStyle{
		cfg: cfg,
	}
}

// ApplyToContainerOption sets a given option for Container.
func (c *CustomPrettyStyle) ApplyToContainerOption(cfg *ContainerOptions) {
	if c != nil && cfg != nil {
		cfg.PrettyOptions = append(cfg.PrettyOptions, c)
	}
}

// ApplyToPrettyOption sets a given option to Pretty printer.
func (c *CustomPrettyStyle) ApplyToPrettyOption(cfg *PrettyOptions) {
	if c != nil && cfg != nil {
		cfg.RenderConfig = c.cfg
	}
}

// PrettyPostRenderHook provides an option to set post render function across multiple constructors.
type PrettyPostRenderHook struct {
	fn PrettyPostRenderFunc
}

// WithPrettyPostRenderHook sets post render function.
func WithPrettyPostRenderHook(fn PrettyPostRenderFunc) *PrettyPostRenderHook {
	return &PrettyPostRenderHook{
		fn: fn,
	}
}

// ApplyToContainerOption sets a given option for Container.
func (c *PrettyPostRenderHook) ApplyToContainerOption(cfg *ContainerOptions) {
	cfg.PrettyOptions = append(cfg.PrettyOptions, c)
}

// ApplyToPrettyOption sets a given option to Pretty printer.
func (c *PrettyPostRenderHook) ApplyToPrettyOption(cfg *PrettyOptions) {
	cfg.PostRenderFunc = c.fn
}

// PostHook provides an option to set post execution function across multiple constructors.
type PostHook struct {
	fn PostHookFunc
}

// WithPostHook sets post execution function.
func WithPostHook(fn PostHookFunc) *PostHook {
	return &PostHook{
		fn: fn,
	}
}

// ApplyToContainerOption sets a given option for Container.
func (c *PostHook) ApplyToContainerOption(cfg *ContainerOptions) {
	cfg.PostHookFunc = c.fn
}

// EnableUpgradeNotice provides an option to enable upgrade notice across multiple constructors.
type EnableUpgradeNotice struct {
	upgradeOpts []upgrade.Options
	repo        string
	owner       string
}

// WithUpgradeNotice enables upgrade notice.
func WithUpgradeNotice(owner, repo string, opts ...upgrade.Options) *EnableUpgradeNotice {
	return &EnableUpgradeNotice{
		owner:       owner,
		repo:        repo,
		upgradeOpts: opts,
	}
}

// ApplyToContainerOption sets a given option for Container.
func (c *EnableUpgradeNotice) ApplyToContainerOption(cfg *ContainerOptions) {
	cfg.UpgradeNotice = upgrade.NewGitHubDetector(c.owner, c.repo, c.upgradeOpts...)
}

// WithPrettyStyleFromEnv Load a custom style from environment variable
func WithPrettyStyleFromEnv(envVariable string) (*CustomPrettyStyle, error) {
	path := os.Getenv(envVariable)
	options, err := parseConfigFile(path)

	return options, err
}

// WithPrettyStyleFile Load a custom style from file
func WithPrettyStyleFile(path string) (*CustomPrettyStyle, error) {
	options, err := parseConfigFile(path)

	return options, err
}

func parseConfigFile(fileName string) (*CustomPrettyStyle, error) {
	if fileName == "" {
		return nil, nil
	}

	fileName = filepath.Clean(fileName)
	body, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	styles := PrettyDefaultRenderConfig()
	extension := filepath.Ext(fileName)
	switch extension {
	case ".yml", ".yaml":
		err = yaml.Unmarshal(body, styles)
	case ".json":
		err = json.Unmarshal(body, styles)
	}

	if err != nil {
		return nil, err
	}
	return &CustomPrettyStyle{
		cfg: styles,
	}, err
}

// ExcludedFields provides an option to store the version fields that are excluded by the user
type ExcludedFields struct {
	Fields version.Field
}

// ApplyToContainerOption sets a given option for Container.
func (c *ExcludedFields) ApplyToContainerOption(cfg *ContainerOptions) {
	cfg.ExcludedField = c.Fields
}

// WithOmitUnset excludes version fields that have not been set.
// version fields with the values `N/A` or `(devel)“ will be excluded
func WithOmitUnset() *ExcludedFields {
	return &ExcludedFields{
		Fields: version.UnsetFields(),
	}
}

// WithExcludedFields sets the excluded fields for a Container.
func WithExlcudedFields(field version.Field) *ExcludedFields {
	return &ExcludedFields{
		Fields: field,
	}
}
