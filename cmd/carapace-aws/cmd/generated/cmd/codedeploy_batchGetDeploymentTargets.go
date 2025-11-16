package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_batchGetDeploymentTargetsCmd = &cobra.Command{
	Use:   "batch-get-deployment-targets",
	Short: "Returns an array of one or more targets associated with a deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_batchGetDeploymentTargetsCmd).Standalone()

	codedeploy_batchGetDeploymentTargetsCmd.Flags().String("deployment-id", "", "The unique ID of a deployment.")
	codedeploy_batchGetDeploymentTargetsCmd.Flags().String("target-ids", "", "The unique IDs of the deployment targets.")
	codedeploy_batchGetDeploymentTargetsCmd.MarkFlagRequired("deployment-id")
	codedeploy_batchGetDeploymentTargetsCmd.MarkFlagRequired("target-ids")
	codedeployCmd.AddCommand(codedeploy_batchGetDeploymentTargetsCmd)
}
