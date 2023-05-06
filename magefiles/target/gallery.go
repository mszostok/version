package target

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/mattn/go-shellwords"
	"github.com/samber/lo"
)

const (
	yellow            = "\033[1;33m"
	nc                = "\033[0m" // No Color
	screenshotCommand = "screencapture -x -R0,25,1285,650 %s"
)

var (
	repoRootDir      = lo.Must(os.Getwd())
	currentDir       = filepath.Join(repoRootDir, "magefiles", "hack")
	assetExamplesDir = fmt.Sprintf("%s/docs/assets/examples", repoRootDir)
)

var (
	bigScreenCapture = []Program{
		{Name: "plain"},
		{Name: "cobra", Args: Args{"version"}},
		{Name: "printer", Args: Args{"", "-oyaml", "-ojson", "-oshort"}},
		{Name: "custom-formatting"},
	}

	mediumScreenCapture = []Program{
		{Name: "upgrade-notice-sub-cmd", Args: Args{"version check"}},
		{Name: "upgrade-notice-custom", Args: Args{"version"}},
		{Name: "upgrade-notice-standalone"},
		{Name: "printer-post-hook"},
		{Name: "upgrade-notice-cobra", Args: Args{"version", "version -ojson"}},
		{Name: "custom-layout"},
		{Name: "custom-layout",
			Suffix: "-env-style",
			Envs:   fmt.Sprintf("CLI_STYLE=%s", filepath.Join(repoRootDir, "examples/style.yaml"))},
		{Name: "custom-renderer"},
		{Name: "custom-fields"},
	}
	smallScreenCapture = []Program{
		{Name: "cobra", Args: Args{"version -h"}},
		{Name: "cobra-alias", Args: Args{"version -h"}},
	}
)

type Program struct {
	Name   string
	Args   Args
	Suffix string
	Envs   string
}
type Args []string

func (p *Program) Run() {
	if p == nil {
		return
	}
	if p.Args == nil {
		capture(p.Name, "", p.Envs, p.Suffix)
		return
	}
	for _, arg := range p.Args {
		capture(p.Name, arg, p.Envs, p.Suffix)
	}
}

func Gallery() {
	remove(assetExamplesDir)
	lo.Must0(os.MkdirAll(assetExamplesDir, os.ModePerm))

	// Big
	setup("version-cmd")
	for _, program := range bigScreenCapture {
		program.Run()
	}

	// Medium
	setup("custom-pretty-cmd")
	for _, program := range mediumScreenCapture {
		program.Run()
	}

	// Small
	setup("help-cmd")
	for _, program := range smallScreenCapture {
		program.Run()
	}

}

func setup(profile string) {
	lo.Must0(os.Chdir(currentDir))
	fmt.Printf("\033]50;SetProfile=%s\a", profile)
	run("osascript resize_window.scpt")
	run("clear")
	time.Sleep(500 * time.Millisecond) // wait for the setup to establish
}

func remove(path string) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		lo.Must0(os.RemoveAll(path))
	}
}

func capture(program string, arg string, env string, suffix string) {
	run("clear")

	lo.Must0(os.Chdir(filepath.Join(repoRootDir, "examples")))
	run(`go install -ldflags "-X 'go.szostok.io/version.buildDate=%s' -X 'go.szostok.io/version.version=v0.6.1'" ./%s`, time.Now(), program)
	lo.Must0(os.Chdir(os.Getenv("HOME")))
	fmt.Printf("▲ %s%s%s %s\n", yellow, program, nc, arg)

	toRun := fmt.Sprintf("%s %s", program, arg)
	if env != "" {
		toRun = fmt.Sprintf("%s %s", env, toRun)
	}

	run(toRun)

	fileName := fmt.Sprintf("%s/screen-%s-%s%s.png", assetExamplesDir, program, strings.ReplaceAll(arg, " ", "_"), suffix)
	remove(fileName)
	run(screenshotCommand, fileName)
}

func run(format string, a ...interface{}) {
	rawCmd := fmt.Sprintf(format, a...)

	envs, args := lo.Must2(shellwords.ParseWithEnvs(rawCmd))

	c := exec.Command(args[0], args[1:]...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Env = append(os.Environ(), envs...)

	lo.Must0(c.Run())
}
