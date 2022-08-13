# From inspiration to creation

Have you found your muse? Are you looking to align with your cool branding?

**We've got your back!**

## Go believers

With dedicated functional options, you can achieve what you need, all in your favourite language. To name just a few:

- `printer.WithPrettyFormatting` lets you override only the format
- `printer.WithPrettyLayout` lets you override only the layout
- `upgrade.WithMinElapseTimeForRecheck` lets you override the minimum time that must elapse before checking for a new release

## I come from the K8s ecosystem

!!! coming-soon "Coming soon"

    See the [mszostok/version#13](https://github.com/mszostok/version/issues/13) issue for reference. If you'd like to see it included in a future release, add 👍 under the issue.

So YAMLs (and JSONs), then? Oh yes!

You can configure the `pretty` output style in a few ways:

- Enable loading your custom style from an environment variable via `printer.WithPrettyStyleFromEnv("ENV_NAME_FOR_FILE_PATH")`
- Load a style file directly using the Go function: `printer.WithPrettyStyleFile`
