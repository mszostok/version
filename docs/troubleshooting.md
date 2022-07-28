# Troubleshooting

This page lists solutions to problems you might encounter with `go.szostok.io/version`.

## Build binary doesn't have default version data

If your binary doesn't have default version data, for example:

```bash
  Version             (devel)
  Git Commit          N/A
  Build Date          28 Jul 22 22:07 CEST (2 seconds ago)
  Commit Date
  Dirty Build         no
  Go Version          1.18.2
  Compiler            gc
  Platform            darwin/amd64
```

You can see that git commit and commit date are not populated. The problem might be related to the build process. Ensure that you don't specify the `main.go` file directly.
Instead of running:

```bash
go build -o example ./cmd/client/main.go
```

Run

```bash
go build -o example ./cmd/client # NOTE: only folder specified.
```

Difference

```diff
-go build -o example ./cobra/main.go
+go build -o example ./cobra/
```
