package main

import (
	"github.com/spf13/cobra"
	"log"
	"server01.jz/yaoshuai_dev/luna/cmd/luna/v2/internal/project"
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
