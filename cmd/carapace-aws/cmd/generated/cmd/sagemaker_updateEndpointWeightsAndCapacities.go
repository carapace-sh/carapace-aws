package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateEndpointWeightsAndCapacitiesCmd = &cobra.Command{
	Use:   "update-endpoint-weights-and-capacities",
	Short: "Updates variant weight of one or more variants associated with an existing endpoint, or capacity of one variant associated with an existing endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateEndpointWeightsAndCapacitiesCmd).Standalone()

	sagemaker_updateEndpointWeightsAndCapacitiesCmd.Flags().String("desired-weights-and-capacities", "", "An object that provides new capacity and weight values for a variant.")
	sagemaker_updateEndpointWeightsAndCapacitiesCmd.Flags().String("endpoint-name", "", "The name of an existing SageMaker endpoint.")
	sagemaker_updateEndpointWeightsAndCapacitiesCmd.MarkFlagRequired("desired-weights-and-capacities")
	sagemaker_updateEndpointWeightsAndCapacitiesCmd.MarkFlagRequired("endpoint-name")
	sagemakerCmd.AddCommand(sagemaker_updateEndpointWeightsAndCapacitiesCmd)
}
