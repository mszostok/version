<h1>
    <img alt="logo" src="./docs/assets/logo-small.png" width="28px" />
    <code>version</code> - contributing
</h1>

Thanks for your interest in the `version` project!

This document contains contribution guidelines for this repository. Read it before you start contributing.

## Contributing

Before proposing or adding changes, check the [existing issues](https://github.com/mszostok/version/issues) and make sure the discussion/work has not already been started to avoid duplication.

If you'd like to see a new feature implemented, use this [feature request template](https://github.com/mszostok/version/issues/new?assignees=&labels=&template=feature_request.md) to create an issue.

Similarly, if you spot a bug, use this [bug report template](https://github.com/mszostok/version/issues/new?labels=bug&template=bug_report.md) to let me know!

## Ready for action? Start developing!

To start contributing, follow these steps:

1. [Fork the `version` repository](https://github.com/mszostok/version/fork).

2. Clone the repository locally.

   > **Note**
   > This project uses Go modules, so you can check it out locally wherever you want. It doesn't need to be checked out in `$GOPATH`.

3. Set the `version` repository as upstream:

   ```bash
     git remote add upstream git@github.com:mszostok/version.git
   ```

4. Fetch all the remote branches for this repository:

   ```bash
     git fetch --all
   ```

5. Set the `main` branch to point to upstream:

   ```bash
     git branch -u upstream/main main
   ```

You're all set! 🚀

### Go style guidelines

This project adheres to the [Go official](https://github.com/golang/go/wiki/CodeReviewComments) and [Uber](https://github.com/uber-go/guide/blob/master/style.md) guidelines.

### Documentation

- [pip](https://pypi.org/project/pip/)
- [Magefile](https://magefile.org/)

  > **Note**
  > Run `mage -l` to see all possible targets.

#### Build the site locally

- Install `mkdocs` and the `mkdocs-material` theme

  ```sh
  pip install -r requirements.txt
  ```

- Start local server

  ```sh
  mkdocs serve -D
  ```

  Site can be viewed at http://localhost:8000

- Find the documentation page file (`.md` file) under `docs/` and edit it.
