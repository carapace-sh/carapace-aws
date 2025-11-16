package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_batchGetDeploymentInstancesCmd = &cobra.Command{
	Use:   "batch-get-deployment-instances",
	Short: "This method works, but is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_batchGetDeploymentInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_batchGetDeploymentInstancesCmd).Standalone()

		codedeploy_batchGetDeploymentInstancesCmd.Flags().String("deployment-id", "", "The unique ID of a deployment.")
		codedeploy_batchGetDeploymentInstancesCmd.Flags().String("instance-ids", "", "The unique IDs of instances used in the deployment.")
		codedeploy_batchGetDeploymentInstancesCmd.MarkFlagRequired("deployment-id")
		codedeploy_batchGetDeploymentInstancesCmd.MarkFlagRequired("instance-ids")
	})
	codedeployCmd.AddCommand(codedeploy_batchGetDeploymentInstancesCmd)
}
