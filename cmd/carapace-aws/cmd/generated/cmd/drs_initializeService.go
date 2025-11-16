package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_initializeServiceCmd = &cobra.Command{
	Use:   "initialize-service",
	Short: "Initialize Elastic Disaster Recovery.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_initializeServiceCmd).Standalone()

	drsCmd.AddCommand(drs_initializeServiceCmd)
}
