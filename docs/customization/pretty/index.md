# From inspiration to creation

Did you find your muse? or wants to align with your cool branding?

**We've got your back!**

## Options

### Go believers

You can define both the formatting and layout using Go code. With dedicated functions:

- `version.WithPrettyFormat` to override only the format
- `version.WithPrettyLayout` to override only the layout
- `version.WithPrettyStyle` to override both the format and layout

you can achieve what you need without leaving you favourite language.

### I come from K8s ecosystem

So YAMLs, then? Yes! However, we are not so strict, you can also provide a JSON if you want.

You can configure that in different ways:

1. Enable loading style from environment variable via `version.WithPrettyStyleFromEnv("ENV_NAME_FOR_FILE_PATH")`,
2. Load a style file directly using Go function, `version.WithPrettyStyleFile`.
