package extension_test

import (
	"bytes"
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

	cmd := extension.NewVersionCobraCmd(
		extension.WithAliasesOptions("v", "vv"),
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
}
