package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_associateUserCmd = &cobra.Command{
	Use:   "associate-user",
	Short: "Associates the user to an EC2 instance to utilize user-based subscriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_associateUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManagerUserSubscriptions_associateUserCmd).Standalone()

		licenseManagerUserSubscriptions_associateUserCmd.Flags().String("domain", "", "The domain name of the Active Directory that contains information for the user to associate.")
		licenseManagerUserSubscriptions_associateUserCmd.Flags().String("identity-provider", "", "The identity provider for the user.")
		licenseManagerUserSubscriptions_associateUserCmd.Flags().String("instance-id", "", "The ID of the EC2 instance that provides the user-based subscription.")
		licenseManagerUserSubscriptions_associateUserCmd.Flags().String("tags", "", "The tags that apply for the user association.")
		licenseManagerUserSubscriptions_associateUserCmd.Flags().String("username", "", "The user name from the identity provider.")
		licenseManagerUserSubscriptions_associateUserCmd.MarkFlagRequired("identity-provider")
		licenseManagerUserSubscriptions_associateUserCmd.MarkFlagRequired("instance-id")
		licenseManagerUserSubscriptions_associateUserCmd.MarkFlagRequired("username")
	})
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_associateUserCmd)
}
