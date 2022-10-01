```go hl_lines="19-20"
--8<--
examples/cobra/main.go
dff
--8<--
```

You can customize almost all aspects:

- Enable upgrade notice:

  ```go
  extension.NewVersionCobraCmd(
      // 2. Explicit turn on upgrade notice
      extension.WithUpgradeNotice("mszostok", "codeowners-validator"),
  ),
  ```

  It prints the notice on standard error channel ([`stderr`](<https://en.wikipedia.org/wiki/Standard_streams#Standard_error_(stderr)>)). As a result, output processing, such as executing `<cli> version -ojson | jq .gitCommit`, works properly even if the upgrade notice is displayed.

- Define pre-hook function:

  ```go
  extension.WithPreHook(func(ctx context.Context) error {
   // function body
  })
  ```

- Define post-hook function:

  !!! note ""

      It's executed only if version print was successful.

  ```go
  extension.WithPostHook(func(ctx context.Context) error {
   // function body
  })
  ```

- Define custom aliases:

  ```go
  extension.WithAliasesOptions("v", "vv"),
  ```
