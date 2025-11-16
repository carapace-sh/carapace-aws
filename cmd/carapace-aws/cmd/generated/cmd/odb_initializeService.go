package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_initializeServiceCmd = &cobra.Command{
	Use:   "initialize-service",
	Short: "Initializes the ODB service for the first time in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_initializeServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_initializeServiceCmd).Standalone()

	})
	odbCmd.AddCommand(odb_initializeServiceCmd)
}
