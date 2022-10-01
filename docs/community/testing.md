# Package testing

The `version` package is tested both via unit and e2e tests.

## Unit tests

Unit tests focus more on the corner cases that are hard to reproduce using e2e testing. Unit tests are executed on CI via [**Testing**](https://github.com/mszostok/version/actions/workflows/testing.yml) workflow.

- All tests are executed with the latest Go version on all platforms, using GitHub Action job strategy:
  ```yaml
  strategy:
  	matrix:
  		os: [ ubuntu-latest, macos-latest, windows-latest ]
  ```
- All tests are run both on pull-requests and the `main` branch
- The tests' coverage is uploaded to [coveralls.io/github/mszostok/version](https://coveralls.io/github/mszostok/version)

## E2E tests

The e2e tests build a Go binary, run it, and compare with [golden files](https://github.com/mszostok/version/tree/main/tests/e2e/testdata). E2E tests are executed on CI via [**Testing**](https://github.com/mszostok/version/actions/workflows/testing.yml) workflow.

As a result, e2e test focus on:

- Building Go binaries
- Overriding version information via `ldflags`
- Running binary on operating system
- Testing if color output for non-tty output streams is disabled automatically
- Ensuring that all [examples](https://github.com/mszostok/version/tree/main/examples) are runnable
- Executing a real call against GitHub API
- Executing binaries on all platforms, using GitHub Action job strategy:
	```yaml
	strategy:
		matrix:
			os: [ ubuntu-latest, macos-latest, windows-latest ]
	```

Each time a new functionality is implemented, a dedicated [test case](https://github.com/mszostok/version/blob/main/tests/e2e/e2e_test.go#L31) is added.

!!! note

    Currently, there is no easy way to calculate the coverage based on the e2e tests (built and executed binaries). However, this will be enabled once the [golang#51430](https://github.com/golang/go/issues/51430) issue will be implemented.
