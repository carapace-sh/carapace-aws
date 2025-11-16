package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_deleteLicenseConfigurationCmd = &cobra.Command{
	Use:   "delete-license-configuration",
	Short: "Deletes the specified license configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_deleteLicenseConfigurationCmd).Standalone()

	licenseManager_deleteLicenseConfigurationCmd.Flags().String("license-configuration-arn", "", "ID of the license configuration.")
	licenseManager_deleteLicenseConfigurationCmd.MarkFlagRequired("license-configuration-arn")
	licenseManagerCmd.AddCommand(licenseManager_deleteLicenseConfigurationCmd)
}
