package target

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/muesli/reflow/indent"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"

	"go.szostok.io/version"
	"go.szostok.io/version/style"
)

const (
	formatMDFilePath = "./docs/customization/pretty/format.md"
	layoutMDFilePath = "./docs/customization/pretty/layout.md"
)

func EmbedDefaultPrettyFormatting() {
	formatting := map[string]any{
		"formatting": style.DefaultFormatting(),
	}

	replacePrettyExample(formatMDFilePath, "Format", formatting)
}

func EmbedDefaultPrettyLayout() {
	layout := map[string]any{
		"layout": style.Layout{GoTemplate: version.PrettyKVLayoutGoTpl},
	}

	replacePrettyExample(layoutMDFilePath, "Layout", layout)
}

func replacePrettyExample(fileName, sectionName string, in interface{}) {
	// YAML
	var buff bytes.Buffer
	buff.WriteString("```yaml\n")
	enc := yaml.NewEncoder(&buff)
	enc.SetIndent(2)
	lo.Must0(enc.Encode(in))
	buff.WriteString("```\n")

	Replace(fileName, "YAML"+sectionName, buff.Bytes())

	// JSON
	buff.Reset()
	buff.WriteString("```json\n")
	buff.Write(lo.Must(json.MarshalIndent(in, "", "  ")))
	buff.WriteString("\n```\n")

	Replace(fileName, "JSON"+sectionName, buff.Bytes())
}
func EmbedDefaultPrettyStyle() {

}

func Replace(fileName, section string, data []byte) {
	var (
		sectionStart = regexp.MustCompile(fmt.Sprintf("<!-- %s start -->", section))
		sectionEnd   = regexp.MustCompile(fmt.Sprintf("<!-- %s end -->", section))
	)

	file := lo.Must(os.Open(fileName))

	tocBlock := false

	var buff bytes.Buffer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if sectionEnd.Match(scanner.Bytes()) {
			tocBlock = false
		}

		line := scanner.Bytes()

		if !tocBlock {
			buff.Write(scanner.Bytes())
			buff.WriteString("\n")
		}

		if sectionStart.Match(line) {
			tocBlock = true
			buff.Write(indent.Bytes(data, 4))
		}
	}
	lo.Must0(os.WriteFile(fileName, buff.Bytes(), 0o644))
}
