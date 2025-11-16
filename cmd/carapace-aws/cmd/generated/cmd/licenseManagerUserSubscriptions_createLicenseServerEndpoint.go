package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_createLicenseServerEndpointCmd = &cobra.Command{
	Use:   "create-license-server-endpoint",
	Short: "Creates a network endpoint for the Remote Desktop Services (RDS) license server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_createLicenseServerEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManagerUserSubscriptions_createLicenseServerEndpointCmd).Standalone()

		licenseManagerUserSubscriptions_createLicenseServerEndpointCmd.Flags().String("identity-provider-arn", "", "The Amazon Resource Name (ARN) that identifies the `IdentityProvider` resource that contains details about a registered identity provider.")
		licenseManagerUserSubscriptions_createLicenseServerEndpointCmd.Flags().String("license-server-settings", "", "The `LicenseServerSettings` resource to create for the endpoint.")
		licenseManagerUserSubscriptions_createLicenseServerEndpointCmd.Flags().String("tags", "", "The tags that apply for the license server endpoint.")
		licenseManagerUserSubscriptions_createLicenseServerEndpointCmd.MarkFlagRequired("identity-provider-arn")
		licenseManagerUserSubscriptions_createLicenseServerEndpointCmd.MarkFlagRequired("license-server-settings")
	})
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_createLicenseServerEndpointCmd)
}
