package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listTrialsCmd = &cobra.Command{
	Use:   "list-trials",
	Short: "Lists the trials in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listTrialsCmd).Standalone()

	sagemaker_listTrialsCmd.Flags().String("created-after", "", "A filter that returns only trials created after the specified time.")
	sagemaker_listTrialsCmd.Flags().String("created-before", "", "A filter that returns only trials created before the specified time.")
	sagemaker_listTrialsCmd.Flags().String("experiment-name", "", "A filter that returns only trials that are part of the specified experiment.")
	sagemaker_listTrialsCmd.Flags().String("max-results", "", "The maximum number of trials to return in the response.")
	sagemaker_listTrialsCmd.Flags().String("next-token", "", "If the previous call to `ListTrials` didn't return the full set of trials, the call returns a token for getting the next set of trials.")
	sagemaker_listTrialsCmd.Flags().String("sort-by", "", "The property used to sort results.")
	sagemaker_listTrialsCmd.Flags().String("sort-order", "", "The sort order.")
	sagemaker_listTrialsCmd.Flags().String("trial-component-name", "", "A filter that returns only trials that are associated with the specified trial component.")
	sagemakerCmd.AddCommand(sagemaker_listTrialsCmd)
}
