# Runnable examples

???+ example "Prerequisites"

    To play with the examples:

    1. Clone the repository:
    	 ```bash
    	 gh repo clone mszostok/version
    	 ```
    2. Navigate to the [`examples`](https://github.com/mszostok/version/tree/main/examples) directory.
    3. Run a chosen example.

## [Cobra](https://github.com/mszostok/version/tree/main/examples/cobra/main.go)

![](assets/examples/screen-cobra-version_-h.png)
![](assets/examples/screen-cobra-version.png)

!!! run-example "Run in terminal"

    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.buildDate=`date`'" -o example ./cobra

    # Try out
    ./example version -h
    ./example version
    ```

## [Cobra Upgrade Notice](https://github.com/mszostok/version/tree/main/examples/upgrade-notice-cobra/main.go)

![](assets/examples/screen-upgrade-notice-cobra-version.png)
![](assets/examples/screen-upgrade-notice-cobra-version_-ojson.png)

!!! run-example "Run in terminal"

    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.version=0.6.0'" -o example ./upgrade-notice

    # Try out
    ./example version
    ./example version -ojson
    ```

## [Printer Upgrade Notice](https://github.com/mszostok/version/tree/main/examples/upgrade-notice-custom/main.go)

![](assets/examples/screen-upgrade-notice-custom-version.png)

!!! run-example "Run in terminal"

    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.version=0.6.0'" -o example ./upgrade-notice-custom

    # Try out
    ./example version
    ```

## [Upgrade Notice sub-command](https://github.com/mszostok/version/tree/main/examples/upgrade-notice-sub-cmd)

![](assets/examples/screen-upgrade-notice-sub-cmd-version_check.png)

!!! run-example "Run in terminal"

    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.version=0.6.0'" -o example ./upgrade-notice-sub-cmd

    # Try out
    ./example
    ```

## [Standalone Upgrade Notice](https://github.com/mszostok/version/tree/main/examples/upgrade-notice-standalone)

![](assets/examples/screen-upgrade-notice-standalone-.png)

!!! run-example "Run in terminal"

    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.version=0.6.0'" -o example ./upgrade-notice-standalone

    # Try out
    ./example version check
    ```

## [Custom Fields](https://github.com/mszostok/version/tree/main/examples/custom-fields/main.go)

![](assets/examples/screen-custom-fields-.png)

!!! run-example "Run in terminal"

    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.buildDate=`date`'" -o example ./custom-fields

    # Try out
    ./example
    ```

## [Printer](https://github.com/mszostok/version/tree/main/examples/printer/main.go)

![](assets/examples/screen-printer-.png)
![](assets/examples/screen-printer--oyaml.png)
![](assets/examples/screen-printer--oshort.png)

!!! run-example "Run in terminal"

    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.buildDate=`date`'" -o example ./printer

    # Try out
    ./example
    ./example -oyaml
    ./example version -oshort
    ```

## [Printer Post Hook](https://github.com/mszostok/version/tree/main/examples/printer-post-hook/main.go)

![](assets/examples/screen-printer-post-hook-.png)

!!! run-example "Run in terminal"

    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.version=0.6.0'" -o example ./printer-post-hook

    # Try out
    ./example
    ```

## [Plain](https://github.com/mszostok/version/tree/main/examples/plain/main.go)

![](assets/examples/screen-plain-.png)

!!! run-example "Run in terminal"

    ```bash
    # Build
    go build  -ldflags "-X 'go.szostok.io/version.buildDate=`date`'" -o example ./plain

    # Try out
    ./example
    ```

## [Custom Formatting](https://github.com/mszostok/version/tree/main/examples/custom-formatting/main.go)

![](assets/examples/screen-custom-formatting-.png)

!!! run-example "Run in terminal"

    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.buildDate=`date`'" -o example ./custom-formatting

    # Try out
    ./example
    ```

## [Custom Layout](https://github.com/mszostok/version/tree/main/examples/custom-layout/main.go)

![](assets/examples/screen-custom-layout-.png)

!!! run-example "Run in terminal"

    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.buildDate=`date`'" -o example ./custom-layout

    # Try out
    ./example
    ```

## [Custom Style from environment variable](https://github.com/mszostok/version/tree/main/examples/custom-layout/main.go)

![](assets/examples/screen-custom-layout--env-style.png)

!!! run-example "Run in terminal"

    ```bash
    # Build
    export CLI_STYLE="${PWD}/style.yaml"
    go build -ldflags "-X 'go.szostok.io/version.buildDate=`date`'" -o example ./custom-layout

    # Try out
    ./example
    ```

## [Custom Renderer](https://github.com/mszostok/version/tree/main/examples/custom-renderer/main.go)

![](assets/examples/screen-custom-renderer-.png)

!!! run-example "Run in terminal"

    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.buildDate=`date`'" -o example ./custom-renderer

    # Try out
    ./example
    ```
