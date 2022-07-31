# Custom Renderer

If the custom formatting and [layout](./layout.md) don't fulfil you needs, you can simply specify your own rendering function.

```go
// NewRoot returns a root cobra.Command for the whole CLI.
func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "example",
		Short: "An example CLI built with github.com/spf13/cobra",
	}

	opts := []upgrade.Options{
		upgrade.WithRenderer(func(in *upgrade.Info) (string, error) {
			return fmt.Sprintf(`
      Version             %q
      New Version         %q
      Published At        %v
   `, in.Version, in.NewVersion, in.PublishedAt), nil
		}),
	}

	cmd.AddCommand(
		// 1. Register 'version' command
		version.NewCobraCmd(
			// 2. Explict turn on upgrade notice
			version.WithUpgradeNotice("mszostok", "codeowners-validator", opts...)),
	)

	return cmd
}
```
