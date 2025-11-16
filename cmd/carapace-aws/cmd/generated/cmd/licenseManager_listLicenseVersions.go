package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listLicenseVersionsCmd = &cobra.Command{
	Use:   "list-license-versions",
	Short: "Lists all versions of the specified license.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listLicenseVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_listLicenseVersionsCmd).Standalone()

		licenseManager_listLicenseVersionsCmd.Flags().String("license-arn", "", "Amazon Resource Name (ARN) of the license.")
		licenseManager_listLicenseVersionsCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
		licenseManager_listLicenseVersionsCmd.Flags().String("next-token", "", "Token for the next set of results.")
		licenseManager_listLicenseVersionsCmd.MarkFlagRequired("license-arn")
	})
	licenseManagerCmd.AddCommand(licenseManager_listLicenseVersionsCmd)
}
