package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_updateKxEnvironmentCmd = &cobra.Command{
	Use:   "update-kx-environment",
	Short: "Updates information for the given kdb environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_updateKxEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_updateKxEnvironmentCmd).Standalone()

		finspace_updateKxEnvironmentCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspace_updateKxEnvironmentCmd.Flags().String("description", "", "A description of the kdb environment.")
		finspace_updateKxEnvironmentCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
		finspace_updateKxEnvironmentCmd.Flags().String("name", "", "The name of the kdb environment.")
		finspace_updateKxEnvironmentCmd.MarkFlagRequired("environment-id")
	})
	finspaceCmd.AddCommand(finspace_updateKxEnvironmentCmd)
}
