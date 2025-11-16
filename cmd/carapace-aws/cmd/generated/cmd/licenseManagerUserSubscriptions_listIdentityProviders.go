package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_listIdentityProvidersCmd = &cobra.Command{
	Use:   "list-identity-providers",
	Short: "Lists the Active Directory identity providers for user-based subscriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_listIdentityProvidersCmd).Standalone()

	licenseManagerUserSubscriptions_listIdentityProvidersCmd.Flags().String("filters", "", "You can use the following filters to streamline results:")
	licenseManagerUserSubscriptions_listIdentityProvidersCmd.Flags().String("max-results", "", "The maximum number of results to return from a single request.")
	licenseManagerUserSubscriptions_listIdentityProvidersCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_listIdentityProvidersCmd)
}
