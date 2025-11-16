package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_createLicenseConversionTaskForResourceCmd = &cobra.Command{
	Use:   "create-license-conversion-task-for-resource",
	Short: "Creates a new license conversion task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_createLicenseConversionTaskForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_createLicenseConversionTaskForResourceCmd).Standalone()

		licenseManager_createLicenseConversionTaskForResourceCmd.Flags().String("destination-license-context", "", "Information that identifies the license type you are converting to.")
		licenseManager_createLicenseConversionTaskForResourceCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the resource you are converting the license type for.")
		licenseManager_createLicenseConversionTaskForResourceCmd.Flags().String("source-license-context", "", "Information that identifies the license type you are converting from.")
		licenseManager_createLicenseConversionTaskForResourceCmd.MarkFlagRequired("destination-license-context")
		licenseManager_createLicenseConversionTaskForResourceCmd.MarkFlagRequired("resource-arn")
		licenseManager_createLicenseConversionTaskForResourceCmd.MarkFlagRequired("source-license-context")
	})
	licenseManagerCmd.AddCommand(licenseManager_createLicenseConversionTaskForResourceCmd)
}
