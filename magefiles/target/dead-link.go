package target

import (
	"fmt"
	"strings"
	"time"

	"go.szostok.io/magex/printer"
	"go.szostok.io/magex/shx"
)

var excludedLinks = []string{
	"https://github.com/mszostok/version.*", // repo is private for now
	".*/projects/version/static/.*",
	"https://linkedin.com/in/mszostok",
	"https://fonts.gstatic.com",
	"https://twitter.com/m_szostok",
}

func CheckDeadLinks() error {
	printer.Title("Checking for dead links in docs...")

	mkdocsSvr := shx.MustAsyncCmdf("mkdocs serve -a localhost:60123")
	mkdocsSvr.MustStart()
	defer mkdocsSvr.MustStop()

	time.Sleep(time.Second) // mkdocs needs some time

	return shx.MustCmdf("./bin/muffet http://localhost:60123 --skip-tls-verification --rate-limit=50 %s -v", exclude(excludedLinks)).Run()
}

func exclude(in []string) string {
	var buff strings.Builder
	for _, name := range in {
		buff.WriteString(fmt.Sprintf("--exclude='%s' ", name))
	}
	return buff.String()
}
