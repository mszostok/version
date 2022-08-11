package main

import (
	"fmt"

	"go.szostok.io/version"
)

func main() {
	info := version.Get()
	fmt.Println()
	fmt.Println("Version: ", info.Version)
	fmt.Println("Git Commit: ", info.GitCommit)
	fmt.Println("Build Date: ", info.BuildDate)
	fmt.Println("Commit Date: ", info.CommitDate)
	fmt.Println("Dirty Build: ", info.DirtyBuild)
	fmt.Println("Go Version: ", info.GoVersion)
	fmt.Println("Compiler: ", info.Compiler)
	fmt.Println("Platform: ", info.Platform)
}
