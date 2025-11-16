package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_getLicenseManagerReportGeneratorCmd = &cobra.Command{
	Use:   "get-license-manager-report-generator",
	Short: "Gets information about the specified report generator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_getLicenseManagerReportGeneratorCmd).Standalone()

	licenseManager_getLicenseManagerReportGeneratorCmd.Flags().String("license-manager-report-generator-arn", "", "Amazon Resource Name (ARN) of the report generator.")
	licenseManager_getLicenseManagerReportGeneratorCmd.MarkFlagRequired("license-manager-report-generator-arn")
	licenseManagerCmd.AddCommand(licenseManager_getLicenseManagerReportGeneratorCmd)
}
