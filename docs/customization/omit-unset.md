!!! coming-soon "Coming soon"

    See the [mszostok/version#12](https://github.com/mszostok/version/issues/12) issue for a reference. If you want to see it, please add 👍 under the issue.

## Examples

- Explicitly exclude a given set of version fields:

    ```go
    // excludedFields defines preset for fields that should be excluded in output.
    const excludedFields = printer.FieldCompiler | printer.FieldPlatform

    p := printer.New(printer.WithExcludedFields(excludedFields))
    if err := p.Print(os.Stdout); err != nil {
    	log.Fatal(err)
    }
    ```

- Don't display empty(`""`) and unset(`N/A`) version fields:

    ```go
    p := printer.New(printer.WithOmitUnset(excludedFields))
    if err := p.Print(os.Stdout); err != nil {
    	log.Fatal(err)
    }
    ```
