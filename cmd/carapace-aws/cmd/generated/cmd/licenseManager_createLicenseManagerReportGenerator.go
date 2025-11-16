package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_createLicenseManagerReportGeneratorCmd = &cobra.Command{
	Use:   "create-license-manager-report-generator",
	Short: "Creates a report generator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_createLicenseManagerReportGeneratorCmd).Standalone()

	licenseManager_createLicenseManagerReportGeneratorCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	licenseManager_createLicenseManagerReportGeneratorCmd.Flags().String("description", "", "Description of the report generator.")
	licenseManager_createLicenseManagerReportGeneratorCmd.Flags().String("report-context", "", "Defines the type of license configuration the report generator tracks.")
	licenseManager_createLicenseManagerReportGeneratorCmd.Flags().String("report-frequency", "", "Frequency by which reports are generated.")
	licenseManager_createLicenseManagerReportGeneratorCmd.Flags().String("report-generator-name", "", "Name of the report generator.")
	licenseManager_createLicenseManagerReportGeneratorCmd.Flags().String("tags", "", "Tags to add to the report generator.")
	licenseManager_createLicenseManagerReportGeneratorCmd.Flags().String("type", "", "Type of reports to generate.")
	licenseManager_createLicenseManagerReportGeneratorCmd.MarkFlagRequired("client-token")
	licenseManager_createLicenseManagerReportGeneratorCmd.MarkFlagRequired("report-context")
	licenseManager_createLicenseManagerReportGeneratorCmd.MarkFlagRequired("report-frequency")
	licenseManager_createLicenseManagerReportGeneratorCmd.MarkFlagRequired("report-generator-name")
	licenseManager_createLicenseManagerReportGeneratorCmd.MarkFlagRequired("type")
	licenseManagerCmd.AddCommand(licenseManager_createLicenseManagerReportGeneratorCmd)
}
