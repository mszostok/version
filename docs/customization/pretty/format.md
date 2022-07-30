# Format

!!! note ""
    Formatting focuses on the style of displayed pretty version's data.

Format allows you to define your own theme and adjust output for your branding color. In general, add underscores, bold, italic, and text and background colors.

## Go

!!! tip

    Want to try? See the [custom formatting](/examples#custom-formatting) example!

Example usage:

```go
func main() {
	format := style.Formatting{
		Header: style.Header{
			Prefix: "💡 ",
			FormatPrimitive: style.FormatPrimitive{
				Color:   "magenta",
				Options: []string{"underscore"},
			},
		},
	}
	version.NewPrinter(version.WithPrettyFormatting(format))
}
```

Check the [`style.Formatting`](https://github.com/mszostok/version/blob/main/style/formatting.go#L4) struct for all possible settings.


## Config file

The config file can be loaded by:

- enabling loading style from environment variable via `version.WithPrettyStyleFromEnv("ENV_NAME_FOR_FILE_PATH")`,
- or using `version.WithPrettyStyleFile` function directly.

=== "YAML"

    <!-- YAMLFormat start -->
    ```yaml
    formatting:
      header:
        prefix: '▓▓▓ '
        color: magenta
        background: ""
        options: []
        name: ""
      key:
        color: gray
        background: ""
        options:
          - bold
      val:
        color: white
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
          "color": "magenta",
          "background": "",
          "options": null,
          "name": ""
        },
        "key": {
          "color": "gray",
          "background": "",
          "options": [
            "bold"
          ]
        },
        "val": {
          "color": "white",
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
