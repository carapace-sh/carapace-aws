package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createEdgeDeploymentStageCmd = &cobra.Command{
	Use:   "create-edge-deployment-stage",
	Short: "Creates a new stage in an existing edge deployment plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createEdgeDeploymentStageCmd).Standalone()

	sagemaker_createEdgeDeploymentStageCmd.Flags().String("edge-deployment-plan-name", "", "The name of the edge deployment plan.")
	sagemaker_createEdgeDeploymentStageCmd.Flags().String("stages", "", "List of stages to be added to the edge deployment plan.")
	sagemaker_createEdgeDeploymentStageCmd.MarkFlagRequired("edge-deployment-plan-name")
	sagemaker_createEdgeDeploymentStageCmd.MarkFlagRequired("stages")
	sagemakerCmd.AddCommand(sagemaker_createEdgeDeploymentStageCmd)
}
