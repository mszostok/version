# From inspiration to creation

Did you find your muse? or wants to align with your cool branding?

**We've got your back!**

## Go believers

With dedicated options functions you can achieve what you need without leaving you favourite language. Just to name a few:

- `printer.WithPrettyFormatting` to override only the format.
- `printer.WithPrettyLayout` to override only the layout.
- `upgrade.WithMinElapseTimeForRecheck` to the minimum time that must elapse before checking for a new release.


## I come from the K8s ecosystem

!!! coming-soon "Coming soon"

    See the [mszostok/version#13](https://github.com/mszostok/version/issues/13) issue for a reference. If you want to see it, please add 👍 under the issue.

So YAMLs (and JSONs), then? Oh yes!

You can configure that in different ways:

1. Enable loading style from environment variable via `printer.WithPrettyStyleFromEnv("ENV_NAME_FOR_FILE_PATH")`,
2. Load a style file directly using Go function, `printer.WithPrettyStyleFile`.
