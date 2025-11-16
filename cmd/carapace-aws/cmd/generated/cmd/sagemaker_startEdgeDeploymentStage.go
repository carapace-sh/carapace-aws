package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_startEdgeDeploymentStageCmd = &cobra.Command{
	Use:   "start-edge-deployment-stage",
	Short: "Starts a stage in an edge deployment plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_startEdgeDeploymentStageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_startEdgeDeploymentStageCmd).Standalone()

		sagemaker_startEdgeDeploymentStageCmd.Flags().String("edge-deployment-plan-name", "", "The name of the edge deployment plan to start.")
		sagemaker_startEdgeDeploymentStageCmd.Flags().String("stage-name", "", "The name of the stage to start.")
		sagemaker_startEdgeDeploymentStageCmd.MarkFlagRequired("edge-deployment-plan-name")
		sagemaker_startEdgeDeploymentStageCmd.MarkFlagRequired("stage-name")
	})
	sagemakerCmd.AddCommand(sagemaker_startEdgeDeploymentStageCmd)
}
