package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_cancelDeploymentCmd = &cobra.Command{
	Use:   "cancel-deployment",
	Short: "Cancels a deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_cancelDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_cancelDeploymentCmd).Standalone()

		greengrassv2_cancelDeploymentCmd.Flags().String("deployment-id", "", "The ID of the deployment.")
		greengrassv2_cancelDeploymentCmd.MarkFlagRequired("deployment-id")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_cancelDeploymentCmd)
}
