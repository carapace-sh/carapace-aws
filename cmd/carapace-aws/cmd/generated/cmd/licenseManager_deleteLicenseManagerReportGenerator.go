package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_deleteLicenseManagerReportGeneratorCmd = &cobra.Command{
	Use:   "delete-license-manager-report-generator",
	Short: "Deletes the specified report generator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_deleteLicenseManagerReportGeneratorCmd).Standalone()

	licenseManager_deleteLicenseManagerReportGeneratorCmd.Flags().String("license-manager-report-generator-arn", "", "Amazon Resource Name (ARN) of the report generator to be deleted.")
	licenseManager_deleteLicenseManagerReportGeneratorCmd.MarkFlagRequired("license-manager-report-generator-arn")
	licenseManagerCmd.AddCommand(licenseManager_deleteLicenseManagerReportGeneratorCmd)
}
