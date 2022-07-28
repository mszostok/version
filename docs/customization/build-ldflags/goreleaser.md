# GoReleaser

The build customization is nicely described at [goreleaser.com/customization/build](https://goreleaser.com/customization/build). To make it work with `go.szostok.io/version`, adjust your `builds[*].ldfags` sections:

```yaml
# .goreleaser.yaml
builds:
  -
    # .. your settings ..

    ldflags:
      - -s -w
      - -X go.szostok.io/version.version={{.Version}}
      - -X go.szostok.io/version.buildDate={{.Date}}
```

The rest properties are set based on the built-in data. However, if you want to have a full customization, add:

```yaml
# .goreleaser.yaml
builds:
  -
    # .. your settings ..

    ldflags:
      - -s -w
      - -X go.szostok.io/version.version={{.Version}}
      - -X go.szostok.io/version.buildDate={{.Date}}
      - -X go.szostok.io/commit={{.FullCommit}}
      - -X go.szostok.io/commitDate={{.CommitDate}}
      - -X go.szostok.io/dirtyBuild=false
```
