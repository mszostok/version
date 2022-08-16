# Get Started

The `version` package was built to remove the repetitiveness in implementing the *version* command.
As such, it was designed in the form of building blocks, which you can mix and match based on your needs, and integrate with different frameworks.

For example:

- Use the [upgrade notice](upgrade-notice.md) component. The usage is not limited to the CLI only. You can embed this into all of your binaries.
- Use `version`'s core `info` object with data collected automatically via [`version.Get()`](./usage/plain).
- Use `version`'s core [printer](./usage/printer) component with `version.Get()` or with the `Info` object that you filled yourself.
- Set up [Cobra's `version` command](./usage/cobra) with the default printer and an enabled upgrade notice as shown in the [quick-start](../quick-start.md).

To get inspired, see also other [examples](../examples).

Each component can be customized. See the [customization](../customization) documentation to learn more.
