package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_getLicenseConversionTaskCmd = &cobra.Command{
	Use:   "get-license-conversion-task",
	Short: "Gets information about the specified license type conversion task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_getLicenseConversionTaskCmd).Standalone()

	licenseManager_getLicenseConversionTaskCmd.Flags().String("license-conversion-task-id", "", "ID of the license type conversion task to retrieve information on.")
	licenseManager_getLicenseConversionTaskCmd.MarkFlagRequired("license-conversion-task-id")
	licenseManagerCmd.AddCommand(licenseManager_getLicenseConversionTaskCmd)
}
