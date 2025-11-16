package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_listLicenseServerEndpointsCmd = &cobra.Command{
	Use:   "list-license-server-endpoints",
	Short: "List the Remote Desktop Services (RDS) License Server endpoints",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_listLicenseServerEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManagerUserSubscriptions_listLicenseServerEndpointsCmd).Standalone()

		licenseManagerUserSubscriptions_listLicenseServerEndpointsCmd.Flags().String("filters", "", "You can use the following filters to streamline results:")
		licenseManagerUserSubscriptions_listLicenseServerEndpointsCmd.Flags().String("max-results", "", "The maximum number of results to return from a single request.")
		licenseManagerUserSubscriptions_listLicenseServerEndpointsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	})
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_listLicenseServerEndpointsCmd)
}
