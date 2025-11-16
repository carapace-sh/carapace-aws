package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_getEnvironmentCmd = &cobra.Command{
	Use:   "get-environment",
	Short: "Returns the FinSpace environment object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_getEnvironmentCmd).Standalone()

	finspace_getEnvironmentCmd.Flags().String("environment-id", "", "The identifier of the FinSpace environment.")
	finspace_getEnvironmentCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_getEnvironmentCmd)
}
