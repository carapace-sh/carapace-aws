package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_disassociateUserCmd = &cobra.Command{
	Use:   "disassociate-user",
	Short: "Disassociates the user from an EC2 instance providing user-based subscriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_disassociateUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManagerUserSubscriptions_disassociateUserCmd).Standalone()

		licenseManagerUserSubscriptions_disassociateUserCmd.Flags().String("domain", "", "The domain name of the Active Directory that contains information for the user to disassociate.")
		licenseManagerUserSubscriptions_disassociateUserCmd.Flags().String("identity-provider", "", "An object that specifies details for the Active Directory identity provider.")
		licenseManagerUserSubscriptions_disassociateUserCmd.Flags().String("instance-id", "", "The ID of the EC2 instance which provides user-based subscriptions.")
		licenseManagerUserSubscriptions_disassociateUserCmd.Flags().String("instance-user-arn", "", "The Amazon Resource Name (ARN) of the user to disassociate from the EC2 instance.")
		licenseManagerUserSubscriptions_disassociateUserCmd.Flags().String("username", "", "The user name from the Active Directory identity provider for the user.")
	})
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_disassociateUserCmd)
}
