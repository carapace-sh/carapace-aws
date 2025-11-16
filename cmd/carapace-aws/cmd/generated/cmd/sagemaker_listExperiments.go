package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listExperimentsCmd = &cobra.Command{
	Use:   "list-experiments",
	Short: "Lists all the experiments in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listExperimentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listExperimentsCmd).Standalone()

		sagemaker_listExperimentsCmd.Flags().String("created-after", "", "A filter that returns only experiments created after the specified time.")
		sagemaker_listExperimentsCmd.Flags().String("created-before", "", "A filter that returns only experiments created before the specified time.")
		sagemaker_listExperimentsCmd.Flags().String("max-results", "", "The maximum number of experiments to return in the response.")
		sagemaker_listExperimentsCmd.Flags().String("next-token", "", "If the previous call to `ListExperiments` didn't return the full set of experiments, the call returns a token for getting the next set of experiments.")
		sagemaker_listExperimentsCmd.Flags().String("sort-by", "", "The property used to sort results.")
		sagemaker_listExperimentsCmd.Flags().String("sort-order", "", "The sort order.")
	})
	sagemakerCmd.AddCommand(sagemaker_listExperimentsCmd)
}
