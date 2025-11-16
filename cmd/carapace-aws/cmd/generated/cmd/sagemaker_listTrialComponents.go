package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listTrialComponentsCmd = &cobra.Command{
	Use:   "list-trial-components",
	Short: "Lists the trial components in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listTrialComponentsCmd).Standalone()

	sagemaker_listTrialComponentsCmd.Flags().String("created-after", "", "A filter that returns only components created after the specified time.")
	sagemaker_listTrialComponentsCmd.Flags().String("created-before", "", "A filter that returns only components created before the specified time.")
	sagemaker_listTrialComponentsCmd.Flags().String("experiment-name", "", "A filter that returns only components that are part of the specified experiment.")
	sagemaker_listTrialComponentsCmd.Flags().String("max-results", "", "The maximum number of components to return in the response.")
	sagemaker_listTrialComponentsCmd.Flags().String("next-token", "", "If the previous call to `ListTrialComponents` didn't return the full set of components, the call returns a token for getting the next set of components.")
	sagemaker_listTrialComponentsCmd.Flags().String("sort-by", "", "The property used to sort results.")
	sagemaker_listTrialComponentsCmd.Flags().String("sort-order", "", "The sort order.")
	sagemaker_listTrialComponentsCmd.Flags().String("source-arn", "", "A filter that returns only components that have the specified source Amazon Resource Name (ARN).")
	sagemaker_listTrialComponentsCmd.Flags().String("trial-name", "", "A filter that returns only components that are part of the specified trial.")
	sagemakerCmd.AddCommand(sagemaker_listTrialComponentsCmd)
}
