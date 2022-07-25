package main

import "github.com/Delta456/box-cli-maker/v2"

func main() {
	Box := box.New(box.Config{Px: 0, Py: 0, Type: "Round", Color: "Cyan", ContentAlign: "Left"})
	Box.TitlePos = "Top"
	Box.Print("gimme", `
  Version             1.42.0
  Git Commit          6809811
  Build Date          25 Jul 22 22:55 CEST(13 minutes ago)
  Commit Date         23 Jul 22 14:03 CEST(2 days ago)
  Dirty Build         yes
  Go Version          1.18.2
  Compiler            gc
  Platform            darwin/amd64
`)
}
