package extension

import (
	"go.szostok.io/version/printer"
	"go.szostok.io/version/upgrade"
)

// Inspired by https://github.com/kubernetes-sigs/controller-runtime/blob/v0.12.3/pkg/client/options.go
// It allows having the same functional opt func across all constructors. For example:
//
// - extension.NewVersionCobraCmd(extension.WithUpgradeNotice(..))
// - extension.NewVersionCobraCmd(extension.WithPrinterOptions(..))

type (
	// CobraOption provides an option to set a Cobra options.
	CobraOption interface {
		// ApplyToCobraOption sets a given option for Cobra command.
		ApplyToCobraOption(*CobraOptions)
	}

	// CobraOptions holds Cobra command possible customization settings.
	CobraOptions struct {
		PrinterOptions []printer.ContainerOption
		Aliases        []string
	}
)

// CustomPrinterOptions provides an option to set a custom printer related options across multiple constructors.
type CustomPrinterOptions struct {
	printerOptions []printer.ContainerOption
}

// WithPrinterOptions sets a custom printer related options.
func WithPrinterOptions(opts ...printer.ContainerOption) *CustomPrinterOptions {
	return &CustomPrinterOptions{
		printerOptions: opts,
	}
}

// ApplyToCobraOption sets a given option for Cobra.
func (c *CustomPrinterOptions) ApplyToCobraOption(options *CobraOptions) {
	options.PrinterOptions = append(options.PrinterOptions, c.printerOptions...)
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

// ApplyToCobraOption sets a given option for Cobra.
// It's a syntax sugar for:
//
//	extension.NewVersionCobraCmd(
//		extension.WithPrinterOptions(
//			printer.WithUpgradeNotice("mszostok", "codeowners-validator"),
//		),
//	)
//
// so you can just do:
// extension.NewVersionCobraCmd(
//
//		extension.WithUpgradeNotice("mszostok", "codeowners-validator"),
//	)
func (c *EnableUpgradeNotice) ApplyToCobraOption(options *CobraOptions) {
	options.PrinterOptions = append(options.PrinterOptions, printer.WithUpgradeNotice(c.owner, c.repo, c.upgradeOpts...))
}

// AliasesOptions provides an option to set a custom printer related options across multiple constructors.
type AliasesOptions struct {
	aliases []string
}

// WithAliasesOptions sets a given aliases for the 'version' command.
func WithAliasesOptions(aliases ...string) *AliasesOptions {
	return &AliasesOptions{
		aliases: aliases,
	}
}

// ApplyToCobraOption sets a given option for Cobra.
func (c *AliasesOptions) ApplyToCobraOption(options *CobraOptions) {
	options.Aliases = c.aliases
}
