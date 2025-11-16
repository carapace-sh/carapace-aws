package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_getKxEnvironmentCmd = &cobra.Command{
	Use:   "get-kx-environment",
	Short: "Retrieves all the information for the specified kdb environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_getKxEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_getKxEnvironmentCmd).Standalone()

		finspace_getKxEnvironmentCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
		finspace_getKxEnvironmentCmd.MarkFlagRequired("environment-id")
	})
	finspaceCmd.AddCommand(finspace_getKxEnvironmentCmd)
}
