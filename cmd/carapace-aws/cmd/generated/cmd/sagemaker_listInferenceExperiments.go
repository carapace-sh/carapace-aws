package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listInferenceExperimentsCmd = &cobra.Command{
	Use:   "list-inference-experiments",
	Short: "Returns the list of all inference experiments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listInferenceExperimentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listInferenceExperimentsCmd).Standalone()

		sagemaker_listInferenceExperimentsCmd.Flags().String("creation-time-after", "", "Selects inference experiments which were created after this timestamp.")
		sagemaker_listInferenceExperimentsCmd.Flags().String("creation-time-before", "", "Selects inference experiments which were created before this timestamp.")
		sagemaker_listInferenceExperimentsCmd.Flags().String("last-modified-time-after", "", "Selects inference experiments which were last modified after this timestamp.")
		sagemaker_listInferenceExperimentsCmd.Flags().String("last-modified-time-before", "", "Selects inference experiments which were last modified before this timestamp.")
		sagemaker_listInferenceExperimentsCmd.Flags().String("max-results", "", "The maximum number of results to select.")
		sagemaker_listInferenceExperimentsCmd.Flags().String("name-contains", "", "Selects inference experiments whose names contain this name.")
		sagemaker_listInferenceExperimentsCmd.Flags().String("next-token", "", "The response from the last list when returning a list large enough to need tokening.")
		sagemaker_listInferenceExperimentsCmd.Flags().String("sort-by", "", "The column by which to sort the listed inference experiments.")
		sagemaker_listInferenceExperimentsCmd.Flags().String("sort-order", "", "The direction of sorting (ascending or descending).")
		sagemaker_listInferenceExperimentsCmd.Flags().String("status-equals", "", "Selects inference experiments which are in this status.")
		sagemaker_listInferenceExperimentsCmd.Flags().String("type", "", "Selects inference experiments of this type.")
	})
	sagemakerCmd.AddCommand(sagemaker_listInferenceExperimentsCmd)
}
