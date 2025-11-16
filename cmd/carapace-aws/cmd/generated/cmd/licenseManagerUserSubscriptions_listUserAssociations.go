package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_listUserAssociationsCmd = &cobra.Command{
	Use:   "list-user-associations",
	Short: "Lists user associations for an identity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_listUserAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManagerUserSubscriptions_listUserAssociationsCmd).Standalone()

		licenseManagerUserSubscriptions_listUserAssociationsCmd.Flags().String("filters", "", "You can use the following filters to streamline results:")
		licenseManagerUserSubscriptions_listUserAssociationsCmd.Flags().String("identity-provider", "", "An object that specifies details for the identity provider.")
		licenseManagerUserSubscriptions_listUserAssociationsCmd.Flags().String("instance-id", "", "The ID of the EC2 instance, which provides user-based subscriptions.")
		licenseManagerUserSubscriptions_listUserAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return from a single request.")
		licenseManagerUserSubscriptions_listUserAssociationsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
		licenseManagerUserSubscriptions_listUserAssociationsCmd.MarkFlagRequired("identity-provider")
		licenseManagerUserSubscriptions_listUserAssociationsCmd.MarkFlagRequired("instance-id")
	})
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_listUserAssociationsCmd)
}
