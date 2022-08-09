# Upgrade notice

![](../../assets/examples/screen-upgrade-notice-cobra-version.png)

!!! tip ""
    Currently, it works only for GitHub releases.

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

You can customize almost all aspect of the upgrade check:

- Set maximum duration time for update check operation (default 10s).

    ```go
    upgrade.WithUpdateCheckTimeout(30*time.Second)
    ```

- Set a custom function to compare release versions (default SemVer check).

    ```go
    upgrade.WithIsVersionGreater(func(current string, new string) bool {
      	//.. compare current with new ..
      	return true
    })
    ```

- Set the minimum time that must elapse before checking for a new release (default 24h).

    ```go
    upgrade.WithMinElapseTimeForRecheck(time.Second)
    ```

- Change formatting.

    ```go
    upgrade.WithFormatting(&style.Formatting{
			Header: style.Header{},
			Key:    style.Key{},
			Val:    style.Val{},
			Date:   style.Date{},
		})
    ```

- Change [layout](./layout.md).

    ```go
    upgrade.WithLayout(&style.Layout{
    			GoTemplate: forBoxLayoutGoTpl,
    		})
    ```

- Change both formatting and layout.

    ```go
    upgrade.WithStyle(&style.Config{})
    ```

- Define [custom renderer](./custom-renderer.md).

    ```go
    upgrade.WithRenderer(func(in *upgrade.Info) (string, error) {
    	return fmt.Sprintf(`
    		Version             %q
    		New Version         %q
    		Published At        %v
    	`, in.Version, in.NewVersion, in.PublishedAt), nil
    })
    ```

- Add post render hook.

    ```go
    upgrade.WithPostRenderHook(func(body string) (string, error) {
    	return body + "\ncustom footer", nil
    })
    ```
