package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listAppsCmd = &cobra.Command{
	Use:   "list-apps",
	Short: "Lists apps.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listAppsCmd).Standalone()

	sagemaker_listAppsCmd.Flags().String("domain-id-equals", "", "A parameter to search for the domain ID.")
	sagemaker_listAppsCmd.Flags().String("max-results", "", "This parameter defines the maximum number of results that can be return in a single response.")
	sagemaker_listAppsCmd.Flags().String("next-token", "", "If the previous response was truncated, you will receive this token.")
	sagemaker_listAppsCmd.Flags().String("sort-by", "", "The parameter by which to sort the results.")
	sagemaker_listAppsCmd.Flags().String("sort-order", "", "The sort order for the results.")
	sagemaker_listAppsCmd.Flags().String("space-name-equals", "", "A parameter to search by space name.")
	sagemaker_listAppsCmd.Flags().String("user-profile-name-equals", "", "A parameter to search by user profile name.")
	sagemakerCmd.AddCommand(sagemaker_listAppsCmd)
}
