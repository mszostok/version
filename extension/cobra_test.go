package extension_test

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.szostok.io/version"
	"go.szostok.io/version/extension"
	"go.szostok.io/version/printer"
)

const helpMsg = `
cobra version
cobra version -o=json
cobra version -o=yaml
cobra version -o=short
`

func TestCobra(t *testing.T) {
	// given
	cmd := extension.NewVersionCobraCmd()

	// then
	assert.Equal(t, "version", cmd.Use)
	assert.Equal(t, "Print the CLI version", cmd.Short)

	normalizedExample := strings.ReplaceAll(cmd.Example, os.Args[0], "cobra")
	assert.Equal(t, helpMsg, normalizedExample)

	require.Len(t, cmd.Aliases, 1)
	assert.Equal(t, "ver", cmd.Aliases[0])

	output, err := cmd.Flags().GetString("output")
	require.NoError(t, err)
	assert.Equal(t, "pretty", output)
}

func TestCobraOptions(t *testing.T) {
	// given
	const expMsg = "testing"
	root := cobra.Command{}
	var buff bytes.Buffer
	root.SetOut(&buff)

	preHookCalledBeforePost := false
	preHookCalled := false
	postHookCalled := false

	cmd := extension.NewVersionCobraCmd(
		extension.WithAliasesOptions("v", "vv"),
		extension.WithPreHook(func(ctx context.Context) error {
			if !postHookCalled {
				preHookCalledBeforePost = true
			}
			preHookCalled = true
			return nil
		}),
		extension.WithPostHook(func(ctx context.Context) error {
			postHookCalled = true
			return nil
		}),
		extension.WithPrinterOptions(
			printer.WithPrettyRenderer(func(in *version.Info, isSmartTerminal bool) (string, error) {
				return expMsg, nil
			}),
		),
	)

	// then
	require.Len(t, cmd.Aliases, 2)
	assert.Equal(t, "v", cmd.Aliases[0])
	assert.Equal(t, "vv", cmd.Aliases[1])

	// execute command
	assert.NotNil(t, cmd.RunE)
	err := cmd.RunE(&root, []string{})
	require.NoError(t, err)
	assert.Equal(t, expMsg, buff.String())

	// hooks
	assert.True(t, preHookCalledBeforePost)
	assert.True(t, preHookCalled)
	assert.True(t, postHookCalled)
}

func TestCobraHookOptions(t *testing.T) {
	// given
	root := cobra.Command{}

	preHookCalledBeforePost := false
	preHookCalled := false
	postHookCalled := false

	cmd := extension.NewVersionCobraCmd(
		extension.WithPreHook(func(ctx context.Context) error {
			if !postHookCalled {
				preHookCalledBeforePost = true
			}
			preHookCalled = true
			return nil
		}),
		extension.WithPostHook(func(ctx context.Context) error {
			postHookCalled = true
			return nil
		}),
	)

	// execute command
	assert.NotNil(t, cmd.RunE)
	err := cmd.RunE(&root, []string{})
	require.NoError(t, err)

	// hooks
	assert.True(t, preHookCalledBeforePost)
	assert.True(t, preHookCalled)
	assert.True(t, postHookCalled)
}
