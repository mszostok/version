# Custom Renderer

If the custom [formatting](./format.md) and [layout](./layout.md) don't fulfill you needs, you can simply specify your own rendering function.

!!! tip

    Want to try? See the [custom renderer](/examples#custom-renderer) example!

```go
func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "example",
		Short: "An example CLI built with github.com/spf13/cobra",
	}

	renderFn := func(in *version.Info) (string, error) {
		return fmt.Sprintf(`
      Version             %q
      Git Commit          %.4s
   `, in.Version, in.GitCommit), nil
	}

	cmd.AddCommand(
		// you just need to add this, and you are done.
		version.NewCobraCmd(version.WithPrettyRenderer(renderFn)),
	)

	return cmd
}
```
