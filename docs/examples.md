# Runnable examples

???+ example "Prerequisites"

    To play with it:

    1. Clone the repository:
    	 ```bash
    	 gh repo clone mszostok/version
    	 ```
    2. Navigate to [`example`](https://github.com/mszostok/version/tree/main/examples) directory.
    3. Run a given example.


## [Cobra](https://github.com/mszostok/version/tree/main/examples/cobra/main.go)

![](assets/examples/screen-cobra-version_-h.png)
![](assets/examples/screen-cobra-version.png)

!!! run-example "Run in terminal"
    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.buildDate=`date`'" -o example ./cobra

    # Showcase
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

    # Showcase
    ./example version
    ./example version -ojson
    ```

## [Custom Upgrade Notice](https://github.com/mszostok/version/tree/main/examples/upgrade-notice-custom/main.go)

![](assets/examples/screen-upgrade-notice-custom-version.png)

!!! run-example "Run in terminal"
    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.version=0.6.0'" -o example ./upgrade-notice-custom

    # Showcase
    ./example version
    ```

## [Upgrade Notice sub-command](https://github.com/mszostok/version/tree/main/examples/upgrade-notice-sub-cmd)

![](assets/examples/screen-upgrade-notice-sub-cmd-version_check.png)

!!! run-example "Run in terminal"
    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.version=0.6.0'" -o example ./upgrade-notice-sub-cmd

    # Showcase
    ./example version check
    ```

## [Printer](https://github.com/mszostok/version/tree/main/examples/printer/main.go)

![](assets/examples/screen-printer-.png)
![](assets/examples/screen-printer--oyaml.png)
![](assets/examples/screen-printer--oshort.png)

!!! run-example "Run in terminal"
    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.buildDate=`date`'" -o example ./printer

    # Showcase
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

    # Showcase
    ./example
    ```

## [Plain](https://github.com/mszostok/version/tree/main/examples/plain/main.go)

![](assets/examples/screen-plain-.png)

!!! run-example "Run in terminal"

    ```bash
    # Build
    go build  -ldflags "-X 'go.szostok.io/version.buildDate=`date`'" -o example ./plain

    # Showcase
    ./example
    ```

## [Custom Formatting](https://github.com/mszostok/version/tree/main/examples/custom-formatting/main.go)

![](assets/examples/screen-custom-formatting-.png)

!!! run-example "Run in terminal"
    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.buildDate=`date`'" -o example ./custom-formatting

    # Showcase
    ./example
    ```

## [Custom Layout](https://github.com/mszostok/version/tree/main/examples/custom-layout/main.go)

![](assets/examples/screen-custom-layout-.png)

!!! run-example "Run in terminal"
    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.buildDate=`date`'" -o example ./custom-layout

    # Showcase
    ./example
    ```

## [Custom Renderer](https://github.com/mszostok/version/tree/main/examples/custom-renderer/main.go)

![](assets/examples/screen-custom-renderer-.png)

!!! run-example "Run in terminal"
    ```bash
    # Build
    go build -ldflags "-X 'go.szostok.io/version.buildDate=`date`'" -o example ./custom-renderer

    # Showcase
    ./example
    ```
