package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listFailuresForLicenseConfigurationOperationsCmd = &cobra.Command{
	Use:   "list-failures-for-license-configuration-operations",
	Short: "Lists the license configuration operations that failed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listFailuresForLicenseConfigurationOperationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_listFailuresForLicenseConfigurationOperationsCmd).Standalone()

		licenseManager_listFailuresForLicenseConfigurationOperationsCmd.Flags().String("license-configuration-arn", "", "Amazon Resource Name of the license configuration.")
		licenseManager_listFailuresForLicenseConfigurationOperationsCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
		licenseManager_listFailuresForLicenseConfigurationOperationsCmd.Flags().String("next-token", "", "Token for the next set of results.")
		licenseManager_listFailuresForLicenseConfigurationOperationsCmd.MarkFlagRequired("license-configuration-arn")
	})
	licenseManagerCmd.AddCommand(licenseManager_listFailuresForLicenseConfigurationOperationsCmd)
}
