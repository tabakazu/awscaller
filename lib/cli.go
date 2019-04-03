package awscaller

import (
	"github.com/urfave/cli"
)

const appName string = "awscaller"
const appUsage string = "Shows aws api caller infomations"
const appVersion string = "0.0.3"

// CliApp return *cli.App
func CliApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion
	app.Action = actionMain
	app.Commands = []cli.Command{
		actionAccount(),
		actionAttachedPolicies(),
		actionUserID(),
		actionUserName(),
	}
	return app
}
