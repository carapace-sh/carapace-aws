package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateIdentityPropagationConfigCmd = &cobra.Command{
	Use:   "update-identity-propagation-config",
	Short: "Adds or updates services and authorized targets to configure what the Quick Sight IAM Identity Center application can access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateIdentityPropagationConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateIdentityPropagationConfigCmd).Standalone()

		quicksight_updateIdentityPropagationConfigCmd.Flags().String("authorized-targets", "", "Specifies a list of application ARNs that represent the authorized targets for a service.")
		quicksight_updateIdentityPropagationConfigCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the identity propagation configuration that you want to update.")
		quicksight_updateIdentityPropagationConfigCmd.Flags().String("service", "", "The name of the Amazon Web Services service that contains the authorized targets that you want to add or update.")
		quicksight_updateIdentityPropagationConfigCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateIdentityPropagationConfigCmd.MarkFlagRequired("service")
	})
	quicksightCmd.AddCommand(quicksight_updateIdentityPropagationConfigCmd)
}
