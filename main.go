package main

import (
	"os"
	awscaller "tabakazu/awscaller/lib"
)

func main() {
	awscaller.CliApp().Run(os.Args)
}
