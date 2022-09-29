package main

import (
	"github.com/leoay/luna/cmd/luna/v2/internal/project"
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
	rootCmd.AddCommand(upgrade.CmdUpgrade)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
