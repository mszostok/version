# GoReleaser

The build customization is nicely described at [`goreleaser.com/customization/build`](https://goreleaser.com/customization/build). To make it work with `go.szostok.io/version`, adjust your `builds[*].ldfags` sections:

```yaml
# .goreleaser.yaml
builds:
  - # .. your settings ..

    ldflags:
      - -s -w
      - -X go.szostok.io/version.version={{.Version}}
      - -X go.szostok.io/version.buildDate={{.Date}}
```

The remaining properties are set based on the built-in data. However, for full customization, use:

```yaml
# .goreleaser.yaml
builds:
  - # .. your settings ..

    ldflags:
      - -s -w
      - -X go.szostok.io/version.version={{.Version}}
      - -X go.szostok.io/version.buildDate={{.Date}}
      - -X go.szostok.io/version.commit={{.FullCommit}}
      - -X go.szostok.io/version.commitDate={{.CommitDate}}
      - -X go.szostok.io/version.dirtyBuild=false
```
