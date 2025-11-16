package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createEdgeDeploymentPlanCmd = &cobra.Command{
	Use:   "create-edge-deployment-plan",
	Short: "Creates an edge deployment plan, consisting of multiple stages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createEdgeDeploymentPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createEdgeDeploymentPlanCmd).Standalone()

		sagemaker_createEdgeDeploymentPlanCmd.Flags().String("device-fleet-name", "", "The device fleet used for this edge deployment plan.")
		sagemaker_createEdgeDeploymentPlanCmd.Flags().String("edge-deployment-plan-name", "", "The name of the edge deployment plan.")
		sagemaker_createEdgeDeploymentPlanCmd.Flags().String("model-configs", "", "List of models associated with the edge deployment plan.")
		sagemaker_createEdgeDeploymentPlanCmd.Flags().String("stages", "", "List of stages of the edge deployment plan.")
		sagemaker_createEdgeDeploymentPlanCmd.Flags().String("tags", "", "List of tags with which to tag the edge deployment plan.")
		sagemaker_createEdgeDeploymentPlanCmd.MarkFlagRequired("device-fleet-name")
		sagemaker_createEdgeDeploymentPlanCmd.MarkFlagRequired("edge-deployment-plan-name")
		sagemaker_createEdgeDeploymentPlanCmd.MarkFlagRequired("model-configs")
	})
	sagemakerCmd.AddCommand(sagemaker_createEdgeDeploymentPlanCmd)
}
