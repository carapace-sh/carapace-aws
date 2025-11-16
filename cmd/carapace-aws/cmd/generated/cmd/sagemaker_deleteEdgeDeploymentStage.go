package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteEdgeDeploymentStageCmd = &cobra.Command{
	Use:   "delete-edge-deployment-stage",
	Short: "Delete a stage in an edge deployment plan if (and only if) the stage is inactive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteEdgeDeploymentStageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteEdgeDeploymentStageCmd).Standalone()

		sagemaker_deleteEdgeDeploymentStageCmd.Flags().String("edge-deployment-plan-name", "", "The name of the edge deployment plan from which the stage will be deleted.")
		sagemaker_deleteEdgeDeploymentStageCmd.Flags().String("stage-name", "", "The name of the stage.")
		sagemaker_deleteEdgeDeploymentStageCmd.MarkFlagRequired("edge-deployment-plan-name")
		sagemaker_deleteEdgeDeploymentStageCmd.MarkFlagRequired("stage-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteEdgeDeploymentStageCmd)
}
