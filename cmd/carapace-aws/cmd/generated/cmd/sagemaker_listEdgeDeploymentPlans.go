package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listEdgeDeploymentPlansCmd = &cobra.Command{
	Use:   "list-edge-deployment-plans",
	Short: "Lists all edge deployment plans.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listEdgeDeploymentPlansCmd).Standalone()

	sagemaker_listEdgeDeploymentPlansCmd.Flags().String("creation-time-after", "", "Selects edge deployment plans created after this time.")
	sagemaker_listEdgeDeploymentPlansCmd.Flags().String("creation-time-before", "", "Selects edge deployment plans created before this time.")
	sagemaker_listEdgeDeploymentPlansCmd.Flags().String("device-fleet-name-contains", "", "Selects edge deployment plans with a device fleet name containing this name.")
	sagemaker_listEdgeDeploymentPlansCmd.Flags().String("last-modified-time-after", "", "Selects edge deployment plans that were last updated after this time.")
	sagemaker_listEdgeDeploymentPlansCmd.Flags().String("last-modified-time-before", "", "Selects edge deployment plans that were last updated before this time.")
	sagemaker_listEdgeDeploymentPlansCmd.Flags().String("max-results", "", "The maximum number of results to select (50 by default).")
	sagemaker_listEdgeDeploymentPlansCmd.Flags().String("name-contains", "", "Selects edge deployment plans with names containing this name.")
	sagemaker_listEdgeDeploymentPlansCmd.Flags().String("next-token", "", "The response from the last list when returning a list large enough to need tokening.")
	sagemaker_listEdgeDeploymentPlansCmd.Flags().String("sort-by", "", "The column by which to sort the edge deployment plans.")
	sagemaker_listEdgeDeploymentPlansCmd.Flags().String("sort-order", "", "The direction of the sorting (ascending or descending).")
	sagemakerCmd.AddCommand(sagemaker_listEdgeDeploymentPlansCmd)
}
