package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_stopDeploymentCmd = &cobra.Command{
	Use:   "stop-deployment",
	Short: "Attempts to stop an ongoing deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_stopDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_stopDeploymentCmd).Standalone()

		codedeploy_stopDeploymentCmd.Flags().String("auto-rollback-enabled", "", "Indicates, when a deployment is stopped, whether instances that have been updated should be rolled back to the previous version of the application revision.")
		codedeploy_stopDeploymentCmd.Flags().String("deployment-id", "", "The unique ID of a deployment.")
		codedeploy_stopDeploymentCmd.MarkFlagRequired("deployment-id")
	})
	codedeployCmd.AddCommand(codedeploy_stopDeploymentCmd)
}
