package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeEdgeDeploymentPlanCmd = &cobra.Command{
	Use:   "describe-edge-deployment-plan",
	Short: "Describes an edge deployment plan with deployment status per stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeEdgeDeploymentPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeEdgeDeploymentPlanCmd).Standalone()

		sagemaker_describeEdgeDeploymentPlanCmd.Flags().String("edge-deployment-plan-name", "", "The name of the deployment plan to describe.")
		sagemaker_describeEdgeDeploymentPlanCmd.Flags().String("max-results", "", "The maximum number of results to select (50 by default).")
		sagemaker_describeEdgeDeploymentPlanCmd.Flags().String("next-token", "", "If the edge deployment plan has enough stages to require tokening, then this is the response from the last list of stages returned.")
		sagemaker_describeEdgeDeploymentPlanCmd.MarkFlagRequired("edge-deployment-plan-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeEdgeDeploymentPlanCmd)
}
