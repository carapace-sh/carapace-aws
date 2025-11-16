package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_deleteDeploymentCmd = &cobra.Command{
	Use:   "delete-deployment",
	Short: "Deletes a deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_deleteDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_deleteDeploymentCmd).Standalone()

		greengrassv2_deleteDeploymentCmd.Flags().String("deployment-id", "", "The ID of the deployment.")
		greengrassv2_deleteDeploymentCmd.MarkFlagRequired("deployment-id")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_deleteDeploymentCmd)
}
