package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_updateEnvironmentCmd = &cobra.Command{
	Use:   "update-environment",
	Short: "Update your FinSpace environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_updateEnvironmentCmd).Standalone()

	finspace_updateEnvironmentCmd.Flags().String("description", "", "The description of the environment.")
	finspace_updateEnvironmentCmd.Flags().String("environment-id", "", "The identifier of the FinSpace environment.")
	finspace_updateEnvironmentCmd.Flags().String("federation-mode", "", "Authentication mode for the environment.")
	finspace_updateEnvironmentCmd.Flags().String("federation-parameters", "", "")
	finspace_updateEnvironmentCmd.Flags().String("name", "", "The name of the environment.")
	finspace_updateEnvironmentCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_updateEnvironmentCmd)
}
