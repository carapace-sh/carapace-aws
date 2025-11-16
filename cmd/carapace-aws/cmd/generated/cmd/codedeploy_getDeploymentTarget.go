package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_getDeploymentTargetCmd = &cobra.Command{
	Use:   "get-deployment-target",
	Short: "Returns information about a deployment target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_getDeploymentTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_getDeploymentTargetCmd).Standalone()

		codedeploy_getDeploymentTargetCmd.Flags().String("deployment-id", "", "The unique ID of a deployment.")
		codedeploy_getDeploymentTargetCmd.Flags().String("target-id", "", "The unique ID of a deployment target.")
		codedeploy_getDeploymentTargetCmd.MarkFlagRequired("deployment-id")
		codedeploy_getDeploymentTargetCmd.MarkFlagRequired("target-id")
	})
	codedeployCmd.AddCommand(codedeploy_getDeploymentTargetCmd)
}
