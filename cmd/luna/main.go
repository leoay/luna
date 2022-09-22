package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "luna",
	Short:   "Kratos: An elegant toolkit for Go microservices.",
	Long:    `Kratos: An elegant toolkit for Go microservices.`,
	Version: release,
}

func main() {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "my test program",
	}
}
