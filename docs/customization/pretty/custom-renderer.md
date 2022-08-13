# Custom Renderer

If the custom [formatting](./format.md) and [layout](./layout.md) don't meet your needs, you can simply specify your own rendering function.

!!! tip

    Want to try? See the [custom renderer](/examples#custom-renderer) example!

```go
func main() {
	renderFn := func(in *version.Info, isSmartTerminal bool) (string, error) {
		return fmt.Sprintf(`
      Version             %q
      Git Commit          %.4s
   `, in.Version, in.GitCommit), nil
	}

	p := printer.New(printer.WithPrettyRenderer(renderFn))
	if err := p.Print(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
```
