//go:build e2e

package integration

import (
	"bytes"
	"context"
	"io"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/Netflix/go-expect"
	"github.com/mattn/go-shellwords"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.szostok.io/magex/shx"
)

const exampleDir = "../../examples"

func buildBinaryAllLDFlags(t *testing.T, dir string) string {
	t.Helper()

	var (
		buildDate  = time.Date(2022, time.April, 1, 12, 22, 14, 0, time.UTC).Format("2006-01-02T15:04:05Z0700")
		commitDate = time.Date(2022, time.March, 28, 15, 32, 14, 0, time.UTC).Format("2006-01-02T15:04:05Z0700")
		commit     = "324d022c190ce49e0440e6bdac6383e4874c7c70"
		dirtyBuild = "false"
		binary     = "example"
	)

	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	ran, code, err := shx.MustCmdf(`go build -ldflags="-X go.szostok.io/version.version=0.6.1 -X 'go.szostok.io/version.buildDate=%s' -X go.szostok.io/version.commit=%s -X go.szostok.io/version.commitDate=%s -X go.szostok.io/version.dirtyBuild=%s -X go.szostok.io/version.name=%s" -o %s . `,
		buildDate,
		commit,
		commitDate,
		dirtyBuild,
		binary,
		binary,
	).
		In(filepath.Join(exampleDir, dir)).
		Exec()

	require.NoError(t, err)
	assert.True(t, ran)
	assert.Equal(t, 0, code)
	return filepath.Join(exampleDir, dir, binary)
}

type Executor struct {
	binaryPath string
	cmd        string
}

func Exec(binary, cmd string) *Executor {
	return &Executor{
		binaryPath: binary,
		cmd:        cmd,
	}
}

type ExecuteOutput struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

func (s *Executor) AwaitResultAtMost(timeout time.Duration) (*ExecuteOutput, error) {
	var buffOut, buffErr bytes.Buffer

	exitCode, err := s.execute(timeout, &buffOut, &buffErr)
	if err != nil {
		return nil, err
	}

	return &ExecuteOutput{
		ExitCode: exitCode,
		Stdout:   buffOut.String(),
		Stderr:   buffErr.String(),
	}, nil
}

func (s *Executor) AwaitColorResultAtMost(timeout time.Duration) (*ExecuteOutput, error) {
	wait := make(chan struct{})
	c, err := expect.NewConsole()
	if err != nil {
		return nil, err
	}

	var (
		out    string
		outErr error
	)

	go func() {
		out, outErr = c.ExpectEOF()
		close(wait)
	}()

	exitCode, err := s.execute(timeout, c.Tty(), c.Tty())
	if err != nil {
		return nil, errors.Wrap(err, "while executing binary")
	}

	if err := c.Close(); err != nil {
		return nil, err
	}
	<-wait

	if outErr != nil {
		return nil, outErr
	}

	return &ExecuteOutput{
		ExitCode: exitCode,
		Stdout:   out,
	}, nil
}
func (s *Executor) execute(timeout time.Duration, stdout, stderr io.Writer) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	args, err := shellwords.Parse(s.cmd)
	if err != nil {
		return 0, err
	}

	cmd := exec.CommandContext(ctx, s.binaryPath, args...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err = cmd.Run()
	if err != nil {
		return 0, err
	}
	exitCode := cmd.ProcessState.ExitCode()

	return exitCode, nil
}
