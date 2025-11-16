package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopEdgeDeploymentStageCmd = &cobra.Command{
	Use:   "stop-edge-deployment-stage",
	Short: "Stops a stage in an edge deployment plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopEdgeDeploymentStageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_stopEdgeDeploymentStageCmd).Standalone()

		sagemaker_stopEdgeDeploymentStageCmd.Flags().String("edge-deployment-plan-name", "", "The name of the edge deployment plan to stop.")
		sagemaker_stopEdgeDeploymentStageCmd.Flags().String("stage-name", "", "The name of the stage to stop.")
		sagemaker_stopEdgeDeploymentStageCmd.MarkFlagRequired("edge-deployment-plan-name")
		sagemaker_stopEdgeDeploymentStageCmd.MarkFlagRequired("stage-name")
	})
	sagemakerCmd.AddCommand(sagemaker_stopEdgeDeploymentStageCmd)
}
