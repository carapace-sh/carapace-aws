package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_deleteEnvironmentCmd = &cobra.Command{
	Use:   "delete-environment",
	Short: "Delete an FinSpace environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_deleteEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_deleteEnvironmentCmd).Standalone()

		finspace_deleteEnvironmentCmd.Flags().String("environment-id", "", "The identifier for the FinSpace environment.")
		finspace_deleteEnvironmentCmd.MarkFlagRequired("environment-id")
	})
	finspaceCmd.AddCommand(finspace_deleteEnvironmentCmd)
}
