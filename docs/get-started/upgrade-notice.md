# Upgrade notice

![](../assets/examples/screen-upgrade-notice-cobra-version.png)

!!! tip ""

    Currently, the upgrade notice works only for GitHub releases.

!!! tip

    Want to try? See the [upgrade notice](../../examples#cobra-upgrade-notice) examples!

The upgrade notice is disabled by default. You can easily enable it based on your usage:

- Printer

  ```go
  p := printer.New(
    printer.WithUpgradeNotice("mszostok", "codeowners-validator", upgradeOpts...),
  )
  ```

  It prints the notice to the standard error channel ([`stderr`](<https://en.wikipedia.org/wiki/Standard_streams#Standard_error_(stderr)>)). As a result, output processing, such as executing `<cli> -ojson | jq .gitCommit`, works properly even if the upgrade notice is displayed.

- Cobra CLI

  ```go
  extension.NewVersionCobraCmd(
      // 2. Explicit turn on upgrade notice
      extension.WithUpgradeNotice("mszostok", "codeowners-validator"),
  ),
  ```

  It prints the notice on standard error channel ([`stderr`](<https://en.wikipedia.org/wiki/Standard_streams#Standard_error_(stderr)>)). As a result, output processing, such as executing `<cli> version -ojson | jq .gitCommit`, works properly even if the upgrade notice is displayed.

- Standalone

  ```go
  notice := upgrade.NewGitHubDetector("mszostok", "codeowners-validator")
  err := notice.PrintIfFoundGreater(os.Stderr, "0.5.4")
  ```

Once enabled, each execution checks for new releases but only once every 24 hours. If a newer version has been found, it displays an upgrade notice for each output format to the standard
error channel ([`stderr`](<https://en.wikipedia.org/wiki/Standard_streams#Standard_error_(stderr)>)).

You can customize almost all aspects of the upgrade check. See [customization](../../customization/upgrade-notice) for more details.
