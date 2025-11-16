package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_updateLicenseManagerReportGeneratorCmd = &cobra.Command{
	Use:   "update-license-manager-report-generator",
	Short: "Updates a report generator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_updateLicenseManagerReportGeneratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_updateLicenseManagerReportGeneratorCmd).Standalone()

		licenseManager_updateLicenseManagerReportGeneratorCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		licenseManager_updateLicenseManagerReportGeneratorCmd.Flags().String("description", "", "Description of the report generator.")
		licenseManager_updateLicenseManagerReportGeneratorCmd.Flags().String("license-manager-report-generator-arn", "", "Amazon Resource Name (ARN) of the report generator to update.")
		licenseManager_updateLicenseManagerReportGeneratorCmd.Flags().String("report-context", "", "The report context.")
		licenseManager_updateLicenseManagerReportGeneratorCmd.Flags().String("report-frequency", "", "Frequency by which reports are generated.")
		licenseManager_updateLicenseManagerReportGeneratorCmd.Flags().String("report-generator-name", "", "Name of the report generator.")
		licenseManager_updateLicenseManagerReportGeneratorCmd.Flags().String("type", "", "Type of reports to generate.")
		licenseManager_updateLicenseManagerReportGeneratorCmd.MarkFlagRequired("client-token")
		licenseManager_updateLicenseManagerReportGeneratorCmd.MarkFlagRequired("license-manager-report-generator-arn")
		licenseManager_updateLicenseManagerReportGeneratorCmd.MarkFlagRequired("report-context")
		licenseManager_updateLicenseManagerReportGeneratorCmd.MarkFlagRequired("report-frequency")
		licenseManager_updateLicenseManagerReportGeneratorCmd.MarkFlagRequired("report-generator-name")
		licenseManager_updateLicenseManagerReportGeneratorCmd.MarkFlagRequired("type")
	})
	licenseManagerCmd.AddCommand(licenseManager_updateLicenseManagerReportGeneratorCmd)
}
