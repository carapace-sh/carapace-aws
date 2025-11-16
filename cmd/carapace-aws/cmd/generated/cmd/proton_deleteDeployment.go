package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_deleteDeploymentCmd = &cobra.Command{
	Use:   "delete-deployment",
	Short: "Delete the deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_deleteDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_deleteDeploymentCmd).Standalone()

		proton_deleteDeploymentCmd.Flags().String("id", "", "The ID of the deployment to delete.")
		proton_deleteDeploymentCmd.MarkFlagRequired("id")
	})
	protonCmd.AddCommand(proton_deleteDeploymentCmd)
}
