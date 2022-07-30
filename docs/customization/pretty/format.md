# Format

!!! tip
    Formatting is the style of a pretty version's format.

Format allows you to define your own theme and adjust output for your branding color. In general, add underscores, bold, italic, and text and background colors.

## Go

Example usage:

```go
func main() {
	format := style.Formatting{
		Header: style.Header{
			Prefix: "▓▓▓ ",
			FormatPrimitive: style.FormatPrimitive{
				Color:   "magenta",
				Options: []string{"bold"},
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
