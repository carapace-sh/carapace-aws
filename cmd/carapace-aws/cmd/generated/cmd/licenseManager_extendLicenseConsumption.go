package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_extendLicenseConsumptionCmd = &cobra.Command{
	Use:   "extend-license-consumption",
	Short: "Extends the expiration date for license consumption.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_extendLicenseConsumptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_extendLicenseConsumptionCmd).Standalone()

		licenseManager_extendLicenseConsumptionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request.")
		licenseManager_extendLicenseConsumptionCmd.Flags().String("license-consumption-token", "", "License consumption token.")
		licenseManager_extendLicenseConsumptionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request.")
		licenseManager_extendLicenseConsumptionCmd.MarkFlagRequired("license-consumption-token")
		licenseManager_extendLicenseConsumptionCmd.Flag("no-dry-run").Hidden = true
	})
	licenseManagerCmd.AddCommand(licenseManager_extendLicenseConsumptionCmd)
}
