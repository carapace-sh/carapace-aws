package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listStudioLifecycleConfigsCmd = &cobra.Command{
	Use:   "list-studio-lifecycle-configs",
	Short: "Lists the Amazon SageMaker AI Studio Lifecycle Configurations in your Amazon Web Services Account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listStudioLifecycleConfigsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listStudioLifecycleConfigsCmd).Standalone()

		sagemaker_listStudioLifecycleConfigsCmd.Flags().String("app-type-equals", "", "A parameter to search for the App Type to which the Lifecycle Configuration is attached.")
		sagemaker_listStudioLifecycleConfigsCmd.Flags().String("creation-time-after", "", "A filter that returns only Lifecycle Configurations created on or after the specified time.")
		sagemaker_listStudioLifecycleConfigsCmd.Flags().String("creation-time-before", "", "A filter that returns only Lifecycle Configurations created on or before the specified time.")
		sagemaker_listStudioLifecycleConfigsCmd.Flags().String("max-results", "", "The total number of items to return in the response.")
		sagemaker_listStudioLifecycleConfigsCmd.Flags().String("modified-time-after", "", "A filter that returns only Lifecycle Configurations modified after the specified time.")
		sagemaker_listStudioLifecycleConfigsCmd.Flags().String("modified-time-before", "", "A filter that returns only Lifecycle Configurations modified before the specified time.")
		sagemaker_listStudioLifecycleConfigsCmd.Flags().String("name-contains", "", "A string in the Lifecycle Configuration name.")
		sagemaker_listStudioLifecycleConfigsCmd.Flags().String("next-token", "", "If the previous call to ListStudioLifecycleConfigs didn't return the full set of Lifecycle Configurations, the call returns a token for getting the next set of Lifecycle Configurations.")
		sagemaker_listStudioLifecycleConfigsCmd.Flags().String("sort-by", "", "The property used to sort results.")
		sagemaker_listStudioLifecycleConfigsCmd.Flags().String("sort-order", "", "The sort order.")
	})
	sagemakerCmd.AddCommand(sagemaker_listStudioLifecycleConfigsCmd)
}
