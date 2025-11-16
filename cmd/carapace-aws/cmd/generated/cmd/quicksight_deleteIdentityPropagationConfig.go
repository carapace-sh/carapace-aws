package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteIdentityPropagationConfigCmd = &cobra.Command{
	Use:   "delete-identity-propagation-config",
	Short: "Deletes all access scopes and authorized targets that are associated with a service from the Quick Sight IAM Identity Center application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteIdentityPropagationConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteIdentityPropagationConfigCmd).Standalone()

		quicksight_deleteIdentityPropagationConfigCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that you want to delete an identity propagation configuration from.")
		quicksight_deleteIdentityPropagationConfigCmd.Flags().String("service", "", "The name of the Amazon Web Services service that you want to delete the associated access scopes and authorized targets from.")
		quicksight_deleteIdentityPropagationConfigCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteIdentityPropagationConfigCmd.MarkFlagRequired("service")
	})
	quicksightCmd.AddCommand(quicksight_deleteIdentityPropagationConfigCmd)
}
