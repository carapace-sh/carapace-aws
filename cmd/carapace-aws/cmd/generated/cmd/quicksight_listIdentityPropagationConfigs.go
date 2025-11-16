package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listIdentityPropagationConfigsCmd = &cobra.Command{
	Use:   "list-identity-propagation-configs",
	Short: "Lists all services and authorized targets that the Quick Sight IAM Identity Center application can access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listIdentityPropagationConfigsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listIdentityPropagationConfigsCmd).Standalone()

		quicksight_listIdentityPropagationConfigsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contain the identity propagation configurations of.")
		quicksight_listIdentityPropagationConfigsCmd.Flags().String("max-results", "", "The maximum number of results to be returned.")
		quicksight_listIdentityPropagationConfigsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
		quicksight_listIdentityPropagationConfigsCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_listIdentityPropagationConfigsCmd)
}
