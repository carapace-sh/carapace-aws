package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteEdgeDeploymentPlanCmd = &cobra.Command{
	Use:   "delete-edge-deployment-plan",
	Short: "Deletes an edge deployment plan if (and only if) all the stages in the plan are inactive or there are no stages in the plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteEdgeDeploymentPlanCmd).Standalone()

	sagemaker_deleteEdgeDeploymentPlanCmd.Flags().String("edge-deployment-plan-name", "", "The name of the edge deployment plan to delete.")
	sagemaker_deleteEdgeDeploymentPlanCmd.MarkFlagRequired("edge-deployment-plan-name")
	sagemakerCmd.AddCommand(sagemaker_deleteEdgeDeploymentPlanCmd)
}
