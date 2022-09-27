package main

import (
	"github.com/leoay/Luna/cmd/luna/v2/internal/project"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:     "luna",
	Short:   "Luna: An elegant toolkit for Go microservices.",
	Long:    `Luna: An elegant toolkit for Go microservices.`,
	Version: release,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	//rootCmd.AddCommand(upgrade.CmdUpgrade)
	//rootCmd.AddCommand(change.CmdChange)
	//rootCmd.AddCommand(run.CmdRun)
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
