package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listAppImageConfigsCmd = &cobra.Command{
	Use:   "list-app-image-configs",
	Short: "Lists the AppImageConfigs in your account and their properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listAppImageConfigsCmd).Standalone()

	sagemaker_listAppImageConfigsCmd.Flags().String("creation-time-after", "", "A filter that returns only AppImageConfigs created on or after the specified time.")
	sagemaker_listAppImageConfigsCmd.Flags().String("creation-time-before", "", "A filter that returns only AppImageConfigs created on or before the specified time.")
	sagemaker_listAppImageConfigsCmd.Flags().String("max-results", "", "The total number of items to return in the response.")
	sagemaker_listAppImageConfigsCmd.Flags().String("modified-time-after", "", "A filter that returns only AppImageConfigs modified on or after the specified time.")
	sagemaker_listAppImageConfigsCmd.Flags().String("modified-time-before", "", "A filter that returns only AppImageConfigs modified on or before the specified time.")
	sagemaker_listAppImageConfigsCmd.Flags().String("name-contains", "", "A filter that returns only AppImageConfigs whose name contains the specified string.")
	sagemaker_listAppImageConfigsCmd.Flags().String("next-token", "", "If the previous call to `ListImages` didn't return the full set of AppImageConfigs, the call returns a token for getting the next set of AppImageConfigs.")
	sagemaker_listAppImageConfigsCmd.Flags().String("sort-by", "", "The property used to sort results.")
	sagemaker_listAppImageConfigsCmd.Flags().String("sort-order", "", "The sort order.")
	sagemakerCmd.AddCommand(sagemaker_listAppImageConfigsCmd)
}
