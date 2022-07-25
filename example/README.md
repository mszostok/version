# Examples

Runnable examples. To play with it:
1. Clone the repository:
    ```bash
    gh repo clone mszostok/version
    ```
2. Navigate to [`example`](.) directory.
3. Download dependencies:
    ```bash
    go mod download
    ```
## Table of content

<!-- toc -->

- [Usage](#usage)
  * [Plain](#plain)
  * [Cobra](#cobra)

<!-- tocstop -->

## Usage

### Plain

```go mdox-exec="sed -n '9,21p' plain/main.go"
func main() {
	version.CollectFromBuildInfo()

	info := version.Get()
	fmt.Println("Version: ", info.Version)
	fmt.Println("Git Commit: ", info.GitCommit)
	fmt.Println("Build Date: ", info.BuildDate)
	fmt.Println("Commit Date: ", info.CommitDate)
	fmt.Println("Dirty Build: ", info.DirtyBuild)
	fmt.Println("Go Version: ", info.GoVersion)
	fmt.Println("Compiler: ", info.Compiler)
	fmt.Println("Platform: ", info.Platform)
}
```
[_source_](./plain/main.go)

**Test**
```bash
go build  -ldflags "-X 'github.com/mszostok/version.buildDate=`date`'" -o example ./plain
./example
```

### Cobra

```go mdox-exec="sed -n '12,24p' cobra/main.go"
func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "example",
		Short: "An example CLI built with github.com/spf13/cobra",
	}

	cmd.AddCommand(
		// you just need to add this, and you are done.
		version.NewCobraCmd(),
	)

	return cmd
}
```
[_source_](./cobra/main.go)

**Test**
```bash
go build  -ldflags "-X 'github.com/mszostok/version.buildDate=`date`'" -o example ./cobra
./example version -h
./example version
```
