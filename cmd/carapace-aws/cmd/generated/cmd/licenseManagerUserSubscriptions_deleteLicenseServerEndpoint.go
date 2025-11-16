package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_deleteLicenseServerEndpointCmd = &cobra.Command{
	Use:   "delete-license-server-endpoint",
	Short: "Deletes a `LicenseServerEndpoint` resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_deleteLicenseServerEndpointCmd).Standalone()

	licenseManagerUserSubscriptions_deleteLicenseServerEndpointCmd.Flags().String("license-server-endpoint-arn", "", "The Amazon Resource Name (ARN) that identifies the `LicenseServerEndpoint` resource to delete.")
	licenseManagerUserSubscriptions_deleteLicenseServerEndpointCmd.Flags().String("server-type", "", "The type of License Server that the delete request refers to.")
	licenseManagerUserSubscriptions_deleteLicenseServerEndpointCmd.MarkFlagRequired("license-server-endpoint-arn")
	licenseManagerUserSubscriptions_deleteLicenseServerEndpointCmd.MarkFlagRequired("server-type")
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_deleteLicenseServerEndpointCmd)
}
