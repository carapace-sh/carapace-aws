package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_getDeploymentInstanceCmd = &cobra.Command{
	Use:   "get-deployment-instance",
	Short: "Gets information about an instance as part of a deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_getDeploymentInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_getDeploymentInstanceCmd).Standalone()

		codedeploy_getDeploymentInstanceCmd.Flags().String("deployment-id", "", "The unique ID of a deployment.")
		codedeploy_getDeploymentInstanceCmd.Flags().String("instance-id", "", "The unique ID of an instance in the deployment group.")
		codedeploy_getDeploymentInstanceCmd.MarkFlagRequired("deployment-id")
		codedeploy_getDeploymentInstanceCmd.MarkFlagRequired("instance-id")
	})
	codedeployCmd.AddCommand(codedeploy_getDeploymentInstanceCmd)
}
