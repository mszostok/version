# Format

!!! note ""

    Formatting focuses on the style of your displayed pretty version data.

Format lets you define your own theme and adjust the output to your branding colors. In general, you can add underscores, bold and italic formatting, text, and background colors.

## Go

!!! tip

    Want to try? See the [custom formatting](../../../examples#custom-formatting) example!

Example usage:

```go
func main() {
	format := style.Formatting{
		Header: style.Header{
			Prefix: "💡 ",
			FormatPrimitive: style.FormatPrimitive{
				Color:   "Magenta",
				Options: []string{"Underline"},
			},
		},
	}
	printer.New(printer.WithPrettyFormatting(&format))
}
```

Check the [`style.Formatting`](https://github.com/mszostok/version/blob/main/style/formatting.go#L4) struct for all possible settings.

## Config file

To load the config file, you can:

- Enable loading your custom style from an environment variable via `printer.WithPrettyStyleFromEnv("ENV_NAME_FOR_FILE_PATH")`
- Use the `printer.WithPrettyStyleFile("file_path")` function directly

=== "YAML"

    <!-- YAMLFormat start -->
    ```yaml
    formatting:
      header:
        prefix: '▓▓▓ '
        color: Magenta
        background: ""
        options: []
      key:
        color: Gray
        background: ""
        options:
          - Bold
      val:
        color: White
        background: ""
        options: []
      date:
        enableHumanizedSuffix: true
    ```
    <!-- YAMLFormat end -->

=== "JSON"

    <!-- JSONFormat start -->
    ```json
    {
      "formatting": {
        "header": {
          "prefix": "▓▓▓ ",
          "color": "Magenta",
          "background": "",
          "options": null
        },
        "key": {
          "color": "Gray",
          "background": "",
          "options": [
            "Bold"
          ]
        },
        "val": {
          "color": "White",
          "background": "",
          "options": null
        },
        "date": {
          "enableHumanizedSuffix": true
        }
      }
    }
    ```
    <!-- JSONFormat end -->
