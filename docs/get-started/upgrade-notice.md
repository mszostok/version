# Upgrade notice

![](../assets/examples/screen-upgrade-notice-cobra-version.png)

!!! tip ""
    Currently, it works only for GitHub releases.

!!! tip

    Want to try? See the [upgrade notice](../../examples#cobra-upgrade-notice) examples!

The upgrade notice is disabled by default. You can easily enable it based on your usage:

1. Printer:
	  ```go
	  p := printer.New(
	  	printer.WithUpgradeNotice("mszostok", "codeowners-validator", upgradeOpts...),
	  )
	  ```
	 It prints the notice on standard error. As a result, executing e.g. `<cli> -ojson | jq .gitCommit` works properly even if upgrade notice is displayed.

2. Cobra CLI:
	  ```go
	  extension.NewVersionCobraCmd(
	  	// 2. Explict turn on upgrade notice
	  	extension.WithUpgradeNotice("mszostok", "codeowners-validator"),
	  ),
	  ```
	 It prints the notice on standard error. As a result, executing e.g. `<cli> version -ojson | jq .gitCommit` works properly even if upgrade notice is displayed.

3. Standalone:

    ```go
    notice := upgrade.NewGitHubDetector("mszostok", "codeowners-validator")
    err := notice.PrintIfFoundGreater(os.Stderr, "0.5.4")
    ```


Once enabled, each execution checks for new releases, but only once every 24 hours. If a newer version was found, displays upgrade notice for each output format on standard
error.

You can customize almost all aspect of upgrade check. Please see [customization](../../customization/upgrade-notice).
