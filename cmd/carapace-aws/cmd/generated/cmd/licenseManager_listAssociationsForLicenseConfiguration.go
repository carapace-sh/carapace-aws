package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listAssociationsForLicenseConfigurationCmd = &cobra.Command{
	Use:   "list-associations-for-license-configuration",
	Short: "Lists the resource associations for the specified license configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listAssociationsForLicenseConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_listAssociationsForLicenseConfigurationCmd).Standalone()

		licenseManager_listAssociationsForLicenseConfigurationCmd.Flags().String("license-configuration-arn", "", "Amazon Resource Name (ARN) of a license configuration.")
		licenseManager_listAssociationsForLicenseConfigurationCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
		licenseManager_listAssociationsForLicenseConfigurationCmd.Flags().String("next-token", "", "Token for the next set of results.")
		licenseManager_listAssociationsForLicenseConfigurationCmd.MarkFlagRequired("license-configuration-arn")
	})
	licenseManagerCmd.AddCommand(licenseManager_listAssociationsForLicenseConfigurationCmd)
}
