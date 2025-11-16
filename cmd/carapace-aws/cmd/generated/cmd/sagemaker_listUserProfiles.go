package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listUserProfilesCmd = &cobra.Command{
	Use:   "list-user-profiles",
	Short: "Lists user profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listUserProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listUserProfilesCmd).Standalone()

		sagemaker_listUserProfilesCmd.Flags().String("domain-id-equals", "", "A parameter by which to filter the results.")
		sagemaker_listUserProfilesCmd.Flags().String("max-results", "", "This parameter defines the maximum number of results that can be return in a single response.")
		sagemaker_listUserProfilesCmd.Flags().String("next-token", "", "If the previous response was truncated, you will receive this token.")
		sagemaker_listUserProfilesCmd.Flags().String("sort-by", "", "The parameter by which to sort the results.")
		sagemaker_listUserProfilesCmd.Flags().String("sort-order", "", "The sort order for the results.")
		sagemaker_listUserProfilesCmd.Flags().String("user-profile-name-contains", "", "A parameter by which to filter the results.")
	})
	sagemakerCmd.AddCommand(sagemaker_listUserProfilesCmd)
}
